package additional

import (
	"fmt"
	"github.com/briandowns/spinner"
	"os/exec"
	"time"
)

// 清理系统垃圾
func CleanUpGarbage()  {
	fmt.Println()
	str := []string{"[>>>                >]", "[]>>>>              []", "[]  >>>>            []", "[]    >>>>          []", "[]      >>>>        []", "[]        >>>>      []",  "[]            >>>>  []", "[]              >>>>[]", "[>                >>>]"}

	// 新建一个微调器
	//s:=spinner.New(spinner.CharSets[43],100*time.Millisecond)
	s:=spinner.New(str,100*time.Millisecond)

	// 为微调器添加前缀
	s.Prefix = "【附加】正在安全清理: "


	// 为微调器添加后缀
	//s.Suffix="00"
	// 执行完成后的提示
	s.FinalMSG="【附加】系统垃圾已清理完毕！                                       "

	// 设置颜色
	s.Color("green")

	go func(){
		// 开始
		s.Start()

		// 需要执行的代码
		cmd := "del /f /s /q %systemdrive%\\*.tmp"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %systemdrive%\\*._mp"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %systemdrive%\\*.gid"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %systemdrive%\\*.chk"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %systemdrive%\\*.old"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %systemdrive%\\recycled\\*.*"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /a /q %systemdrive%\\*.sqm"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %windir%\\*.bak"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %windir%\\temp\\*.*"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %windir%\\SoftwareDistribution\\Download\\*.*"
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q %userprofile%\\cookies\\*.* "
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q \"%userprofile%\\Local Settings\\Temporary Internet Files\\*.*\""
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q \"%userprofile%\\Local Settings\\Temp\\*.*\""
		exec.Command("cmd.exe", "/c", cmd).Run()

		cmd = "del /f /s /q \"%userprofile%\\recent\\*.*\""
		exec.Command("cmd.exe", "/c", cmd).Run()

		// 结束
		s.Stop()
		fmt.Println()
	}()

}

