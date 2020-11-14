package MagicSpider

import (
	"sync"
	"time"
)

//提交任务
type Schedule struct {
	httpRequests chan Request
	parseResult  chan ParseResult
	items        chan Item
}

func NewSchedule() *Schedule {
	return &Schedule{
		httpRequests: make(chan Request),
		parseResult:  make(chan ParseResult),
		items:        make(chan Item),
	}
}
func (s *Schedule) SubmitTask(r Request) {
	s.httpRequests <- r
}

func (s *Schedule) SubmitRes(res ParseResult) {
	s.parseResult <- res
}
func (s *Schedule) SubmitItems(item Item) {
	s.items <- item
}
func (s *Schedule) HttpRequests() chan Request {
	return s.httpRequests
}

//通信
func (s *Schedule) Communicate(wg *sync.WaitGroup) {
	for {
		timeout := time.After(2 * time.Second)
		select {
		case res := <-s.parseResult:
			//处理requests
			for _, r := range res.Requests {
				go s.SubmitTask(r)
			}
			for _, item := range res.Items {
				go s.SubmitItems(item)
			}
		case <-timeout:
			wg.Done()
			return
		}
	}
}

//处理items
func (s *Schedule) Process(wg *sync.WaitGroup) {
	for {
		timeout := time.After(2 * time.Second)
		select {
		case item := <-s.items:
			item.Process()
		case <-timeout:
			wg.Done()
			return
		}
	}
}
