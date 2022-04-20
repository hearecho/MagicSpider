package MagicSpider

import (
	"github.com/bits-and-blooms/bloom/v3"
	"sync"
)

type schedule interface {
	Put(request Request, requests chan Request)
}

// 队列实现
type ScheduleQueue struct {
	Queue []Request
	Filter *bloom.BloomFilter
	sync.Mutex
}

func NewScheduleQueue() *ScheduleQueue {
	return &ScheduleQueue{
		Queue:  []Request{},
		Filter: bloom.NewWithEstimates(10000,0.01),
	}
}

func (s *ScheduleQueue) Put(request Request, requests chan Request)  {
	s.Lock()
	temp := []byte(request.Url)
	if !s.Filter.Test(temp) {
		s.Filter.Add(temp)
		s.Queue = append(s.Queue, request)
	}
	s.Unlock()
	next := s.nextRequest()
	if next != nil {
		requests <- *next	
	}
}

func (s *ScheduleQueue) nextRequest() *Request{
	s.Lock()
	defer s.Unlock()
	if len(s.Queue) <= 0 {
		return nil
	}
	next := s.Queue[0]
	s.Queue = s.Queue[1:]
	return &next
}

