package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
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
	// 先强制关闭进程
	exec.Command("taskkill", "/f", "/t", "/im", "2345PicViewer.exe").Run()
	exec.Command("taskkill", "/f", "/t", "/im", "2345PicTool.exe").Run()
	exec.Command("taskkill", "/f", "/t", "/im", "Pic_2345Svc.exe").Run()
	exec.Command("taskkill", "/f", "/t", "/im", "Update_2345Pic.exe").Run()

	// 卸载服务 这里"/c"，表示执行完关闭窗口
	exec.Command("cmd", "/c", "sc", "delete", "ViewPic_2345Svc").Run()

	// 得到安装路径，使用狸猫换太子净化文件
	path := tools.Get2345PicInstallPath()
	// 根目录
	clearFileContents(path + `\2345PicHomePage.exe`)
	clearFileContents(path + `\2345PicTool.exe`)
	_ = os.Remove(path + `\2345PicUpdate.exe`) // 直接删除，因为会检测文件是否正确 clearFileContents(path + `\2345PicUpdate.exe`)
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
