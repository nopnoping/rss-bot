package rsspull

import (
	"log"
	"testing"
	"time"
)

func TestAtomPull(t *testing.T) {
	pull := NewRssPull()
	feed := pull.Pull("https://www.ruanyifeng.com/blog/atom.xml")
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}

func TestRSSPull(t *testing.T) {
	pull := NewRssPull()
	feed := pull.Pull("https://feeds.appinn.com/appinns/")
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}

func TestTime(t *testing.T) {
	date := "Mon, 04 Dec 2023 07:02:27 +0000"
	parse, _ := time.Parse(time.RFC1123, date)
	log.Println(parse.Unix())
}
