package rsspull

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"rssbot/config"
	"rssbot/rsspull/parse"
	"strconv"
	"strings"
	"sync"
	"time"
)

var timeFormat = []string{
	"Mon, 02 Jan 2006 15:04:05 MST",
	"2006-01-02T15:04:05Z",
	"02 Jan 06 15:04 MST",
	"Mon, 02 Jan 2006 15:04:05",
	time.RFC3339,
}

type RssPull struct {
	client *rssClient
}

func NewRssPull() *RssPull {
	var c *http.Client
	if config.RssClientProxyURL == "" {
		c = &http.Client{
			Timeout: config.RssClientTimeOut,
		}
	} else {
		proxyURL, err := url.Parse(config.RssClientProxyURL)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return nil
		}
		c = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
			Timeout: config.RssClientTimeOut,
		}
	}

	return &RssPull{
		client: &rssClient{
			client: c,
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

	items := make([]*parse.FeedItem, 0)
itemLoops:
	for _, item := range feed.Items {
		for _, f := range timeFormat {
			if t, err := time.Parse(f, item.PubDate); err == nil {
				item.PubDate = strconv.FormatInt(t.Unix(), 10)
				items = append(items, item)
				continue itemLoops
			}
		}
		log.Printf("unsupported date format. date:%s.\n", item.PubDate)
	}
	feed.Items = items

	return feed
}

var once sync.Once
var defaultRssPull *RssPull

func GetDefaultRssPull() *RssPull {
	once.Do(func() {
		defaultRssPull = NewRssPull()
	})
	return defaultRssPull
}
