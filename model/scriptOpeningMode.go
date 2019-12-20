package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"golang.org/x/sys/windows/registry"
	"strings"
)

// 主要实现jsx脚本打开方式

// 写入key下的字符串，默认字符串使用``写入
func writeKeyString(k registry.Key, path, name, value string) {
	key, _, _ := registry.CreateKey(k, path, registry.ALL_ACCESS)
	defer key.Close()

	// 写入：字符串 name 留空会写入默认的字符串
	_ = key.SetStringValue(name, value)
}

// 读取key下的字符串，默认字符串使用``读取
func readKeyString(k registry.Key, path, name string) (string, bool) {
	// 打开一个key，路径是 Photoshop.PSDTFile.140\DefaultIcon，权限是只读
	key, err := registry.OpenKey(k, path, registry.READ)
	if err != nil {
		return "", false
	}
	defer key.Close()

	// 读取指定字符串名字的值
	s, _, _ := key.GetStringValue(name)

	// 将读到的值按切割成字符串切片 D:\Program Files (x86)\Adobe\Adobe Photoshop 2020\Photoshop.exe,58
	slice := strings.Split(s, ",")
	return slice[0], true
}

// 取消脚本执行警告，这是第一次重构
func psUserConfig(edition int) {
	home, _ := tools.Home()

	var psPath string
	if edition > 2019{
		// 2020版本的ps路径不一样
		psPath = home + fmt.Sprintf("\\AppData\\Roaming\\Adobe\\Adobe Photoshop %d\\Adobe Photoshop %d Settings",edition,edition)
	}else {
		psPath = home + fmt.Sprintf("\\AppData\\Roaming\\Adobe\\Adobe Photoshop CC %d\\Adobe Photoshop CC %d Settings",edition,edition)
	}
	// 警告文件位置
	txtPath := psPath+`\PSUserConfig.txt`

	// 路径是否存在，不存在就返回
	exPath, _ := tools.IsPathExists(psPath)
	if !exPath {
		return
	}

	// 路径是否存在，不存在就提示
	exTxT, _ := tools.IsPathExists(txtPath)
	if !exTxT {
		fmt.Println(fmt.Sprintf("【警告】监测到你已升级至ps%d，已修复不必要的安全弹窗警告，请重启ps生效！\n",edition))
	}

	// 先强制生成的文本写覆盖入目标文件
	tools.CreateFile(txtPath, "WarnRunningScripts 0")
}

// jsx文件的打开方式
func OpenMode()  {
	// 140的代号是2020版本的ps，如果2020版本的ps存在
	if val, ok := readKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.140\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			writeKeyString(registry.CURRENT_USER,`SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.140")
			psUserConfig(2020) // 取消脚本执行警告
			return
		}
	}

	if val, ok := readKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.130\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			writeKeyString(registry.CURRENT_USER,`SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.130")
			psUserConfig(2019) // 取消脚本执行警告
			return
		}
	}

	if val, ok := readKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.120\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			writeKeyString(registry.CURRENT_USER,`SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.120")
			return
		}
	}

	if val, ok := readKeyString(registry.CLASSES_ROOT, `Photoshop.PSDTFile.110\DefaultIcon`, ``); ok {
		// 单判断注册表不够，同时再判断文件exe文件是否存在
		ok, _ := tools.IsPathExists(val)
		if ok {
			// 写入注册表
			writeKeyString(registry.CURRENT_USER,`SOFTWARE\Classes\.jsx`, ``, "Photoshop.PSDTFile.110")
			return
		}
	}
	fmt.Println("\n【错误】未找到ps或ps版本超出支持范围2017~2020版本，脚本嵌入可能出现异常！")
}






