package tools

import (
	"bytes"
	"github.com/go-toast/toast"
	"golang.org/x/sys/windows/registry"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// IsExeRuning : 程序是否运行，strKey:包含此名字的所有程序， strExeName:二次校验路径不确定就留空
// 例如 IsExeRuning("Photoshop.exe", "Adobe")
func IsExeRuning(strKey string, strExeName string) bool {
	buf := bytes.Buffer{}
	cmd := exec.Command("wmic", "process", "get", "name,executablepath")
	cmd.Stdout = &buf
	cmd.Run()

	cmd2 := exec.Command("findstr", strKey)
	cmd2.Stdin = &buf
	data, err := cmd2.CombinedOutput()
	if err != nil && err.Error() != "exit status 1" {
		//XBLog.LogF("ServerMonitor", "IsExeRuning CombinedOutput error, err:%s", err.Error())
		return false
	}

	strData := string(data)
	//fmt.Printf("%#v",strData)
	if strings.Contains(strData, strExeName) {
		return true
	} else {
		return false
	}
}

// WinNotification win 通知
func WinNotification(title, message string) {
	// 转换成绝对路径
	icon, _ := filepath.Abs("config\\static\\img\\logo7-1.png")

	notification := toast.Notification{
		AppID: "GoCutting App",
		Title: title,
		//ActivationArguments: "Start Notepad",
		Message: message,
		Icon:    icon, // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "我知道了", ""},
		},
	}

	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

// Get2345PicInstallPath 获取2345看图王安装路径
func Get2345PicInstallPath() string {
	// 定义一个变量获取指定路径注册表的值 例如 NameSpace 的
	key, _ := registry.OpenKey(registry.CLASSES_ROOT, `2345Pic.jpg\DefaultIcon`, registry.ALL_ACCESS)
	defer key.Close()

	// 读取：字符串
	str, _, _ := key.GetStringValue(``)

	// 去掉尾部 \icon\jpg.ico
	str = strings.TrimSuffix(str, `\icon\jpg.ico`)
	return str
}
