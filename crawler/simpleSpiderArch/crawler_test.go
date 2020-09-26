package simpleSpiderArch

import (
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/engine"
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/parse"
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/types"
	"testing"
)

func TestFetch(t *testing.T)  {
	r := types.Request{
		Url:   "https://so.gushiwen.cn/gushi/tangshi.aspx",
		Parse: parse.NameParse,
		Depth: 1,
	}
	engine.Run(r)
}
