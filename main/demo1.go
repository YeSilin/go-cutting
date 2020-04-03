package main

import (
	"github.com/go-toast/toast"
	"log"
)

func main0() {
	notification := toast.Notification{
		//AppID:   "Microsoft.Windows.Shell.RunDialog",
		AppID:   "Example App",

		Title:   "标题",
		Message: "这是消息内容，等等。。。",
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "按钮1", "https://www.google.com/"},
			{"protocol", "按钮2", "https://github.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}