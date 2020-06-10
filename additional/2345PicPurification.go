package additional

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/registry"
	"os"
	"strings"
)

// 读取2345看图王安装路径
func read2345PicInstallationPath() string {
	// 定义一个变量获取指定路径注册表的值 例如 NameSpace 的
	key, _ := registry.OpenKey(registry.CLASSES_ROOT, `2345Pic.jpg\DefaultIcon`, registry.ALL_ACCESS)
	defer key.Close()

	// 读取：字符串
	s, _, _ := key.GetStringValue(``)

	// 去掉尾部 \icon\jpg.ico
	s = strings.TrimSuffix(s, `\icon\jpg.ico`)
	return s
}

// 清理文件
func cleanUpFile(filePath string) {
	// 写入模式，如果可能，打开时清空文件
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString("去广告勿删")
}

// 开始净化2345看图王
func CleanUp2345Pic() {
	// 得到安装路径
	path := read2345PicInstallationPath()

	// 根目录
	cleanUpFile(path + `\2345PicUpdate.exe`)
	cleanUpFile(path + `\Pic_2345Upgrade.exe`)

	// Protect
	cleanUpFile(path + `\Protect\2345PicAssistant.exe`)
	cleanUpFile(path + `\Protect\2345PicMiniPage.exe`)
	cleanUpFile(path + `\Protect\Helper_2345Pic.exe`)
	cleanUpFile(path + `\Protect\Pic_2345Svc.exe`)
	cleanUpFile(path + `\Protect\ServiceManager.exe`)
	cleanUpFile(path + `\Protect\Tool_Uninstall.exe`)

	// tool
	cleanUpFile(path + `\tool\Update_2345Pic.exe`)

	fmt.Println("\n:: 一键去广告执行完毕，若失败请先确保执行前2345看图王已完全退出...")
}
