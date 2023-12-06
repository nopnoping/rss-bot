package main

import (
	"rssbot/bot"
	"rssbot/db"
)

func taskInit() {
	t := &db.Task{
		TaskId:    1,
		StartTime: 0,
		Period:    1,
	}
	db.CreateTask(t)
}

func main() {
	//config.BotProxyURL = ""
	taskInit()

	b := bot.NewBot()
	go b.Start()

	select {}
}
