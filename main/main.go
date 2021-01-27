package main

import (
	"fmt"
	"github.com/yesilin/go-cutting/cli"
	"github.com/yesilin/go-cutting/gui"
	"github.com/yesilin/go-cutting/initialize"
	"github.com/yesilin/go-cutting/logs"
	"github.com/yesilin/go-cutting/settings"
	"github.com/yesilin/go-cutting/tools"
	"sync"
	"time"
)

func init() {
	logs.InitLog()                // 初始化日志
	settings.InitSetting()        //初始化设置
	initialize.InitNetwork()      // 没有网络不让使用
	initialize.InitNotification() // ps 未运行就进行通知
	initialize.InitFolder()       // 创建必须提前存在的文件夹
	initialize.InitScript()       // 创建必须提前准备的脚本
	initialize.InitCipherList()   // 判断是否打开暗号列表
}

func main() {
	// 提示信息 剩余使用天数
	var tips string
	// 使用权限
	var power bool

	// 这是版本信息
	const version = 1.001064

	// 限制软件使用 2019.7.19
	// 定义私密文件路径
	PrivateFile, _ := tools.Home()
	PrivateFile = fmt.Sprintf("%s\\Documents\\Adobe\\Config.chx", PrivateFile)
	power, tips = initialize.RestrictingSoftwareUse2(PrivateFile, version, tools.GetNtpTime(), 50) // 这里改版本信息！！！！！！！！！！！！！！！！！！！！
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
	go cli.Start(tips, version)

	// 运行gui
	gui.Start()

	wg.Wait()
}
