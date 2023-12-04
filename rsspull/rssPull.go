package rsspull

import (
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

func (r *RssPull) Pull(url string) (feed *parse.FeedInfo) {
	body, header, _ := r.client.get(url)
	if strings.Contains(header.Get("Content-Type"), "xml") {
		feed = parseFeed(body, "xml")
	}
	return feed
}
