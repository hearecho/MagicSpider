package utils

import (
	"sync"
	"time"
)

/**
请求限速
*/
type LimitRate struct {
	rate       int
	interval   time.Duration
	lastAction time.Time
	lock       sync.Mutex
}

func (l *LimitRate)Limit() bool {
	result := false
	for  {
		l.lock.Lock()
		//判断最后一次执行时间与当前时间是否大于限制速率
		if time.Now().Sub(l.lastAction) > l.interval {
			l.lastAction = time.Now()
			result = true
		}
		l.lock.Unlock()
		if result {
			return result
		}
		time.Sleep(l.interval)
	}
}

func (l *LimitRate)SetRate(r int)  {
	l.rate = r
	l.interval = time.Microsecond * time.Duration(1000*1000/l.rate)
}

func (l *LimitRate)GetRate() int  {
	return l.rate
}


