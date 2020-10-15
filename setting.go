package MagicSpider

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

/**
从配置文件中读取配置
*/
type Setting struct {
	//爬虫名称
	SpiderName string
	//访问超时
	TimeOut time.Duration
	//runtime 路径
	RuntimePath string
	//MAxDepth 爬取最大深度
	MaxDepth int
	//爬取速率 1S爬取几次,默认速率1s 1w次(基本等于不限速)
	Rate int
	//返回文件类型 html,json 默认json
	DocType string
}

var S = &Setting{
	SpiderName:  "spider",
	TimeOut:     3 * time.Second,
	RuntimePath: "runtime/",
	MaxDepth:    2,
	Rate:        10000,
	DocType:     "html",
}

func InitSetting() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Printf("Fatal error config file: %s will use default configuration \n", err)
		return
	}
	if viper.IsSet("base.maxDepth") {
		S.MaxDepth = viper.GetInt("base.maxDepth")
	}
	if viper.IsSet("base.spiderName") {
		S.SpiderName = viper.GetString("base.spiderName")
	}
	if viper.IsSet("base.runtimePath") {
		S.RuntimePath = viper.GetString("base.runtimePath")
	}
	if viper.IsSet("base.timout") {
		S.TimeOut = viper.GetDuration("base.timout") * time.Second
	}
	if viper.IsSet("base.rate") {
		S.Rate = viper.GetInt("base.rate")
	}
	if viper.IsSet("base.docType") {
		S.DocType = viper.GetString("base.docType")
	}
	log.Printf("Use personal config file%v\n", S)
}
