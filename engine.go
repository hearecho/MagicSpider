package MagicSpider

import (
	"fmt"
	"sync"
	"time"
)

type Engine struct {
	//协程个数
	WorkerCount int
	//起始请求
	StartRequests []Request
	//Schedule调度器
	S *Schedule
}

func (e *Engine)Go()  {
	//设置waitgroup
	wg := &sync.WaitGroup{}
	wg.Add(e.WorkerCount+2)
	//创建worker
	for i:=0;i<e.WorkerCount;i++ {
		createWorker(e,wg)
	}
	//将起始请求放入channel中
	for _,r := range e.StartRequests {
		e.S.SubmitTask(r)
	}
	//处理Res
	go e.S.Communicate(wg)
	go e.S.Process(wg)
	wg.Wait()
	fmt.Println("爬取结束")
}

//创建worker
func createWorker(e *Engine,wg *sync.WaitGroup)  {
	go worker(e,wg)
}

//worker的运行逻辑，负责处理传入的requests，并得到item传回engine
func worker(e *Engine,wg *sync.WaitGroup)  {
	for {
		timeout := time.After(2*time.Second)
		select {
		case httpRequest := <- e.S.HttpRequests():
			httpResp,err := Fetch(httpRequest)
			if err != nil{
				panic(err)
			}
			res := httpRequest.Parse(httpResp)
			//将res添加到通道中
			e.S.SubmitRes(*res)
		case <-timeout:
			wg.Done()
			return
		}
	}
}

