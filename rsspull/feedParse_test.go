package rsspull

import (
	"log"
	"testing"
)

func TestRssJson(t *testing.T) {
	data := []byte(`
{
  "rss": {
    "@version": "2.0",
    "channel": {
      "title": "Example RSS Feed",
      "link": "http://www.example.com",
      "description": "This is an example RSS feed.",
      "items": [
        {
          "title": "Item 1",
          "link": "http://www.example.com/item1",
          "description": "This is the first item in the feed.",
          "pubDate": "Mon, 01 Dec 2023 12:15:00 GMT"
        },
        {
          "title": "Item 2",
          "link": "http://www.example.com/item2",
          "description": "This is the second item in the feed.",
          "pubDate": "Mon, 01 Dec 2023 13:30:00 GMT"
        }
      ]
    }
  }
}
`)

	feed, _ := parseFeed(data, "json")
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}

func TestAtomJson(t *testing.T) {
	data := []byte(`
{
  "feed": {
    "title": "Example Atom Feed",
    "subtitle": "This is an example Atom feed.",
    "link": [
      { "rel": "self", "href": "http://www.example.com/feed" },
      { "href": "http://www.example.com" }
    ],
    "updated": "2023-12-01T12:00:00Z",
    "id": "urn:uuid:12345",
    "author": { "name": "John Doe", "email": "johndoe@example.com" },
    "entry": [
      {
        "title": "Item 1",
        "link": [{ "href": "http://www.example.com/item1" }],
        "id": "urn:uuid:67890",
        "updated": "2023-12-01T12:15:00Z",
        "summary": "This is the first item in the feed."
      },
      {
        "title": "Item 2",
        "link": [{ "href": "http://www.example.com/item2" }],
        "id": "urn:uuid:ABCDE",
        "updated": "2023-12-01T13:30:00Z",
        "summary": "This is the second item in the feed."
      }
    ]
  }
}
`)
	feed, _ := parseFeed(data, "json")
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}
