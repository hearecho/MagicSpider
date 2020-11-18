# MagicSpider
go语言实现一个并发爬虫框架,目前仍然在完善中。初步具有一个爬虫应该具有的功能。

#### 1.如何开始
##### 下载依赖包
```shell script
go get github.com/hearecho/MagicSpider
```
##### 启动程序
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
##### 编写解析函数以及存储使用的Item
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
- [x] 增加多输出log
- [ ] 增加数据库存储
- [ ] 增加请求失败，请求超时，请求代理等等




