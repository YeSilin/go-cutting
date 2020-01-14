package additional

import (
	"fmt"

	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"time"
)

// 批量文件改名
func haoZipRename() {
	dataPath := "Config/EXE/HaoZipRename/HaoZipRename.exe"
	exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	// 获取主目录
	haozipPath, _ := tools.Home()
	haozipPath = fmt.Sprintf("%s/AppData/Roaming/HaoZip", haozipPath)
	// 删除文件残留
	time.Sleep(1 * time.Second) // 休眠1秒
	del := os.RemoveAll(haozipPath)
	if del != nil {
		fmt.Println(del)
	}
}

func Additional() {
	model.EnglishTitle("Additional", 79)
	fmt.Println("\n【附加】这里提供一些实用的附加功能，例如常见问题与简单的系统优化一键式解决方案！")

	for {
		fmt.Println("\n【附加】[1]激活win10系统\t[2]微信QQ防撤回\t\t[3]取得文件所有权")
		fmt.Println("\n【附加】[4]净化设备驱动器\t[5]解决黑屏卡死\t\t[6]功能暂未开发")
		fmt.Println("\n【附加】[7]定时18点关机\t\t[8]取消18点关机\t\t[9]PSD文件修复")

		var help = model.Input("\n【附加】请选择需要使用的功能：", false)
		switch help {
		case "1": //激活win10系统
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/W10DigitalActivation.exe /activate"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
			fmt.Println("\n【注意】win10系统已激活，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
			tools.PrintLine(2)

		case "2": //UI3简单测试
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/RevokeMsgPatcher.exe"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
			tools.PrintLine(2)

		case "3": // 取得文件所有权

			// 创建一个协程使用cmd启动外部程序
			cmd := exec.Command("cmd.exe", "/c", "regedit /s .\\Config\\takeOwnership.reg")
			go cmd.Run()
			fmt.Println("\n【注意】右键菜单已添加，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
			tools.PrintLine(2)

		case "4":
			// 净化设备驱动器
			CURRENT_USER_NoNameSpace()
			LOCAL_MACHINE_NoNameSpace()

		case "5":
			// 重启资源管理器
			// 创建一个协程使用cmd启动外部程序
			go func() {
				exec.Command("cmd.exe", "/c", "taskkill /f /im explorer.exe & start explorer.exe").Run()
			}()

		case "6":


		case "7":
			// 定时关机
			s, Unsigned := tools.DistanceIsEighteen()
			if Unsigned { // 无符号就设置
				cmd := fmt.Sprintf("shutdown /s /t %d", s)
				go exec.Command("cmd.exe", "/c", cmd).Run()
				fmt.Printf("\n【注意】定时关机设置成功，%d秒后将自动关机！\n", s)
			} else {
				fmt.Println("\n【注意】定时关机设置失败，已超过18点，此条命令不生效！")
			}
			tools.PrintLine(2)

		case "8":
			// 取消定时关机
			go exec.Command("cmd.exe", "/c", "shutdown /a").Run()
			fmt.Println("\n【注意】定时关机已关闭，18点后不会自动关机！")
			tools.PrintLine(2)

		case "9": //UI3简单测试
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/PSDRepairKit/PSDRepairKit.exe"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case "-":
			goto FLAG //跳出循环

		default:
			continue
		}
	}

FLAG: // 跳到函数结束
}
