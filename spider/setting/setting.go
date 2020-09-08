package setting

import (
	"encoding/json"
	"github.com/hearecho/MagicSpider/spider/log"
	"io/ioutil"
	"os"
)

/**
setting 配置文件
*/
var (
	//finish crawl number
	Count int64
	//the limit of crawl
	TotalCount int64
	//request headers
	Headers map[string]string
	//存储类型
)
type Setting struct {
	DBusername string `json:"username"`
	DBpassword string `json:"password"`
	DBip       string `json:"ip"`
	DBname     string `json:"dbname"`
	Count      int64    `json:"count"`
	TotalCount int64    `json:"total_count"`
	Headers    map[string]string `json:"headers"`
}
var S Setting

func InitSetting() {
	data, err := ioutil.ReadFile("spider/setting/config.json")
	if err != nil {
		log.Error.Println("读取配置文件出错！")
		os.Exit(3)
	}
	err = json.Unmarshal(data, &S)
	if err != nil {
		log.Error.Println("初始化配置出错")
		os.Exit(3)
	}
}

