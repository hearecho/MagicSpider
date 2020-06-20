package fetch

import (
	"io/ioutil"
	"net/http"
	"strings"
)


func Fetch(url,method string,headers map[string]string,form string) ([]byte,error) {
	client := &http.Client{}
	method = strings.ToUpper(method)
	req,err := http.NewRequest(method,url,strings.NewReader(form))
	if err != nil {
		return nil,nil
	}
	if strings.EqualFold(method,"POST") {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for k,v := range headers {
		req.Header.Set(k,v)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK{
		return nil,nil
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,err
}


