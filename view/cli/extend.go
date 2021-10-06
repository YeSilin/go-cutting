package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/additional"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

// 文件修复工具选择
func (c *CLI) fileRepairToolSelection() {
OuterLoop:
	for {
		tools.EnglishTitle("File repair tool selection", 74)
		text := `
:: 以下为从其他网站收集来的恢复工具，解决一些突发情况，请低调使用请勿传播！

   [1]PSD文件修复              [2]XLS文件修复              [3]刷新文件时间`
		fmt.Println(text)
		c.key, c.info = input.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch c.key {
		case "1":
			additional.RunPSDRepairKit() //PSD文件修复
		case "2":
			additional.RunAdvancedExcelRepair() // 运行XLS修复工具
		case "3":
			// 刷新文件时间
			additional.RefreshFileTime(viper.GetString("picture"), viper.GetString("picture")+"/Result")
		case "-":
			break OuterLoop
		default:
			if len(c.info) != 0 {
				fmt.Println(c.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(c.key, ctc.ForegroundGreen))
			}
		}
	}
}

// 定时关机工具选择
func (c *CLI) timingShutdownToolSelection() {
OuterLoop:
	for {
		tools.EnglishTitle("Timing shutdown tool selection", 74)
		text := `
:: 以下小功能主要用于关闭远程服务器，提前预定关机，避免忘记关闭远程服务器！

   [1]定时十八点关机           [2]取消十八点关机           [3]功能暂未开发`
		fmt.Println(text)
		c.key, c.info = input.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch c.key {
		case "1":
			// 定时关机
			s, Unsigned := tools.DistanceIsEighteen()
			if Unsigned { // 无符号就设置
				cmd := fmt.Sprintf("shutdown /s /t %d", s)
				go exec.Command("cmd.exe", "/c", cmd).Run()
				fmt.Printf("\n:: 定时关机设置成功，%d秒后将自动关机！\n", s)
			} else {
				fmt.Println("\n:: 定时关机设置失败，已超过18点，此条命令不生效！")
			}
		case "2":
			// 取消定时关机
			go exec.Command("cmd.exe", "/c", "shutdown /a").Run()
			fmt.Println("\n:: 定时关机已关闭，18点后不会自动关机！")
		case "3":

		case "-":
			break OuterLoop
		default:
			if len(c.info) != 0 {
				fmt.Println(c.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(c.key, ctc.ForegroundGreen))
			}
		}
	}
}

func (c *CLI) extendChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Extend", 74)
		fmt.Println("\n:: 这里提供一些实用的扩展功能，例如常见问题与简单的系统优化一键式解决方案！")
		tips := `
   [1]激活WIN10系统           [2]微信QQ防撤回           [3]取得文件所有权

   [4]净化设备驱动器          [5]解决黑屏卡死           [6]文件修复工具.

   [7]定时关机工具.           [8]看图王去广告           [9]新建文本修复`
		fmt.Println(tips)

		c.key, c.info = input.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch c.key {
		case "1":
			additional.RunW10DigitalActivation() //激活win10系统
		case "2":
			additional.RunRevokeMsgPatcher() // 微信QQ防撤回
		case "3":
			additional.ImportTakeOwnership() // 取得文件所有权
		case "4":
			// 净化设备驱动器
			tools.CleanUpExplorer()
			fmt.Println("\n:: 已清除恶意盘符，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
		case "5":
			// 重启资源管理器
			// 创建一个协程使用cmd启动外部程序
			go func() {
				exec.Command("cmd.exe", "/c", "taskkill /f /im explorer.exe & start explorer.exe").Run()
			}()
		case "6":
			c.fileRepairToolSelection() // 文件修复工具选择
		case "7":
			c.timingShutdownToolSelection() // 定时关机工具选择
		case "8":
			model.CleanUp2345Pic() // 净化2345看图王
		case "9":
			additional.ImportNewTextFile() // 注册表导入右键新建文本文档
		case "-":
			break OuterLoop //跳出循环
		default:
			if len(c.info) != 0 {
				fmt.Println(c.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(c.key, ctc.ForegroundGreen))
			}
		}
	}
}
