package engine

/**
请求结构:主要包括请求所使用的URl，以及相应的解析函数
解析结果:新得到的URL,以及解析得到的数据
函数也作为一个参数传递。以便用户自定义而且可能不同的界面解析函数不同
 */
type Request struct {
	Url string
	ParseFunc func([]byte,map[string]string) ParseRes
}

type ParseRes struct {
	//解析得到的新的Url，构建的Request
	Requests []Request
	Items []interface{}
}



