package engine

import (
	"fmt"
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/fetch"
	"github.com/hearecho/MagicSpider/crawler/simpleSpiderArch/types"
)

func Run(starts ...types.Request)  {
	var requests []types.Request
	for _,r := range starts {
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		resp,err := fetch.Fetch(r)
		if err != nil {
			continue
		}
		//解析
		result := r.Parse(resp)
		//添加新的request到其中
		requests = append(requests,result.Requests...)
		//处理得到的结果
		for _,i := range result.Items {
			fmt.Printf("crawled item %v\n",i)
		}
	}
}
