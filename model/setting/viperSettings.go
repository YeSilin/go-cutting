// 这是一个名为毒蛇的配置
package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
	"strings"
)

// 初始化
func init() {
	// 初始化的 Viper实例
	//v := viper.New()

	// 设置默认值
	viper.SetDefault("memory", false)
	viper.SetDefault("openPs", true)
	viper.SetDefault("blackEdge", true)
	viper.SetDefault("prefix", "")
	viper.SetDefault("reserve", 5)

	//  设置配置文件名，不带后缀
	viper.SetConfigName("settings")

	// 第一个搜索路径
	viper.AddConfigPath("./Config/")

	//设置配置文件类型
	viper.SetConfigType("yaml")

	// 安全保存配置文件，如果没有配置文件就保存
	_ = viper.SafeWriteConfig()

	// 搜索路径，并读取配置数据
	err := viper.ReadInConfig()
	if err != nil {
		// 如果读取失败，就保存当前默认配置文件
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("v.WriteConfig err: ", err)
			return
		}
	}

	// 监视配置文件，重新读取配置数据
	viper.WatchConfig()

	// 显示更新信息，不稳
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("【提示】配置文件已更新，来自：", e.Name)
	//})
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

// 当前状态
func current() {
	model.EnglishTitle("Settings", 79)
	fmt.Println("\n【设置】这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

	var memoryStr string
	switch viper.GetBool("memory") {
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
	switch viper.GetBool("openPs") {
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

	// 自动黑边状态
	var blackEdgeStr string
	switch viper.GetBool("blackEdge") {
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

	// 自定前缀状态
	var prefixStr string
	if viper.GetString("prefix") != "" {
		prefixStr = "已定义"
		prefixStr = model.ColourString(prefixStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	} else {
		prefixStr = "未定义"
		prefixStr = model.ColourString(prefixStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 带颜色的预留画布提示
	var reserveStr = fmt.Sprintf("%.2fcm", viper.GetFloat64("reserve"))
	reserveStr = model.ColourString(reserveStr, ctc.ForegroundGreen) // 设置带颜色的字符串

	fmt.Printf("\n【状态】[1]记忆框架：%s\t[2]自动新建：%s\t[3]自动黑边：%s\n", memoryStr, openPsStr, blackEdgeStr)
	fmt.Printf("\n【状态】[4]自定前缀：%s\t[5]切布预留：%s\t[6]恢复默认出厂设置\n", prefixStr, reserveStr)
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
				viper.Set("memory", true)
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 保存最新配置
				_ = viper.WriteConfig()
				continue
			case "2":
				viper.Set("memory", false)
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 保存最新配置
				_ = viper.WriteConfig()
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "2":
			var tempOpenPs = isStringInput("\n【更改】是否自动新建切图文档，[1]是，[2]否：")
			switch tempOpenPs {
			case "1":
				viper.Set("openPs", true)
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 保存最新配置
				_ = viper.WriteConfig()
			case "2":
				viper.Set("openPs", false)
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 保存最新配置
				_ = viper.WriteConfig()
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "3":
			var tempBlackEdge = isStringInput("\n【更改】是否切图自动添加黑边，[1]是，[2]否：")
			switch tempBlackEdge {
			case "1":
				viper.Set("blackEdge", true)
				fmt.Println("\n【提示】设置成功 - 功能已开启！")
				// 保存最新配置
				_ = viper.WriteConfig()
				generate.Tailor("") // 根据配置更新通用裁剪
			case "2":
				viper.Set("blackEdge", false)
				fmt.Println("\n【提示】设置成功 - 功能已关闭！")
				// 保存最新配置
				_ = viper.WriteConfig()
				generate.Tailor("") // 根据配置更新通用裁剪
			case "-":
				fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 36) + "\n")
				goto FLAG // 跳到循环结束
			}
		case "4":
			fmt.Println("\n【提示】自定义前缀可以在使用【-1】暗号时自动添加，例如定义为【沐：】为前缀！")
			fmt.Println("\n此功能未开发，设置无效")
			var tempPrefixStr = model.Input("\n【提示】请输入最新的切图前缀：", false)

			switch tempPrefixStr {
			case "-":
				goto FLAG // 跳到循环结束
			case "0": // 直接回车代表删除前缀
				// 设置前缀
				viper.Set("prefix","")
			default:
				// 设置前缀
				viper.Set("prefix", tempPrefixStr)
				fmt.Printf("\n【提示】切图前缀已更改成 【%s】，输入内容为空代表删除！\n", tempPrefixStr)
				// 保存最新配置
				_ = viper.WriteConfig()
			}

		case "5":
			fmt.Println("\n【警告】修改此项将直接影响最终的切图结果，如未出现特殊情况请勿修改")
			var tempReserve = model.Input("\n【警告】请输入最新的切图预留：", false)

			switch tempReserve {
			case "-":
				goto FLAG // 跳到循环结束
			default:
				// 转成64位浮点数再赋值
				reserve64, _ := strconv.ParseFloat(tempReserve, 64)
				viper.Set("reserve", reserve64)
				fmt.Printf("\n【提示】切图预留已更改成 %.2fcm，一切后果自负\n", reserve64)
				// 保存最新配置
				_ = viper.WriteConfig()
			}
		case "6":
			tools.CallClear() // 清屏
			fmt.Println("\n【提示】已恢复默认设置成功，配置信息已重新加载并生效！")

			// 重置为默认参数
			viper.Set("memory", false)
			viper.Set("openPs", true)
			viper.Set("blackEdge", true)
			viper.Set("prefix", "")
			viper.Set("reserve", 5)

			// 保存最新配置
			_ = viper.WriteConfig()
			generate.Tailor("") // 根据配置更新通用裁剪
		case "-":
			goto FLAG // 跳到循环结束
		default:
			continue
		}
	}
FLAG: //为了跳出for循环
	// 保存最新配置
	_ = viper.WriteConfig()
}
