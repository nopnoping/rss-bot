package config

import "time"

var RssClientProxyURL string
var BotProxyURL string
var Token string
var DbPath string

var RssClientTimeOut = 5 * time.Second
var PushTaskPeriod = 5 * time.Minute
var TaskPeriodScale = PushTaskPeriod / time.Second

var BotPushChCap = 20
var BotUpdateOffset = 0
var BotUpdateTimeout = 60
var RssTimeFormat = []string{
	"Mon, 02 Jan 2006 15:04:05 MST",
	"2006-01-02T15:04:05Z",
	"02 Jan 06 15:04 MST",
	"Mon, 02 Jan 2006 15:04:05",
	time.RFC3339,
}
