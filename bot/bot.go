package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	pushtask "rssbot/push-task"
	"strings"
	"sync"
)

type Bot struct {
	botClient *tgbotapi.BotAPI
	msgCh     chan *pushtask.PushMsg
	once      sync.Once
	stop      chan struct{}
}

var DefaultProxyUrl = "http://127.0.0.1:7890"

func NewBot() *Bot {
	proxyURL, err := url.Parse(DefaultProxyUrl)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	b, err := tgbotapi.NewBotAPIWithClient("6807409395:AAFW90O-9G-1rPzhDv-q1ZsGUAYvBx9v74s", client)
	if err != nil {
		log.Printf("connect bot err:%v\n", err)
		return nil
	}
	return &Bot{
		botClient: b,
		msgCh:     make(chan *pushtask.PushMsg, 30),
		stop:      make(chan struct{}),
	}
}

func (b *Bot) Stop() {
	b.once.Do(func() {
		close(b.stop)
	})
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
		case <-b.stop:
			return
		case pMsg := <-b.msgCh:
			go b.handlePush(pMsg)
		case rMsg := <-receive:
			go b.handleReq(rMsg)
		}
	}
}

func (b *Bot) handlePush(pMsg *pushtask.PushMsg) {

}

func (b *Bot) handleReq(rMsg tgbotapi.Update) {
	txt := rMsg.Message.Text
	if len(txt) == 0 {
		return
	}

	// parse cmd
	cmdStart := strings.Index(txt, "/")
	if cmdStart == -1 {
		return
	}
	cmdEnd := strings.Index(txt, " ")
	if cmdEnd == -1 {
		cmdEnd = len(txt)
	}

	switch txt[cmdStart:cmdEnd] {
	case "/sub":
		b._sub(rMsg)
	case "/unsub":
		b._unsub(rMsg)
	case "/rss":
		b._rss(rMsg)
	case "/start":
		b._start(rMsg)
	case "/batch":
		b._batch(rMsg)
	default:
		b._start(rMsg)
	}
	//if rMsg.Message != nil {
	//	log.Printf("[%s] %s\n", rMsg.Message.From.UserName, rMsg.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, rMsg.Message.Text)
	//	msg.ReplyToMessageID = rMsg.Message.MessageID
	//
	//	m, _ := b.botClient.Send(msg)
	//	time.Sleep(time.Second)
	//	eM := tgbotapi.NewEditMessageText(rMsg.Message.Chat.ID, m.MessageID, "Done")
	//	b.botClient.Send(eM)
	//}
}
