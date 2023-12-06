package config

import "time"

var RssClientTimeOut = 5 * time.Second
var RssClientProxyURL = "http://127.0.0.1:7890"
var PushTaskPeriod = 5 * time.Minute
var TaskPeriodScale = PushTaskPeriod / time.Second
var DbPath = "./rssbot.db"

var BotProxyURL = "http://127.0.0.1:7890"
var Token = "6807409395:AAFW90O-9G-1rPzhDv-q1ZsGUAYvBx9v74s"
var BotPushChCap = 20
var BotUpdateOffset = 0
var BotUpdateTimeout = 60
