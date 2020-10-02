package engine

import (
	"fmt"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/fetch"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/schedule"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/types"
)

type Engine struct {
	//协程个数
	WorkerCount int
	//起始请求
	StartRequests []types.Request
	//Schedule调度器
	S *schedule.Schedule
}

func (e *Engine)Go()  {
	//创建worker
	for i:=0;i<e.WorkerCount;i++ {
		createWorker(e.S)
	}
	//将起始请求放入channel中
	for _,r := range e.StartRequests {
		e.S.SubmitTask(r)
	}
	for item := range e.S.Items() {
		fmt.Println("get item: ",item)
	}
}

//创建worker
func createWorker(s *schedule.Schedule)  {
	go worker(s)
}

//worker的运行逻辑，负责处理传入的requests，并得到item传回engine
func worker(s *schedule.Schedule)  {
	for {
		httpRequest := <- s.HttpRequests()
		httpResp,err := fetch.Fetch(httpRequest)
		if err != nil{
			panic(err)
		}
		res := httpRequest.Parse(httpResp)
		for _,r := range res.Requests {
			//后续可以更改为schedule进行任务提交
			s.SubmitTask(r)
		}
		//提交items
		for _,item := range res.Items {
			s.SubmitItem(item)
		}
	}
}