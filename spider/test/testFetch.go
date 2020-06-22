package main

import (
	"MagicSpider/spider/fetch"
	"MagicSpider/spider/parse"
	"fmt"
)

func main() {
	resp,_ := fetch.Fetch(
		"https://www.douban.com/group/shanghaizufang/discussion",
		"get",
		"")
	re := map[string]string{
		"itemRe":`<a href="(.*?)" title="(.*?)"[\s\S]*?<td nowrap="nowrap" class="time">(.*?)</td>`,
		"linkRe":`<a href="(.*?)" >\d</a>`,
	}
	item := parse.ParesLink(resp,re)
	for _,v := range item.Items {
		fmt.Println(v)
	}
	//fmt.Print(string(resp))
}


