package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/gui"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/model/additional"
	"github.com/yesilin/go-cutting/model/layout"
	"github.com/yesilin/go-cutting/model/setting"
	"github.com/yesilin/go-cutting/tools"
	"github.com/zserge/webview"
	"strings"
	"time"
)

func init() {
	// 运行web服务器
	go gui.RunWebServer()

	go func() {
		// 导入注册表 使用正确的打开方式，并且取消脚本执行警告
		model.OpenMode()

		// 创建jsx文件夹
		_ = tools.CreateMkdirAll("config/jsx/temp")
		generate.SelectTailor()                         // 生成裁剪选择脚本备用
		generate.Tailor("")                             // 生成通用裁剪脚本备用
		generate.ClearMetadata()                        // 生成 -3 要用的清除元数据脚本备用
		generate.ClearMetadataNoPopUp()                 // 生成我自己动作要用的清除元数据脚本备用
		generate.BlackEdge()                            // 生成添加黑边脚本备用
		generate.SizeMarks()                            // 生成 将矩形选框转换为标记测量标志
		generate.SaveForWeb(viper.GetString("picture")) // 生成详情页指定保存位置
		generate.SaveAsJPEG()                           // 生成带自带清除元数据的另存脚本

		// 管理员取得所有权
		//generate.TakeOwnership()

		// 创建历史记录文件夹
		now := time.Now().Format("2006-01")
		_ = tools.CreateMkdirAll(fmt.Sprintf("Config/History/%s", now))

		// 创建套图文件夹
		_ = tools.CreateMkdirAll("config/Picture")

		// 创建备份文件夹
		_ = tools.CreateMkdirAll("config/Backups")
	}()

	// 实现快捷键 -1
	//go model.NegativeOne()
}

func main() {
	// 提示信息 剩余使用天数
	var tips string
	// 使用权限
	var power bool

	// 进行三次网络请求，都失败就当没有网络，并且不让使用
	for i := 0; i < 4; i++ {
		tools.CallClear() // 清屏

		// 有网直接退出循环
		if tools.IsNetwork() {
			color.LightCyan.Println("【验证】网络已连接服务器获取使用权限成功，请尽量不关闭软件，避免断网时无法使用！")
			break
		}

		// 第四次获取网络就当没网络
		if i == 3 {
			color.LightCyan.Println("【验证】网络已断开无法向服务器请求使用权限，软件将在五秒内自动关闭...")
			time.Sleep(5 * time.Second) // 休眠五秒
			return
		}

		color.LightCyan.Printf("【验证】第 %d 次网络连接失败，正在重新向服务器获取使用权限请稍等...", i+1)
		time.Sleep(2 * time.Second) // 休眠2秒
	}

	// 限制软件使用 2019.7.19
	// 定义私密文件路径
	PrivateFile, _ := tools.Home()
	PrivateFile = fmt.Sprintf("%s\\Documents\\Adobe\\Config.chx", PrivateFile)
	power, tips = model.RestrictingSoftwareUse2(PrivateFile, 1.001001, tools.GetNtpTime(), 30) // 这里改版本信息！！！！！！！！！！！！！！！！！！！！
	// 如果权限不是true
	if !power {
		fmt.Println(tips)
		time.Sleep(5 * time.Second) // 休眠五秒
		return
	}

	for {
		fmt.Println(tips) // 提示信息
		color.LightCyan.Println("\n " + (strings.Repeat("-", 20)) + " Welcome to the GoCutting v1.1.1 app " + strings.Repeat("-", 20))
		fmt.Println("\n【更新】添加新暗号【--】返回上一次输入，例如镂空大小输错，返回重新输入镂空大小！")

		tips := `
【菜单】[1]快捷切图        [2]快捷贴图        [3]快捷效果        [4]自动套图

【菜单】[5]附加功能        [6]暗号列表        [7]设置中心        [8]帮助信息`
		fmt.Println(tips)
		factory := model.Input("\n【菜单】请选择上方的菜单功能：", false)

		switch factory {
		case "1":
			tools.CallClear()      // 清屏
			model.OldFrameChoice() // 切图
		case "2":
			tools.CallClear()      // 清屏
			model.MapFrameChoice() // 贴图
		case "3":
			tools.CallClear() // 清屏
			model.Choice()    // 效果图
		case "4":
			tools.CallClear() // 清屏
			// 搭建web窗口
			//go webview.Open("GoCutting", "http://localhost:9090/autoNestingPictures", 350, 600, true)
			layout.Choice() // 套图
		case "5":
			tools.CallClear()       // 清屏
			additional.Additional() // 附加
		case "6":
			tools.CallClear() // 清屏
			// 启动gui
			// 搭建web窗口
			go webview.Open("GoCutting", "http://localhost:9090/index", 350, 600, true)
			//go gui.RunWebview()

		case "7":
			tools.CallClear()       // 清屏
			setting.ModifySetting() // 设置
		case "8":
			tools.CallClear() // 清屏
			model.Help()      // 帮助
		case "-":
			fmt.Println("【提示】已经是最顶层菜单了，无需再返回，输入其他数字试下其他功能吧！")
			continue
		default:
			tools.CallClear() // 清屏
			fmt.Printf("\n【错误】输入的 [%s] 不是已知的功能选项，请重新输入！\n", model.ColourString(factory, ctc.ForegroundGreen))
			continue
		}
	}
}
