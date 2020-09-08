package fetch

import (
	"github.com/hearecho/MagicSpider/spider/log"
	"github.com/hearecho/MagicSpider/spider/setting"
	"io/ioutil"
	"net/http"
	"strings"
)


func Fetch(url,method string,form string) ([]byte,error) {
	client := &http.Client{}
	method = strings.ToUpper(method)
	req,err := http.NewRequest(method,url,strings.NewReader(form))
	if err != nil {
		log.Error.Printf("construct request fail! err:%s\n",err.Error())
		return nil,nil
	}
	if strings.EqualFold(method,"POST") {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for k,v := range setting.S.Headers {
		req.Header.Set(k,v)
	}
	log.Info.Printf("crawling url:%s\n",req.URL)
	resp, err := client.Do(req)
	if resp == nil {
		log.Error.Printf("crawel error! url:%s\n",url)
	}
	if err != nil || resp.StatusCode != http.StatusOK{
		log.Error.Printf("crawel error! url:%s\t status_code:%d\n",url,resp.StatusCode)
		return nil,nil
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	log.Info.Printf("crawel url:%s finish! download byte:%d byte\n",url, len(body))
	return body,err
}


