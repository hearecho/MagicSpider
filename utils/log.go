package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}
	writers :=[]io.Writer{
		F,os.Stdout,
	}
	mutliOut := io.MultiWriter(writers...)
	logger = log.New(mutliOut, DefaultPrefix, log.LstdFlags)
}

func openLogFile(name string, path string) (*os.File, error) {
	err := MkDir(path)
	f,err:= Open(path+name,os.O_CREATE|os.O_APPEND, 0777)
	return f,err
}

func getLogFileName() string {
	return "log_"+strconv.Itoa(int(time.Now().Unix()))
}

func getLogFilePath() string {
	return "runtime/log/"
}


func Debug(format string,v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format+"\n",v)
}

func Info(format string,v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format+"\n",v)
}

func Warn(format string,v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format+"\n",v)
}

func Error(format string,v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format+"\n",v)
}

func Fatal(format string,v ...interface{}) {
	setPrefix(FATAL)
	logger.Printf(format+"\n",v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
