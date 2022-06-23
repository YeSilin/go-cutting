package cli

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
)

// 选择要修改的配置
func (c *CLI) settingsChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Settings", 74)
		fmt.Println("\n:: 这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

		current() // 显示当前状态

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
			modifyPicturePath() // 自动套图文件夹路径
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
