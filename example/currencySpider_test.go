package example

import (
	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/example/parse"
	"testing"
)

func Test_CurrencySpiderArch(t *testing.T) {
	r := []MagicSpider.Request{
		{Url: "https://so.gushiwen.cn/gushi/tangshi.aspx",
			Parse: parse.NameParse,
			Depth: 1,},
	}
	e := MagicSpider.Engine{
		WorkerCount:   100,
		StartRequests: r,
		S:             MagicSpider.NewSchedule(),
	}
	e.Go()
}
