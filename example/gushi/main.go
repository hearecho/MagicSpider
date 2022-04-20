package main

import (
	"github.com/hearecho/MagicSpider"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace1.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
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
