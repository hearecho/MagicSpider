package main

import (
	"fmt"
	"github.com/hearecho/MagicSpider/spider/db"
	"github.com/hearecho/MagicSpider/spider/engine"
	"github.com/hearecho/MagicSpider/spider/log"
	"github.com/hearecho/MagicSpider/spider/parse"
	"github.com/hearecho/MagicSpider/spider/scheduler"
	"github.com/hearecho/MagicSpider/spider/setting"
)

func init() {
	//初始化日志
	log.InitLogger()
	fmt.Println("=============Init Spider===============")
	fmt.Printf("\n" +
		"\t\t\t      | \n" +
		"\t\t\t    _| \n" +
		"\t\t\t///\\(o_o)/\\\\ \n" +
		"\t\t\t|||  ` '  ||| \n")
	setting.InitSetting()
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
}
