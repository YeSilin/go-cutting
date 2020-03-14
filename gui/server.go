package gui

import (
	"log"
	"net/http"
)

// web 服务器
func RunWebServer() {
	// 文件服务器 返回html,img,css,js
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./config/html"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./config/img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./config/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./config/js"))))

	http.HandleFunc("/index", indexHandle)
	http.HandleFunc("/autoNestingPictures", autoNestingPicturesHandle)

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
