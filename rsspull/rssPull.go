package rsspull

import (
	"log"
	"net/http"
	"rssbot/rsspull/parse"
	"strings"
	"time"
)

type RssPull struct {
	client *rssClient
}

func NewRssPull() *RssPull {
	return &RssPull{
		client: &rssClient{
			client: &http.Client{
				Timeout: time.Second,
			},
		},
	}
}

func (r *RssPull) Pull(url string) *parse.FeedInfo {
	var feed *parse.FeedInfo
	body, header, err := r.client.get(url)
	if err != nil {
		log.Printf("Rss Pull get err:%v\n", err)
		return nil
	}
	if strings.Contains(header.Get("Content-Type"), "json") || strings.HasSuffix(url, "json") {
		if feed, err = parseFeed(body, "json"); err != nil {
			log.Printf("parse json feed error:%v\n", err)
			return nil
		}
	} else {
		if feed, err = parseFeed(body, "xml"); err != nil {
			log.Printf("parse xml feed error:%v\n", err)
			return nil
		}
	}
	return feed
}
