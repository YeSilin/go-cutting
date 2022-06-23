package presenter

import "github.com/yesilin/go-cutting/tools"

//IsPhotoshopRun 判断Photoshop是否已打开运行
func IsPhotoshopRun() bool {
	return tools.IsExeRuning("Photoshop.exe", "Adobe")
}
