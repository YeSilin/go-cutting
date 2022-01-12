// Package controller  启动暗号命令的函数汇总
package controller

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/nested"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)

// Command1 执行暗号-1
func Command1() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/SelectTailor.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	// 每次选择正确的脚本时删除多余备份，最大保留30个
	go tools.DeleteRedundantBackups("Config/JSX/Temp/*", 30)
}

// Command2 执行暗号-2 重建新文档
func Command2() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/newDocument.jsx"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}

// Command3 执行暗号-3 深度清除源数据
func Command3() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/jsx/clearMetadata.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command4 执行暗号-4 快捷新建当天工作目录
func Command4() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now := time.Now().Format("2006-01-02")

	// 高老板
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/全镂空/半透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/全镂空/不透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/无镂空/半透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/无镂空/不透", now))

	// 这里的
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/御尚檀", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/御尚檀", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/岚湘", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/岚湘", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/沐兰", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/沐兰", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/华府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/华府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/木韵阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/木韵阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/金樽府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/金樽府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/怡柟", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/怡柟", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/舍得", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/舍得", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/西厢", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/西厢", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/藏湘阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/藏湘阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/阑若", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/阑若", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/木墨", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/木墨", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/墨屏", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/墨屏", now))

	cmd := exec.Command("cmd.exe", "/c", "start D:\\切图（请移动至个人目录）")
	cmd.Run()
}

// Command5 执行暗号-5 复制并关闭其他文档
func Command5() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/jsx/copyAndCloseOtherDocuments.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command6 执行暗号-6 简单清除元数据
func Command6() {
	dataPath := "config/jsx/clearMetadataNoPopUp.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command7 执行暗号-7 为当前文档添加黑边
func Command7() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/blackEdge.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command9 执行暗号-9 查询历史记录文件
func Command9() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	fileName := time.Now().Format("2006-01-02")
	now := time.Now().Format("2006-01")

	// 储存历史记录路径
	path := fmt.Sprintf("config/History/%s/%s.txt", now, fileName)

	// 先查看是否有历史记录文件
	exists, _ := tools.IsPathExists(path)
	// 如果找不到文件，就创建文件 头
	if !exists {
		//fmt.Println("\n【错误】找不到今天的切图历史记录，可能今天还未开始切图，已自动打开历史文件夹！")

		exec.Command("cmd.exe", "/c", "start config\\History").Run()
		//tools.PrintLine(2)
		return
	}
	// 创建一个协程使用cmd来运行脚本
	cmd := exec.Command("cmd.exe", "/c", "start "+path)
	go cmd.Run()
}

// Command10 执行暗号-10 快捷另存为jpg
func Command10() {
	model.SaveAsJPEG()
}

// Command11 执行暗号-11 快捷另存全部jpg
func Command11() {
	model.SaveAllJPEG()
}

// Command12 执行暗号-12 快捷保存并关闭全部文档
func Command12() {
	model.SaveAndCloseAllDocuments()
}

// Command41 执行暗号-41 快速打开套图文件夹
func Command41() {
	go tools.OpenFolder(viper.GetString("picture"), false)
}

// Command42 执行暗号-42 随机重命名
func Command42() {
	nested.RandomRenameFile(viper.GetString("picture"))
}


// Command97 执行暗号-97
func Command97() {
	nested.ReplaceDetailsPage(viper.GetString("picture")) // 替换详情页
}

// Command98 执行暗号-98
func Command98() {
	model.SaveForWeb()
}

// Command99 执行暗号-99
func Command99() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/software/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}
