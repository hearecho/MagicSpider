package main

import (
	"MagicSpider/spider/db"
	"MagicSpider/spider/engine"
	"MagicSpider/spider/log"
	"MagicSpider/spider/parse"
	"MagicSpider/spider/scheduler"
	"MagicSpider/spider/setting"
	"fmt"
)

func init()  {
	//初始化日志
	log.InitLogger()
	setting.InitSetting()
	fmt.Println("=============Init Spider===============")
	fmt.Printf("\n"+
		"\t\t\t      | \n"+
		"\t\t\t    _| \n"+
		"\t\t\t///\\(o_o)/\\\\ \n"+
		"\t\t\t|||  ` '  ||| \n")
	db.InitDB()
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
	//fmt.Println(db.QuerySetting("count"))
}
