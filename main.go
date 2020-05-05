package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/autoPicture"
	"github.com/yesilin/go-cutting/logs"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/model/additional"
	"github.com/yesilin/go-cutting/model/web"
	"github.com/yesilin/go-cutting/settings"
	"github.com/yesilin/go-cutting/tools"
	"github.com/zserge/webview"
	"strings"
	"time"
)

func init() {
	model.InitNetwork()      // 没有网络不让使用
	model.InitNotification() // ps 未运行就进行通知
	model.InitFolder()       // 创建必须提前存在的文件夹
	model.InitScript()       // 创建必须提前准备的脚本
	go web.RunWebServer()    // 必须提前运行web服务器
	model.InitCipherList()   // 判断是否打开暗号列表
	logs.InitLog()           // 初始化日志
	// 管理员取得所有权
	//generate.TakeOwnership()

	// 实现快捷键 -1
	//go model.NegativeOne()
}

func main() {
	// 提示信息 剩余使用天数
	var tips string
	// 使用权限
	var power bool

	// 限制软件使用 2019.7.19
	// 定义私密文件路径
	PrivateFile, _ := tools.Home()
	PrivateFile = fmt.Sprintf("%s\\Documents\\Adobe\\Config.chx", PrivateFile)
	power, tips = model.RestrictingSoftwareUse2(PrivateFile, 1.001017, tools.GetNtpTime(), 30) // 这里改版本信息！！！！！！！！！！！！！！！！！！！！
	// 如果权限不是true
	if !power {
		fmt.Println(tips)
		time.Sleep(5 * time.Second) // 休眠五秒
		return
	}

	for {
		fmt.Println(tips) // 提示信息
		color.LightCyan.Println("\n " + (strings.Repeat("-", 18)) + " Welcome to the GoCutting v1.1.17 app " + strings.Repeat("-", 17))
		fmt.Println("\n:: 添加新暗号【--】返回上一次输入，例如镂空大小输错，返回重新输入镂空大小！")

		tips := `
   [1]快捷切图         [2]快捷贴图         [3]快捷效果         [4]自动套图

   [5]附加功能         [6]暗号列表         [7]设置中心         [8]帮助信息`
		fmt.Println(tips)
		factory := model.Input("\n:: 请选择上方的菜单功能：", false)
		tools.CallClear() // 清屏
		switch factory {
		case "1":
			model.OldFrameChoice() // 切图
		case "2":
			model.MapFrameChoice() // 贴图
		case "3":
			model.Choice() // 效果图
		case "4":
			autoPicture.Choice() // 套图
		case "5":
			additional.Additional() // 附加
		case "6":
			// 启动gui
			// 搭建web窗口
			go webview.Open("GoCutting", "http://localhost:12110/index", 350, 600, true)
			//go gui.RunWebview()
		case "7":
			settings.ModifySetting() // 设置
		case "8":
			model.Help() // 帮助
		case "-", "--":
			fmt.Println(":: 已经是最顶层菜单了，无需再返回，输入其他数字试下其他功能吧！")
			continue
		default:
			fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入！\n", model.ColourString(factory, ctc.ForegroundGreen))
			continue
		}
	}
}
