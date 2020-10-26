package example

import (
	"testing"

	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/example/parse/gushi"
)

func Test_CurrencySpiderArch(t *testing.T) {
	r := []MagicSpider.Request{
		MagicSpider.Request{
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
