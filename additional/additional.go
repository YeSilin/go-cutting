package additional

import (
	"fmt"
	"os/exec"
)

// 运行PSD修复工具
func RunPSDRepairKit() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/extendedSoftware/PSDRepairKit/PSDRepairKit.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具PSD文件修复，请稍后...")
}

// 运行XLS修复工具
func RunAdvancedExcelRepair() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/extendedSoftware/AdvancedExcelRepair/AER.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具XLS文件修复，请稍后...")
}

// 运行QQ微信QQ防撤回工具
func RunRevokeMsgPatcher() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/extendedSoftware/RevokeMsgPatcher.exe"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	fmt.Println("\n:: 正在打开附加工具微信QQ防撤回，请稍后...")
}

// 运行win10数字激活工具
func RunW10DigitalActivation() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/extendedSoftware/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
	fmt.Println("\n:: win10系统已激活，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}

// 注册表导入取得文件所有权
func ImportTakeOwnership() {
	// 创建一个协程使用cmd启动外部程序
	cmd := exec.Command("cmd.exe", "/c", "regedit /s .\\config\\regedit\\takeOwnership.reg")
	go cmd.Run()
	fmt.Println("\n:: 右键菜单已添加，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}

// 注册表导入右键新建文本文档
func ImportNewTextFile() {
	// 创建一个协程使用cmd启动外部程序
	cmd := exec.Command("cmd.exe", "/c", "regedit /s .\\config\\regedit\\newTextFile.reg")
	go cmd.Run()
	fmt.Println("\n:: 右键菜单已添加，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}
