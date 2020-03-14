// 关于主图修改
package layout

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"image"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// 通用主图第二版
func UniversalMasterGraph(originalPath string, delete bool) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.png", originalPath))

	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 && len(pngSlice) < 1 {
		fmt.Println("\n【提示】转换失败，因为套图文件夹下没有 jpg 或 png 格式图片！")
		return
	}
	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir(originalPath, "Config/Backups/")

		// 完成后的目标路径
		resultPath := fmt.Sprintf("%s/主图", originalPath)

		// 创建目标路径文件夹
		_ = tools.CreateMkdirAll(resultPath)

		// jpg修改全部大小为800
		for _, v := range jpgSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, originalPath, resultPath, 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// png修改全部大小为800
		for _, v := range pngSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, originalPath, resultPath, 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// 删除复制的文件
		if delete {
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
		}

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("Config/Backups/*", 10)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}

// 带水印主图，水印路径，是否有白底图
func WatermarkMasterGraph(originalPath, watermarkPath string, delete bool) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.png", originalPath))

	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 && len(pngSlice) < 1 {
		fmt.Println("\n【提示】转换失败，因为套图文件夹下没有 jpg 或 png 格式图片！")
		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", originalPath)).Run()
		return
	}

	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir(originalPath, "Config/Backups/")

		// 完成后的目标路径
		resultPath := fmt.Sprintf("%s/主图", originalPath)

		// 创建目标路径文件夹
		_ = tools.CreateMkdirAll(resultPath)

		// 获取白底图
		minImage, isMinImage := tools.MinWhiteBackground(fmt.Sprintf("%s/*.jpg", originalPath))

		// 获取最小图片 - 白底图
		//minImage := tools.GetMinFile("Config\\Picture\\*.jpg")

		// 如果有白底图
		if isMinImage {
			// jpg修改全部大小为800，并且加水印
			for _, v := range jpgSlice {
				if v == minImage { // 只改大小不加水印
					dstPath := strings.Replace(minImage, originalPath, resultPath, 1)
					tools.ImageResize(minImage, dstPath, 800, 800, 99)
					continue
				}
				srcPath := v                                                      // 源图像
				savePath := strings.Replace(srcPath, originalPath, resultPath, 1) // 保存路径
				// 修改源图像大小，并且加水印
				// savePath, originalPath string, width, height int, watermarkPath string, x, y int
				tools.ImageResizeAndWatermark(savePath, srcPath, 800, 800, watermarkPath, 0, 0)
			}
		} else { // 没白底图的时候
			// jpg修改全部大小为800，并且加水印
			for _, v := range jpgSlice {
				srcPath := v                                                      // 源图像
				savePath := strings.Replace(srcPath, originalPath, resultPath, 1) // 保存路径
				// 修改源图像大小，并且加水印
				// savePath, originalPath string, width, height int, watermarkPath string, x, y int
				tools.ImageResizeAndWatermark(savePath, srcPath, 800, 800, watermarkPath, 0, 0)
			}
		}

		// png修改全部大小为800，png不加水印只改大小
		for _, v := range pngSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, originalPath, resultPath, 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// 删除复制的文件
		if delete {
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
		}

		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", originalPath)).Run()

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("config/Backups/*", 10)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}

// 全部修改为72ppi
func PixelsPerInchChangedTo72(originalPath string) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	files, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 如果jpg文件小于1个，就不执行
	if len(files) < 1 {
		fmt.Println("\n【提示】转换失败，因为套图文件夹下没有 jpg 格式图片！")
		return
	}

	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir(originalPath, "Config/Backups/")

		// 全部一起改名为temp
		for _, v := range files {
			_ = os.Rename(v, strings.Replace(v, ".jpg", ".gcTemp", 1))
		}

		// 重新获取所有扩展名是gcTemp的文件名，类型是字符串切片
		files2, _ := filepath.Glob(fmt.Sprintf("%s/*.gcTemp", originalPath))
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

// 全部修改为300ppi
func PixelsPerInchChangedTo300(originalPath string) {
	// 返回绝对路径
	originalPath ,err := filepath.Abs(originalPath)
	if err != nil {
		fmt.Println("filepath.Abs err:",err)
		return
	}
	// 全部换成正斜杠
	originalPath = strings.Replace(originalPath, "\\", "/", -1)


	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))

	// 如果jpg文件小于1个，就不执行
	if len(jpgSlice) < 1 {
		fmt.Println("\n【提示】转换失败，因为套图文件夹下没有 jpg 格式图片！")
		return
	}

	go func() {
		// 为了防止文件丢失，在重命名之前先备份一次文件
		_ = tools.CopyDir(originalPath, "Config/Backups/")

		// 解析指定文件生成模板对象
		tmpl, err := template.ParseFiles("config/jsx/template/pixelsPerInchChangedTo300.gohtml")
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}

		// 创建文件，返回两个值，一是创建的文件，二是错误信息
		f, err := os.Create("config/jsx/pixelsPerInchChangedTo300.jsx")
		if err != nil { // 如果有错误，打印错误，同时返回
			fmt.Println("创建文件错误 =", err)
			return
		}

		// 修改成js脚本可以看懂的路径
		for i := 0; i < len(jpgSlice); i++ {
			jpgSlice[i] = strings.Replace(jpgSlice[i], "\\", "/", -1)
			jpgSlice[i] = strings.Replace(jpgSlice[i], ":", "", 1)
			jpgSlice[i] = "/"+jpgSlice[i]
		}

		// 利用给定数据渲染模板，并将结果写入f
		_ = tmpl.Execute(f, tools.StrToJsArray("srcArray", jpgSlice))

		// 关闭文件
		f.Close()

		// 创建一个协程使用cmd来运行脚本
		dataPath := "Config/jsx/pixelsPerInchChangedTo300.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("Config/Backups/*", 10)
	}()

	fmt.Println("\n【提示】脚本注入成功，正在转成 300PPI 若文件丢失，备份文件在上级目录 Backups！")
}
