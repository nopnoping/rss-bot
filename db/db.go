package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"rssbot/config"
	"time"
)

var database *gorm.DB

func init() {
	var err error
	if database, err = gorm.Open(sqlite.Open(config.DbPath), &gorm.Config{}); err != nil {
		log.Fatalf("database connect err:%v\n", err)
	}

	if err = database.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("automigrate task table err:%v\n", err)
	}

	if err = database.AutoMigrate(&User{}); err != nil {
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
		//task.StartTime = time.Now().Unix() + int64((time.Duration)(task.Period)*5*60)
		task.StartTime = time.Now().Unix() + int64(config.TaskPeriodScale)
		taskIds = append(taskIds, task.TaskId)
	}
	UpdateTask(tasks)

	users := GetUsersByTaskIds(taskIds)
	if len(users) == 0 {
		return nil
	}
	return users
}
