package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/model/quickCipher"
	"github.com/yesilin/go-cutting/tools"
	"github.com/zserge/webview"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// 负责写入到浏览器的函数
func responseWriter(w http.ResponseWriter, file string) {
	f, err := os.Open(file)
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
		if err == io.EOF {
			break
		}
		_, _ = w.Write(buf[:n])
		//fmt.Print(string(buf[:n]))
	}
}


// 首页回调函数
func index(w http.ResponseWriter, r *http.Request) {
	// r:代表跟请求相关的所有内容

	//获取请求的方法
	if method := r.Method; method == "GET" {
		// 返回页面

		responseWriter(w, "Config/html/index.html")
	} else {
		_ = r.ParseForm() // 解析
		// 获取表单中的数据
		//fmt.Printf("%v\n", r.Form)

		// 获取浏览器提交的数据
		cipher1 := r.Form.Get("裁剪快捷键")
		cipher2 := r.Form.Get("重建新文档")
		cipher3 := r.Form.Get("深度清理PSD")
		cipher4 := r.Form.Get("快捷文件夹")
		cipher6 := r.Form.Get("快速清理PSD")
		cipher7 := r.Form.Get("自动加黑边")
		cipher9 := r.Form.Get("到切图历史")
		cipher10 := r.Form.Get("优化版另存")
		cipher98 := r.Form.Get("快速导出图片")
		cipher99 := r.Form.Get("激活win10系统")

		switch {
		case cipher1 == "true":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/SelectTailor.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
			// 每次选择正确的脚本时删除多余备份，最大保留30个
			go tools.DeleteRedundantBackups("Config/JSX/Temp/*", 30)
		case cipher2 == "true":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
			fmt.Println("\n【注意】已重建新文档，参数来自最近一次切图记录！")
			tools.PrintLine(2)
		case cipher3 == "true":
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/JSX/ClearMetadataJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher4 == "true":
			quickCipher.Work() // 工作目录
		case cipher6 == "true":
			// 简单清除元数据
			dataPath := "Config/JSX/ClearMetadataNoPopUpJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher7 == "true":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/BlackEdgeJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher9 == "true":
			//获取当前时间，进行格式化 2006-01-02 15:04:05
			fileName := time.Now().Format("2006-01-02")
			now := time.Now().Format("2006-01")

			// 储存历史记录路径
			path := fmt.Sprintf("Config/History/%s/%s.txt", now, fileName)

			// 先查看是否有历史记录文件
			exists, _ := tools.IsPathExists(path)
			// 如果找不到文件，就创建文件 头
			if !exists {
				exec.Command("cmd.exe", "/c", "start Config\\History").Run()
				return
			}
			// 创建一个协程使用cmd来运行脚本
			cmd := exec.Command("cmd.exe", "/c", "start "+path)
			go cmd.Run()
		case cipher10 == "true":
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/JSX/SaveAsJPEG.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher98 == "true":
			// 创建套图文件夹
			_ = tools.CreateMkdirAll("Config/Picture/主图")
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/SaveForWeb.jsx"
			exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

			time.Sleep(time.Second) // 停一秒
			// 如果存在images就打开
			if ok, _ := tools.IsPathExists("Config/Picture/主图/images"); ok {
				go exec.Command("cmd.exe", "/c", "start Config\\Picture\\主图\\images").Run()
			} else {
				go exec.Command("cmd.exe", "/c", "start Config\\Picture\\主图").Run()
			}
		case cipher99 == "true":
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/W10DigitalActivation.exe /activate"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}
	}
}


func automaticLayout(w http.ResponseWriter, r *http.Request) {
	// 返回页面
	responseWriter(w, "config/html/automaticLayout.html")
}



// web 服务器
func runWebServer() {
	// 文件服务器 返回html,img,css,js
	http.Handle("/html/", http.StripPrefix("/html/",http.FileServer(http.Dir("./config/html"))))
	http.Handle("/img/", http.StripPrefix("/img/",http.FileServer(http.Dir("./config/img"))))
	http.Handle("/css/", http.StripPrefix("/css/",http.FileServer(http.Dir("./config/css"))))
	http.Handle("/js/", http.StripPrefix("/js/",http.FileServer(http.Dir("./config/js"))))


	http.HandleFunc("/index", index)
	http.HandleFunc("/automaticLayout", automaticLayout)

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func GuiMain() {
	// 运行web服务器
	go runWebServer()

	// 搭建web窗口
	_ = webview.Open("GoCutting", "http://localhost:9090/index", 350, 600, true)
}

