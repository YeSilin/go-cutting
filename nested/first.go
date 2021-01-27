//自动套图 - 对主图的操作
package nested

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 拷贝所有TXT文件到指定目录
func copyAllTXT(srcPath, dstPath string) {
	// 获取所有扩展名是txt的文件名，类型是字符串切片
	txtSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.txt", srcPath))

	// 复制所有文本到指定文件夹
	for i := 0; i < len(txtSlice); i++ {
		// 生成目标文件路径
		dst := strings.Replace(txtSlice[i], srcPath, dstPath, 1)
		tools.CopyFile(txtSlice[i], dst)
	}
}

// 删除所有指定路径文件
func delAllFiles(files []string) {
	// 删除之前复制的jpg
	for i := range files {
		err := os.Remove(files[i])
		if err != nil {
			logrus.Error(err)
		}
	}
}

// 修改所有图片大小到指定目录
func modifyAllImgSize(images []string, dstPath string) {
	// jpg修改全部大小为800，到主图文件夹
	for i := range images {
		// 得到文件名
		_, file := filepath.Split(images[i])
		ret := dstPath + "/" + file
		tools.ImageResize(images[i], ret, 800, 800, 98)
	}
}

// 通用主图第二版
func UniversalMainImage(originalPath string, delete bool) {
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
		_ = tools.CopyDir(originalPath, "config/Backups/")

		// 完成后的目标路径
		resultPath := fmt.Sprintf("%s/主图", originalPath)

		// 创建目标路径文件夹
		_ = tools.CreateMkdirAll(resultPath)

		// jpg修改全部大小为800，到主图文件夹
		modifyAllImgSize(jpgSlice, resultPath)

		// png修改全部大小为800，到主图文件夹
		modifyAllImgSize(pngSlice, resultPath)

		// 复制所有文本到主图文件夹
		copyAllTXT(originalPath, resultPath)

		// 删除复制的文件
		if delete {
			delAllFiles(jpgSlice) // 删除之前复制的jpg
			delAllFiles(pngSlice) // 删除之前复制的png
		}

		// 删除多余备份，最大保留15个
		tools.DeleteRedundantBackups("config/Backups/*", 15)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}

// 带水印主图，水印路径，是否有白底图
func WatermarkMainImage(originalPath, watermarkPath string, delete bool) {
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

		// 完成后的主图路径
		resultPath := fmt.Sprintf("%s/主图", originalPath)
		// 完成后的主图无水印路径
		noWatermarkPath := resultPath + "/无水印"

		// 创建必须存在的文件夹
		_ = tools.CreateMkdirAll(noWatermarkPath)

		// 无水印主图操作，jpg修改全部大小为800，到 主图/无水印 文件夹
		modifyAllImgSize(jpgSlice, noWatermarkPath)

		// 带水印操作，获取白底图
		minImage, isMinImage := tools.MinWhiteBackground(fmt.Sprintf("%s/*.jpg", originalPath))

		// 如果有白底图
		if isMinImage {
			// jpg修改全部大小为800，并且加水印，到主图文件夹
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
			// jpg修改全部大小为800，并且加水印，到主图文件夹
			for _, v := range jpgSlice {
				srcPath := v                                                      // 源图像
				savePath := strings.Replace(srcPath, originalPath, resultPath, 1) // 保存路径
				// 修改源图像大小，并且加水印
				// savePath, originalPath string, width, height int, watermarkPath string, x, y int
				tools.ImageResizeAndWatermark(savePath, srcPath, 800, 800, watermarkPath, 0, 0)
			}
		}

		// png修改全部大小为800，png不加水印只改大小，到主图文件夹
		modifyAllImgSize(pngSlice, resultPath)

		// 复制所有文本到主图文件夹
		copyAllTXT(originalPath, resultPath)

		// 删除复制的文件
		if delete {
			delAllFiles(jpgSlice) // 删除之前复制的jpg
			delAllFiles(pngSlice) // 删除之前复制的png
		}

		// 删除多余备份，最大保留10个
		tools.DeleteRedundantBackups("config/Backups/*", 15)
	}()
	fmt.Println("\n【提示】已转成 800*800 如果文件丢失，备份文件夹在上级目录下的 Backups！")
}
