package utils

import (
	"fmt"
	"jixiang/models"
	"time"
)

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *models.Task                                      // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	job := NewCommandJob(task)
	job.task = task
	job.Concurrent = task.Concurrent == 1
	return job, nil
}

func NewCommandJob(task *models.Task) *Job {
	job := &Job{
		id:   task.Id,
		name: task.TaskName,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
	//	arges := strings.Split(task.Arge,",")
	//	thread ,_ := strconv.Atoi(arges[1])
	//	NewPortBrute(arges[0],1,1,thread,task.Id)
		return "", "", nil,true
	}
	return job
}

func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("执行任务失败", err)
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	j.status++
	defer func() {
		j.status--
	}()

	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	_, _, err, _ := j.runFunc(timeout)

	if err !=nil{
		fmt.Println("eccadcjads",err)
	}


	// 如果是只执行一次的任务，修改任务状态为0，暂停任务
	if j.task.TaskCycle == "now"{
		data := make(map[string]interface{})
		data["status"] = 0
		models.EditTask(j.task.Id,data)
	}

}




