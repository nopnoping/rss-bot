package rsspull

import (
	"encoding/json"
	"errors"
	"github.com/beevik/etree"
	"log"
	"rssbot/rsspull/parse"
	"strings"
)

// The format of data could be json/xml
// if format is empty, it will judge data format by prefix
func parseFeed(data []byte, format string) (*parse.FeedInfo, error) {
	switch format {
	case "xml":
		return parseXML(data)
	case "json":
		return parseJson(data)
	default:
		return nil, nil
	}
}

func parseXML(data []byte) (*parse.FeedInfo, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		log.Printf("Doc Read bytes err:%v\n", err)
		return nil, err
	}

	root := doc.Root()
	if root == nil {
		log.Println("Parse Feed err, Don't find root element!")
		return nil, errors.New("not xml format data")
	}

	key := ""
	switch root.Tag {
	case "rdf:RDF":
		for _, attr := range root.Attr {
			if attr.Key == "xmlns" {
				vs := strings.Split(attr.Value, "/")
				key = "rdf:RDF-" + vs[len(vs)-1]
				break
			}
		}
	case "feed":
		for _, attr := range root.Attr {
			if attr.Key == "version" {
				key = "atom-" + attr.Value
				break
			}
			if attr.Key == "xmlns" && attr.Value == "http://www.w3.org/2005/Atom" {
				key = "atom-1.0"
			}
		}
	case "rss":
		for _, attr := range root.Attr {
			if attr.Key == "version" {
				key = "rss-" + attr.Value
				break
			}
		}
	default:
		log.Printf("Parse Feed err, No matched Root Tag!. tag:%s\n", root.Tag)
		return nil, errors.New("unsupported version")
	}

	var p parse.Parser
	var ok bool
	if p, ok = parse.ParserMap[key]; !ok {
		log.Printf("Parse Feed err, No matched parser!. key:%s\n", key)
		return nil, errors.New("unsupported version")
	}
	return p.Parse(root), nil
}

func parseJson(data []byte) (*parse.FeedInfo, error) {
	feed := &parse.FeedInfo{
		Channel: &parse.FeedChannel{},
		Items:   make([]*parse.FeedItem, 0),
	}
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		log.Printf("Deserialized err:%v\n", err)
		return nil, err
	}

	if _, ok := temp["rss"]; ok {
		if err := parseRSSJson(temp, feed); err != nil {
			return nil, err
		}
	} else if _, ok := temp["feed"]; ok {
		if err := parseAtomJson(temp, feed); err != nil {
			return nil, err
		}
	} else {
		log.Println("Unsupported json format.")
		return nil, errors.New("unsupported json format")
	}

	return feed, nil
}

func parseRSSJson(temp map[string]interface{}, feed *parse.FeedInfo) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("parse format error")
		}
	}()
	for k, v := range temp["rss"].(map[string]interface{}) {
		if k == "channel" {
			for k1, v1 := range v.(map[string]interface{}) {
				switch k1 {
				case "title":
					feed.Channel.Title = v1.(string)
				case "link":
					feed.Channel.Link = v1.(string)
				case "items":
					for _, v2 := range v1.([]interface{}) {
						t := v2.(map[string]interface{})
						item := &parse.FeedItem{
							Title:   t["title"].(string),
							Link:    t["link"].(string),
							PubDate: t["pubDate"].(string),
						}
						feed.Items = append(feed.Items, item)
					}
				}
			}
		}
	}
	return
}

func parseAtomJson(temp map[string]interface{}, feed *parse.FeedInfo) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("parse format error")
		}
	}()
	for k, v := range temp["feed"].(map[string]interface{}) {
		switch k {
		case "title":
			feed.Channel.Title = v.(string)
		case "link":
			for _, v1 := range v.([]interface{}) {
				t := v1.(map[string]interface{})
				if val, ok := t["rel"]; ok && val == "self" {
					continue
				}
				feed.Channel.Link = t["href"].(string)
				break
			}
		case "entry":
			for _, v2 := range v.([]interface{}) {
				t := v2.(map[string]interface{})
				item := &parse.FeedItem{
					Title:   t["title"].(string),
					Link:    t["link"].([]interface{})[0].(map[string]interface{})["href"].(string),
					PubDate: t["updated"].(string),
				}
				feed.Items = append(feed.Items, item)
			}
		}
	}
	return
}
