package main

import (
	"MagicSpider/spider/engine"
	"MagicSpider/spider/parse"
	"MagicSpider/spider/scheduler"
)

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
