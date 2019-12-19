package layout

import (
	"fmt"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/tools"
	"image"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)



// 通用主图第一版
func picture01() {
	// 为了防止文件丢失，在重命名之前先备份一次文件
	//_ = tools.CopyDir("Config/Picture", "Config/Backups/")

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	files, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 如果jpg文件小于1个，就不执行
	if len(files) < 1 {
		return
	}
	// 创建套图文件夹
	_ = tools.CreateMkdirAll("Config/Picture/主图")
	generate.UniversalMasterGraph(len(files))

	// 多出几个文件数量就循环几次，如果是负数自然就不循环
	//for i := 0; i < len(files); i++ {
	//	// 删除文件，Go中删除文件和删除文件夹同一个函数
	//	err := os.Remove(files[i]) // 由于windows 系统获取到的文件名字，默认是升序，于是可以不用排序
	//	// 打印被删除的文件
	//	//fmt.Println(files[i])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	// 删除多余备份，最大保留10个
	//tools.DeleteRedundantBackups("Config/Backups/*",10)
	// 打开套图文件夹
	exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
}

// 通用主图第二版
func UniversalMasterGraph() {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(".\\Config\\Picture\\*.png")

	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 && len(pngSlice) < 1 {
		fmt.Println("\n【提示】转换失败，因为 Picture 文件夹下没有 jpg 或 png 格式图片！")
		return
	}
	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir("Config/Picture", "Config/Backups/")

		// 创建套图文件夹
		_ = tools.CreateMkdirAll("Config/Picture/主图")

		// jpg修改全部大小为800
		for _, v := range jpgSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, "\\Picture", "\\Picture\\主图", 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// png修改全部大小为800
		for _, v := range pngSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, "\\Picture", "\\Picture\\主图", 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// 删除之前复制的jpg
		for _, v := range jpgSlice {
			err := os.Remove(v)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 删除之前复制的png
		for _, v := range pngSlice {
			err := os.Remove(v)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("Config/Backups/*", 10)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}

// 带水印主图，水印路径，是否有白底图
func WatermarkMasterGraph(watermarkPath string) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(".\\Config\\Picture\\*.png")

	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 && len(pngSlice) < 1 {
		fmt.Println("\n【提示】转换失败，因为 Picture 文件夹下没有 jpg 或 png 格式图片！")
		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
		return
	}

	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir("Config/Picture", "Config/Backups/")

		// 创建套图文件夹
		_ = tools.CreateMkdirAll("Config/Picture/主图")

		// 获取白底图
		minImage, isMinImage := tools.MinWhiteBackground("Config\\Picture\\*.jpg")

		// 获取最小图片 - 白底图
		//minImage := tools.GetMinFile("Config\\Picture\\*.jpg")

		// 如果有白底图
		if isMinImage {
			// jpg修改全部大小为800，并且加水印
			for _, v := range jpgSlice {
				if v == minImage { // 只改大小不加水印
					dstPath := strings.Replace(minImage, "\\Picture", "\\Picture\\主图", 1)
					tools.ImageResize(minImage, dstPath, 800, 800, 99)
					continue
				}
				originalPath := v                                                          // 源图像
				savePath := strings.Replace(originalPath, "\\Picture", "\\Picture\\主图", 1) // 保存路径
				// 修改源图像大小，并且加水印
				// savePath, originalPath string, width, height int, watermarkPath string, x, y int
				tools.ImageResizeAndWatermark(savePath, originalPath, 800, 800, watermarkPath, 0, 0)
			}
		} else { // 没白底图的时候
			// jpg修改全部大小为800，并且加水印
			for _, v := range jpgSlice {
				originalPath := v                                                          // 源图像
				savePath := strings.Replace(originalPath, "\\Picture", "\\Picture\\主图", 1) // 保存路径
				// 修改源图像大小，并且加水印
				// savePath, originalPath string, width, height int, watermarkPath string, x, y int
				tools.ImageResizeAndWatermark(savePath, originalPath, 800, 800, watermarkPath, 0, 0)
			}
		}

		// png修改全部大小为800，png不加水印只改大小
		for _, v := range pngSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, "\\Picture", "\\Picture\\主图", 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// 删除之前复制的jpg
		for _, v := range jpgSlice {
			err := os.Remove(v)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 删除之前复制的png
		for _, v := range pngSlice {
			err := os.Remove(v)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("Config/Backups/*", 10)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}

// 全部修改为72ppi
func AllResolution72() {

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	files, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 如果jpg文件小于1个，就不执行
	if len(files) < 1 {
		fmt.Println("\n【提示】转换失败，因为 Picture 文件夹下没有 jpg 格式图片！")
		return
	}

	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir("Config/Picture", "Config/Backups/")

		// 全部一起改名为temp
		for _, v := range files {
			_ = os.Rename(v, strings.Replace(v, ".jpg", ".gcTemp", 1))
		}

		// 重新获取所有扩展名是gcTemp的文件名，类型是字符串切片
		files2, _ := filepath.Glob(".\\Config\\Picture\\*.gcTemp")
		// jpg修改全部大小为72ppi
		for i, srcPath := range files2 {
			// 打开文件
			file, _ := os.Open(srcPath)
			// 获取文件信息
			c, _, _ := image.DecodeConfig(file)
			width := c.Width
			height := c.Height
			file.Close() // 必须关闭文件
			// 修改大小
			tools.ImageResize(srcPath, files[i], width, height, 100)
		}

		// 删除之前自动生成的gcTemp
		for _, v := range files2 {
			err := os.Remove(v)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("Config/Backups/*", 10)
	}()

	fmt.Println("\n【提示】已转成 72ppi 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}
