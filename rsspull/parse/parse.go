package parse

import "github.com/beevik/etree"

type FeedInfo struct {
	channel *FeedChannel
	items   []*FeedItem
}

type FeedChannel struct {
	title string
	link  string
}

type FeedItem struct {
	title   string
	link    string
	pubDate string
}

var ParserMap = make(map[string]Parser)

type Parser interface {
	Parse(root *etree.Element) *FeedInfo
}
