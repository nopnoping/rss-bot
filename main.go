package main

import (
	"flag"
	"log"
	"rssbot/bot"
	"rssbot/config"
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
	flag.StringVar(&config.RssClientProxyURL, "rssproxy", "", "rss client proxy url")
	flag.StringVar(&config.Token, "token", "", "telegram bot token")
	flag.StringVar(&config.BotProxyURL, "botproxy", "", "bot proxy url")
	flag.StringVar(&config.DbPath, "db", "./rssbot.db", "sqlite db path")
	flag.Parse()

	taskInit()

	if config.Token == "" {
		log.Fatalf("token shoudn't be empty!")
	}

	b := bot.NewBot()
	go b.Start()

	select {}
}
