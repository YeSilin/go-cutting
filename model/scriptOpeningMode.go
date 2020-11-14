// 主要实现jsx脚本打开方式
package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"golang.org/x/sys/windows/registry"
)

// 获取PS路径
func getPsPath() (str string, ok bool) {
	str, ok = tools.ReadKeyString(registry.CURRENT_USER, `Software\Classes\adbps\DefaultIcon`, "")
	return
}

// 取消脚本执行警告，这是第一次重构
func psUserConfig(edition int) {
	home, _ := tools.Home()

	var psPath string
	if edition > 2019 {
		// 2020版本开始的ps路径不一样
		psPath = home + fmt.Sprintf("\\AppData\\Roaming\\Adobe\\Adobe Photoshop %d\\Adobe Photoshop %d Settings", edition, edition)
	} else {
		psPath = home + fmt.Sprintf("\\AppData\\Roaming\\Adobe\\Adobe Photoshop CC %d\\Adobe Photoshop CC %d Settings", edition, edition)
	}
	// 警告文件位置
	txtPath := psPath + `\PSUserConfig.txt`

	// 路径是否存在，不存在就返回
	exPath, _ := tools.IsPathExists(psPath)
	if !exPath {
		return
	}

	// 路径是否存在，不存在就提示
	exTxT, _ := tools.IsPathExists(txtPath)
	if !exTxT {
		fmt.Println(fmt.Sprintf("\n:: 监测到你已升级至ps%d，已修复不必要的安全弹窗警告，请重启ps生效！", edition))
	}

	// 先强制生成的文本写覆盖入目标文件
	tools.CreateFile(txtPath, "WarnRunningScripts 0")
}

// jsx文件的打开方式
func OpenMode() {
	// 160的代号是2022版本的ps，如果2022版本的ps存在
	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.160\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.160")
			psUserConfig(2022) // 取消脚本执行警告
			return
		}
	}

	// 150的代号是2021版本的ps，如果2021版本的ps存在
	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.150\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.150")
			psUserConfig(2021) // 取消脚本执行警告
			return
		}
	}

	// 140的代号是2020版本的ps，如果2020版本的ps存在
	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.140\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.140")
			psUserConfig(2020) // 取消脚本执行警告
			return
		}
	}

	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.130\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.130")
			psUserConfig(2019) // 取消脚本执行警告
			return
		}
	}

	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.120\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.120")
			return
		}
	}

	if val, ok := tools.ReadKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.110\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			tools.WriteKeyString(registry.CURRENT_USER, `SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.110")
			return
		}
	}
	fmt.Println("\n:: 未找到ps或超出版本2017~2021支持范围，脚本嵌入可能出现异常！")
}
