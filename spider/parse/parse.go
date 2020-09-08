package parse

import (
	"github.com/hearecho/MagicSpider/spider/engine"
	"github.com/hearecho/MagicSpider/spider/log"
	"github.com/hearecho/MagicSpider/spider/model"
	"github.com/hearecho/MagicSpider/spider/setting"
	"regexp"
	"sync/atomic"
)

/**
linksReg  和itemReg
eg: itemReg: `<a href="(.*?)" title="(.*?)".*<td nowrap="nowrap" class="time">(.*?)</td>`
	linkRe": `<a href="(.*?)">\d</a>`
 */
const lRe = `<a href="(.*?)" >\d</a>`
const itemRe  =  `<a href="(.*?)" title="(.*?)"[\s\S]*?<td nowrap="nowrap" class="time">(.*?)</td>`

func ParesLink(body [] byte,reg map[string]string) engine.ParseRes{
	linkRe := regexp.MustCompile(lRe)
	zfRe := regexp.MustCompile(itemRe)
	result := engine.ParseRes{}

	//解析租房具体信息
	zfMatches := zfRe.FindAllSubmatch(body,-1)
	for _,m := range zfMatches  {
		url := string(m[1])
		title := string(m[2])
		time := string(m[3])
		item := model.Item{
			Url:   url,
			Title: title,
			Time:  time,
		}
		//存储
		result.Items = append(result.Items,item)
		log.Info.Printf("parse resp result:%v\t [num]:%d\n",item,setting.Count)
		atomic.AddInt64(&setting.S.Count,1)
		//新的Requests,
		result.Requests = append(result.Requests,engine.Request{
			Url:       url,
			ParseFunc: nil,
		})
	}

	//解析下一页
	nextMatches := linkRe.FindAllSubmatch(body,-1)
	for _,m := range nextMatches {
		result.Requests = append(result.Requests,engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParesLink,
		})
	}
	return result
}

