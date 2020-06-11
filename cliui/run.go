package cliui

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"github.com/zserge/webview"
	"strings"
)

// 开始运行主体
func Run(tips string, version float64) {
	for {
		fmt.Println(tips) // 提示信息
		// 装换版本为字符串
		versionStr := fmt.Sprintf("%d.%d.%d", int(version), int64(version*1000000/1000)%1000, tools.Float64ToInt64(version*1000000)%1000)
		color.LightCyan.Println("\n " + (strings.Repeat("-", 18)) + fmt.Sprintf(" Welcome to the GoCutting v%s app ", versionStr) + strings.Repeat("-", 17))
		fmt.Println("\n:: 添加新暗号【--】返回上一次输入，例如镂空大小输错，返回重新输入镂空大小！")

	//	tips := `
   //[1]快捷切图         [2]快捷贴图         [3]快捷效果         [4]自动套图
   //
   //[5]附加功能         [6]暗号列表         [7]设置中心         [8]帮助信息`
		tips := `
   [1]快捷切图.        [2]快捷贴图.        [3]快捷效果.        [4]自动套图.

   [5]附加功能.        [6]暗号列表         [7]设置中心.        [8]帮助信息.`
		fmt.Println(tips)

		//factory, info := model.Input("\n:: 请选择上方的菜单功能：", false, true)
		factory, info := model.InputMenuSelection("\n:: 请选择上方的菜单功能：")

		tools.CallClear() // 清屏
		switch factory {
		case "1":
			oldFrameChoice() // 切图
		case "2":
			mapFrameChoice() // 贴图
		case "3":
			temporaryChoice() // 效果图
		case "4":
			autoPictureChoice() // 套图
		case "5":
			additionalChoice() // 附加
		case "6":
			// 启动gui
			// 搭建web窗口
			go webview.Open("GoCutting", "http://localhost:12110/index", 350, 600, true)
			//go gui.RunWebview()
		case "7":
			ModifySetting() // 设置
		case "8":
			Help() // 帮助
		case "-":
			fmt.Println("\n:: 已经是最顶层菜单了，无需再返回，输入其他数字试下其他功能吧！")
			continue
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else{
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(factory, ctc.ForegroundGreen))
			}
		}
	}
}
