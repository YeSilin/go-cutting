package tools

import "time"

// 返回当前时间
func NowTime() (now string) {
	// 获取当前时间，进行格式化 2006-01-02 15:04:05
	return time.Now().Format("060102150405")
}
