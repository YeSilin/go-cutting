package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 创建文件并写入数据的函数
func CreateFile(path, data string) {
	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(path)
	if err != nil { // 如果有错误，打印错误，同时返回
		fmt.Println("创建文件错误 =", err)
		return
	}
	defer f.Close() // 在退出整个函数时，关闭文件

	// 第一个参数返回数据长度，第二个，错误信息
	_, err1 := f.WriteString(data)
	if err1 != nil { // 如果有错误，打印错误
		fmt.Println("写入文件错误 =", err1)
	}
}

// 追加写入数据
func WriteFile(path, data string) {
	// 打开文件,追加模式，读权限与写权限
	f, err := os.OpenFile(path, os.O_APPEND, 6)
	if err != nil {
		fmt.Println("打开文件错误 =", err)
		return
	}

	// 写入数据
	_, err1 := f.WriteString(data)
	if err1 != nil {
		fmt.Println("写入文件错误 =", err1)
	}
	// 在退出整个函数时，关闭文件
	defer f.Close()
}

// 读取文件
func ReadFile(filePath string) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件错误 =", err)
	}

	// 关闭文件
	defer f.Close()

	// 定义字节切片大小，2K大小
	buf := make([]byte, 1024*2)
	// n代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { // 文件出错，同时没有到结尾
		fmt.Println("读取文件错误 =", err1)
	}

	fmt.Println("buf = ", string(buf[:n]))
}

// 用行读取文件
func ReadLine(name, filePath string) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("行读取打开文件错误 =", err)
	}
	defer f.Close() // 最后关闭文件

	// 创建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	// 循环读取文件中的内容，直到文件的末尾
	for {
		// 遇到‘\n’结束读取，但是‘\n’也读取进入
		buf, err := r.ReadBytes('\n')
		if err == io.EOF { // 文件读取错误
			break
		}

		// 先将字节转成字符串，然后再使用字符串替换的方法
		bufs := strings.ReplaceAll(string(buf), "\\", "\\\\")
		bufs = strings.ReplaceAll(bufs, "\"", "\\"+"\"")

		bufs = fmt.Sprintf("%s.WriteString(\"%s", name, bufs)
		// 在末尾插入右括号，左移两位是因为结尾有\n
		index := len(bufs) - 2
		bufs = bufs[:index] + "\\n\")" + bufs[index:]

		fmt.Printf("%s", bufs)
	}
}

// 用go格式，行读取文件(目前这个函数用来生成js脚本字符串)
func FormatReadLine(name, filePath string) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("行读取打开文件错误 =", err)
	}
	defer f.Close() // 最后关闭文件

	// 创建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	// 循环读取文件中的内容，直到文件的末尾
	for {
		// 遇到‘\n’结束读取，但是‘\n’也读取进入
		buf, err := r.ReadBytes('\n')

		// 先将字节转成字符串，然后再使用字符串替换的方法
		bufStr := fmt.Sprintf("%q",string(buf))

		// 添加前后缀
		bufStr = fmt.Sprintf("%s.WriteString(%s)", name, bufStr)

		// 打印出来
		fmt.Println(bufStr)

		if  err == io.EOF { // 文件读取结束
			break
		}
	}
}


// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建单个文件夹
func CreateMkdir(filePath string) {
	exist, err := IsPathExists(filePath)
	if err != nil {
		//fmt.Printf("get dir error![%v]\n", err)
		// 创建失败，因为判断是否有该文件夹时发生了错误
		return
	}

	// 判断文件夹是否已存在
	if exist {
		//fmt.Printf("has dir![%v]\n", filePath)
		// 创建失败，因为该文件夹已存在
	} else {
		// fmt.Printf("no dir![%v]\n", filePath)
		// 创建文件夹
		err := os.Mkdir(filePath, os.ModePerm)
		if err != nil {
			//fmt.Printf("mkdir failed![%v]\n", err)
			// 创建失败，因为创建文件夹时发生了错误
		} else {
			// fmt.Printf("mkdir success!\n")
			// 创建成功
		}
	}
}

// 调用os.MkdirAll递归创建文件夹
func CreateMkdirAll(filePath string) error {
	exist, err := IsPathExists(filePath)
	if err != nil {
		fmt.Printf("创建失败，因为判断是否有该文件夹时发生了错误[%v]\n", err)
		return err
	}
	// 如果文件夹不存在
	if !exist {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}


// 获取指定目录的最小文件名，带完整路径
func GetMinFile(pattern string) (minFile string) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	files, _ := filepath.Glob(pattern)
	// 如果jpg文件小于1个，就不执行
	if len(files) < 1 {
		return
	}

	// 先定义一个切片去装每个文件的大小信息
	filesSize := make([]int64, len(files))
	// 将所有文件的大小信息装进切片
	for i := 0; i < len(filesSize); i++ {
		// 获取文件全部信息
		fileInfo, _ := os.Stat(files[i])
		// 获取指定文件大小信息，返回的是字节
		nowSize := fileInfo.Size()
		// 按文件名的顺序装入
		filesSize[i] = nowSize
	}
	// 打印储存文件大小的切片
	//fmt.Println(filesSize)

	// 预定义最小文件的索引是0
	minFileIndex := 0
	for i := 1; i < len(filesSize); i++ {
		// 如果最小文件的索引大于 i 的索引
		if filesSize[minFileIndex] > filesSize[i] {
			// 那么就更改最小文件的索引
			minFileIndex = i
		}
	}
	// 打印最小的索引
	//fmt.Println(minFileIndex)

	// 打印最小的文件
	//fmt.Println(files[minFileIndex])
	return files[minFileIndex]
}