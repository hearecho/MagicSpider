package MagicSpider

import (
	"fmt"
	"sync"
	"time"

	"github.com/hearecho/MagicSpider/utils"
)

type engine struct {
	//协程个数
	workerCount int
	//起始请求
	startRequests []Request
	//Schedule_bak调度器
	//s *Schedule_bak
	s schedule
	requests chan Request
	parseResults chan ParseResult
	items chan Item
}

func NewEngine(wokerCount int, seeds []Request) *engine {
	return &engine{
		workerCount:   wokerCount,
		startRequests: seeds,
		s:             NewScheduleQueue(),
	}
}
func (e *engine) Go() {
	start := time.Now().UnixNano() / 1e6
	//读取配置
	InitSetting()
	//设置waitgroup
	wg := &sync.WaitGroup{}
	wg.Add(e.workerCount + 2)
	lr := &utils.LimitRate{}
	lr.SetRate(S.Rate)
	// 初始化chan
	e.requests = make(chan Request, 100)
	e.parseResults = make(chan ParseResult, 100)
	e.items = make(chan Item, 100)


	for _, r := range e.startRequests {
		e.s.Put(r, e.requests)
	}
	//创建worker
	for i := 0; i < e.workerCount; i++ {
		go downloader(e, wg, lr)
	}
	//处理Res
	go e.Communicate(wg)
	go e.itemPipeline(wg)
	wg.Wait()
	utils.Info(fmt.Sprintf("crawl end. use time:%dms", time.Now().UnixNano()/1e6-start))
}

// Communicate 主要是负责管控数据流
func (e *engine)Communicate(wg *sync.WaitGroup) {
	for {
		timeout := time.After(2 * time.Second)
		select {
		case res := <-e.parseResults:
			//处理requests
			for _, r := range res.Requests {
				go e.s.Put(r, e.requests)
			}
			for _, item := range res.Items {
				i := item
				go func() {
					e.items <- i
				}()
			}
		case <-timeout:
			wg.Done()
			return
		}
	}
}

// itemPipeline item pipeline组件
func (e *engine) itemPipeline(wg *sync.WaitGroup) {
	for {
		timeout := time.After(2 * time.Second)
		select {
		case item := <-e.items:
			item.Process()
		case <-timeout:
			wg.Done()
			return
		}
	}
}
