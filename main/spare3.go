package main

import (
	"fmt"
	"github.com/zserge/webview"
	"io"
	"log"
	"net/http"
	"os"
)

// 回调函数
func handler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()  // 解析参数，默认是不会解析的
	//fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	//fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的

	// 返回设置页面
	f, err := os.Open("Config/HTML/Settings.html")
	if err != nil {
		fmt.Println("os.Open,err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("f.Read,err", err)
			return
		}
		if err == io.EOF{
			break
		}
		_, _ = w.Write(buf[:n])
		//fmt.Print(string(buf[:n]))
	}
}

// web 服务器
func runWebServer() {
	http.HandleFunc("/", handler)            // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main333() {
	// 运行web服务器
	go runWebServer()

	// 搭建web窗口
	_ = webview.Open("GoCutting", "http://localhost:9090", 400, 600, true)

}
