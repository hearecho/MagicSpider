package MagicSpider

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/hearecho/MagicSpider/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)



//worker的运行逻辑，负责处理传入的requests，并得到item传回engine
func downloader(e *engine, wg *sync.WaitGroup, lr *utils.LimitRate) {
	for {
		if lr.Limit(){
			timeout := time.After(2 * time.Second)
			select {
			case httpRequest := <-e.requests:
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
				// 这里是解析
				// 解析部分也可以开启协程
				res := httpRequest.Parse(httpResp)
				usedTime := time.Now().UnixNano()/1e6 - start
				utils.Info(fmt.Sprintf("parse web:%s content used time:%dms",httpRequest.Url,usedTime))
				//将res添加到通道中
				go func() {
					e.parseResults <- res
				}()
			case <-timeout:
				wg.Done()
				return
			}
		}
	}
}

func Fetch(r Request) (*Response, error) {
	//固定深度直接终止
	if r.Depth > S.MaxDepth {
		return &Response{}, nil
	}
	start := time.Now().UnixNano()/ 1e6
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.Url, strings.NewReader(""))
	if err != nil {
		utils.Error(fmt.Sprintf("%v",err))
		return &Response{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)

	if err != nil {
		utils.Error(fmt.Sprintf("%v",err))
		return &Response{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return &Response{}, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	//防止出现乱码情况，如果解析不出来则默认utf8
	bodyReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, encode.NewDecoder())
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		utils.Error(fmt.Sprintf("%v",err))
		return &Response{}, err
	}
	usedTime := time.Now().UnixNano()/ 1e6-start
	utils.Info(fmt.Sprintf("crawl url:%s\tuse time:%dms",r.Url,usedTime))
	return &Response{Body: body, Common: Common{
		Depth: r.Depth,
		Meta:  r.Meta,
	}}, nil
}

//确定读入
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil && S.DocType != "json"{
		utils.Error(fmt.Sprintf("%v",err))
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}
