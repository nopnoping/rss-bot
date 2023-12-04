package parse

import (
	"github.com/beevik/etree"
)

type rssV2_0 struct{}

var _ Parser = (*rssV2_0)(nil)

func (r rssV2_0) Parse(root *etree.Element) *FeedInfo {
	feed := &FeedInfo{
		channel: &FeedChannel{},
		items:   make([]*FeedItem, 0),
	}

	if channel := root.SelectElement("channel"); channel != nil {
		if title := channel.SelectElement("title"); title != nil {
			feed.channel.title = title.Text()
		}
		if link := channel.SelectElement("link"); link != nil {
			feed.channel.link = link.Text()
		}
		for _, i := range channel.SelectElements("item") {
			item := &FeedItem{}
			if title := i.SelectElement("title"); title != nil {
				item.title = title.Text()
			}
			if link := i.SelectElement("link"); link != nil {
				item.link = link.Text()
			}
			if pubDate := i.SelectElement("pubDate"); pubDate != nil {
				item.pubDate = pubDate.Text()
			}
			feed.items = append(feed.items, item)
		}
	}

	return feed
}

type rssV1_0 struct{}

var _ Parser = (*rssV1_0)(nil)

func (r rssV1_0) Parse(root *etree.Element) *FeedInfo {
	feed := &FeedInfo{
		channel: &FeedChannel{},
		items:   make([]*FeedItem, 0),
	}

	if channel := root.SelectElement("channel"); channel != nil {
		if title := channel.SelectElement("title"); title != nil {
			feed.channel.title = title.Text()
		}
		if link := channel.SelectElement("link"); link != nil {
			feed.channel.link = link.Text()
		}
	}

	for _, i := range root.SelectElements("item") {
		item := &FeedItem{}
		if title := i.SelectElement("title"); title != nil {
			item.title = title.Text()
		}
		if link := i.SelectElement("link"); link != nil {
			item.link = link.Text()
		}
		if pubDate := i.SelectElement("dc:date"); pubDate != nil {
			item.pubDate = pubDate.Text()
		}
		feed.items = append(feed.items, item)
	}

	return feed
}

func init() {
	ParserMap["rss-2.0"] = &rssV2_0{}
	ParserMap["rss-1.0"] = &rssV1_0{}
	ParserMap["rss-0.94"] = &rssV2_0{}
	ParserMap["rss-0.93"] = &rssV2_0{}
	ParserMap["rss-0.92"] = &rssV2_0{}
	ParserMap["rss-0.91"] = &rssV1_0{}
	ParserMap["rss-0.9"] = &rssV1_0{}
}
