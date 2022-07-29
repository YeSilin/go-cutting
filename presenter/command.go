// Package presenter  启动暗号命令的函数汇总
package presenter

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)

// Command1 执行暗号-1
func Command1() {
	model.RunLoadSaveScript()
}

// Command2 执行暗号-2 重建新文档
func Command2() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "resources/jsx/newDocument.jsx"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}

// Command3 执行暗号-3 深度清除源数据
func Command3() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "resources/jsx/clearMetadata.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command5 执行暗号-5 复制并关闭其他文档
func Command5() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "resources/jsx/copyAndCloseOtherDocuments.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// Command6 执行暗号-6 简单清除元数据
func Command6() {
	model.RunClearMetadataStd()
}

// Command7 执行暗号-7 为当前文档添加黑边
func Command7() {
	model.RunAddBlackEdge()
}

// Command9 执行暗号-9 查询历史记录文件
func Command9() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	fileName := time.Now().Format("2006-01-02")
	now := time.Now().Format("2006-01")

	// 储存历史记录路径
	path := fmt.Sprintf("resources/History/%s/%s.txt", now, fileName)

	// 先查看是否有历史记录文件
	exists, _ := tools.IsPathExists(path)
	// 如果找不到文件，就创建文件 头
	if !exists {
		//fmt.Println("\n【错误】找不到今天的切图历史记录，可能今天还未开始切图，已自动打开历史文件夹！")

		exec.Command("cmd.exe", "/c", "start resources\\History").Run()
		//tools.PrintLine(2)
		return
	}
	// 创建一个协程使用cmd来运行脚本
	cmd := exec.Command("cmd.exe", "/c", "start "+path)
	go cmd.Run()
}

// Command10 执行暗号-10 快捷另存为jpg
func Command10() {
	model.RunSaveAsJPEG()
}

// Command11 执行暗号-11 快捷另存全部jpg
func Command11() {
	model.RunSaveAllJPEG()
}

// Command12 执行暗号-12 快捷保存并关闭全部文档
func Command12() {
	model.RunSaveAndCloseAllDocuments()
}

// Command41 执行暗号-41 快速打开套图文件夹
func Command41() {
	go tools.OpenFolder(viper.GetString("picture"), false)
}

// Command42 执行暗号-42 随机重命名
func Command42() {
	RandomRenameFile(viper.GetString("picture"))
}

// Command48 执行暗号-48
func Command48() {
	ReplaceSmartObjects(viper.GetString("picture")) // 替换详情页
}

// Command49 执行暗号-49
func Command49() {
	model.RunSaveForWeb()
}

// Command98 执行暗号-98
func Command98() {
	model.RunSaveForWeb()
}

// Command99 执行暗号-99
func Command99() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "resources/software/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}

// SelectCommand 选择一个暗号执行，并且返回一些提示信息
func SelectCommand(Command string) (ok bool, info string) {
	// 开始指定功能
	switch Command {
	case "-1":
		Command1()
		return true, ":: 检测到输入的内容为隐藏暗号，正在调用快捷裁剪..."
	case "-2":
		Command2()
		return true, ":: 检测到输入的内容为隐藏暗号，正在重建新文档..."
	case "-3":
		Command3() // 深度清除源数据
		return true, ":: 检测到输入的内容为隐藏暗号，正在深度清理PSD..."
	case "-4":
		return false, ""
	case "-5":
		Command5() // 复制并关闭其他文档
		return true, ":: 检测到输入的内容为隐藏暗号，正在复制并关闭其他文档..."
	case "-6":
		Command6() // 简单清除元数据
		return true, ":: 检测到输入的内容为隐藏暗号，正在快速清理PSD..."
	case "-7":
		Command7() // 为当前文档添加黑边
		return true, ":: 检测到输入的内容为隐藏暗号，正在为当前文档添加黑边..."
	case "-8":
		tools.CallClear() // 清屏
		return true, ""
	case "-9":
		Command9() // 打开历史记录
		return true, ":: 检测到输入的内容为隐藏暗号，正在打开切图历史..."
	case "-10":
		Command10() // 快捷另存为jpg
		return true, ""
	case "-11":
		Command11() // 快捷另存全部打开的文件
		return true, ":: 检测到输入的内容为隐藏暗号，正在另存全部打开的文件..."
	case "-12":
		Command12() // 快捷保存并关闭全部文档
		return true, ":: 检测到输入的内容为隐藏暗号，正在保存并关闭全部文档..."
	case "-41":
		Command41() // 快捷保存并关闭全部文档
		return true, ":: 检测到输入的内容为隐藏暗号，已打开套图文件夹..."
	case "-42":
		Command42()
		return true, ":: 检测到输入的内容为隐藏暗号，已执行随机重命名..."
	case "-48":
		Command48()
		return true, ":: 检测到输入的内容为隐藏暗号，正在替换名字以 [DP] 开头的智能对象图层..."
	case "-49":
		Command49()
		return true, ":: 检测到输入的内容为隐藏暗号，正在导出为Web所用格式..."
	case "-97":
		return true, ""
	case "-98":
		Command98()
		return true, ":: 检测到输入的内容为隐藏暗号，正在导出为Web所用格式..."
	case "-99":
		Command99()
		return true, ":: 检测到输入的内容为隐藏暗号，正在后台激活Win10系统..."
	default:
		return false, ""
	}
}
