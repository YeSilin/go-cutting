// 注册表操作函数
package tools

import (
	"golang.org/x/sys/windows/registry"
	"strings"
)




// 写入key下的字符串，默认字符串使用``写入
func WriteKeyString(k registry.Key, path, name, value string) {
	key, _, _ := registry.CreateKey(k, path, registry.ALL_ACCESS)
	defer key.Close()

	// 写入：字符串 name 留空会写入默认的字符串
	_ = key.SetStringValue(name, value)
}

// 读取key下的字符串，默认字符串使用``读取
func ReadKeyString(k registry.Key, path, name string) (string, bool) {
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