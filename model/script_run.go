package model

// 放一些jsx文件直接运行的函数

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

// RunGlobalHotkey 运行易语言写的辅助工具
func RunGlobalHotkey() {
	dataPath := "data/software/AuxiliaryTools/AuxiliaryTools.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunAutoCreateDocuments 是否打开自动新建Photoshop文档
func RunAutoCreateDocuments() {
	if viper.GetBool("openPs") { // 是否自动新建ps文档
		// 创建一个协程使用cmd来运行脚本
		dataPath := "data/jsx/newDocument.jsx"
		go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	}
}

// RunLoadSaveScript 根据当前文档名选择正确的快捷裁剪脚本 执行暗号-1
func RunLoadSaveScript() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "data/jsx/loadSaveScript.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	// 每次选择正确的脚本时删除多余备份，最大保留30个
	go tools.DeleteRedundantBackups("data/jsx/Temp/*", 100)
}

// RunClearMetadataStd 执行暗号-6 简单清除元数据
func RunClearMetadataStd() {
	dataPath := "data/jsx/clearMetadataStd.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunAddBlackEdge 执行暗号-7 为当前文档添加黑边
func RunAddBlackEdge() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "data/jsx/addBlackEdge.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunSaveAsJPEG 储存为jpeg格式的调用  暗号-10的实现
func RunSaveAsJPEG() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "data/JSX/SaveAsJPEG.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunSaveAllJPEG 将所有打开的文档储存为jpeg格式的调用 暗号-11
func RunSaveAllJPEG() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "data/JSX/saveAllJPEG.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunSaveAndCloseAllDocuments 保存并关闭全部文档的调用 暗号-12
func RunSaveAndCloseAllDocuments() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "data/jsx/saveAndCloseAllDocuments.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunReplaceSmartObjects 运行替换智能对象脚本 暗号-48
func RunReplaceSmartObjects() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "data/jsx/replaceSmartObjects.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// RunSaveForWeb 导出web格式的调用 暗号-49的实现
func RunSaveForWeb() {
	go func() {
		// 自动套图工作路径
		picturePath := viper.GetString("picture")

		// 创建套图文件夹
		_ = tools.CreateMkdirAll(fmt.Sprintf("%s/主图", picturePath))

		// 创建一个协程使用cmd来运行脚本
		dataPath := "data/jsx/saveForWeb.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

		//time.Sleep(time.Second) // 停一秒
		//
		//// 如果存在images就打开
		//if ok, _ := tools.IsPathExists(fmt.Sprintf("%s/主图/images", picturePath)); ok {
		//	exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图\\images", picturePath)).Run()
		//} else {
		//	exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图", picturePath)).Run()
		//}
	}()
}
