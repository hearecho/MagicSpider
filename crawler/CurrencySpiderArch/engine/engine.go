package engine

import (
	"fmt"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/fetch"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/types"
)

type Engine struct {
	//协程个数
	WorkerCount int
	//起始请求
	StartRequests []types.Request
	//设置爬取深度
}

func (e *Engine)Go()  {
	httpRequests := make(chan types.Request)
	items := make(chan interface{})
	//创建worker
	for i:=0;i<e.WorkerCount;i++ {
		createWorker(httpRequests,items)
	}
	//将起始请求放入channel中
	for _,r := range e.StartRequests {
		httpRequests <- r
	}
	for item := range items {
		fmt.Println("get item: ",item)
	}
}

//创建worker
func createWorker(httpRequests chan types.Request,items chan interface{})  {
	go worker(httpRequests,items)
}

//worker的运行逻辑，负责处理传入的requests，并得到item传回engine
func worker(httpRequests chan types.Request,items chan interface{})  {
	for {
		httpRequest := <- httpRequests
		httpResp,err := fetch.Fetch(httpRequest)
		if err != nil{
			panic(err)
		}
		res := httpRequest.Parse(httpResp)
		for _,r := range res.Requests {
			//后续可以更改为schedule进行任务提交
			httpRequests <- r
		}
		//提交items
		for _,item := range res.Items {
			items <- item
		}
	}
}