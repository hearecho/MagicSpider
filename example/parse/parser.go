package parse

import (
	"fmt"
	"github.com/hearecho/MagicSpider"
	"github.com/hearecho/MagicSpider/utils"
	"os"
	"regexp"
	"strings"
)

type Item struct {
	Title   string
	Author  string
	Content string
}

func (i *Item)Process()  {
	utils.IsNotExistMkDir("runtime/")
	f, _ := utils.Open("runtime/result.txt",os.O_CREATE|os.O_APPEND,0777)
	item := fmt.Sprintf("【TiTle】:%v\t 【Author】:%v\t 【Content】:%v\n",i.Title,i.Author,strings.Trim(i.Content,"\n"))
	f.WriteString(item)
	f.Close()
}
func NameParse(r *MagicSpider.Response) *MagicSpider.ParseResult {
	//使用re进行
	nameRe := `<a href="(.*?)" target="_blank">([^<]+)</a>\((.*?)\)`
	re, _ := regexp.Compile(nameRe)
	result := re.FindAllSubmatch(r.Body, -1)
	res := &MagicSpider.ParseResult{}
	for _, item := range result {
		//新增url
		request := MagicSpider.Request{Url: "https://so.gushiwen.cn" + string(item[1]),
			Parse: ContentParse,
			Depth: r.Depth + 1}
		res.Requests = append(res.Requests, request)
		//处理item
		//title := Item{Title: string(item[2]), Author: string(item[3])}
		//res.Items = append(res.Items, title)
	}
	return res
}

func ContentParse(r *MagicSpider.Response) *MagicSpider.ParseResult {
	contentRe := `<div class="contson"[^>]+>([\s\S]*?)</div>`
	re, _ := regexp.Compile(contentRe)
	result := re.FindSubmatch(r.Body)
	res := &MagicSpider.ParseResult{}
	content := Item{Content:string(result[1])}
	res.Items = append(res.Items, &content)
	return res

}
