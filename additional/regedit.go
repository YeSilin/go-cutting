package additional

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

// 干掉百度网盘，wps，腾讯微云等一系列花里胡哨的磁盘驱动器旁边的图标
func CURRENT_USER_NoNameSpace() {

	fmt.Print("\n:: CURRENT_USER：")

	// 定义一个变量获取指定路径注册表的值 例如 NameSpace 的
	key, _ := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MyComputer\\NameSpace", registry.ALL_ACCESS)
	defer key.Close()

	// 读取：一个项下的所有子项
	keys, _ := key.ReadSubKeyNames(0)
	for _, key_subkey := range keys {
		// 输出所有子项的名字
		//fmt.Printf("\n【附加】GUID：%s",key_subkey)

		// 定义一个变量获取指定路径注册表的值
		//path := "Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MyComputer\\NameSpace\\" + key_subkey
		//key2,_:=registry.OpenKey(registry.CURRENT_USER, path, registry.ALL_ACCESS)
		key2, _ := registry.OpenKey(key, key_subkey, registry.ALL_ACCESS)
		defer key2.Close()

		// 读取：字符串
		s, _, _ := key2.GetStringValue(``)
		fmt.Printf("%s ", s)

		// 删除：子项
		registry.DeleteKey(key, key_subkey)
	}

	fmt.Println("等恶意盘符！")
}

func LOCAL_MACHINE_NoNameSpace() {
	fmt.Print("\n:: LOCAL_MACHINE：")

	// 定义一个变量获取指定路径注册表的值 例如 NameSpace 的
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Explorer\\MyComputer\\NameSpace", registry.ALL_ACCESS)
	defer key.Close()

	// 读取：一个项下的所有子项
	keys, _ := key.ReadSubKeyNames(0)
	for _, key_subkey := range keys {
		// 输出所有子项的名字
		//fmt.Printf("\n【附加】GUID：%s",key_subkey)

		// 定义一个变量获取指定路径注册表的值
		key2, _ := registry.OpenKey(key, key_subkey, registry.ALL_ACCESS)
		defer key2.Close()

		// 读取：字符串
		s, _, _ := key2.GetStringValue(``)
		fmt.Printf("%s ", s)

		// 删除：子项
		registry.DeleteKey(key, key_subkey)
	}

	fmt.Println("等恶意盘符！")

	fmt.Println("\n:: 已清除恶意盘符，此项附加功能的命令需要右键管理员身份运行本软件方可生效！")
}
