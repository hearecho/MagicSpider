package MagicSpider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"sync"
	"time"

	"github.com/hearecho/MagicSpider/utils"
)

type Engine struct {
	//协程个数
	WorkerCount int
	//起始请求
	StartRequests []Request
	//Schedule调度器
	S *Schedule
}

func (e *Engine) Go() {
	start := time.Now().UnixNano() / 1e6
	//读取配置
	InitSetting()
	//设置waitgroup
	wg := &sync.WaitGroup{}
	wg.Add(e.WorkerCount + 2)
	lr := &utils.LimitRate{}
	lr.SetRate(S.Rate)
	//创建worker
	for i := 0; i < e.WorkerCount; i++ {
		go worker(e, wg, lr)
	}
	//将起始请求放入channel中
	for _, r := range e.StartRequests {
		e.S.SubmitTask(r)
	}
	//处理Res
	go e.S.Communicate(wg)
	go e.S.Process(wg)
	wg.Wait()
	utils.Info(fmt.Sprintf("crawl end. use time:%dms",time.Now().UnixNano()/1e6-start))
}


//worker的运行逻辑，负责处理传入的requests，并得到item传回engine
func worker(e *Engine, wg *sync.WaitGroup, lr *utils.LimitRate) {
	for {
		if lr.Limit(){
			timeout := time.After(2 * time.Second)
			select {
			case httpRequest := <-e.S.HttpRequests():
				httpResp, err := Fetch(httpRequest)
				if err != nil {
					utils.Error(fmt.Sprintf("%v",err))
					break
				}
				//根据Doctype设置Doc
				if S.DocType == "html" {
					httpResp.Doc, err = htmlquery.Parse(bytes.NewReader(httpResp.Body))
				} else {
					err = json.Unmarshal(httpResp.Body, &httpResp.Doc)
				}
				if err != nil {
					utils.Error(fmt.Sprintf("%v",err))
					break
				}
				start := time.Now().UnixNano()/1e6
				res := httpRequest.Parse(httpResp)
				usedTime := time.Now().UnixNano()/1e6 - start
				utils.Info(fmt.Sprintf("parse web:%s content used time:%dms",httpRequest.Url,usedTime))
				//将res添加到通道中
				e.S.SubmitRes(res)
			case <-timeout:
				wg.Done()
				return
			}
		}
	}
}
