package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/tools"
)



/**快捷键说明*/
func key() {
	tips := `
:: 目前已实现的全局暗号如下：

   [-]返回上一级             [-1]裁剪快捷键             [-2]重建新文档

   [-3]深度清理PSD           [-4]快捷文件夹             [-5]复制其他层

   [-6]快速清理PSD           [-7]自动加黑边             [-8]清屏快捷键

   [-9]到切图历史            [-10]单文档另存            [-11]全文档另存

   [-12]全文档关存           [-95]功能未开发            [-96]功能未开发

   [-97]替换详情页           [-98]导出详情页            [-99]激活win10系统`
	fmt.Println(tips)
}

/**切图注意事项*/
func careful() {
	fmt.Println("\n:: 新旧厂切图需要注意的规则如下：")
	fmt.Println("\n   1. 切图的单位是 厘米；")
	fmt.Println("\n   2. 分辨率是 100像素/英寸；")
	fmt.Println("\n   3. 颜色模式是 CMYK；")
	fmt.Println("\n   4. 颜色配置文件是 工作中的CMYK:Japan color 2011 Coated；")
	fmt.Println("\n   5. 切图时动物和文字都不能被切到，并且太阳和月亮这种正圆不能变形；")
	fmt.Println("\n   6. 四周是纯白底色的时候要加 0.1厘米的黑色描边 快捷键是 Alt + Ctrl + C；")
	fmt.Println("\n   7. 切图时半透或透光不透影最大148的宽，不透最大180的宽；")
	fmt.Println("\n   8. 切图遇到不透画布并且双面图案不一样时，每张需额外备注：“印一张”；")
	fmt.Printf("\n   9. 目前软件公式中，旧厂订布预留是 %.2f 厘米！。\n", viper.GetFloat64("reserve"))

	fmt.Println("\n\n:: 将已切好的图片发送给以下人员：")
	fmt.Println("\n   大部分切图不透与半透 --> 直接发群里\t卷帘画布 --> 单独发给厂长。")
}

/**软件使用技巧*/
func skill() {
	tips := `
:: 大部分暗号可以任意调用，但有些暗号使用前需满足一些条件：

   [-]这是一个全局通用的返回功能，不管在任何界面都将返回上一级菜单

   [--]返回上一次输入，例如镂空大小输错，返回重新输入镂空大小

   [-1]优化另存快速裁剪，折屏等复杂框架切图核心，是主要优化的一个功能

   [-2]重建新文档功能，只能重建前一次使用本软件所创建的ps文档

   [-3]清除元数据功能，顾名思义就是清除垃圾并且不损失清晰度

   [-4]快捷文件夹功能，此功能创建的文件夹在桌面可以看到

   [-5]复制其他文档的全部图层并关闭其他文档，关闭时不保存修改

   [-7]自动加黑边功能，现在不管是打开任何ps文档都支持加黑边了

   [-10]另存为JPG前进行一次元数据清除，并自动选择最佳参数

   [-11]另存全部打开的文档为JPG，文件名带白底图三字会多保存一张PNG

   [-12]保存并关闭全部打开的文档，请勿在有打开原图的情况下使用

   [-98]导出Web所用格式，用于详情页快速导出切片使用
	
   [-99]激活win10系统，这是一个集成大神之作的功能，只能激活win10系统`
	fmt.Println(tips)
}

/**初始化*/
func (c *CLI) help() {
OuterLoop:
	for {
		tools.EnglishTitle("Help info", 74)
		fmt.Println("\n:: 此项目是通过注入JS脚本对PS进行短暂的间接控制，非实时监控，资源消耗极低！")
		fmt.Println("\n   [1]查看快捷暗号.            [2]查看切图规则.            [3]查看功能说明.")
		help, info := input.InputMenuSelection("\n:: 请选择需要查看的帮助：")
		tools.CallClear() // 清屏
		switch help {
		case "1":
			key() // 快捷键
		case "2":
			careful() // 注意事项
		case "3":
			skill() // 软件使用技巧
		case "-":
			break OuterLoop //跳出循环
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(help, ctc.ForegroundGreen))
			}
		}
	}
}