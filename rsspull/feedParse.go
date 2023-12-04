package rsspull

import (
	"github.com/beevik/etree"
	"log"
	"rssbot/rsspull/parse"
	"strings"
)

// The format of data could be json/xml
// if format is empty, it will judge data format by prefix
func parseFeed(data []byte, format string) *parse.FeedInfo {
	switch format {
	case "xml":
		return parseXML(data)
	case "json":
		return nil
	default:
		return nil
	}
}

func parseXML(data []byte) *parse.FeedInfo {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		log.Fatalf("Doc Read bytes err:%v", err)
		return nil
	}

	root := doc.Root()
	if root == nil {
		log.Fatal("Parse Feed err, Don't find root element!")
		return nil
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
				key = "feed-" + attr.Value
				break
			}
			if attr.Key == "xmlns" && attr.Value == "http://www.w3.org/2005/Atom" {
				key = "feed-1.0"
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
		log.Fatalf("Parse Feed err, No matched Root Tag!. tag:%s\n", root.Tag)
		return nil
	}

	var p parse.Parser
	var ok bool
	if p, ok = parse.ParserMap[key]; !ok {
		log.Fatalf("Parse Feed err, No matched parser!. key:%s\n", key)
		return nil
	}
	return p.Parse(root)
}
