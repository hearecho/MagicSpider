package setting

/**
setting 配置文件
 */
var (
	//finish crawl number
	Count uint64
	//the limit of crawl
	TotalCount uint64
	//request headers
	Headers map[string]string
	//存储类型
	//db连接
	DBurl string
)

func InitSetting()  {
	Count = 1
	TotalCount = 2000
	Headers = map[string]string{
		"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
	}
	DBurl = "root:123456@tcp(127.0.0.1:3306)/spider?charset=utf8"
}
