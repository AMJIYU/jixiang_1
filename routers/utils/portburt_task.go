package utils
//端口爆破
import (
	"crypto/md5"
	"fmt"
	"github.com/fatih/color"
	"io"
	"jixiang/common"
	"jixiang/models"
	"jixiang/plugins"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)


var (
	mutex       sync.Mutex
	successHash map[string]bool
	bruteResult map[string]models.Service
)


func NewPortBrute( ipprotocol string,thread,taskid int) {

	ipList := make([]models.IpAddr, 0)
	var port int
	var protocol string
	if ipprotocol == "mysql" {
		port = 3306
		protocol = "MYSQL"
	} else if ipprotocol == "ssh" {
		port = 22
		protocol = "SSH"
	}else if ipprotocol == "ftp" {
		port = 21
		protocol = "FTP"
	}else if ipprotocol == "smb" {
		port = 445
		protocol = "SMB"
	}else if ipprotocol == "mssql" {
		port = 1433
		protocol = "MSSQL"
	}else if ipprotocol == "postgresql" {
		port = 5432
		protocol = "POSTGRESQL"
	}else if ipprotocol == "mongodb" {
		port = 27017
		protocol = "MONGODB"
	}else if ipprotocol == "redis" {
		port = 6379
		protocol = "REDIS"
	}
	// 从数据库取结果
	datalists := models.GetIplistBrute(port, protocol)
	for _, target := range datalists {
			tmpPort, _ := strconv.Atoi(target.Port)
			addr := models.IpAddr{Ip: target.Ip, Port: tmpPort, Protocol: protocol}
			ipList = append(ipList, addr)
		}


	userDict, uErr := common.Readfile("./configs/user.txt")
	passDict, pErr := common.Readfile("./configs/pass.txt")
	taskruntime := time.Now().Format("20060102150405")

	if uErr == nil && pErr == nil {
		scanTasks := GenerateTask(ipList, userDict, passDict)
		//爆破之前记录任务id到日志
		data := make(map[string]interface{})
		data["taskid"] = taskid
		data["created_time"] = taskruntime
		data["all_num"] = len(ipList)
		data["userdict"] = 0
		data["passdict"] = 0
		models.AddTaskLog(data)
		RunTask(scanTasks, thread, taskid,taskruntime)
	} else {
		fmt.Println("Read File Err!")
	}
	//根据任务id 取出爆破成功数量
	maps := make(map[string]interface{})
	maps["task_id"] = taskid
	succesnum := models.GetPortBruteResTotal(maps)
	data := make(map[string]interface{})
	data["vuln_num"] = succesnum
	models.EditTask(taskid, data)

	mapsTaskLog := make(map[string]interface{})
	mapsTaskLog["task_id"] = taskid
	succesnumTaskLog := models.GetPortBruteResTotal(mapsTaskLog)

	dataTaskLog := make(map[string]interface{})
	dataTaskLog["vuln_num"] = succesnumTaskLog
	models.EditTaskLogTaskTime(taskid, taskruntime,dataTaskLog)

}


//为生产者做准备，常见端口爆破，返回多个Service类型任务
func GenerateTask(addr []models.IpAddr, userList []string, passList []string) (scanTasks []models.Service) {
	//每个都生成一个空的账号密码，用于爆破空账号密码
	scanTasks = make([]models.Service, 0)
	for _, ip := range addr {
		if ip.Protocol == "REDIS" || ip.Protocol == "FTP" || ip.Protocol == "POSTGRESQL" || ip.Protocol == "SSH" {
			scanTask := models.Service{Ip: ip.Ip, Port: ip.Port, Protocol: ip.Protocol, UserName: "", PassWord: ""}
			scanTasks = append(scanTasks, scanTask)
		}
	}
	//更具用户名与密码字典，爆破账号密码
	for _, u := range userList {
		for _, p := range passList {
			for _, ip := range addr {
				scanTask := models.Service{Ip: ip.Ip, Port: ip.Port, Protocol: ip.Protocol, UserName: u, PassWord: p}
				scanTasks = append(scanTasks, scanTask)
			}
		}
	}

	return
}



