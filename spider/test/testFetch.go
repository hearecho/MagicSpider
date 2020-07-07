package main

import (
	"MagicSpider/spider/log"
	"MagicSpider/spider/setting"
	"fmt"
)

func main() {
	log.InitLogger()
	setting.InitSetting()
	fmt.Println(setting.S.DBname)
}


