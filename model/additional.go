package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"path/filepath"
	"strings"
)

// RunPSDRepairKit 运行PSD修复工具
func RunPSDRepairKit() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/software/PSDRepairKit/PSDRepairKit.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具PSD文件修复，请稍后...")
}

// RunAdvancedExcelRepair 运行XLS修复工具
func RunAdvancedExcelRepair() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/software/AdvancedExcelRepair/AER.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具XLS文件修复，请稍后...")
}

// RunRevokeMsgPatcher 运行QQ微信QQ防撤回工具
func RunRevokeMsgPatcher() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/software/RevokeMsgPatcher.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具微信QQ防撤回，请稍后...")
}

// RunW10DigitalActivation 运行win10数字激活工具
func RunW10DigitalActivation() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/software/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
	fmt.Println("\n:: win10系统已激活，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}

// ImportTakeOwnership 注册表导入取得文件所有权
func ImportTakeOwnership() {
	// 创建一个协程使用cmd启动外部程序
	cmd := exec.Command("cmd.exe", "/c", "regedit /s .\\config\\regedit\\takeOwnership.reg")
	go cmd.Run()
	fmt.Println("\n:: 右键菜单已添加，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}

// ImportNewTextFile 注册表导入右键新建文本文档
func ImportNewTextFile() {
	// 创建一个协程使用cmd启动外部程序
	cmd := exec.Command("cmd.exe", "/c", "regedit /s .\\config\\regedit\\newTextFile.reg")
	go cmd.Run()
	fmt.Println("\n:: 右键菜单已添加，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}

// RefreshFileTime 刷新文件属性的时间信息，例如创建时间，最后修改时间与访问时间
func RefreshFileTime(srcPath, dstPath string) {
	// 获取所有文件名，类型是字符串切片
	files, _ := filepath.Glob(fmt.Sprintf("%s/*", srcPath))

	// 没有文件就不执行
	if len(files) < 1 {
		fmt.Println("\n:: 刷新文件时间失败，因为工作文件夹下没有任何文件！")
		return
	}

	// 创建输出的目录
	_ = tools.CreateMkdirAll(dstPath)

	// 先准备要储存的位置
	newFiles := make([]string, len(files))
	// 遍历出新路径
	for i, v := range files {
		newFiles[i] = strings.Replace(v, srcPath, dstPath, 1)
	}

	// 直接复制到结果目录
	for i, v := range newFiles {
		tools.CopyFile(files[i], v)
	}

	fmt.Println("\n:: 文件属性时间已刷新，结果文件在工作文件夹下的 Result 目录下！")
}
