package db

type Task struct {
	TaskId    uint  `gorm:"primaryKey;column:task_id"`
	StartTime int64 `gorm:"column:start_time"`
	Period    uint  `gorm:"column:period;unique"`
}

func CreateTask(task *Task) {
	db.Create(task)
}
