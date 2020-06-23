// 初始化
package model

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"time"
)

// 初始化网络
func InitNetwork() {
	// 进行三次网络请求，都失败就当没有网络，并且不让使用
	for i := 0; i < 4; i++ {
		tools.CallClear() // 清屏

		// 有网直接退出循环
		if tools.IsNetwork() {
			color.LightCyan.Println(":: 网络已连接服务器获取使用权限成功，请尽量不关闭软件，避免断网时无法使用！")
			break
		}

		// 第四次获取网络就当没网络
		if i == 3 {
			color.LightCyan.Println(":: 网络已断开无法向服务器请求使用权限，软件将在五秒内自动关闭...")
			time.Sleep(5 * time.Second) // 休眠五秒
			os.Exit(1)
		}

		color.LightCyan.Printf(":: 第 %d 次网络连接失败，正在重新向服务器获取使用权限请稍等...", i+1)
		time.Sleep(2 * time.Second) // 休眠2秒
	}
}

// 初始化通知
func InitNotification() {
	// ps 未运行就进行通知
	go func() {
		if ok := tools.IsExeRuning("Photoshop.exe", "Adobe"); !ok {
			tools.WinNotification("Photoshop 未运行", "大部分功能依赖于它，建议打开")
		}
	}()
}

// 初始化文件夹
func InitFolder() {
	go func() {
		// 创建jsx文件夹
		_ = tools.CreateMkdirAll("config/jsx/temp")
		// 创建历史记录文件夹
		now := time.Now().Format("2006-01")
		_ = tools.CreateMkdirAll(fmt.Sprintf("Config/History/%s", now))
		// 创建套图文件夹
		_ = tools.CreateMkdirAll("config/Picture")

		// 创建备份文件夹
		_ = tools.CreateMkdirAll("config/Backups")
	}()
}

// 初始化脚本
func InitScript() {
	go func() {
		OpenMode()                                      // 导入注册表 使用正确的打开方式，并且取消脚本执行警告
		generate.SelectTailor()                         // 生成裁剪选择脚本备用
		generate.GeneralCutting("")                     // 生成通用裁剪脚本备用
		generate.ClearMetadata()                        // 生成 -3 要用的清除元数据脚本备用
		generate.ClearMetadataNoPopUp()                 // 生成我自己动作要用的清除元数据脚本备用
		generate.BlackEdge()                            // 生成添加黑边脚本备用
		generate.SaveForWeb(viper.GetString("picture")) // 生成详情页指定保存位置
		generate.SaveAsJPEG()                           // 生成带自带清除元数据的另存脚本
		generate.SaveAllJPEG()                          // 生成另存全部文件脚本
		generate.SaveAndCloseAllDocuments()             // 生成保存并关闭全部文档的脚本
		generate.CopyAndCloseOtherDocuments()           // 生成复制并关闭其他文档脚本
	}()
}

// 初始化暗号列表
func InitCipherList() {

	go func() {
		// 是否自动开启暗号列表
		if viper.GetBool("cipherList") {
			// 搭建web窗口
			//webview.Open("GoCutting", "http://localhost:12110/index", 350, 600, true)
			RunWebview()
		}
	}()
}
