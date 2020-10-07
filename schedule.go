package MagicSpider

import (
	"sync"
	"time"
)

//提交任务
type Schedule struct {
	httpRequests chan Request
	parseResult chan ParseResult
	items chan Item
}

func NewSchedule() *Schedule  {
	return &Schedule{
		httpRequests: make(chan Request),
		parseResult:        make(chan ParseResult),
		items:make(chan Item),
	}
}
func (s *Schedule)SubmitTask(r Request)  {
	go func() {s.httpRequests <- r}()
}

func (s *Schedule)SubmitRes(res ParseResult)  {
	go func() {
		s.parseResult <- res
	}()
}
func	(s *Schedule)SubmitItems(item Item)  {
	go func() {
		s.items <- item
	}()
}

func (s *Schedule)HttpRequests()  chan Request {
	return s.httpRequests
}
func (s *Schedule)Result()  chan ParseResult {
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
			item.Process()
		case <- timeout:
			wg.Done()
			return
		}
	}
}

