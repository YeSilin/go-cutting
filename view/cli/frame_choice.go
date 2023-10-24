package cli

// 快捷切图的框架选择

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
)

// 单个座屏选择，备用
func (c *CLI) singleSeatScreenChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Single Seat Screen Choice", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]常规座屏.                  [2]拉布座屏.                  [3]暂未开发`
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
			frame1()                      // 小座屏
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "2":
			frame8to2()                   // 拉布座屏
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的边框类型，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 上下座屏的选择
func (c *CLI) upperAndLowerScreenChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Upper And Lower Screen Choice", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]上下镂空.                  [2]上下画布.                  [3]暂未开发`
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
			frame4to1()                   // 上下镂空
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "2":
			frame4to2()                   // 上下画布
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的边框类型，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 卷帘拉布选择
func (c *CLI) rollerShutterClothChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Roller Shutter Cloth Choice", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]卷帘座屏.                  [2]拉布座屏.                  [3]拉布折屏.`
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
			frame8to1()                   // 卷帘座屏
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "2":
			frame8to2()                   // 拉布座屏
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "3":
			frame8to3()                   // 拉布折屏
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的边框类型，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 旧厂各种框架的选择汇总
func (c *CLI) oldFrameChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Cutting", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]常规座屏.                  [2]左右镂空.                  [3]左右画布.

   [4]上下座屏.                  [5]顶天立地.                  [6]传统折屏.

   [7]多个座屏.                  [8]卷帘拉布.                  [9]补切画布.`
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
			frame1() // 常规座屏
		case "2":
			frame2() // 左右镂空
		case "3":
			frame3() // 左右画布
		case "4":
			c.upperAndLowerScreenChoice() // 上下座屏框架选择
		case "5":
			frame5() // 顶天立地
		case "6":
			frame6() // 常规折屏
		case "7":
			frame7() // 多座屏
		case "8":
			c.rollerShutterClothChoice() // 卷帘拉布框架选择
		case "9":
			frame9() // 补切画布
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的边框类型，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}
