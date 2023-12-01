package rsspull

import (
	"io"
	"log"
	"net/http"
)

type rssClient struct {
	// some http request config
	client *http.Client
}

func (c *rssClient) get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Build %s request occur error:%v\n", url, err)
		return nil, err
	}

	// Add Header
	request.Header.Add("", "")

	// Send request
	response, err := c.client.Do(request)
	if err != nil {
		log.Printf("Send %s request occur error:%v\n", url, err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Read %s Body occur error:%v\n", url, err)
		return nil, err
	}
	return body, nil
}
