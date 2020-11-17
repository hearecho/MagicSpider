package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"os"
	"regexp"
	"strings"

	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/utils"
)

type Item struct {
	Title   string
	Author  string
	Content string
}

func (i *Item) Process() {
	_ = utils.IsNotExistMkDir(MagicSpider.S.RuntimePath)
	f, _ := utils.Open(MagicSpider.S.RuntimePath+"result.csv", os.O_CREATE|os.O_APPEND, 0777)
	item := fmt.Sprintf("%v,%v,%v\n", i.Title, i.Author, strings.Trim(i.Content, "\n"))
	f.WriteString(item)
	f.Close()
}
func NameParse(r *MagicSpider.Response) MagicSpider.ParseResult {
	//使用re进行
	nameRe := `<a href="(.*?)" target="_blank">([^<]+)</a>\((.*?)\)`
	re, _ := regexp.Compile(nameRe)
	result := re.FindAllSubmatch(r.Body, -1)
	res := &MagicSpider.ParseResult{}
	for _, item := range result {
		//新增url
		request := MagicSpider.Request{Url: "https://so.gushiwen.cn" + string(item[1]),
			Parse: ContentParse,
			Common: MagicSpider.Common{
				Depth: r.Depth + 1,
				Meta:  &Item{Title: string(item[2]), Author: string(item[3])},
			}}
		res.Requests = append(res.Requests, request)
	}
	return *res
}

func ContentParse(r *MagicSpider.Response) MagicSpider.ParseResult {
	res := &MagicSpider.ParseResult{}
	//contentRe := `<div class="contson"[^>]+>([\s\S]*?)</div>`
	//re, _ := regexp.Compile(contentRe)
	//result := re.FindSubmatch(r.Body)
	node := r.Doc.(*html.Node)
	content := htmlquery.Find(node,"//div[@class='contson']")[0]

	r.Meta.(*Item).Content = htmlquery.InnerText(content)
	res.Items = append(res.Items, r.Meta.(*Item))
	return *res

}
