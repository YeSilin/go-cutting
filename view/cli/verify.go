package cli

import (
	"github.com/gookit/color"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"time"
)

// VerifyNetwork 验证网络没有网络不让使用
func VerifyNetwork() {
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