package push_task

import (
	"fmt"
	"log"
	"rssbot/config"
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
	User *db.User
	Info *parse.FeedInfo
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
	log.Println("push task start......")
	tick := time.Tick(config.PushTaskPeriod)

	for {
		select {
		case <-p.shutdown:
			log.Println("push task shutdown!")
			return
		case <-tick:
			log.Println("push task kick a tick......")
			users := db.GetCurrentCanPullUserAndUpdateTask()
			if users == nil {
				break
			}

			// use url to group
			urlPull := make(map[string][]*db.User)
			for _, user := range users {
				if _, ok := urlPull[user.Url]; !ok {
					urlPull[user.Url] = make([]*db.User, 0)
				}
				urlPull[user.Url] = append(urlPull[user.Url], user)
			}

			for url, users := range urlPull {
				p.wg.Add(1)
				go func(url string, users []*db.User) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("push task handle pull err:%v\n", r)
						}
					}()
					defer p.wg.Done()

					if feed := rsspull.GetDefaultRssPull().Pull(url); feed != nil {
						for _, user := range users {
							f := &parse.FeedInfo{
								Channel: feed.Channel,
								Items:   make([]*parse.FeedItem, 0),
							}
							for _, item := range feed.Items {
								if num, err := strconv.ParseInt(item.PubDate, 10, 64); err == nil && num >= user.PrevSendTime {
									f.Items = append(f.Items, item)
								}
							}
							if len(f.Items) > 0 {
								user.PrevSendTime = time.Now().Unix()
								p.msgCh <- &PushMsg{User: user, Info: f}
							}
						}
					}
				}(url, users)
			}
			// save
			p.wg.Wait()
		}
	}
}
