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

func (r *RssPull) Pull(url string) (feed *parse.FeedInfo, err error) {
	body, header, err := r.client.get(url)
	if err != nil {
		return
	}
	if strings.Contains(header.Get("Content-Type"), "json") || strings.HasSuffix(url, "json") {
		feed = parseFeed(body, "json")
	} else {
		feed = parseFeed(body, "xml")
	}
	return
}
