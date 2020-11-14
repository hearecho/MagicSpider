package MagicSpider

//定义Request等等信息

type Request struct {
	Url     string                       //url
	Parse   func(*Response) ParseResult //解析器,一个接口
	Headers map[string]string
	Common
}

type ParseResult struct {
	Requests []Request //新的Request
	Items    []Item    //得到的Item
}

type Response struct {
	Body []byte //爬取内容
	Doc  interface{}
	Common
}

func NIlParser(*Response) ParseResult {
	return ParseResult{}
}

type Common struct {
	Depth int         //爬取深度
	Meta  interface{} //在不同层次的requests中传递信息
}