func RunTask(scanTasks []models.Service, thread,taskid int,taskruntime string) {
	start := time.Now()
	wg := &sync.WaitGroup{}
	successHash = make(map[string]bool)
	bruteResult = make(map[string]models.Service)

	// 创建一个buffer为thread * 2的channel
	taskChan := make(chan models.Service, thread*2)

	// 创建Thread个协程
	for i := 0; i < thread; i++ {
		go runBrute(taskChan,taskid,taskruntime, wg)

	}

	// 生产者，不断的把生产要扫描的数据，存放到 channel，直到channel阻塞
	for _, task := range scanTasks {
		wg.Add(1)
		taskChan <- task
	}

	// 生产完成后，从生产方关闭task
	close(taskChan)
//等待一组线程全部执行完成
	wg.Wait()
	waitTimeout(wg, 3*time.Second)

	color.Red("Scan complete. %d vulnerabilities found! \n", len(bruteResult))
	//end := time.Now()
	costTime := time.Since(start)
	data := make(map[string]interface{})

	// 获取这个任务id扫描成功的数量
	maps := make(map[string]interface{})
	maps["task_id"] = taskid
	maps["task_time"] = taskruntime
	succesnum := models.GetPortBruteResTotal(maps)
	fmt.Println("succesnum:",succesnum)
	fmt.Println(" end.Sub(start):", costTime)

	data["succes_num"] = succesnum
	data["run_time"] = fmt.Sprintf("%s",costTime)
	data["status"] = 2
	models.EditTaskLogTaskId(taskid,data)

}

// 消费者 每个协程都从生产者的channel中读取数据后，开启扫描
func runBrute(taskChan chan models.Service,taskid int, taskruntime string, wg *sync.WaitGroup) {
//遍历切片需要异常处理
	for target := range taskChan {
		//fmt.Println("now is :runBrute ",target)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("错误:runBrute error", target,err)
			}
		}()

		protocol := strings.ToUpper(target.Protocol)

		var k string
		if protocol == "REDIS" || protocol == "FTP" || protocol == "SNMP" || protocol == "POSTGRESQL" || protocol == "SSH" {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.UserName)
		}

		// 生成唯一hask
		h := MakeTaskHash(k)

		if checkTashHash(h) {
			wg.Done()
			continue
		}
		fmt.Fprintf(os.Stdout, "Now is %s %s %s\r", target.Ip,target.UserName, target.PassWord)
		err, res := plugins.ScanFuncMap[protocol](target.Ip, strconv.Itoa(target.Port), target.UserName, target.PassWord)
		if err == nil && res == true {
			saveRes(target,  h,taskruntime,taskid)
		} else {
			//fmt.Println("插件爆破时错误:", err)
		}
		wg.Done()
	}

}

//设置超时时间
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // 超时未响应
	}
}


func MD5(s string) (m string) {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
//给每个任务名变成hash
func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}

// 标记特定服务的特定用户是否破解成功，成功的话不再尝试破解该用户
//SuccessHash map[string]bool hash唯一
func checkTashHash(hash string) bool {
	_, ok := successHash[hash]
	return ok
}


func setTaskHask(hash string) () {
	mutex.Lock()
	successHash[hash] = true
	mutex.Unlock()
}

func saveRes(target models.Service, h ,taskruntime string ,taskid int  ) {

	setTaskHask(h)
	_, ok := bruteResult[h]
	if !ok {
		mutex.Lock()
		//爆破结果写入数据库中
		color.Cyan("[+] %s %d %s %s \n", target.Ip, target.Port, target.UserName, target.PassWord)
		data := make(map[string]interface{})
		data["ip"] = target.Ip
		data["port"] = target.Port
		data["protocol"] = target.Protocol
		data["user"] = target.UserName
		data["pass"] = target.PassWord
		data["taskid"] = taskid
		//data["vulntype"] = protocol // todo:漏洞类型要判断下
		//data["task_id"] = protocol
		data["task_time"] = taskruntime
		models.AddPortBruteRes(data)

		bruteResult[h] = models.Service{Ip: target.Ip, Port: target.Port, Protocol: target.Protocol, UserName: target.UserName, PassWord: target.PassWord}
		mutex.Unlock()
	}
}
