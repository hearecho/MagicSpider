package bilibili

import (
	"fmt"

	"github.com/hearecho/MagicSpider"
)

type Item struct {
}

func (i *Item) Process() {

}

func Parser(resp *MagicSpider.Response) *MagicSpider.ParseResult {
	res := &MagicSpider.ParseResult{}
	m := resp.Doc.(map[string]interface{})
	fmt.Println(m)
	return res
}
