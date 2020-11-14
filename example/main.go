package main

import (
	"example/parse/gushi"
	"github.com/hearecho/MagicSpider"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	r := []MagicSpider.Request{{
			Url:   "https://so.gushiwen.cn/gushi/tangshi.aspx",
			Parse: gushi.NameParse,
			Common: MagicSpider.Common{
				Depth: 1,
				Meta:  &gushi.Item{},
			},
		}}
	e := MagicSpider.Engine{
		WorkerCount:   10,
		StartRequests: r,
		S:             MagicSpider.NewSchedule(),
	}
	e.Go()
}
