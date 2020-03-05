package main

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"path/filepath"
	"strings"
)



// 获取工作目录中其中一个文件夹的所有文件，返回全部文件与文件夹名
func FilesInFolder(originalPath string) (files []string, folder string) {
	// 获取所有文件
	files, err := filepath.Glob(originalPath)

	if err != nil {
		fmt.Println("filepath.Glob err: ",err)
		return
	}



	// 定义一个原图路径
	var fromPath string
	for i := 0; i < len(files); i++ {
		// 先把所有反斜杠修改成正斜杠
		files[i] = strings.Replace(files[i],"\\", "/", -1)
		// 主图文件夹就忽略
		if files[i] == "config/Picture/主图" {
			continue
		}

		// 不是文件夹也忽略
		if !tools.IsDir(files[i]) {
			continue
		}

		// 只认文件名值最大的一个文件夹
		fromPath = files[i]
		folder = filepath.Base(fromPath) //获取路径中的文件名test.txt
	}

	// 如果一个文件夹也没有
	if fromPath == ""{
		fmt.Println("请先拷贝文件夹到工作目录")
		return
	}

	// 对字符串进行拼接
	fromPath = fmt.Sprintf("%s/*",fromPath)

	// 重新获取所有文件
	files, _ = filepath.Glob(fromPath)
	for i := 0; i < len(files); i++ {
		// 先把所有反斜杠修改成正斜杠
		files[i] = strings.Replace(files[i],"\\", "/", -1)
	}
	return
}


func main11() {
	s,n := FilesInFolder("config/Picture/*.txt")
	fmt.Println(s)
	fmt.Println(n)

}
