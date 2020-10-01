package types

//定义Request等等信息

type Request struct {
	Url   string                    //url
	Parse func(*Response) *ParseResult //解析器,一个接口
	Depth int                       //爬取深度
}

type ParseResult struct {
	Requests []Request     //新的Request
	Items    []interface{} //得到的Item
}

type Response struct {
	Depth int //抓取深度
	Body []byte //爬取内容
}

