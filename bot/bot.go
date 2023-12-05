package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	pushtask "rssbot/push-task"
	"rssbot/rsspull/parse"
)

type Bot struct {
	botClient *tgbotapi.BotAPI
	msgCh     chan *PushMsg
}

type PushMsg struct {
	TwitterId int64
	Info      *parse.FeedInfo
}

func NewBot() *Bot {
	b, err := tgbotapi.NewBotAPI("6874948067:AAFIhFfrL1tsIe1S8t5FnzEgTMK4WK9QE_I\n")
	if err != nil {
		log.Printf("connect bot err:%v\n", err)
		return nil
	}
	return &Bot{
		botClient: b,
		msgCh:     make(chan *PushMsg),
	}
}

func (b *Bot) Start() {
	pushTask := pushtask.NewPushTask(b.msgCh)
	go pushTask.Start()

	// 设置消息处理器
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	receive, _ := b.botClient.GetUpdatesChan(updateConfig)

	for {
		select {
		case pMsg := <-b.msgCh:
			fmt.Println(pMsg)
		case rMsg := <-receive:
			fmt.Println(rMsg)
		}
	}
}
