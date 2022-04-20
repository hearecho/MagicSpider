package MagicSpider

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/hearecho/MagicSpider/utils"
	"sync"
)

type schedule interface {
	Put(request Request, requests chan Request, lr *utils.LimitRate)
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

func (s *ScheduleQueue) Put(request Request, requests chan Request, lr *utils.LimitRate)  {
	s.Lock()
	temp := []byte(request.Url)
	if !s.Filter.Test(temp) {
		s.Filter.Add(temp)
		s.Queue = append(s.Queue, request)
	}
	s.Unlock()
	next := s.nextRequest(lr)
	if next != nil {
		requests <- *next	
	}
}

func (s *ScheduleQueue) nextRequest(lr *utils.LimitRate) *Request{
	s.Lock()
	defer s.Unlock()
	// 限速器在scheduler中使用
	if lr.Limit() {
		if len(s.Queue) <= 0 {
			return nil
		}
		next := s.Queue[0]
		s.Queue = s.Queue[1:]
		return &next	
	} else {
		return nil
	}
}

