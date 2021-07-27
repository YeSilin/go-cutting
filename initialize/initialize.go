// Package initialize 初始化函数
package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"time"
)



// InitNotification 初始化通知
func InitNotification() {
	// ps 未运行就进行通知
	go func() {
		if ok := tools.IsExeRuning("Photoshop.exe", "Adobe"); !ok {
			tools.WinNotification("Photoshop 未运行", "大部分功能依赖于它，建议打开")
		}
	}()
}

// InitFolder 初始化文件夹
func InitFolder() {
	go func() {
		// 创建jsx文件夹
		_ = tools.CreateMkdirAll("config/jsx/temp")
		// 创建历史记录文件夹
		now := time.Now().Format("2006-01")
		_ = tools.CreateMkdirAll(fmt.Sprintf("config/History/%s", now))
		// 创建套图文件夹
		_ = tools.CreateMkdirAll("config/Picture")

		// 创建备份文件夹
		_ = tools.CreateMkdirAll("config/backups")
	}()
}

// InitScript 初始化脚本
func InitScript() {
	go func() {
		OpenMode()                                       // 导入注册表 使用正确的打开方式，并且取消脚本执行警告
		generate.SelectTailor()                          // 生成裁剪选择脚本备用
		generate.GeneralCutting("")                      // 生成通用裁剪脚本备用
		generate.ClearMetadata()                         // 生成 -3 要用的清除元数据脚本备用
		generate.ClearMetadataNoPopUp()                  // 生成我自己动作要用的清除元数据脚本备用
		generate.BlackEdge()                             // 生成添加黑边脚本备用
		model.SaveForWebInit(viper.GetString("picture")) // 生成详情页指定保存位置
		model.SaveAsJPEGInit()                           // 生成带自带清除元数据的另存脚本
		model.SaveAllJPEGInit()                          // 生成另存全部文件脚本
		model.SaveAndCloseAllDocumentsInit()          // 生成保存并关闭全部文档的脚本
		generate.CopyAndCloseOtherDocuments()            // 生成复制并关闭其他文档脚本
	}()
}

// InitCipherList 初始化暗号列表
func InitCipherList() {

	go func() {
		// 是否自动开启暗号列表
		if viper.GetBool("cipherList") {
			// 搭建web窗口
			//webview.Open("GoCutting", "http://localhost:12110/index", 350, 600, true)
			//unclassified.RunWebview()
		}
	}()
}
