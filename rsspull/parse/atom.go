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
		Channel: &FeedChannel{},
		Items:   make([]*FeedItem, 0),
	}

	if element := root.SelectElement(attr.title); element != nil {
		feed.Channel.Title = element.Text()
	}

	for _, link := range root.SelectElements(attr.link) {
		href := ""
		for _, attr := range link.Attr {
			if attr.Key == "rel" && attr.Value == "self" {
				continue
			}
			if attr.Key == "href" {
				// todo: relative
				href = attr.Value
			}
		}
		feed.Channel.Link = href
	}

	for _, entry := range root.SelectElements(attr.item) {
		item := &FeedItem{}
		if element := entry.SelectElement(attr.itemTitle); element != nil {
			item.Title = element.Text()
		}
		if link := entry.SelectElement(attr.itemLink); link != nil {
			for _, attr := range link.Attr {
				if attr.Key == "href" {
					// todo: relative
					item.Link = attr.Value
				}
			}
		}
		if date := entry.SelectElement(attr.itemUpdate); date != nil {
			// todo: format
			item.PubDate = date.Text()
		}
		feed.Items = append(feed.Items, item)
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
