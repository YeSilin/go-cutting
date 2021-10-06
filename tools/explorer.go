package tools

import (
	"golang.org/x/sys/windows/registry"
)

// CleanUpExplorer 清理资源管理器，干掉百度网盘，wps，腾讯微云等一系列花里胡哨的磁盘驱动器旁边的图标
func CleanUpExplorer() {
	// 定义一个变量获取指定路径注册表的值 例如 NameSpace 的
	key, _ := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MyComputer\\NameSpace", registry.ALL_ACCESS)
	defer key.Close()

	// 获取所有子项，0表示获取全部
	keys, _ := key.ReadSubKeyNames(0)
	for _, subKey := range keys {
		// 删除子项
		_ = registry.DeleteKey(key, subKey)
	}

	// 另一个目录的垃圾一起清理到
	key, _ = registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MyComputer\\NameSpace", registry.ALL_ACCESS)
	defer key.Close()

	// 读取：一个项下的所有子项
	keys, _ = key.ReadSubKeyNames(0)
	for _, subKey := range keys {
		// 删除子项
		_ = registry.DeleteKey(key, subKey)
	}
}