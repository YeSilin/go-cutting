// 修改所有设置

package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
	"strings"
)

// 验证输入的内容是不是有效数据
// @param text: 传入用户输入提示信息
// @return: 返回有效的配置信息
func isStringInput(text string, isPath bool) string {
	var receive string
	for {
		fmt.Print(text)
		_, _ = fmt.Scanln(&receive) // 储存用户输入的值

		// 如果本次是为了判断路径
		if isPath {
			// 删除首尾连续的的空白字符。
			receive = strings.TrimSpace(receive)
			// 把所有反斜杠修改成正斜杠
			//receive = strings.Replace(receive, "\\", "/", -1)
			tools.CallClear() // 清屏
			return receive
		}

		// 只有这三个字符才能传出，其他字符则一直循环
		if (receive == "1") || (receive == "2") || (receive == "-") {
			tools.CallClear() // 清屏
			return receive
		}

	}
}

// 修改是否记忆框架
func modifyMemoryFrame() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Memory Frame", 79)
	fmt.Println("\n【提示】在快捷切图中是否记住框架的选择，如选择普通座屏就会一直停留在普通座屏")
	var tempMemory = isStringInput("\n【更改】是否记住框架的选择，[1]是，[2]否：", false)
	switch tempMemory {
	case "1":
		viper.Set("memory", true)
		fmt.Println("\n【提示】设置成功 - 记住框架的选择已开启！")
	case "2":
		viper.Set("memory", false)
		fmt.Println("\n【提示】设置成功 - 记住框架的选择已关闭！")

	case "-":
		//fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
		return
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改是否自动新建切图文档
func modifyAutomaticCreateDocuments() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Automatic Create Documents", 79)
	fmt.Println("\n【提示】在快捷切图中是否在输入好尺寸信息后自动调用 PS 新建此文档")

	var tempOpenPs = isStringInput("\n【更改】是否自动新建切图文档，[1]是，[2]否：", false)
	switch tempOpenPs {
	case "1":
		viper.Set("openPs", true)
		fmt.Println("\n【提示】设置成功 - 自动新建切图文档已开启！")
	case "2":
		viper.Set("openPs", false)
		fmt.Println("\n【提示】设置成功 - 自动新建切图文档已关闭！")
	case "-":
		//fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
		return
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改是否自动添加黑边
func modifyAutomaticAddBlackEdge() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Automatic Add Black Edge", 79)
	fmt.Println("\n【提示】在快捷切图中是否在使用【-1】保存时自动为当前文档添加黑边")

	var tempBlackEdge = isStringInput("\n【更改】是否切图自动添加黑边，[1]是，[2]否：", false)
	switch tempBlackEdge {
	case "1":
		viper.Set("blackEdge", true)
		fmt.Println("\n【提示】设置成功 - 切图自动添加黑边已开启！")
		generate.Tailor("") // 根据配置更新通用裁剪
	case "2":
		viper.Set("blackEdge", false)
		fmt.Println("\n【提示】设置成功 - 切图自动添加黑边已关闭！")
		generate.Tailor("") // 根据配置更新通用裁剪
	case "-":
		//fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
		return
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改最新的切图预留
func modifyLatestCanvasReservation() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Latest Canvas Reservation", 79)
	fmt.Println("\n【警告】修改此项将直接影响最终的切图结果，如未出现特殊情况请勿修改")
	var tempReserve = model.Input("\n【警告】请输入最新的切图预留：", false)

	switch tempReserve {
	case "-":
		return
	default:
		// 转成64位浮点数再赋值
		reserve64, _ := strconv.ParseFloat(tempReserve, 64)
		viper.Set("reserve", reserve64)
		fmt.Printf("\n【提示】切图预留已更改成 %.2fcm，一切后果自负\n", reserve64)
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改是否自动打开暗号列表
func modifyCipherList() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Cipher List", 79)
	fmt.Println("\n【提示】在每次打开切图软件时，是否同时打开暗号列表的 UI 操作界面")
	var tempMemory = isStringInput("\n【更改】是否自启动暗号列表，[1]是，[2]否：", false)
	switch tempMemory {
	case "1":
		viper.Set("cipherList", true)
		fmt.Println("\n【提示】设置成功 - 同时打开暗号列表已开启！")
	case "2":
		viper.Set("cipherList", false)
		fmt.Println("\n【提示】设置成功 - 同时打开暗号列表已关闭！")
	case "-":
		return
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改自动套图路径
func modifyPicturePath() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Picture Path", 79)
	fmt.Printf("\n【提示】正在修改自动套图文件夹路径，当前路径是【%s】\n", viper.GetString("picture"))
	// 套图文件夹位置
	var tempPictureStr = isStringInput("\n【提示】请输入最新的套图文件夹位置：", true)
	switch tempPictureStr {
	case "-":
		return
	// 直接回车代表恢复默认路径
	case "":
		// 设置套图文件夹位置
		viper.Set("picture", "config\\picture")
		tools.CallClear() // 清屏
		fmt.Println("\n【提示】直接回车会恢复默认设置，现已恢复默认设置！")
	default:
		tools.CallClear() // 清屏

		// 判断路径是否存在
		if ok, _ := tools.IsPathExists(tempPictureStr); !ok {
			tools.CreateMkdirAll(tempPictureStr)
			fmt.Println("\n【提示】输入的文件夹未创建，已成功创建该文件夹！")
		}
		// 设置套图文件夹位置
		viper.Set("picture", tempPictureStr)

		fmt.Printf("\n【提示】套图文件位置已更改成 【%s】，输入内容为空代表删除！\n", tempPictureStr)
	}

	// 生成详情页指定保存位置
	go func() {
		generate.SaveForWeb(viper.GetString("picture"))
		// 保存最新配置
		viper.WriteConfig()
	}()
}

// 自动主图时删除来源
func modifyAutomaticDeletion() {
	tools.CallClear() // 清屏
	model.EnglishTitle("Modify Automatic Deletion", 79)
	fmt.Println("\n【提示】在自动套图中是否在使用一键主图的时候，自动删除转换之前的源文件")

	var tempAutomaticDeletion = isStringInput("\n【更改】是否自动删除源文件，[1]是，[2]否：", false)
	switch tempAutomaticDeletion {
	case "1":
		viper.Set("automaticDeletion", true)
		fmt.Println("\n【提示】设置成功 - 转换后将自动删除源文件！")
		generate.Tailor("") // 根据配置更新通用裁剪
	case "2":
		viper.Set("automaticDeletion", false)
		fmt.Println("\n【提示】设置成功 - 转换后将不会删除源文件！")
		generate.Tailor("") // 根据配置更新通用裁剪
	case "-":
		//fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
		return
	}
	// 保存最新配置
	go viper.WriteConfig()
}

// 修改为默认设置
func modifyToDefaultSetting() {
	tools.CallClear() // 清屏
	fmt.Println("\n【提示】已恢复默认设置成功，配置信息已重新加载并生效！")

	go func() {
		// 重置为默认参数
		viper.Set("memory", false)
		viper.Set("openPs", true)
		viper.Set("blackEdge", true)
		viper.Set("prefix", "")
		viper.Set("reserve", 5)
		viper.Set("picture", "config\\picture")
		viper.Set("automaticDeletion", false)

		// 保存最新配置
		_ = viper.WriteConfig()
		generate.Tailor("") // 根据配置更新通用裁剪
	}()
}

// 选择要修改的配置
func ModifySetting() {
	for {
		current() // 显示当前状态

		var modify = model.Input("\n【设置】请选择需要修改的设置：", false)
		switch modify {
		case "1":
			modifyMemoryFrame() // 是否记住框架
		case "2":
			modifyAutomaticCreateDocuments() // 是否自动新建切图文档
		case "3":
			modifyAutomaticAddBlackEdge() // 是否切图自动添加黑边
		// 最新的切图前缀
		case "4":
			fmt.Println("\n【提示】自定义前缀可以在使用【-1】暗号时自动添加，例如定义为【沐：】为前缀！")
			fmt.Println("\n此功能未开发，设置无效")
			var tempPrefixStr = model.Input("\n【提示】请输入最新的切图前缀：", false)

			switch tempPrefixStr {
			case "-":
				goto FLAG // 跳到循环结束
			case "0": // 直接回车代表删除前缀
				// 设置前缀
				viper.Set("prefix", "")
			default:
				// 设置前缀
				viper.Set("prefix", tempPrefixStr)
				fmt.Printf("\n【提示】切图前缀已更改成 【%s】，输入内容为空代表删除！\n", tempPrefixStr)
				// 保存最新配置
				_ = viper.WriteConfig()
			}
		case "5":
			modifyLatestCanvasReservation() // 最新的切图预留
		case "6":
			modifyCipherList() // 自动开启暗号列表
		case "7":
			modifyPicturePath() // 自动套图文件夹路径
		case "8":
			modifyAutomaticDeletion() // 自动主图时删除来源
		case "9":
			modifyToDefaultSetting() // 恢复默认设置
		case "-":
			goto FLAG // 跳到循环结束
		default:
			continue
		}
	}
FLAG: //为了跳出for循环
}
