package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	"rssbot/config"
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

func NewBot() *Bot {
	var botClient *tgbotapi.BotAPI
	if config.BotProxyURL == "" {
		var err error
		botClient, err = tgbotapi.NewBotAPI(config.Token)
		if err != nil {
			log.Printf("connect bot err:%v\n", err)
			return nil
		}
	} else {
		proxyURL, err := url.Parse(config.BotProxyURL)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return nil
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}

		botClient, err = tgbotapi.NewBotAPIWithClient(config.Token, client)
		if err != nil {
			log.Printf("connect bot err:%v\n", err)
			return nil
		}
	}

	return &Bot{
		botClient: botClient,
		msgCh:     make(chan *pushtask.PushMsg, config.BotPushChCap),
		stop:      make(chan struct{}),
	}
}

func (b *Bot) Stop() {
	b.once.Do(func() {
		close(b.stop)
	})
}

func (b *Bot) Start() {
	log.Println("bot start.......")
	pushTask := pushtask.NewPushTask(b.msgCh)
	go pushTask.Start()

	// 设置消息处理器
	updateConfig := tgbotapi.NewUpdate(config.BotUpdateOffset)
	updateConfig.Timeout = config.BotUpdateTimeout

	receive, _ := b.botClient.GetUpdatesChan(updateConfig)

	for {
		select {
		case <-b.stop:
			return
		case pMsg := <-b.msgCh:
			log.Printf("get pushMsg. title:%s.", pMsg.Info.Channel.Title)
			go b.handlePush(pMsg)
		case rMsg := <-receive:
			go b.handleReq(rMsg)
		}
	}
}

func (b *Bot) handlePush(pMsg *pushtask.PushMsg) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("handlePush occur err:%v\n", r)
		}
	}()
	if len(pMsg.Info.Items) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintln(pMsg.Info.Channel.Title))

	for _, i := range pMsg.Info.Items {
		sb.WriteString(fmt.Sprintf("<a href=\"%s\">%s</a>\n", i.Link, i.Title))
	}

	msg := tgbotapi.NewMessage(pMsg.ChatId, sb.String())
	msg.ParseMode = "HTML"
	msg.DisableWebPagePreview = true
	if _, err := b.botClient.Send(msg); err != nil {
		log.Printf("send handlePush reply err:%v\n", err)
	}
}

func (b *Bot) handleReq(rMsg tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("handleReq occur err:%v\n", r)
		}
	}()
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
}
