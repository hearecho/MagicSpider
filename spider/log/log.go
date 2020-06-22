package log

import (
	"log"
	"os"
)

var (
	Info *log.Logger
	Error *log.Logger
)

func InitLogger()  {
	Info = log.New(os.Stdout,"[Info]:",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(os.Stdout,"[Error]:",log.Ldate | log.Ltime | log.Lshortfile)
}
