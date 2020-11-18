package main

import (
	"fmt"
	"github.com/hearecho/MagicSpider"
)

func main() {
	requests := []MagicSpider.Request{}
	for i:=4564;i<4700;i++ {
		requests = append(requests,MagicSpider.Request{
			Url:     fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%d",i),
			Parse:   ResParse,
			Headers: nil,
			Common:  MagicSpider.Common{
				Depth: 1,
				Meta:  &Item{},
			},
		})
	}
	e := MagicSpider.NewEngine(10,requests)
	e.Go()
}
