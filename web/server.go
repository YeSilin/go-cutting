package web

import (
	"log"
	"net/http"
)

// web 服务器
func RunWebServer() {
	// 文件服务器 返回html,img,css,js
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./config/web"))))

	http.HandleFunc("/index", indexHandle)
	//http.HandleFunc("/autoNestingPictures", autoNestingPicturesHandle)

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 新版本不好用
//func RunWebview() {
//	w := webview.New(false)
//	defer w.Destroy()
//	w.SetTitle("Minimal webview example")
//	w.SetSize(300, 600, webview.HintNone)
//	w.Navigate("http://localhost:9090/index")
//	w.Run()
//}
