package model

// 2345看图王净化
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
	exec.Command("taskkill", "/f", "/t", "/im", "2345PicWorker.exe").Run()

	// 卸载服务 这里"/c"，表示执行完关闭窗口
	exec.Command("cmd", "/c", "sc", "delete", "ViewPic_2345Svc").Run() // 卸载服务 服务名称好像变了 先留着
	exec.Command("cmd", "/c", "sc", "stop", "2345Pic").Run()           // 先关闭后卸载
	exec.Command("cmd", "/c", "sc", "delete", "2345Pic").Run()         // 文件路径 protect\PicService.exe

	// 得到安装路径，使用狸猫换太子净化文件
	path := tools.Get2345PicInstallPath()

	// 根目录
	clearFileContents(path + `\2345PicHomePage.exe`)
	clearFileContents(path + `\2345PicTool.exe`)
	clearFileContents(path + `\2345PicWorker.exe`)
	_ = os.Remove(path + `\2345PicUpdate.exe`) // 直接删除，因为会检测文件是否正确 clearFileContents(path + `\2345PicUpdate.exe`)
	clearFileContents(path + `\2345PicDumper.exe`)
	clearFileContents(path + `\2345PicFeedback.exe`)
	clearFileContents(path + `\2345PicHelper.exe`)
	clearFileContents(path + `\2345PicLoader.exe`)

	// protect
	clearFileContents(path + `\protect\PicService.exe`)

	// 根据环境变量找到2345偷偷备份的文件并全部删掉
	_ = os.RemoveAll(os.Getenv("AppData") + `\2345Pic\`) // 不能全部删 待优化

	fmt.Println("\n:: 一键去广告执行完毕，若失败请先确保执行前2345看图王已完全退出...")
}
