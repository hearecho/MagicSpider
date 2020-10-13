package example

import (
	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/example/parse/bilibili"
	"github.com/hearecho/MagicSpider/example/parse/gushi"
	"strconv"
	"testing"
)

func Test_CurrencySpiderArch(t *testing.T) {
	r := []MagicSpider.Request{}
	for i:=123345;i<465789;i++ {
		r = append(r,MagicSpider.Request{
			Url:    "https://api.bilibili.com/x/space/upstat?mid="+strconv.Itoa(i),
			Parse:  bilibili.Parser,
			Common: MagicSpider.Common{
				Depth: 1,
				Meta:  &gushi.Item{},
			},
			Headers: map[string]string{
					"Host":"api.bilibili.com",
			},
		})
	}
	e := MagicSpider.Engine{
		WorkerCount:   10,
		StartRequests: r,
		S:             MagicSpider.NewSchedule(),
	}
	e.Go()
}
