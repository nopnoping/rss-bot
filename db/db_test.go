package db

import (
	"log"
	"testing"
	"time"
)

func dbInit() {
	// clear
	Db.Where("1 = 1").Delete(&User{})
	Db.Where("1 = 1").Delete(&Task{})
}
func TestDb(t *testing.T) {
	dbInit()
	cur := time.Now().Format("2006-01-02 15:04:05")
	log.Println(cur)

	task1 := &Task{
		TaskId:    1,
		StartTime: time.Now().Unix() - 100,
		Period:    1,
	}
	task2 := &Task{
		TaskId:    2,
		StartTime: time.Now().Unix() - 100,
		Period:    2,
	}

	user1 := &User{
		TwitterId: 123131,
		Url:       "www.luexp.com",
		TaskId:    1,
	}
	user2 := &User{
		TwitterId: 123131,
		Url:       "www.luexp.com",
		TaskId:    2,
	}
	Db.Create(task1)
	Db.Create(task2)
	Db.Create(user1)
	Db.Create(user2)

	u := GetCurrentCanPullUserAndUpdateTask()
	for _, t := range u {
		log.Println(t)
	}

	tasks := make([]Task, 0)
	Db.Find(&tasks)
	for _, task := range tasks {
		log.Println(time.Unix(task.StartTime, 0).Format("2006-01-02 15:04:05"))
	}
}
