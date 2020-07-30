package main

import (
	"fmt"
	"github.com/yesilin/go-cutting/cliui"
	"github.com/yesilin/go-cutting/gui"
	"github.com/yesilin/go-cutting/logs"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/settings"
	"github.com/yesilin/go-cutting/tools"
	"sync"
	"time"
)

func init() {
	logs.InitLog()           // 初始化日志
	settings.Init()          //初始化设置
	model.InitNetwork()      // 没有网络不让使用
	model.InitNotification() // ps 未运行就进行通知
	model.InitFolder()       // 创建必须提前存在的文件夹
	model.InitScript()       // 创建必须提前准备的脚本
	go model.RunWebServer()  // 必须提前运行web服务器
	model.InitCipherList()   // 判断是否打开暗号列表

	// 实现快捷键 -1
	//go model.NegativeOne()
}

func main() {
	// 提示信息 剩余使用天数
	var tips string
	// 使用权限
	var power bool

	// 这是版本信息
	const version = 1.001047

	// 限制软件使用 2019.7.19
	// 定义私密文件路径
	PrivateFile, _ := tools.Home()
	PrivateFile = fmt.Sprintf("%s\\Documents\\Adobe\\Config.chx", PrivateFile)
	power, tips = model.RestrictingSoftwareUse2(PrivateFile, version, tools.GetNtpTime(), 30) // 这里改版本信息！！！！！！！！！！！！！！！！！！！！
	// 如果权限不是true
	if !power {
		fmt.Println(tips)
		time.Sleep(5 * time.Second) // 休眠五秒
		return
	}

	// 不让自动退出
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 运行主体
	go cliui.Start(tips, version)

	// 运行gui
	gui.Start()

	wg.Wait()
}
