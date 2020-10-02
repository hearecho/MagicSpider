package schedule

import (
	"fmt"
	"github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/types"
	"sync"
	"time"
)

//提交任务
type Schedule struct {
	httpRequests chan types.Request
	parseResult chan types.ParseResult
	items chan interface{}
}

func NewSchedule() *Schedule  {
	return &Schedule{
		httpRequests: make(chan types.Request),
		parseResult:        make(chan types.ParseResult),
		items:make(chan interface{}),
	}
}
func (s *Schedule)SubmitTask(r types.Request)  {
	go func() {s.httpRequests <- r}()
}

func (s *Schedule)SubmitRes(res types.ParseResult)  {
	go func() {
		s.parseResult <- res
	}()
}
func	(s *Schedule)SubmitItems(item interface{})  {
	go func() {
		s.items <- item
	}()
}

func (s *Schedule)HttpRequests()  chan types.Request {
	return s.httpRequests
}
func (s *Schedule)Result()  chan types.ParseResult {
	return s.parseResult
}

//通信
func (s *Schedule)Communicate(wg *sync.WaitGroup)  {
	for {
		timeout := time.After(2*time.Second)
		select {
		case res := <- s.parseResult:
			//处理requests
			for _,r := range res.Requests {
				s.SubmitTask(r)
			}
			for _,item := range res.Items {
				s.SubmitItems(item)
			}
		case <- timeout:
			wg.Done()
			return
		}
	}
}

//处理items
func (s *Schedule)Process(wg *sync.WaitGroup)  {
	i := 0
	for {
		timeout := time.After(2*time.Second)
		select {
		case item := <- s.items:
			i++
			//process
			fmt.Println(i,item)
		case <- timeout:
			wg.Done()
			return
		}
	}
}

