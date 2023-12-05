package db

type Task struct {
	TaskId    uint  `gorm:"primaryKey;column:task_id"`
	StartTime int64 `gorm:"column:start_time"`
	Period    uint  `gorm:"column:period;unique"`
}

func CreateTask(task *Task) {
	db.Create(task)
}

func GetTasksByTime(time int64) []*Task {
	task := make([]*Task, 0)
	db.Where("start_time <= ?", time).Find(&task)
	return task
}

func UpdateTask(task []*Task) {
	db.Save(&task)
}
