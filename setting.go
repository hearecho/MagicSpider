package MagicSpider

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
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
}

var S  = &Setting{
	SpiderName:  "spider",
	TimeOut:     3*time.Second,
	RuntimePath: "runtime/",
	MaxDepth:    2,
}

func InitSetting()  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	path,_ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Printf("Fatal error config file: %s will use default configuration \n", err)
		return
	}
	S.MaxDepth = viper.GetInt("base.maxDepth")
	S.SpiderName = viper.GetString("base.spiderName")
	S.RuntimePath = viper.GetString("base.runtimePath")
	S.TimeOut = viper.GetDuration("base.timout")*time.Second
	log.Println("Use personal config file")
}
