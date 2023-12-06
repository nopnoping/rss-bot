package push_task

import (
	"log"
	"rssbot/db"
	"rssbot/rsspull"
	"rssbot/rsspull/parse"
	"strconv"
	"sync"
	"time"
)

type PushTask struct {
	msgCh    chan *PushMsg
	wg       sync.WaitGroup
	once     sync.Once
	shutdown chan struct{}
}

type PushMsg struct {
	ChatId int64
	Info   *parse.FeedInfo
}

func NewPushTask(msgCh chan *PushMsg) *PushTask {
	return &PushTask{msgCh: msgCh, shutdown: make(chan struct{})}
}

func (p *PushTask) Close() {
	p.once.Do(func() {
		close(p.shutdown)
	})
}

func (p *PushTask) Start() {
	tick := time.Tick(time.Minute * 5)
	for {
		select {
		case <-p.shutdown:
			log.Println("push task shutdown!")
			return
		case <-tick:
			users := db.GetCurrentCanPullUserAndUpdateTask()
			if users == nil {
				break
			}

			for _, user := range users {
				p.wg.Add(1)
				go func(user *db.User) {
					defer p.wg.Done()

					if feed := rsspull.DefaultRssPull.Pull(user.Url); feed != nil {

						items := make([]*parse.FeedItem, 0)
						for _, item := range feed.Items {
							if num, err := strconv.ParseInt(item.PubDate, 10, 64); err == nil && num >= user.PrevPullTime {
								items = append(items, item)
							}
						}
						feed.Items = items

						if len(feed.Items) > 0 {
							p.msgCh <- &PushMsg{ChatId: user.ChatId, Info: feed}
						}
					}

					user.PrevPullTime = time.Now().Unix()
				}(user)
			}
			// save
			p.wg.Wait()
			db.UpdateUsersPrevPullTime(users)
		}
	}
}
