package main

import (
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
			Parse: NameParse,
			Common: MagicSpider.Common{
				Depth: 1,
				Meta:  &Item{},
			},
		}}
	e := MagicSpider.NewEngine(10,r)
	e.Go()
}
