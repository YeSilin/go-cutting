package view

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/tools"
	"strings"
)

type CliView struct {

	//定义必要字段
	key  string //接收用户输入...
	info string //接收输入后返回的具体执行的方案的文本提示

	Version float64 //版本号 暂时先继续使用浮点数，后期改成字符串
	Expire int64 // 限制最长使用时间
	Tips    string  //提示信息 剩余使用天数
	Power   bool    //使用权限
}

//NewCliView 视图的构造方法
func NewCliView() *CliView {
	return &CliView{
		Version: 1.000000,
		Power:   false,
	}

}




//MainMenu 显示主菜单
func (v *CliView) MainMenu() {

	for {
		fmt.Println(v.Tips) // 提示信息
		// 装换版本为字符串
		versionStr := fmt.Sprintf("%d.%d.%d", int(v.Version), int64(v.Version*1000000/1000)%1000, tools.Float64ToInt64(v.Version*1000000)%1000)
		color.LightCyan.Println("\n " + (strings.Repeat("-", 18)) + fmt.Sprintf(" Welcome to the GoCutting v%s app ", versionStr) + strings.Repeat("-", 17))
		fmt.Println("\n:: 添加新暗号【--】返回上一次输入，例如镂空大小输错，返回重新输入镂空大小！")

		tips := `
   [1]快捷切图.        [2]快捷贴图.        [3]快捷效果.        [4]自动套图.

   [5]附加功能.        [6]暗号列表         [7]设置中心.        [8]帮助信息.`
		fmt.Println(tips)

		//factory, info := unclassified.Input("\n:: 请选择上方的菜单功能：", false, true)
		v.key, v.info = input.InputMenuSelection("\n:: 请选择上方的菜单功能：")

		tools.CallClear() // 清屏
		switch v.key {
		case "1":
			v.oldFrameChoice() // 切图
		case "2":
			v.mapFrameChoice() // 贴图
		case "3":
			v.temporaryChoice() // 效果图
		case "4":
			v.autoPictureChoice() // 套图
		case "5":
			v.extendChoice() // 附加
		case "6":
			// 启动gui
			//go unclassified.RunWebview()
		case "7":
			v.settingsChoice() // 设置
		case "8":
			v.help() // 帮助
		case "-":
			fmt.Println("\n:: 已经是最顶层菜单了，无需再返回，输入其他数字试下其他功能吧！")
			continue
		default:
			if len(v.info) != 0 {
				fmt.Println(v.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(v.key, ctc.ForegroundGreen))
			}
		}
	}

}
