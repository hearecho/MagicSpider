package singleSpider

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Fetch(url string) []byte {
	client := &http.Client{}
	req,err := http.NewRequest("GET",url,strings.NewReader(""))
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == http.StatusOK {
		//防止出现乱码情况，如果解析不出来则默认utf8
		bodyReader := bufio.NewReader(resp.Body)
		encode := determineEncoding(bodyReader)
		reader := transform.NewReader(bodyReader,encode.NewDecoder())
		body,err := ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		return body
	}
	return nil
}
//确定读入
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes,err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

const NameRe  = `<a href="(.*?)" target="_blank">([^<]+)</a>`
func ParseTitle(bytes []byte)  {
	//使用re进行
	re,_ := regexp.Compile(NameRe)
	result := re.FindAllSubmatch(bytes,-1)
	for _,item := range result {
		fmt.Println(string(item[1]),string(item[2]))
	}
}


func Run()  {
	ParseTitle(Fetch("https://so.gushiwen.cn/gushi/tangshi.aspx"))
}
