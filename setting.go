package MagicSpider

import (
	"fmt"
	"github.com/hearecho/MagicSpider/utils"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"time"
)

// Setting 配置
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
	//日志级别
	LogLevel int
	// 是否是分布式
	Distribute bool
	StaDB      string
}

var S = &Setting{
	SpiderName:  "spider",
	TimeOut:     3 * time.Second,
	RuntimePath: "runtime/",
	MaxDepth:    2,
	Rate:        10000,
	DocType:     "html",
	LogLevel:    1,
	Distribute:  false,
	StaDB:       "",
}

func printSettings() {
	// 通过反射获取结构体中的值
	t := reflect.TypeOf(*S)
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		v := reflect.ValueOf(*S)
		val := v.FieldByName(name)
		fmt.Printf("%s:%v\n", name, val)
	}
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
		utils.Error(fmt.Sprintf("Fatal error config file: will use default configuration \n", err))
		printSettings()
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
	if viper.IsSet("base.logLevel") {
		S.LogLevel = viper.GetInt("base.logLevel")
	}
	if viper.IsSet("base.distribute") {
		// 只有在分布式环境下才会读取base.staDB
		S.Distribute = viper.GetBool("base.distribute")
	}
	if viper.IsSet("base.staDB") {
		S.StaDB = viper.GetString("base.staDB")
	}
	utils.Info("读取配置如下:")
	printSettings()
}
