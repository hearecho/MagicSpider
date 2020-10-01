package CurrencySpiderArch

import (
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/engine"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/parse"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/types"
	"testing"
)

func Test_CurrencySpiderArch(t *testing.T) {
	r := []types.Request{
		{Url: "https://so.gushiwen.cn/gushi/tangshi.aspx",
			Parse: parse.NameParse,
			Depth: 1,},
	}
	e := engine.Engine{
		WorkerCount:   2,
		StartRequests: r,
	}
	e.Go()
}
