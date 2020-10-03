package types

import items2 "github.com/hearecho/MagicSpider/crawler/CurrencySpiderArch/items"

//定义Request等等信息

type Request struct {
	Url   string                    //url
	Parse func(*Response) *ParseResult //解析器,一个接口
	Depth int                       //爬取深度
}

type ParseResult struct {
	Requests []Request     //新的Request
	Items    []items2.Item//得到的Item
}

type Response struct {
	Depth int //抓取深度
	Body []byte //爬取内容
}

//用于在多层爬虫中传递得到的item
type Item struct {
	//每个信息专属的requestId
	RequestID string
	Data interface{}
}

func NIlParser(*Response) *ParseResult {
	return &ParseResult{}
}

