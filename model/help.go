package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/globa"
	"github.com/yesilin/go-cutting/tools"
)

/**初始化*/
func Help() {
	EnglishTitle("Help", 79)
	fmt.Println("\n【帮助】此项目是通过注入JS脚本对PS进行短暂的间接控制，非实时监控，资源消耗极低！")

	for {
		fmt.Println("\n【帮助】[1]查看快捷暗号\t\t[2]查看切图规则\t\t[3]查看功能说明")
		var help = Input("\n【帮助】请选择需要查看的帮助：", false)
		switch help {
		case "1":
			tools.CallClear() // 清屏
			key() // 快捷键
		case "2":
			tools.CallClear() // 清屏
			careful() // 注意事项
		case "3":
			tools.CallClear() // 清屏
			skill() // 软件使用技巧
		case "-":
			goto FLAG //跳出循环

		default:
			continue
		}
	}

FLAG: // 跳到函数结束

}

/**快捷键说明*/
func key() {

	fmt.Println("\n【帮助】目前已实现的全局暗号如下：")

	fmt.Println("\n\t[-]返回上一级\t\t[-1]裁剪快捷键\t\t[-2]重建新文档")
	fmt.Println("\n\t[-3]深度清理PSD\t\t[-4]快捷文件夹\t\t[-5]画选区测距")
	fmt.Println("\n\t[-6]快速清理PSD\t\t[-7]自动加黑边\t\t[-8]清屏快捷键")
	fmt.Println("\n\t[-9]到切图历史\t\t[-11]小座屏\t\t[-12]左右镂空")
	fmt.Println("\n\t[-13]中间大两边小\t[-14]上下镂空\t\t[-15]顶天立地")
	fmt.Println("\n\t[-16]常规折屏\t\t[-17]三座屏\t\t[-96]功能未开发")
	fmt.Println("\n\t[-97]功能未开发\t\t[-98]快速导出图片\t[-99]激活win10系统")

	tools.PrintLine(2)
}

/**切图注意事项*/
func careful() {
	fmt.Println("\n【帮助】新旧厂切图需要注意的规则如下：")
	fmt.Println("\n    1. 切图的单位是 厘米")
	fmt.Println("\n    2. 分辨率是 100像素/英寸")
	fmt.Println("\n    3. 颜色模式是 CMYK")
	fmt.Println("\n    4. 颜色配置文件是 工作中的CMYK:Japan color 2001 Coated")
	fmt.Println("\n    5. 切图时动物和文字都不能被切到")
	fmt.Println("\n    6. 四周是纯白底色的时候要加 0.5厘米的黑色描边 快捷键是 Alt + Ctrl + C")
	fmt.Println("\n    7. 切图时半透最大148的宽，不透最大180的宽")
	fmt.Println("\n    8. 切图遇到不透画布并且双面图案不一样时，每张需额外备注：“印一张”")
	fmt.Printf("\n    9. 目前程序公式中，旧厂订布预留是 %.2f 厘米！\n", globa.NowSetting.Reserve)

	fmt.Println("\n\n【帮助】将已切好的图片发送给以下人员：")
	fmt.Println("\n    大部分切图不透与半透 --> 直接发群里\t卷帘画布 --> 单独发给厂长")
	tools.PrintLine(2)
}

/**软件使用技巧*/
func skill() {
	fmt.Println("\n【帮助】大部分暗号可以任意调用，但有些暗号使用前需满足一些条件：")
	fmt.Println("\n    暗号[-]这是一个全局通用的返回功能，不管在任何界面都将返回上一级菜单")
	fmt.Println("\n    暗号[-1]优化另存快速裁剪，是主要优化的一个功能，快捷键是 Alt+X ")
	fmt.Println("\n    暗号[-2]重建新文档功能，只能重建前一次使用本软件所创建的ps文档")
	fmt.Println("\n    暗号[-3]清除源数据功能，目的是为了让文件之后保存的更加小巧并且不损失清晰度")
	fmt.Println("\n    暗号[-4]快捷文件夹功能，此功能创建的文件夹在桌面可以看到")
	fmt.Println("\n    暗号[-5]先自己画一个矩形选区，然后再输入暗号转为测距标志")
	fmt.Println("\n    暗号[-7]自动加黑边功能，现在不管是打开任何ps文档都支持加黑边了")
	fmt.Println("\n    暗号[-99]激活win10系统，这是一个集成大神之作的功能，只能激活win10系统")
	tools.PrintLine(2)
}
