package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB

func init() {
	var err error
	if Db, err = gorm.Open(sqlite.Open("rssbot.db"), &gorm.Config{}); err != nil {
		log.Fatalf("database connect err:%v\n", err)
	}

	if err = Db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("automigrate task table err:%v\n", err)
	}

	if err = Db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("automigrate user table err:%v\n", err)
	}
}

func GetCurrentCanPullUserAndUpdateTask() []*User {
	tasks := GetTasksByTime(time.Now().Unix())
	if len(tasks) == 0 {
		return nil
	}

	taskIds := make([]uint, len(tasks))
	for _, task := range tasks {
		task.StartTime = time.Now().Unix() + int64((time.Duration)(task.Period)*5*60)
		taskIds = append(taskIds, task.TaskId)
	}
	UpdateTask(tasks)

	users := GetUsersByTaskIds(taskIds)
	if len(users) == 0 {
		return nil
	}
	return users
}
