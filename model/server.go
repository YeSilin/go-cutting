package model

import (
	"github.com/webview/webview"
	"log"
	"net/http"
	"os"
)

// web 服务器
func RunWebServer() {
	// 文件服务器 返回html,img,css,js
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./config/static"))))

	http.HandleFunc("/index", indexHandle)
	//http.HandleFunc("/autoNestingPictures", autoNestingPicturesHandle)

	err := http.ListenAndServe(":12110", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 新版本不好用
func RunWebview() {
	// 删除奔溃产生的文件
	os.RemoveAll(`GoCutting.exe.WebView2`)

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("GoCutting")
	w.SetSize(350, 600, webview.HintFixed )
	w.Navigate("http://localhost:12110/index")

	w.Run()
}
