package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/globa"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
	"strings"
)

/**修改配置*/
func ModifySetting() {
	for {
		current() // 显示当前状态

		var modify = model.Input("\n【设置】请选择需要修改的设置：", false)
		switch modify {
		case "1":
			var tempMemory = isStringInput("\n【更改】是否记住框架的选择，[1]是，[2]否：", false)
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
			var tempOpenPs = isStringInput("\n【更改】是否自动新建切图文档，[1]是，[2]否：", false)
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
			var tempBlackEdge = isStringInput("\n【更改】是否切图自动添加黑边，[1]是，[2]否：", false)
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
				viper.Set("prefix", "")
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
		case "7":
			fmt.Printf("\n【提示】正在修改自动套图文件夹路径，当前路径是【%s】\n", viper.GetString("picture"))
			// 套图文件夹位置
			var tempPictureStr = isStringInput("\n【提示】请输入最新的套图文件夹位置：", true)

			switch tempPictureStr {
			case "-":
				goto FLAG // 跳到循环结束
			case "": // 直接回车代表恢复默认路径
				// 设置套图文件夹位置
				viper.Set("picture", "config/picture")
				tools.CallClear()      // 清屏
				fmt.Println("\n【提示】直接回车会恢复默认设置，现已恢复默认设置！")
			default:
				// 设置套图文件夹位置
				viper.Set("picture", tempPictureStr)
				fmt.Printf("\n【提示】套图文件位置已更改成 【%s】，输入内容为空代表删除！\n", tempPictureStr)
				// 保存最新配置
				_ = viper.WriteConfig()
			}

		case "9":
			tools.CallClear() // 清屏
			fmt.Println("\n【提示】已恢复默认设置成功，配置信息已重新加载并生效！")

			// 重置为默认参数
			viper.Set("memory", false)
			viper.Set("openPs", true)
			viper.Set("blackEdge", true)
			viper.Set("prefix", "")
			viper.Set("reserve", 5)
			viper.Set("picture", "config/picture")

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
	err := viper.WriteConfig()
	if err != nil {
		// 记录错误日志
		globa.Logger.Println("viper.WriteConfig err:")
	}
}