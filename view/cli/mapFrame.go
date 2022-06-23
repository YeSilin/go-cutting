package cli

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"github.com/yesilin/go-cutting/unclassified"
)

// 贴图框架的选择
func (c *CLI) mapFrameChoice() {

OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)
		tools.EnglishTitle("3ds Max map frame", 74)
		text := `
:: 下方所有框架的切图单位均是像素，小数点位之后值均会被舍弃，并非四舍五入！

   [1]常规座屏贴图.            [2]左右镂空贴图.            [3]左右画布贴图.

   [4]上下镂空贴图.            [5]顶天立地贴图.            [6]各种折屏贴图.

   [7]多个座屏贴图.            [8]卷帘座屏贴图.            [9]不扣补切贴图.`
		fmt.Println(text)

		key := inputString("\n:: 请选择上方的边框类型：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			unclassified.MapFrame1() // 小座屏
		case "2":
			fmt.Println("未开发") // 左右镂空
		case "3":
			fmt.Println("未开发") // 中间大两边小
		case "4":
			fmt.Println("未开发") // 上下镂空
		case "5":
			fmt.Println("未开发") // 顶天立地
		case "6":
			unclassified.MapFrame6() // 常规折屏
		case "7":
			unclassified.MapFrame7() // 多座屏
		case "8":
			fmt.Println("未开发") // 补切画布
		case "9":
			fmt.Println("未开发") // 不扣补切
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}
