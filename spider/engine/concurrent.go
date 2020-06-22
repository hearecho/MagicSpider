package engine

import (
	"MagicSpider/spider/fetch"
	"MagicSpider/spider/setting"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan chan Request)
	WorkerReady(chan Request)
	Run()
}


func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParseRes)
	//创建一个死循环来处理 传递进去的数据
	e.Scheduler.Run()
	//创建WorkerCount个工作协程，每个协程都是从 out通道中取，如果不存在则发生阻塞
	for i :=0;i<e.WorkerCount ;i++  {
		createWorker(out,e.Scheduler)
	}
	//将request 添加到
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}
	//处理输出的结果
	for  setting.Count < setting.TotalCount {
		//从结果中取出存储的数据
		result := <- out
		//获取了信息，以后可以新增存储功能
		//for _, item := range result.Items {
		//	fmt.Printf("get item:%v\n",item)
		//}
		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

/**
将任务添加到工作进程中
 */
func createWorker(out chan ParseRes, s Scheduler) {
	go func() {
		in := make(chan Request)
		for setting.Count < setting.TotalCount{
			s.WorkerReady(in)
			// tell scheduler i'm ready
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/**
工作进程
*/
func  worker(r Request) (ParseRes, error){
	body, err := fetch.Fetch(r.Url,"GET","")
	if err != nil || r.ParseFunc == nil {
		//log.Printf("Fetcher: error " + "fetching url %s: %v", r.Url, err)
		return ParseRes{}, err
	}

	return r.ParseFunc(body,nil), nil
}
