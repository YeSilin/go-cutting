package cliui

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/settings"
	"github.com/yesilin/go-cutting/tools"
)

// 当前状态
func current() {
	tools.EnglishTitle("Settings", 74)
	fmt.Println("\n:: 这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

	var memoryStr string
	switch viper.GetBool("memory") {
	case true:
		memoryStr = "已启用"
		memoryStr = tools.ColourString(memoryStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		memoryStr = "已关闭"
		//memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		memoryStr = "参数错误"
		//memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	var openPsStr string
	switch viper.GetBool("openPs") {
	case true:
		openPsStr = "已启用"
		openPsStr = tools.ColourString(openPsStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		openPsStr = "已关闭"
		//openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		openPsStr = "参数错误"
		//openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 自动黑边状态
	var blackEdgeStr string
	switch viper.GetBool("blackEdge") {
	case true:
		blackEdgeStr = "已启用"
		blackEdgeStr = tools.ColourString(blackEdgeStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		blackEdgeStr = "已关闭"
		//blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		blackEdgeStr = "参数错误"
		//blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 自定前缀状态
	var prefixStr string
	if viper.GetString("prefix") != "" {
		prefixStr = "已定义"
		prefixStr = tools.ColourString(prefixStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	} else {
		prefixStr = "未定义"
		//prefixStr = model.ColourString(prefixStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 带颜色的预留画布提示
	var reserveStr = fmt.Sprintf("%.2fcm", viper.GetFloat64("reserve"))
	reserveStr = tools.ColourString(reserveStr, ctc.ForegroundCyan) // 设置带颜色的字符串

	// 自动开启暗号列表
	var cipherListStr string
	switch viper.GetBool("cipherList") {
	case true:
		cipherListStr = "自启动"
		cipherListStr = tools.ColourString(cipherListStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		cipherListStr = "已关闭"
		//cipherListStr = model.ColourString(cipherListStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		cipherListStr = "参数错误"
		//cipherListStr = model.ColourString(cipherListStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 修改套图文件夹位置
	var pictureStr string
	if viper.GetString("picture") != "config\\picture" {
		pictureStr = "已修改"
		pictureStr = tools.ColourString(pictureStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	} else {
		pictureStr = "默认值"
		//pictureStr = model.ColourString(pictureStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 主图自动删除来源
	var automaticDeletionStr string
	switch viper.GetBool("automaticDeletion") {
	case true:
		automaticDeletionStr = "已启用"
		automaticDeletionStr = tools.ColourString(automaticDeletionStr, ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		automaticDeletionStr = "已关闭"
		//automaticDeletionStr = model.ColourString(automaticDeletionStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		automaticDeletionStr = "参数错误"
		//automaticDeletionStr = model.ColourString(automaticDeletionStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	fmt.Printf("\n   [1]记忆框架：%s       [2]自动新建：%s       [3]自动黑边：%s\n", memoryStr, openPsStr, blackEdgeStr)
	fmt.Printf("\n   [4]自定前缀：%s       [5]切布预留：%s       [6]暗号列表：%s\n", prefixStr, reserveStr, cipherListStr)
	fmt.Printf("\n   [7]套图位置：%s       [8]主图自删：%s       [9]全部恢复默认设置\n", pictureStr, automaticDeletionStr)
}

// 选择要修改的配置
func ModifySetting() {
OuterLoop:
	for {
		current() // 显示当前状态

		modify, info := model.InputMenuSelection("\n:: 请选择需要修改的设置：")
		tools.CallClear() // 清屏
		switch modify {
		case "1":
			settings.ModifyMemoryFrame() // 是否记住框架
		case "2":
			settings.ModifyAutomaticCreateDocuments() // 是否自动新建切图文档
		case "3":
			settings.ModifyAutomaticAddBlackEdge() // 是否切图自动添加黑边
		// 最新的切图前缀
		case "4":

		case "5":
			settings.ModifyLatestCanvasReservation() // 最新的切图预留
		case "6":
			settings.ModifyCipherList() // 自动开启暗号列表
		case "7":
			settings.ModifyPicturePath() // 自动套图文件夹路径
		case "8":
			settings.ModifyAutomaticDeletion() // 自动主图时删除来源
		case "9":
			settings.ModifyToDefaultSetting() // 恢复默认设置
		case "-":
			break OuterLoop // 跳到循环结束
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(modify, ctc.ForegroundGreen))
			}
		}
	}
}
