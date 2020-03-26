// 启动暗号的函数汇总
package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)

// 启动暗号-1
func StartCode1() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/SelectTailor.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	// 每次选择正确的脚本时删除多余备份，最大保留30个
	go tools.DeleteRedundantBackups("Config/JSX/Temp/*", 30)
}

// 启动暗号-2
func StartCode2() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/newDocument.jsx"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}

// 启动暗号-3 深度清除源数据
func StartCode3() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/ClearMetadataJS.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 启动暗号-9
func StartCode9() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	fileName := time.Now().Format("2006-01-02")
	now := time.Now().Format("2006-01")

	// 储存历史记录路径
	path := fmt.Sprintf("Config/History/%s/%s.txt", now, fileName)

	// 先查看是否有历史记录文件
	exists, _ := tools.IsPathExists(path)
	// 如果找不到文件，就创建文件 头
	if !exists {
		fmt.Println("\n【错误】找不到今天的切图历史记录，可能今天还未开始切图，已自动打开历史文件夹！")

		exec.Command("cmd.exe", "/c", "start Config\\History").Run()
		tools.PrintLine(2)
		return
	}
	// 创建一个协程使用cmd来运行脚本
	cmd := exec.Command("cmd.exe", "/c", "start "+path)
	go cmd.Run()
}

// 启动暗号-98
func StartCode98() {
	// 自动套图工作路径
	picturePath := viper.GetString("picture")

	// 创建套图文件夹
	_ = tools.CreateMkdirAll(fmt.Sprintf("%s/主图", picturePath))
	// 创建一个协程使用cmd来运行脚本
	dataPath := "Config/JSX/SaveForWeb.jsx"
	exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	time.Sleep(time.Second) // 停一秒
	// 如果存在images就打开
	if ok, _ := tools.IsPathExists(fmt.Sprintf("%s/主图/images", picturePath)); ok {
		go exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图\\images", picturePath)).Run()
	} else {
		go exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图", picturePath)).Run()
	}
}

// 启动暗号-99
func StartCode99() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/exe/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}
