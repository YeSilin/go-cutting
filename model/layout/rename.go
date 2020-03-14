package layout

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// 一个生成随机数切片的函数，从1开始
func randomNumberSlice(r int) (data []int) {
	//创建随机数种子 进行数据混淆
	rand.Seed(time.Now().UnixNano())

	// 创建一个用于存放数据的切片,长度是r
	data = make([]int, r)

	// 为data切片循环赋值
	for i := 0; i < r; i++ {
		// 生成一个随机数，范围是r
		v := rand.Intn(r) + 1

		// 循环验证是否重复
		for j := 0; j < i; j++ {
			// 如果重复
			if v == data[j] {
				// 重新生成随机数
				v = rand.Intn(r) + 1
				//将j赋值为-1在循环执行到上面是进行++操作后值为0  也就相当于重新比对
				j = -1
			}
		}
		//将没有重复的数字添加到数据切片中
		data[i] = v
	}
	return
}

// 随机重命名文件，不支持带*路径
func randomRename(srcPath, extensionName string) {
	// 获取最小的文件
	min := tools.GetMinFile(fmt.Sprintf("%s/*.%s", srcPath, extensionName))

	// 定义一个临时名字
	tempMin := srcPath + "/GoCuttingTemp.min"
	// 先把最小的改成 .man
	_ = os.Rename(min, tempMin)

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	files, _ := filepath.Glob(fmt.Sprintf("%s/*.%s", srcPath, extensionName))
	// 如果jpg文件小于1个，就不执行
	if len(files) < 1 {
		return
	}

	// 全部一起改名为temp
	for i, v := range files {
		_ = os.Rename(v, fmt.Sprintf("%s/GoCuttingTemp%d.%s", srcPath, i+1, extensionName))
	}

	// 生成随机数切片
	ids := randomNumberSlice(len(files))

	// 第二次获取所有扩展名是jpg的文件名，类型是字符串切片
	files2, _ := filepath.Glob(fmt.Sprintf("%s/*.%s", srcPath, extensionName))

	// 第二次修改为正确名字
	for i, v := range files2 {
		// 重命名为随机数编号
		_ = os.Rename(v, fmt.Sprintf("%s/%d.%s", srcPath, ids[i], extensionName))
	}

	// 记得把最小的 改回 正确的扩展名，把最小的那张固定成最后一张
	min = fmt.Sprintf("%s/%d.%s", srcPath, len(files2)+1, extensionName)
	_ = os.Rename(tempMin, min)
}

// 随机修改文件名为 123456
func Rename(originalPath string) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.png", originalPath))

	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 && len(pngSlice) < 1 {
		fmt.Println("\n【提示】重命名失败，因为 Picture 文件夹下没有 jpg 或 png 格式图片！")
		return
	}

	// 为了防止文件丢失，在重命名之前先备份一次文件
	_ = tools.CopyDir(originalPath, "Config/Backups/")

	// 随机重命名
	randomRename(originalPath, "jpg")

	// 删除多余备份，最大保留10个
	tools.DeleteRedundantBackups("Config/Backups/*", 10)

	fmt.Println("\n【提示】随机重命名成功，现已支持所有尺寸的 jpg 或 png 格式图片！")
}
