package presenter

import (
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"io"
	"log"
	"os"
	"time"
)

// 日志的多个写入器
func logMultiWriter() io.Writer {
	writer1 := os.Stdout

	fileName := time.Now().Format("2006-01")
	fileName = "data/logs/" + fileName + ".log"
	writer2, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	return io.MultiWriter(writer1, writer2)
}

// InitLog 初始化日志
func InitLog() {
	// 设置日志输出级别
	logrus.SetLevel(logrus.InfoLevel)

	// 设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)

	// 创建日志文件夹
	_ = tools.CreateMkdirAll("data/logs")

	// 设置输出位置
	logrus.SetOutput(logMultiWriter())
}

/*
Panic：记录日志，然后panic。
Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
Error：错误日志，需要查看原因；
Warn：警告信息，提醒程序员注意；
Info：关键操作，核心流程的日志；
Debug：一般程序中输出的调试信息；
Trace：很细粒度的信息，一般用不到；
*/
func _() {

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
