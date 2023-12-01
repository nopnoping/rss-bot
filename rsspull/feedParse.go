package rsspull

import "rssbot/rsspull/parse"

// The format of data could be json/xml
// if format is empty, it will judge data format by prefix
func parseChannel(data []byte, format string) *parse.FeedChannel {
	return nil
}

func parseItem(data []byte, format string) []*parse.FeedItem {
	return nil
}
