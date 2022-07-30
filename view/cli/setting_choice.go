package cli

import (
	"fmt"
	"github.com/ncruces/zenity"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
)

// 获取设置的当前状态
func currentState() {
	// 记忆框架
	var memoryStr string
	switch viper.GetBool("memory") {
	case true:
		memoryStr = tools.ColourString("已启用", ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		memoryStr = "已关闭"
	default:
		memoryStr = "参数错误"
	}

	// 自动新建
	var openPsStr string
	switch viper.GetBool("openPs") {
	case true:
		openPsStr = tools.ColourString("已启用", ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		openPsStr = "已关闭"
	default:
		openPsStr = "参数错误"
	}

	// 自动黑边
	var blackEdgeStr string
	switch viper.GetBool("blackEdge") {
	case true:
		blackEdgeStr = tools.ColourString("已启用", ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		blackEdgeStr = "已关闭"
	default:
		blackEdgeStr = "参数错误"
	}

	// 自定前缀
	var prefixStr string
	if viper.GetString("prefix") != "" {
		prefixStr = tools.ColourString("已定义", ctc.ForegroundCyan) // 设置带颜色的字符串
	} else {
		prefixStr = "未定义"
	}

	// 画布预留
	var reserveStr = fmt.Sprintf("%.2fcm", viper.GetFloat64("reserve"))
	reserveStr = tools.ColourString(reserveStr, ctc.ForegroundCyan) // 设置带颜色的字符串

	// 自启管理
	var selfStartingManagementStr string
	switch viper.GetBool("gui") {
	case true:
		selfStartingManagementStr = tools.ColourString("已自启", ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		selfStartingManagementStr = "已关闭"
	default:
		selfStartingManagementStr = "参数错误"
	}

	// 修改套图文件夹位置
	var pictureStr string
	if viper.GetString("picture") != "resources\\picture" {
		pictureStr = tools.ColourString("已修改", ctc.ForegroundCyan) // 设置带颜色的字符串
	} else {
		pictureStr = "默认值"
	}

	// 主图自动删除来源
	var automaticDeletionStr string
	switch viper.GetBool("automaticDeletion") {
	case true:
		automaticDeletionStr = tools.ColourString("已启用", ctc.ForegroundCyan) // 设置带颜色的字符串
	case false:
		automaticDeletionStr = "已关闭"
	default:
		automaticDeletionStr = "参数错误"
	}

	fmt.Printf("\n   [1]记忆框架：%s       [2]自动新建：%s       [3]自动黑边：%s\n", memoryStr, openPsStr, blackEdgeStr)
	fmt.Printf("\n   [4]自定前缀：%s       [5]切布预留：%s       [6]自启管理：%s\n", prefixStr, reserveStr, selfStartingManagementStr)
	fmt.Printf("\n   [7]套图位置：%s       [8]主图自删：%s       [9]全部恢复默认设置\n", pictureStr, automaticDeletionStr)
}

// 修改自动套图路径
func (c *CLI) modifyPicturePathChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)
		tools.EnglishTitle("Modify Picture Path", 74)
		fmt.Printf("\n::自动套图只读取指定路径下的文件，当前工作路径是[%s]\n", tools.ColourString(viper.GetString("picture"), ctc.ForegroundCyan))
		const prompt = `
   [1]设置新的路径.            [2]恢复默认路径.            [3]功能暂未开发.`
		fmt.Println(prompt)

		key := inputString("\n:: 请选择需要修改的设置：")
		tools.CallClear() // 清屏

		switch key {

		case "1":
			path, _ := zenity.SelectFile(
				zenity.Filename("E:\\淘宝美工\\套图汇总"),
				zenity.Directory())
			if path != "" {
				// 设置套图文件夹位置
				viper.Set("picture", path)
				c.info = fmt.Sprintf(":: 套图文件位置已更改成 [%s]", tools.ColourString(path, ctc.ForegroundGreen))
			} else {
				c.info = ":: 未选择正确的目录位置，因此并未修改..."
			}

		case "2":
			// 设置套图文件夹位置
			viper.Set("picture", "resources\\picture")
			c.info = ":: 已恢复默认路径，位置在主程序的 [resources\\picture] 目录下..."
		case "-":
			break OuterLoop // 跳到循环结束
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的设置选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))

		}

	}
	// 生成详情页指定保存位置
	go func() {
		model.SaveForWeb(viper.GetString("picture"))
		// 保存最新配置
		viper.WriteConfig()
	}()
}

// 选择要修改的配置
func (c *CLI) settingsChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Settings", 74)
		fmt.Println("\n:: 这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

		currentState() // 显示当前状态

		key := inputString("\n:: 请选择需要修改的设置：")
		tools.CallClear() // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			modifyMemoryFrame() // 是否记住框架
		case "2":
			modifyAutomaticCreateDocuments() // 是否自动新建切图文档
		case "3":
			modifyAutomaticAddBlackEdge() // 是否切图自动添加黑边
		// 最新的切图前缀
		case "4":

		case "5":
			modifyLatestCanvasReservation() // 最新的切图预留
		case "6":
			settingsSelfStartingManagement() // 自动开启暗号列表
		case "7":
			c.modifyPicturePathChoice() // 自动套图文件夹路径
		case "8":
			modifyAutomaticDeletion() // 自动主图时删除来源
		case "9":
			modifyToDefaultSetting() // 恢复默认设置
		case "-":
			break OuterLoop // 跳到循环结束
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的设置选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}
