package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
)

// 清除文件内容
func clearFileContents(filePath string) {
	// 写入模式，如果可能，打开时清空文件
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString("去广告勿删")
}

// CleanUp2345Pic 开始净化2345看图王
func CleanUp2345Pic() {
	// 得到安装路径
	path := tools.Get2345PicInstallPath()

	// 根目录
	clearFileContents(path + `\2345PicUpdate.exe`)
	clearFileContents(path + `\Pic_2345Upgrade.exe`)

	// Protect
	clearFileContents(path + `\Protect\2345PicAssistant.exe`)
	clearFileContents(path + `\Protect\2345PicMiniPage.exe`)
	clearFileContents(path + `\Protect\Helper_2345Pic.exe`)
	clearFileContents(path + `\Protect\Pic_2345Svc.exe`)
	clearFileContents(path + `\Protect\ServiceManager.exe`)
	clearFileContents(path + `\Protect\Tool_Uninstall.exe`)

	// tool
	clearFileContents(path + `\tool\Update_2345Pic.exe`)

	fmt.Println("\n:: 一键去广告执行完毕，若失败请先确保执行前2345看图王已完全退出...")
}
