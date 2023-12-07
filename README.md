# rss-bot [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> 一个可扩展的轻量级telegram rss订阅机器人

## 安装

```shell
git clone git@github.com:nopnoping/rss-bot.git 
cd rss-bot
go get
go build
```

## 使用示例

```shell
./rssbot -h

#Usage of ./rssbot:
#  -botproxy string
#        bot proxy url
#  -db string
#        sqlite db path (default "./rssbot.db")
#  -rssproxy string
#        rss client proxy url
#  -token string
#        telegram bot token

rssbot -toen=<bot-token> -botproxy=<bot-proxy-url> -rssproxy=<rss-pull-proxy-url>

#2023/12/07 08:12:53 bot start.......
#2023/12/07 08:12:53 push task start......
```

## 文件结构

- rss-bot
  - `bot`
  - `config`
  - `db`
  - `push-task`
  - `rsspull`
  - LICENSE
  - go.mod