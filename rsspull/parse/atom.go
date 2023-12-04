package parse

import (
	"github.com/beevik/etree"
)

type parseTempAttr struct {
	title, link, item               string
	itemTitle, itemLink, itemUpdate string
}

func atomParseTemp(root *etree.Element, attr *parseTempAttr) *FeedInfo {
	feed := &FeedInfo{
		channel: &FeedChannel{},
		items:   make([]*FeedItem, 0),
	}

	if element := root.SelectElement(attr.title); element != nil {
		feed.channel.title = element.Text()
	}
	if link := root.SelectElement(attr.link); link != nil {
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				// todo: relative
				feed.channel.link = attr.Value
			}
		}
	}

	for _, entry := range root.SelectElements(attr.item) {
		item := &FeedItem{}
		if element := entry.SelectElement(attr.itemTitle); element != nil {
			item.title = element.Text()
		}
		if link := entry.SelectElement(attr.itemLink); link != nil {
			for _, attr := range link.Attr {
				if attr.Key == "href" {
					// todo: relative
					item.link = attr.Value
				}
			}
		}
		if date := entry.SelectElement(attr.itemUpdate); date != nil {
			// todo: format
			item.pubDate = date.Text()
		}
		feed.items = append(feed.items, item)
	}
	return feed
}

type atomV0_3 struct{}

var _ Parser = (*atomV0_3)(nil)

func (a atomV0_3) Parse(root *etree.Element) *FeedInfo {
	attr := &parseTempAttr{
		title:      "title",
		link:       "link",
		item:       "entry",
		itemTitle:  "title",
		itemLink:   "link",
		itemUpdate: "created",
	}
	return atomParseTemp(root, attr)
}

type atomV1_0 struct{}

var _ Parser = (*atomV1_0)(nil)

func (a atomV1_0) Parse(root *etree.Element) *FeedInfo {
	attr := &parseTempAttr{
		title:      "title",
		link:       "link",
		item:       "entry",
		itemTitle:  "title",
		itemLink:   "link",
		itemUpdate: "published",
	}
	return atomParseTemp(root, attr)
}

func init() {
	ParserMap["atom-0.3"] = &atomV0_3{}
	ParserMap["atom-1.0"] = &atomV1_0{}
}
