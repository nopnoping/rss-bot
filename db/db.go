package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	if db, err = gorm.Open(sqlite.Open("rssbot.db"), &gorm.Config{}); err != nil {
		log.Fatalf("database connect err:%v\n", err)
	}

	if err = db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("automigrate task table err:%v\n", err)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("automigrate user table err:%v\n", err)
	}
}

func GetCurrentCanPullUserAndUpdateTask() []*User {
	tasks := GetTasksByTime(time.Now().Unix())
	taskIds := make([]uint, len(tasks))
	for _, task := range tasks {
		task.StartTime = time.Now().Unix() + int64((time.Duration)(task.Period)*5*60)
		taskIds = append(taskIds, task.TaskId)
	}
	UpdateTask(tasks)

	users := GetUsersByTaskIds(taskIds)
	return users
}
