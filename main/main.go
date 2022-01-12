package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/controller"
	"github.com/yesilin/go-cutting/initialize"
	"github.com/yesilin/go-cutting/logs"
	"github.com/yesilin/go-cutting/tools"
	"github.com/yesilin/go-cutting/view/cli"
	"github.com/yesilin/go-cutting/view/gui"
	"sync"
	"time"
)

func init() {
	logs.InitLog()                // 初始化日志
	controller.InitSetting()      //初始化设置
	cli.VerifyNetwork()           // 没有网络不让使用
	initialize.InitNotification() // ps 未运行就进行通知
	initialize.InitFolder()       // 创建必须提前存在的文件夹
	initialize.InitScript()       // 创建必须提前准备的脚本
	initialize.InitCipherList()   // 判断是否打开暗号列表
}

func main() {
	//实例一个视图结构体
	cliView := cli.NewCliView()
	cliView.Version = 1.001080 //设置版本号！！！！！！！！！！！！！！！！！！！！！！！
	cliView.Expire = 90        //这里改版本最长有效期！！！！！！！！！！！！！！！！！！！！

	// 限制软件使用 2019.7.19
	// 定义私密文件路径
	PrivateFile, _ := tools.Home()
	PrivateFile = fmt.Sprintf("%s\\Documents\\Adobe\\Config.chx", PrivateFile)
	cliView.Power, cliView.Tips = initialize.RestrictingSoftwareUse2(PrivateFile, cliView.Version, tools.GetNtpTime(), cliView.Expire)

	// 如果权限不是true
	if !cliView.Power {
		fmt.Println(cliView.Tips)
		time.Sleep(5 * time.Second) // 休眠五秒
		return
	}

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
