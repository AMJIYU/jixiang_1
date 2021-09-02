package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	libcron "github.com/robfig/cron/v3"
	. "jixiang/common"
	"jixiang/models"
	"jixiang/routers/middleware/jwt"
	"jixiang/routers/utils"
	"net/http"
	"strconv"
	"time"
)

// 新增定时任务
func AddTask(c *gin.Context) {

	var arge string

	taskname := c.PostForm("taskname")
	description := c.PostForm("description")

	//配合前端灵活输入 最终得到corntab命令， 只有cronspec参数入库
	cronspec := c.PostForm("cronspec")
	day := c.PostForm("day")
	hour := c.PostForm("hour")
	cronspecmd := c.PostForm("cronspecmd")

	//爆破类型
	brute := c.PostForm("brute")
	source := c.PostForm("source")
	sourcecontent := c.PostForm("sourcecontent")

	thread := c.PostForm("thread")
	ipprotocol := c.PostForm("ipprotocol")
	command := "portbrute"
	valid := validation.Validation{}

	var taskcycle string
	if cronspec == "now" {
		nowMin := time.Now().Format("1504")
		min, _ := strconv.Atoi(nowMin[2:])
		min = min + 1
		cronspec = strconv.Itoa(min) + " " + nowMin[:2] + " * * *"
		fmt.Println("执行一次:",cronspec)
		taskcycle = "now"
	} else if cronspec == "day" {
		cronspec = "0 " + hour + " * * *"
		taskcycle = "每天"+hour+"点"
		fmt.Println("cronspec:", cronspec)
	} else if cronspec == "week" {
		cronspec = "0 " + hour + " * * " + day + ""
		fmt.Println("cronspec:", cronspec)
		taskcycle = "每周"+day+"的"+hour+"点"
	} else if cronspec == "cmd" {
		cronspec = cronspecmd
	} else {
	}

	if source == "1" {
		fmt.Println("brute:", brute)
		sourcecontent = brute
	} else if source == "2" {
	}
	//arge = "source=" + source + "&iplist=" + sourcecontent + "&dict=1&thread=" + thread
	arge = sourcecontent+","+thread
	// 输入长度限制
	valid.Required(taskname, "taskname").Message("名称不能为空")

	code := INVALID_PARAMS
	if !valid.HasErrors() {
		if _, err := libcron.ParseStandard(cronspec); err != nil {
			code = ERROR_CRON_SPEC
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  jwt.GetMsg(code),
				"data": make(map[string]string),
			})
			return
		}

		go func() {

			data := make(map[string]interface{})
			data["taskname"] = taskname
			data["description"] = description
			data["cronspec"] = cronspec
			data["command"] = command
			data["arge"] = arge
			data["tasktype"] = brute
			data["taskcycle"] = taskcycle
			taskId := models.AddTask(data)
			thread ,_ := strconv.Atoi(thread)
			utils.NewPortBrute(ipprotocol,thread,taskId)



		}()

		code = SUCCESS

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  jwt.GetMsg(code),
		"data": make(map[string]string),
	})

}
