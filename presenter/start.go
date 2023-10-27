package presenter

// 软件启动一些必要的准备文件

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"time"
)

// InitFolder 初始化文件夹
func InitFolder() {
	go func() {
		// 创建jsx文件夹
		_ = tools.CreateMkdirAll("data/jsx/temp")
		// 创建历史记录文件夹
		now := time.Now().Format("2006-01")
		_ = tools.CreateMkdirAll(fmt.Sprintf("data/History/%s", now))
		// 创建套图文件夹
		_ = tools.CreateMkdirAll("data/Picture")

		// 创建备份文件夹
		_ = tools.CreateMkdirAll("data/backups")
	}()
}

// InitScript 初始化脚本
func InitScript() {
	go func() {
		OpenMode()                                   // 导入注册表 使用正确的打开方式，并且取消脚本执行警告
		model.LoadSaveScript()                       // 生成裁剪选择脚本备用
		model.FrameSaveDef("")                       // 生成通用裁剪脚本备用
		model.ClearMetadata()                        // 生成 -3 要用的清除元数据脚本备用
		model.ClearMetadataStd()                     // 生成我自己动作要用的清除元数据脚本备用
		model.AddBlackEdge()                         // 生成 -7 添加黑边脚本备用
		model.SaveForWeb(viper.GetString("picture")) // 生成详情页指定保存位置
		model.SaveAsJPEG()                           // 生成带自带清除元数据的另存脚本
		model.SaveAllJPEG()                          // 生成另存全部文件脚本
		model.SaveAndCloseAllDocuments()             // 生成保存并关闭全部文档的脚本
		model.CopyAndCloseOtherDocuments()           // 生成 -5 复制并关闭其他文档脚本
		model.CopyOriginalImageStamping()            // 生成 -4 复制原图盖印
	}()
}

// InitHotkey 初始化全局热键
func InitHotkey() {
	hkey := hotkey.New()

	// 注册F1热键
	hkey.Register(hotkey.None, hotkey.F1, func() {
		Command1()
	})
	// 注册F4热键
	hkey.Register(hotkey.None, hotkey.F4, func() {
		Command4()
	})
}
