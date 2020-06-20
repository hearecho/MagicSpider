package scheduler

import "MagicSpider/spider/engine"

//均是没有缓冲区的
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

/**
提交任务到请求队列中
 */
func (s *QueuedScheduler) Submit(req engine.Request)  {
	s.requestChan <- req
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request)  {
	s.workerChan <- w
}
//将request的channel送给调度器
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan chan engine.Request) {
	s.workerChan = c
}

func (s *QueuedScheduler) Run()  {
	//没有缓冲区的
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		//生产请求的队列
		var requestQ   []engine.Request
		//爬取数据的工作队列
		var workerQ []chan engine.Request


		for  {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) >0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ,r)
			case w := <-s.workerChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}

