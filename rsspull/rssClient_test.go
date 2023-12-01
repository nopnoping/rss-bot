package rsspull

import (
	"log"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	client := &rssClient{client: &http.Client{}}
	r, _ := client.get("http://www.baidu.com")
	log.Println((string)(r))
}
