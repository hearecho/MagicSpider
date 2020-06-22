package main

import (
	"MagicSpider/spider/engine"
	"MagicSpider/spider/log"
	"MagicSpider/spider/parse"
	"MagicSpider/spider/scheduler"
	"MagicSpider/spider/setting"
)

func init()  {
	//初始化日志
	log.InitLogger()
	setting.InitSetting()
}
func main() {
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 20,
	}
	e.Run(engine.Request{
		Url:       "https://www.douban.com/group/shanghaizufang/discussion",
		ParseFunc: parse.ParesLink,
	})
}
