package bot

import (
	"fmt"
	"net/http"
	"net/url"
	"rssbot/db"
	"runtime"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	proxyURL, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	request, _ := http.NewRequest("GET", "https://www.google.com/", nil)
	rep, _ := client.Do(request)
	fmt.Println(rep.Status)

}

func TestBotConnect(t *testing.T) {
	runtime.GOMAXPROCS(1)
	bot := NewBot()

	go bot.Start()

	select {
	case <-time.After(time.Minute):
		bot.Stop()
	}

	time.Sleep(2 * time.Second)
}

func TestDb(t *testing.T) {
	num := db.HasThisUrlWithTheChatId(5282246628, "https://droidyue.com/atom.xml")
	fmt.Println(num)
}
