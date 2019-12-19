package setting

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/globa"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
	"strings"
)

/**初始化*/
func init() {
	// 初始化的时候读取本地配置信息，更新为最新的参数
	var err error
	globa.NowSetting, err = JsonSettingDecode("Config/Config.json")
	if err != nil {
		// 如果解码失败就覆盖一个默认json文件 分别是：记忆框架，自动新建，自动黑边，画布预留
		JsonSettingEncoder("Config/Config.json", globa.DefaultSetting)
		// 随后重新赋予给nowSetting
		globa.NowSetting, _ = JsonSettingDecode("Config/Config.json")
	}
	//fmt.Printf("\n%#v\n", nowSetting)
}

// 当前状态
func current() {
	model.EnglishTitle("Settings", 79)
	fmt.Println("\n【设置】这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

	var memoryStr string
	switch globa.NowSetting.Memory {
	case true:
		memoryStr = "已启用"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		memoryStr = "已关闭"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		memoryStr = "参数错误"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	var openPsStr string
	switch globa.NowSetting.OpenPs {
	case true:
		openPsStr = "已启用"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		openPsStr = "已关闭"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		openPsStr = "参数错误"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	var blackEdgeStr string
	switch globa.NowSetting.BlackEdge {
	case true:
		blackEdgeStr = "已启用"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		blackEdgeStr = "已关闭"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		blackEdgeStr = "参数错误"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 带颜色的预留画布提示
	var reserveStr = fmt.Sprintf("%.2fcm", globa.NowSetting.Reserve)
	reserveStr = model.ColourString(reserveStr, ctc.ForegroundGreen) // 设置带颜色的字符串

	fmt.Printf("\n【状态】[1]记忆框架：%s\t[2]自动新建：%s\t[3]自动黑边：%s\n", memoryStr, openPsStr, blackEdgeStr)
	fmt.Printf("\n【状态】[4]暂位预留：%s\t[5]切布预留：%s\t[6]恢复默认出厂设置\n", model.ColourString("未开发", ctc.ForegroundBright), reserveStr)
}

// 验证输入的内容是不是有效数据
// @param text: 传入用户输入提示信息
// @return: 返回有效的配置信息
func isStringInput(text string) string {
	var receive string
	for {
		fmt.Print(text)
		_, _ = fmt.Scanln(&receive) // 储存用户输入的值

		// 只有这三个字符才能传出，其他字符则一直循环
		if (receive == "1") || (receive == "2") || (receive == "-") {
			return receive
		}

	}
}

/**修改配置*/
func ModifySetting() {
	for {
		current() // 当前状态

		var modify = model.Input("\n【设置】请选择需要修改的设置：", false)
		switch modify {
		case "1":
			var tempMemory = isStringInput("\n【更改】是否记住框架的选择，[1]是，[2]否：")
			switch tempMemory {
			case "1":
				globa.NowSetting.Memory = true
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
				continue
			case "2":
				globa.NowSetting.Memory = false
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "2":
			var tempOpenPs = isStringInput("\n【更改】是否自动新建切图文档，[1]是，[2]否：")
			switch tempOpenPs {
			case "1":
				globa.NowSetting.OpenPs = true
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
			case "2":
				globa.NowSetting.OpenPs = false
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "3":
			var tempBlackEdge = isStringInput("\n【更改】是否切图自动添加黑边，[1]是，[2]否：")
			switch tempBlackEdge {
			case "1":
				globa.NowSetting.BlackEdge = true
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
				generate.Tailor() // 同时生成最新通用裁剪备用
			case "2":
				globa.NowSetting.BlackEdge = false
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
				generate.Tailor() // 同时生成最新通用裁剪备用
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "5":
			fmt.Println("\n【警告】修改此项将直接影响最终的切图结果，如未出现特殊情况请勿修改")
			var tempReserve = model.Input("\n【警告】请输入最新的切图预留：", false)

			switch tempReserve {
			case "-":
				goto FLAG // 跳到循环结束
			default:
				// 转成64位浮点数再赋值
				globa.NowSetting.Reserve, _ = strconv.ParseFloat(tempReserve, 64)
				fmt.Printf("\n【提示】切图预留已更改成 %.2fcm，一切后果自负\n", globa.NowSetting.Reserve)
				// 将当前设置编码到json文件
				JsonSettingEncoder("Config/Config.json", globa.NowSetting)
			}
		case "6":
			tools.CallClear() // 清屏
			fmt.Println("\n【提示】已恢复默认设置成功，配置信息已重新加载并生效！")
			//tools.PrintLine(2)
			// 修改成默认参数
			globa.NowSetting = globa.DefaultSetting
			// 将当前设置编码到json文件
			JsonSettingEncoder("Config/Config.json", globa.NowSetting)
			generate.Tailor() // 同时生成最新通用裁剪备用
		case "-":
			goto FLAG // 跳到循环结束
		default:
			continue
		}
	}
FLAG: //为了跳出for循环
	// 将当前设置编码到json文件
	JsonSettingEncoder("Config/Config.json", globa.NowSetting)
}
