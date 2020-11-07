package MagicSpider

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(r Request) (*Response, error) {
	//固定深度直接终止
	if r.Depth > S.MaxDepth {
		return &Response{}, nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", r.Url, strings.NewReader(""))
	if err != nil {
		return &Response{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	log.Printf("start crawle url: %s", r.Url)
	if err != nil {
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
		return &Response{}, err
	}
	log.Printf("end crawle url: %s", r.Url)
	return &Response{Body: body, Common: Common{
		Depth: r.Depth,
		Meta:  r.Meta,
	}}, nil
}

//确定读入
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
