package models

import "time"

type Task struct {
	Id           int    `gorm:"primary_key" json:"id"`
	TaskName     string `form:"task_name"  json:"task_name" binding:"required"`
	TaskType     string `form:"task_type"  json:"task_type"`
	TaskCycle    string `form:"task_cycle"  json:"task_cycle"`
	Description  string `form:"description"  json:"description"`
	CronSpec     string `form:"cron_spec"  json:"cron_spec" binding:"required"`
	Concurrent   int    `form:"concurrent"  json:"concurrent" binding:"gte=0,lte=1"`
	Command      string `form:"command"  json:"command" binding:"required"`
	Arge         string `json:"arge"`
	Status       int    `json:"status"`
	VulnNum      int    `json:"vuln_num"`
	Timeout      int    `form:"timeout"  json:"timeout"`
	ExecuteTimes int
	PrevTime     int64
	CreatedTime  string `json:"created_time"`
	UpdatedTime  string `json:"updated_time"`
}

func AddTask(data map[string]interface{}) int {
	nowTime := time.Now().Format("20060102150405")
	task := Task{
		TaskName:    data["taskname"].(string),
		Description: data["description"].(string),
		CronSpec:    data["cronspec"].(string),
		Command:     data["command"].(string),
		Arge:        data["arge"].(string),
		TaskType:    data["tasktype"].(string),
		TaskCycle:   data["taskcycle"].(string),
		Status:      1,
		//Timeout:      data["protocol"].(string),
		//ExecuteTimes: data["protocol"].(string),
		//PrevTime:     data["protocol"].(string),
		CreatedTime: nowTime,
		UpdatedTime: nowTime,
	}
	db.Create(&task)

	return task.Id

}



func GetTaskById(id int) *Task {
	task := &Task{}
	db.Where("id = ?", id).First(task)
	return task
}


func EditTask(id int, data interface{}) bool {
	db.Model(&Task{}).Where("id = ?", id).Updates(data)
	return true
}
