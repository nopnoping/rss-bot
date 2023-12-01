package parse

type FeedChannel struct {
	title string
	link  string
}

type FeedItem struct {
	title       string
	link        string
	description string
	date        string
}

var parseMap map[string]parse

type parse interface {
	GetChannel() *FeedChannel
	GetItem() []*FeedItem
}
