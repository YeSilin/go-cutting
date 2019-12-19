package model

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)


// 监听暗号 -1
func NegativeOne() {

	for  {
		eve := robotgo.AddEvents("`", "ctrl")

		if eve {
			// fmt.Println("已按下快捷键")
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/SelectTailor.jsx"
			exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
fmt.Println("-------------------")
			// 每次选择正确的脚本时删除多余备份，最大保留30个
			go tools.DeleteRedundantBackups("Config/JSX/Temp/*",30)
		}
		//robotgo.StopEvent()
		// 休眠1秒，太快会奔溃
		time.Sleep(1 * time.Second)
	}
}







