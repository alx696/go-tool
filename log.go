package tool

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var (
	LI *log.Logger
	LW *log.Logger
)

func init() {
	logDir := filepath.Join(os.Getenv("HOME"), "log")
	os.MkdirAll(logDir, fs.ModePerm)

	infoFile, e := os.OpenFile(filepath.Join(logDir, "point-converter-info.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if e != nil {
		log.Fatalln("创建日志文件出错：", e)
	}

	errorFile, e := os.OpenFile(filepath.Join(logDir, "point-converter-warn.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if e != nil {
		log.Fatalln("创建日志文件出错：", e)
	}

	LI = log.New(infoFile, "[INFO ]", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	LW = log.New(errorFile, "[WARN ]", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
