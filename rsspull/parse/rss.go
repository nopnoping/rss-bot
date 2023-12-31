package parse

import (
	"github.com/beevik/etree"
)

type rssV2_0 struct{}

var _ Parser = (*rssV2_0)(nil)

func (r rssV2_0) Parse(root *etree.Element) *FeedInfo {
	feed := &FeedInfo{
		Channel: &FeedChannel{},
		Items:   make([]*FeedItem, 0),
	}

	if channel := root.SelectElement("channel"); channel != nil {
		if title := channel.SelectElement("title"); title != nil {
			feed.Channel.Title = title.Text()
		}
		// Patch: it sometimes has extra atom link element, We should pass it
		for _, link := range channel.SelectElements("link") {
			if link.Space == "" {
				feed.Channel.Link = link.Text()
			}
		}
		for _, i := range channel.SelectElements("item") {
			item := &FeedItem{}
			if title := i.SelectElement("title"); title != nil {
				item.Title = title.Text()
			}
			if link := i.SelectElement("link"); link != nil {
				item.Link = link.Text()
			}
			if pubDate := i.SelectElement("pubDate"); pubDate != nil {
				item.PubDate = pubDate.Text()
			}
			feed.Items = append(feed.Items, item)
		}
	}

	return feed
}

type rssV1_0 struct{}

var _ Parser = (*rssV1_0)(nil)

func (r rssV1_0) Parse(root *etree.Element) *FeedInfo {
	feed := &FeedInfo{
		Channel: &FeedChannel{},
		Items:   make([]*FeedItem, 0),
	}

	if channel := root.SelectElement("channel"); channel != nil {
		if title := channel.SelectElement("Title"); title != nil {
			feed.Channel.Title = title.Text()
		}
		if link := channel.SelectElement("Link"); link != nil {
			feed.Channel.Link = link.Text()
		}
	}

	for _, i := range root.SelectElements("item") {
		item := &FeedItem{}
		if title := i.SelectElement("title"); title != nil {
			item.Title = title.Text()
		}
		if link := i.SelectElement("link"); link != nil {
			item.Link = link.Text()
		}
		if pubDate := i.SelectElement("dc:date"); pubDate != nil {
			item.PubDate = pubDate.Text()
		}
		feed.Items = append(feed.Items, item)
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
