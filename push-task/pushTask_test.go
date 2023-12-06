package push_task

import (
	"testing"
	"time"
)

func TestDb(t *testing.T) {
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
