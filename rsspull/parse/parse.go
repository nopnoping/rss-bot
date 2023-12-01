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

type parse interface {
	GetChannel()
	GetItem()
}
