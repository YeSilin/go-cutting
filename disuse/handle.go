package disuse

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/signal"
	"io"
	"net/http"
	"os"
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
	//fmt.Println(r.Method)

	//获取请求的方法
	if method := r.Method; method == "POST" {
		err := r.ParseForm() // 解析
		if err != nil {
			logrus.Error(err)
		}
		// 获取表单中的数据
		//fmt.Printf("%v\n", r.Form)

		// 获取浏览器提交的数据
		cipher := r.Form.Get("cipher")

		switch cipher {
		case "1":
			signal.ExecuteSignal1()
		case "2":
			signal.ExecuteSignal2()
		case "3":
			signal.ExecuteSignal3() // 深度清除源数据
		case "4":
			signal.ExecuteSignal6() // 简单清除元数据
		case "7":
			signal.ExecuteSignal7() // 为当前文档添加黑边
		case "9":
			signal.ExecuteSignal9() // 打开历史记录
		case "10":
			signal.ExecuteSignal10() // 快捷另存为jpg
		case "98":
			signal.ExecuteSignal98()
		}
	} else {
		// 返回页面
		responseWriter(w, "config/static/html/index.html")
	}
}

// 自动嵌套图片的回调函数
func autoNestingPicturesHandle(w http.ResponseWriter, r *http.Request) {
	// 返回页面
	responseWriter(w, "config/static/html/autoNestingPictures.html")
}
