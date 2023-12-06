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
