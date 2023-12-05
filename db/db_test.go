package db

import (
	"log"
	"testing"
	"time"
)

func dbInit() {
	// clear
	db.Where("1 = 1").Delete(&User{})
	db.Where("1 = 1").Delete(&Task{})
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
	db.Create(task1)
	db.Create(task2)
	db.Create(user1)
	db.Create(user2)

	u := GetCurrentCanPullUserAndUpdateTask()
	for _, t := range u {
		log.Println(t)
	}

	tasks := make([]Task, 0)
	db.Find(&tasks)
	for _, task := range tasks {
		log.Println(time.Unix(task.StartTime, 0).Format("2006-01-02 15:04:05"))
	}
}
