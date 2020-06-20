package main

import (
	"MagicSpider/spider/fetch"
	"MagicSpider/spider/parse"
	"fmt"
)

func main() {
	headers := map[string]string{
		"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
	}
	resp,_ := fetch.Fetch(
		"https://www.douban.com/group/shanghaizufang/discussion",
		"get",
		headers,
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


