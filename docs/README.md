# MagicSpider
go语言实现一个并发爬虫框架,目前仍然在完善中。初步具有一个爬虫应该具有的功能。

#### 1.如何开始
##### 下载依赖包
```shell script
go get github.com/hearecho/MagicSpider
```
##### 启动程序
其中Request结构体中，`Url`就是请求链接，`Parse`为该链接对应的解析函数，`Common`
为中间多层抓取时间中间传递的参数。

`engine`中的构造参数主要是`workerCount`即下载器的协程数量，`startReequests`
即开始抓取时间的`Request`请求,`s`则是指的所使用的调度器。
```go
func main() {
	r := []MagicSpider.Request{
		{Url: "https://so.gushiwen.cn/gushi/tangshi.aspx",
			Parse: parse.NameParse,
			Common: MagicSpider.Common{
				Depth: 1,
				Meta:  &parse.Item{},
			}},
	}
	e := MagicSpider.engine{
		workerCount:   100,
		startRequests: r,
		s:             MagicSpider.NewSchedule(),
	}
	e.Go()
}
```
##### 配置文件
配置文件放置在当前工作环境下即可，配置文件格式为`config.yaml`。默认配置文件如下所示：
```yaml
base:
  spiderName: "spider"
  docType: "html"
  maxDepth: 0
  runtimePath: "runtime/"
  timout: 3
  rate: 10000
  logLevel: 1
```
##### 编写解析函数以及存储使用的Item
`Item`结构体字段可以自定义，但是要实现`Process`方法。`Process`方法用于对抓取到的信息进行存储。

解析函数都需要返回`*MagicSpider.ParseResult`，其字段如下：
```go
type ParseResult struct {
	Requests []Request //新的Request
	Items    []Item    //得到的Item
}
```

```go
type Item struct {
	Title   string
	Author  string
	Content string
}

func (i *Item)Process()  {
	utils.IsNotExistMkDir("runtime/")
	f, _ := utils.Open("runtime/result.txt",os.O_CREATE|os.O_APPEND,0777)
	item := fmt.Sprintf("【TiTle】:%v\t 【Author】:%v\t 【Content】:%v\n",i.Title,i.Author,strings.Trim(i.Content,"\n"))
	f.WriteString(item)
	f.Close()
}
func NameParse(r *MagicSpider.Response) *MagicSpider.ParseResult {
	//使用re进行
	nameRe := `<a href="(.*?)" target="_blank">([^<]+)</a>\((.*?)\)`
	re, _ := regexp.Compile(nameRe)
	result := re.FindAllSubmatch(r.Body, -1)
	res := &MagicSpider.ParseResult{}
	for _, item := range result {
		//新增url
		request := MagicSpider.Request{Url: "https://so.gushiwen.cn" + string(item[1]),
			Parse: ContentParse,
			Common:MagicSpider.Common{
				Depth: r.Depth+1,
				Meta:  &Item{Title: string(item[2]), Author: string(item[3])},
			}}
		res.Requests = append(res.Requests, request)
	}
	return res
}

func ContentParse(r *MagicSpider.Response) *MagicSpider.ParseResult {
	contentRe := `<div class="contson"[^>]+>([\s\s]*?)</div>`
	re, _ := regexp.Compile(contentRe)
	result := re.FindSubmatch(r.Body)
	res := &MagicSpider.ParseResult{}
	//content := Item{Content:string(result[1])}
	r.Meta.(*Item).Content = string(result[1])
	res.Items = append(res.Items, r.Meta.(*Item))
	return res

}
```
#### 2.后续计划

- [x] 实现配置功能
- [x] 增加类似jquery其他解析功能
- [x] 增加item处理工具
- [ ] ....




