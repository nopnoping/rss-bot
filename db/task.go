package db

import "log"

type Task struct {
	TaskId    uint  `gorm:"primaryKey;column:task_id;unique"`
	StartTime int64 `gorm:"column:start_time"`
	Period    uint  `gorm:"column:period;unique"`
}

func CreateTask(task *Task) {
	database.Create(task)
}

func GetTasksByTime(time int64) []*Task {
	task := make([]*Task, 0)
	database.Where("start_time <= ?", time).Find(&task)
	return task
}

func UpdateTask(task []*Task) {
	if len(task) == 0 {
		log.Println("UpdateTask get a empty parameter")
		return
	}
	database.Save(&task)
}
