package schedule

import "github.com/hearecho/MagicSpider/spider/engine"

//提交任务
type Schedule struct {
	requests chan *engine.Request
}

func (s *Schedule)Submit(r *engine.Request)  {
	s.requests <- r
}

