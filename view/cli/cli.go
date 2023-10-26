package cli

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"strings"
	"time"
)

type CLI struct {
	//定义必要字段
	Version string //版本号 暂时先继续使用浮点数，后期改成字符串
	Notice  string //提示信息 从有道云获取
	info    string // 临时通知
}

// NewCliView 视图的构造方法
func NewCliView() *CLI {
	return &CLI{
		Version: "1.0.0", // 默认版本是 1.0.0
	}
}

// VerifyNetwork 验证网络，没有网络不允许使用
func (c *CLI) VerifyNetwork() {
	// 进行三次网络请求，都失败就当没有网络，并且不让使用
	for i := 0; i < 4; i++ {
		tools.CallClear() // 清屏

		// 有网直接退出循环
		if tools.IsNetwork() {
			color.LightCyan.Println(":: 网络已连接服务器获取使用权限成功，请尽量不关闭软件，避免断网时无法使用！\n")
			break
		}

		time.Sleep(2 * time.Second) // 休眠2秒

		// 第四次获取网络就当没网络
		if i == 3 {
			color.LightCyan.Println(":: 网络已断开无法向服务器请求使用权限，软件将在五秒内自动关闭...\n")
			time.Sleep(5 * time.Second) // 休眠五秒
			os.Exit(1)
		}

		color.LightCyan.Printf(":: 第 %d 次网络连接失败，正在重新向服务器获取使用权限请稍等...\n", i+1)
		time.Sleep(2 * time.Second) // 休眠2秒
	}
}

// License 验证许可证，没有许可证不允许使用
func (c *CLI) License() {
	// 成功直接略过
	if presenter.GetLicense(c.Version) {
		return
	}

	// 公告都获取不到说明没有网络
	//if c.Notice == ":: " && time.Now().Format("2006-01-02") == "2022-12-22" {
	//	//fmt.Println(":: 断网临时授权！~")
	//	c.Notice = ":: 断网临时授权！~"
	//	return
	//}

	// 失败打印云端提示信息
	fmt.Println(":: 当前版本已停用，请前往官方群下载最新版本，以下为最新公告：\n\n" + c.Notice)

	time.Sleep(30 * time.Second) // 休眠
	os.Exit(0)
}

// 内部调用，用来显示通知 是否常驻云端通知
func (c *CLI) showNotice(permanent bool) {
	if c.info != "" {
		fmt.Println(c.info)
		c.info = "" // 打印完就清空
		return
	}

	// 是否常驻云端通知
	if permanent {
		fmt.Println(c.Notice) // 显示有道云通知
		return
	}
}

// MainMenu 显示主菜单
func (c *CLI) MainMenu() {
	for {
		// 先显示通知
		c.showNotice(true)

		color.LightCyan.Println("\n" + (strings.Repeat("-", 18)) + fmt.Sprintf(" Welcome to the GoCutting v%s app ", c.Version) + strings.Repeat("-", 17))
		tips := `
:: 输入数字并回车可选择需要进入的功能，输入暗号【-】并回车可返回上一级菜单！

   [1]快捷切图.        [2]快捷贴图.        [3]快捷效果.        [4]自动套图.

   [5]附加功能.        [6]辅助工具         [7]设置中心.        [8]帮助信息.`
		fmt.Println(tips)

		key := inputString("\n:: 请选择上方的菜单功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		// 如果不是暗号就执行菜单
		switch key {
		case "1":
			c.oldFrameChoice() // 切图
		case "2":
			c.mapFrameChoice() // 贴图
		case "3":
			c.temporaryChoice() // 效果图
		case "4":
			c.autoPictureChoice() // 套图
		case "5":
			c.extendChoice() // 附加
		case "6":
			model.RunGlobalHotkey()
		case "7":
			c.settingsChoice() // 设置
		case "8":
			c.help() // 帮助
		case "-":
			c.info = ":: 已经是最顶层菜单了，无需再返回，输入其他数字试下其他功能吧！"
			continue
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}
