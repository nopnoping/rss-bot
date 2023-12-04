package rsspull

import (
	"log"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	client := &rssClient{client: &http.Client{}}
	r, _, _ := client.get("https://sspai.com/feed")
	log.Println((string)(r))
}
