package push_task

import (
	"rssbot/db"
	"runtime"
	"testing"
	"time"
)

func dbInit() {
	// clear
	db.Db.Where("1 = 1").Delete(&db.User{})
	db.Db.Where("1 = 1").Delete(&db.Task{})

	t := &db.Task{
		TaskId:    1,
		StartTime: 0,
		Period:    1,
	}
	db.CreateTask(t)

	u := &db.User{
		ChatId:       111,
		Url:          "https://www.ruanyifeng.com/blog/atom.xml",
		TaskId:       1,
		PrevPullTime: 0,
	}
	db.CreateUser(u)
}

func TestDb(t *testing.T) {
	runtime.GOMAXPROCS(1)
	dbInit()
	db.Db.Debug()
	ch := make(chan *PushMsg)
	push := NewPushTask(ch)
	go push.Start()

	after := time.After(20 * time.Second)

	for {
		select {
		case <-after:
			push.Close()
			time.Sleep(1 * time.Second)
			return
		case msg := <-ch:
			t.Log(msg.ChatId)
			t.Log(msg.Info.Channel.Title, msg.Info.Channel.Title)
			for _, i := range msg.Info.Items {
				t.Log(i.Title, i.Link, i.PubDate)
			}
		}
	}
}
