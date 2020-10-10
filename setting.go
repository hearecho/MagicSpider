package MagicSpider

import "time"

/**
从配置文件中读取配置
*/
type Setting struct {
	//爬虫名称
	SpiderName string
	//访问超时
	TimeOut time.Duration
	//数据存储
	DataBase
}
type DataBase struct {
	Url      string
	UserName string
	Password string
	DbName   string
}

