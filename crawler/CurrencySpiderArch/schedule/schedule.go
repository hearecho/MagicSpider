package schedule

import "github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/types"

//提交任务
type Schedule struct {
	httpRequests chan types.Request
	items chan interface{}
}

func NewSchedule() *Schedule  {
	return &Schedule{
		httpRequests: make(chan types.Request),
		items:        make(chan interface{}),
	}
}
func (s *Schedule)SubmitTask(r types.Request)  {
	s.httpRequests <- r
}

func (s *Schedule)SubmitItem(item interface{})  {
	s.items <- item
}

func (s *Schedule)HttpRequests()  chan types.Request {
	return s.httpRequests
}

func (s *Schedule)Items()  chan interface{} {
	return s.items
}

