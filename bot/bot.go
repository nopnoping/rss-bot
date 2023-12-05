package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"rssbot/rsspull/parse"
)

type Bot struct {
	botClient *tgbotapi.BotAPI
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
	}
}

func (b *Bot) Start() {
}
