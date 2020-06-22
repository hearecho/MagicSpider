package log


import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Error *log.Logger
)

func InitLogger()  {
	crawlerFile,err := os.OpenFile("crawler.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}
	Info = log.New(io.MultiWriter(os.Stdout,crawlerFile),"[Info]:",log.Ldate | log.Ltime | log.Lshortfile)

	Error = log.New(os.Stderr,"[Error]:",log.Ldate | log.Ltime | log.Lshortfile)
}

