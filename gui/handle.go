package gui

import (
	"fmt"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/model/quickCipher"
	"io"
	"net/http"
	"os"
	"os/exec"
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
func indexHandle(w http.ResponseWriter, r *http.Request) {
	// r:代表跟请求相关的所有内容
	//获取请求的方法
	if method := r.Method; method == "POST" {
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
			model.StartCode1()
		case cipher2 == "true":
			model.StartCode2()
		case cipher3 == "true":
			model.StartCode3() // 深度清除源数据
		case cipher4 == "true":
			quickCipher.Work() // 工作目录
		case cipher6 == "true":
			model.StartCode6() // 简单清除元数据
		case cipher7 == "true":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/BlackEdgeJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher9 == "true":
			model.StartCode9() // 打开历史记录
		case cipher10 == "true":
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/JSX/SaveAsJPEG.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case cipher98 == "true":
			model.StartCode98()
		case cipher99 == "true":
			model.StartCode99()
		}

	} else {
		// 返回页面
		responseWriter(w, "config/html/index.html")
	}
}

// 自动嵌套图片的回调函数
func autoNestingPicturesHandle(w http.ResponseWriter, r *http.Request) {
	// 返回页面
	responseWriter(w, "config/html/autoNestingPictures.html")
}
