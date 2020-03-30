package quickCipher

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"strings"
	"time"
)

// 写入历史记录
func History(data string) {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now := time.Now().Format("2006-01")
	now1 := time.Now().Format("2006-01-02")
	now2 := time.Now().Format("15:04:05")

	// 储存历史记录路径
	path := fmt.Sprintf("Config/History/%s/%s.txt", now, now1)

	// 先查看是否有历史记录文件
	exists, _ := tools.IsPathExists(path)
	// 如果找不到文件，就创建文件 头
	if !exists {
		data := fmt.Sprintf("%s  切图历史记录由 GoCutting 自动生成\n", now1)
		tools.CreateFile(path, data)
	}

	// 时间线
	tempData := fmt.Sprintf("\n"+strings.Repeat("-", 40)+" %s "+strings.Repeat("-", 40)+"\n", now2)

	// 时间线加内容
	data = tempData + data

	// 追加写入历史
	tools.WriteFile(path, data)
}
