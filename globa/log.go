package globa

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"log"
	"os"
	"time"
)

// 定义全局日志
var Logger *log.Logger

// 日志的配置
func init() {
	// 创建日志文件夹
	_ = tools.CreateMkdirAll("config/logs")

	// 日志路径和名字
	fileName := fmt.Sprintf("config/logs/error_%s.log", time.Now().Format("20060102"))

	// 打开日志文件
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	// 设置日志输出位置为文件 文件全路径名+行号： /a/b/c/d.go:23   日期：2019/07/09
	Logger = log.New(logFile, "", log.Lshortfile|log.Ldate|log.Ltime)
}
