package autoPicture

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)




// 通用主图第二版
func UniversalMasterGraph(originalPath string, delete bool) {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 获取所有扩展名是png的文件名，类型是字符串切片
	pngSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.png", originalPath))
	// 获取所有扩展名是txt的文件名，类型是字符串切片
	txtSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.txt", originalPath))

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
			tools.ImageResize(srcPath, dstPath, 800, 800, 98)
		}

		// png修改全部大小为800
		for _, v := range pngSlice {
			srcPath := v
			dstPath := strings.Replace(srcPath, originalPath, resultPath, 1)
			tools.ImageResize(srcPath, dstPath, 800, 800, 99)
		}

		// 复制所有文本到主图文件夹
		for i := 0; i < len(txtSlice); i++ {
			// 生成目标文件路径
			dst := strings.Replace(txtSlice[i], originalPath, resultPath, 1)
			tools.CopyFile(txtSlice[i], dst)
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
	// 获取所有扩展名是txt的文件名，类型是字符串切片
	txtSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.txt", originalPath))

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

		// 复制所有文本到主图文件夹
		for i := 0; i < len(txtSlice); i++ {
			// 生成目标文件路径
			dst := strings.Replace(txtSlice[i], originalPath, resultPath, 1)
			tools.CopyFile(txtSlice[i], dst)
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






//  专属主图选择
func mainPictureChoice() {
	for {
		model.EnglishTitle("Main Picture Choice", 79)
		text := `
【主图】专属版主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与LOGO！

【主图】[1]沐兰主图                [2]怡柟家具                [3]御尚家具

【主图】[4]御尚屏风                [5]木韵主图                [6]华府主图

【主图】[7]金尊主图                [8]素梵家具                [9]棠语家具`
		fmt.Println(text)

		layoutType := model.Input("\n【主图】请选择需要使用的功能：", false)

		switch layoutType {
		case "1", "01":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/mulan.png", viper.GetBool("automaticDeletion"))
		case "2", "02":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yinanjj.png", viper.GetBool("automaticDeletion"))
		case "3", "03":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yushantanjj.png", viper.GetBool("automaticDeletion"))
		case "4", "04":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yushantanpf.png", viper.GetBool("automaticDeletion"))
		case "5", "05":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/muyunge.png", viper.GetBool("automaticDeletion"))
		case "6", "06":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/huafu.png", viper.GetBool("automaticDeletion"))
		case "7", "07":
			tools.CallClear() // 清屏
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/jinzunfu.png", viper.GetBool("automaticDeletion"))
		case "8", "08":
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/sufanjj.png", viper.GetBool("automaticDeletion"))
		case "9", "09":
			WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/tangyujj.png", viper.GetBool("automaticDeletion"))
		case "-", "--":
			goto FLAG
		default:
			tools.CallClear() // 清屏
			continue
		}
	}
FLAG:
}