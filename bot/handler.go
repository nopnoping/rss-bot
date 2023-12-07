package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"rssbot/db"
	"rssbot/rsspull"
	"strings"
)

func (b *Bot) _sub(rMsg tgbotapi.Update) {
	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, "handling")
	msg.ReplyToMessageID = rMsg.Message.MessageID

	m, err := b.botClient.Send(msg)
	if err != nil {
		log.Printf("send _sub reply err:%v\n", err)
		return
	}

	words := strings.Fields(rMsg.Message.Text)
	text := ""
	if len(words) < 2 {
		text = "need url"
	} else if db.HasThisUrlWithTheChatId(rMsg.Message.Chat.ID, words[1]) {
		text = "has subscribed!"
	} else {
		url := words[1]
		if feed := rsspull.GetDefaultRssPull().Pull(url); feed == nil {
			text = "can't pull the url"
		} else {
			text = fmt.Sprintf("subscribe %s success!", feed.Channel.Title)
			user := &db.User{
				ChatId:       rMsg.Message.Chat.ID,
				Url:          url,
				PrevSendTime: 0,
				TaskId:       1,
				Title:        feed.Channel.Title,
			}
			db.CreateUser(user)
		}
	}

	eM := tgbotapi.NewEditMessageText(rMsg.Message.Chat.ID, m.MessageID, text)
	_, err = b.botClient.Send(eM)
	if err != nil {
		log.Printf("send _sub reply err:%v\n", err)
		return
	}
}

func (b *Bot) _unsub(rMsg tgbotapi.Update) {
	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, "handling")
	msg.ReplyToMessageID = rMsg.Message.MessageID

	m, err := b.botClient.Send(msg)
	if err != nil {
		log.Printf("send _sub reply err:%v\n", err)
		return
	}

	words := strings.Fields(rMsg.Message.Text)
	text := "Done!"
	if len(words) < 2 {
		text = "need url"
	} else {
		db.DeleteUserByChaiIdAndUrl(rMsg.Message.Chat.ID, words[1])
	}

	eM := tgbotapi.NewEditMessageText(rMsg.Message.Chat.ID, m.MessageID, text)
	_, err = b.botClient.Send(eM)
	if err != nil {
		log.Printf("send _sub reply err:%v\n", err)
		return

	}
}

func (b *Bot) _rss(rMsg tgbotapi.Update) {
	users := db.GetUserSubscribeUrls(rMsg.Message.Chat.ID)
	var sb strings.Builder
	sb.WriteString("subscribe list\n")

	for _, u := range users {
		sb.WriteString(fmt.Sprintf("<a href=\"%s\">%s</a>", u.Url, u.Title))
	}

	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, sb.String())
	msg.ParseMode = "HTML"
	msg.ReplyToMessageID = rMsg.Message.MessageID
	msg.DisableWebPagePreview = true
	if _, err := b.botClient.Send(msg); err != nil {
		log.Printf("send _start reply err:%v\n", err)
	}
}

func (b *Bot) _start(rMsg tgbotapi.Update) {
	info := `/start show manual
/sub subscribe rss web
/unsub unsubscribe rss web
/rss list subscribed rss web
/batch batch subscribe rss webs`

	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, info)
	msg.ReplyToMessageID = rMsg.Message.MessageID
	if _, err := b.botClient.Send(msg); err != nil {
		log.Printf("send _start reply err:%v\n", err)
	}
}

func (b *Bot) _batch(rMsg tgbotapi.Update) {
	info := "wait impl!"
	msg := tgbotapi.NewMessage(rMsg.Message.Chat.ID, info)
	msg.ReplyToMessageID = rMsg.Message.MessageID
	if _, err := b.botClient.Send(msg); err != nil {
		log.Printf("send _start reply err:%v\n", err)
	}
}
