package main

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"path/filepath"
	"strings"
)

func main11() {
	originalPath := `E:\Code\Golang\go-cutting\config\Picture`

	// 完成后的目标路径
	resultPath := fmt.Sprintf("%s/主图", originalPath)

	// 获取所有扩展名是txt的文件名，类型是字符串切片
	txtSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.txt", originalPath))
	fmt.Println(txtSlice)
	// 复制所有文本到主图文件夹
	for i := 0; i < len(txtSlice); i++ {
		dst := strings.Replace(txtSlice[i], originalPath, resultPath, 1)
		tools.CopyFile(txtSlice[i], dst)
	}

}
