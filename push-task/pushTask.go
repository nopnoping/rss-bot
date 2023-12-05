package push_task

import (
	"rssbot/bot"
	"rssbot/db"
	"rssbot/rsspull"
	"rssbot/rsspull/parse"
	"strconv"
	"time"
)

type PushTask struct {
	msgCh chan bot.PushMsg
}

func NewPushTask() *PushTask {
	return &PushTask{}
}

func (p *PushTask) Start() {
	tick := time.Tick(time.Minute * 5)
	for {
		select {
		case <-tick:
			users := db.GetCurrentCanPullUserAndUpdateTask()
			for _, user := range users {
				if feed := rsspull.DefaultRssPull.Pull(user.Url); feed != nil {

					items := make([]*parse.FeedItem, 0)
					for _, item := range feed.Items {
						if num, err := strconv.ParseInt(item.PubDate, 10, 64); err == nil && num >= user.PrevSendTime {
							items = append(items, item)
						}
					}
					feed.Items = items

					p.msgCh <- bot.PushMsg{TwitterId: user.TwitterId, Info: feed}
					user.PrevSendTime = time.Now().Unix()
				}
			}
			// save
			db.UpdateUsers(users)
		}
	}
}
