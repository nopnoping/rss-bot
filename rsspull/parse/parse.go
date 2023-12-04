package parse

import "github.com/beevik/etree"

type FeedInfo struct {
	Channel *FeedChannel
	Items   []*FeedItem
}

type FeedChannel struct {
	Title string
	Link  string
}

type FeedItem struct {
	Title   string
	Link    string
	PubDate string
}

var ParserMap = make(map[string]Parser)

type Parser interface {
	Parse(root *etree.Element) *FeedInfo
}
