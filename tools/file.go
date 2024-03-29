package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

//ReadFile 读取文件，返回字符切片
func ReadFile(Path string) (content []byte, err error) {
	var f *os.File

	// 只读方式打开文件
	f, err = os.Open(Path)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	defer f.Close()

	// 读取内容保存到这个切片
	buf := make([]byte, 4096)
	for {
		var n int
		n, err = f.Read(buf)
		//如果错误不是空，并且也不是文件末尾，那就表示出错了
		if err != nil && err != io.EOF {
			fmt.Println("f.Read err:", err)
			return
		}
		if err == io.EOF {
			break
		}
		// 合并切片
		content = append(content, buf[:n]...)
	}
	return
}

//ReadFileEasy 读取文件，返回字符切片，不返回错误
func ReadFileEasy(Path string) (content []byte) {
	// 只读方式打开文件
	f, err := os.Open(Path)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	defer f.Close()

	// 读取内容保存到这个切片
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		//如果错误不是空，并且也不是文件末尾，那就表示出错了
		if err != nil && err != io.EOF {
			fmt.Println("f.Read err:", err)
			return
		}
		if err == io.EOF {
			break
		}
		// 合并切片
		content = append(content, buf[:n]...)
	}
	return
}

// CreateFile 创建文件并写入数据的函数
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

// WriteFile 追加写入数据
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

// ReadLine 用行读取文件
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

// FormatReadLine 用go格式，行读取文件(目前这个函数用来生成js脚本字符串)
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
		bufStr := fmt.Sprintf("%q", string(buf))

		// 添加前后缀
		bufStr = fmt.Sprintf("%s.WriteString(%s)", name, bufStr)

		// 打印出来
		fmt.Println(bufStr)

		if err == io.EOF { // 文件读取结束
			break
		}
	}
}

// IsPathExists 判断所给路径文件/文件夹是否存在(返回true是存在)
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

// CreateMkdir 创建单个文件夹
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

// CreateMkdirAll 调用os.MkdirAll递归创建文件夹
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

// GetMinFile 获取指定目录的最小文件名，带完整路径
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

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// SubdirectoryMap 获取指定目录中的子目录，以map形式
func SubdirectoryMap(originalPath string) (files map[string][]string) {
	// 字典和切片需要先分配内存空间
	files = make(map[string][]string)

	// 获取所有文件
	f, err := filepath.Glob(originalPath)
	if err != nil {
		fmt.Println("filepath.Glob err: ", err)
		return
	}

	// 开始存进 map 里
	for i := 0; i < len(f); i++ {
		// 先把所有反斜杠修改成正斜杠
		f[i] = strings.Replace(f[i], "\\", "/", -1)

		// 如果是文件夹
		if IsDir(f[i]) {

			// 对字符串进行拼接，获得子文件路径
			fromPath := fmt.Sprintf("%s/*", f[i])

			// 获取子文件夹所有文件
			f2, err := filepath.Glob(fromPath)
			if err != nil {
				fmt.Println("filepath.Glob err: ", err)
				return
			}

			// 得到子文件夹名
			dirName := filepath.Base(f[i])

			// 循环子文件夹内的文件
			for j := 0; j < len(f2); j++ {
				// 先把所有反斜杠修改成正斜杠
				f2[j] = strings.Replace(f2[j], "\\", "/", -1)
				files[dirName] = append(files[dirName], f2[j])
			}

			// 非文件夹存默认 map 里
		} else {
			files["default"] = append(files["default"], f[i])
		}
	}
	return
}

// CopyFile 拷贝单个文件到指定位置
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// OpenFolder 无黑框的方式打开windows文件夹
func OpenFolder(path string, hideWindow bool) {
	//必须要转换下，cmd不支持正斜杆
	path = strings.ReplaceAll(path, "/", "\\")
	cmd := exec.Command("cmd", "/c", "start "+path)

	// 是否隐藏cmd窗口，适合在有gui界面的情况下使用
	if hideWindow {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 不显示黑框
	}
	cmd.Run()
}
