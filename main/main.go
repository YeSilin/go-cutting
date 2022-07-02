package main

import (
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/view/cli"
	"github.com/yesilin/go-cutting/view/gui"
	"sync"
)

func init() {
	presenter.InitLog()     // 初始化日志
	presenter.InitSetting() //初始化设置
	presenter.InitFolder()  // 创建必须提前存在的文件夹
	presenter.InitScript()  // 创建必须提前准备的脚本
}

func main() {
	// 实例一个视图结构体
	cliView := cli.NewCliView()
	// 先验证网络
	cliView.VerifyNetwork()

	// 获取通知
	cliView.Notice = ":: " + presenter.Limit.GetString("notice")
	// 设置版本号
	cliView.Version = "1.1.89" //设置版本号！！！！！！！！！！！！！！！！！！！！！！！

	// 验证许可证
	cliView.License()

	// 不让自动退出
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 运行cli主菜单
	go cliView.MainMenu()

	// 运行gui界面
	if viper.GetBool("gui") {
		guiView := gui.NewGuiView()
		guiView.ShowAndRun()
	}

	wg.Wait()
}
