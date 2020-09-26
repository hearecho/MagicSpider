package parse

import (
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/types"
	"regexp"
)

type Item struct {
	Title   string
	Author  string
	Content string
}

func NameParse(r *types.Response) *types.ParseResult {
	//使用re进行
	nameRe := `<a href="(.*?)" target="_blank">([^<]+)</a>\((.*?)\)`
	re, _ := regexp.Compile(nameRe)
	result := re.FindAllSubmatch(r.Body, -1)
	res := &types.ParseResult{}
	for _, item := range result {
		//新增url
		request := types.Request{Url: "https://so.gushiwen.cn" + string(item[1]),
			Parse: ContentParse,
			Depth: r.Depth + 1}
		res.Requests = append(res.Requests, request)
		//处理item
		title := Item{Title: string(item[2]), Author: string(item[3])}
		res.Items = append(res.Items, title)
	}
	return res
}

func ContentParse(r *types.Response) *types.ParseResult {
	contentRe := `<div class="contson"[^>]+>([\s\S]*?)</div>`
	re, _ := regexp.Compile(contentRe)
	result := re.FindAllSubmatch(r.Body, -1)
	res := &types.ParseResult{}
	for _, item := range result {
		//处理item
		content := Item{Content:string(item[1])}
		res.Items = append(res.Items, content)
	}
	return res

}
