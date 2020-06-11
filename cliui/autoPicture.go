package cliui

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/autoPicture"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

//  家具主图选择
func furnitureMainPictureChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Furniture main picture choice", 74)
		text := `
:: 家具店主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与Logo！

   [1]怡柟家具                   [2]御尚家具                   [3]素梵家具

   [4]棠语家具                   [5]凌轩家具                   [6]暂未开发`
		fmt.Println(text)
		layoutType, info := model.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch layoutType {
		case "1":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yinanjj.png", viper.GetBool("automaticDeletion"))
		case "2":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yushantanjj.png", viper.GetBool("automaticDeletion"))
		case "3":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/sufanjj.png", viper.GetBool("automaticDeletion"))
		case "4":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/tangyujj.png", viper.GetBool("automaticDeletion"))
		case "5":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/lingxuan.png", viper.GetBool("automaticDeletion"))
		case "6":

		case "7":

		case "8":

		case "9":

		case "-":
			break OuterLoop
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(layoutType, ctc.ForegroundGreen))
			}
		}
	}
}

//  屏风主图选择
func screenMainPictureChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Screen main picture choice", 74)
		text := `
:: 屏风店主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与Logo！

   [1]沐兰主图                   [2]华府主图                   [3]木韵主图

   [4]御尚屏风                   [5]金尊主图                   [6]暂未开发`
		fmt.Println(text)

		layoutType, info := model.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch layoutType {
		case "1":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/mulan.png", viper.GetBool("automaticDeletion"))
		case "2":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/huafu.png", viper.GetBool("automaticDeletion"))
		case "3":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/muyunge.png", viper.GetBool("automaticDeletion"))
		case "4":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/yushantanpf.png", viper.GetBool("automaticDeletion"))
		case "5":
			autoPicture.WatermarkMasterGraph(viper.GetString("picture"), "config/static/img/jinzunfu.png", viper.GetBool("automaticDeletion"))
		case "6":
		case "-":
			break OuterLoop
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(layoutType, ctc.ForegroundGreen))
			}
		}
	}
}

// 修改图片分辨率的选择
func modifyResolution() {
OuterLoop:
	for {
		tools.EnglishTitle("Modify resolution", 74)
		text := `
:: 在保持原来长宽像素不变的情况下，将分辨率强行修改至指定数值，请尽量少用！

   [1]全部改为72PPI         [2]全部改为300PPI         [3]全部改为自定义PPI.`
		fmt.Println(text)
		layoutType, info := model.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear()
		switch layoutType {
		case "1":
			autoPicture.PixelsPerInchChangedTo72(viper.GetString("picture")) // 改为72ppi
		case "2":
			autoPicture.PixelsPerInchChangedTo300(viper.GetString("picture"))
		case "3":
			fmt.Println("\n:: 因需求不高，所以暂未开发...")
		case "-", "--":
			break OuterLoop
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(layoutType, ctc.ForegroundGreen))
			}
		}
	}
}

// 套图的选择
func autoPictureChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Auto picture", 74)
		text := `
:: 自动套图可以正确运行的前提是，款式图片均要放在套图文件夹下方可正确运行！

   [1]套图文件夹                [2]随机重命名                [3]备份文件夹

   [4]家具店主图.               [5]通用版主图                [6]屏风店主图.

   [7]修改分辨率.               [8]替换详情页                [9]导出详情页`
		fmt.Println(text)

		layoutType, info := model.InputMenuSelection("\n:: 请选择需要使用的功能：")
		tools.CallClear() // 清屏
		switch layoutType {
		case "1":
			go exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", viper.GetString("picture"))).Run()
			fmt.Println("\n:: 已打开套图文件夹，请复制正方形的 jpg 或 png 高清图片以备自动套图使用！")
		case "2":
			autoPicture.RandomRename(viper.GetString("picture"))
		case "3":
			go exec.Command("cmd.exe", "/c", "start Config\\Backups").Run()
			fmt.Println("\n:: 已打开备份文件夹，为了避免意外丢失文件，目前备份文件最大为 15 份！")
		case "4":
			furnitureMainPictureChoice() // 家具店主图
		case "5":
			autoPicture.UniversalMasterGraph(viper.GetString("picture"), viper.GetBool("automaticDeletion")) // 通用主图
		case "6":
			screenMainPictureChoice() // 屏风店主图
		case "7":
			modifyResolution() // 修改分辨率
		case "8":
			autoPicture.ReplaceDetailsPage(viper.GetString("picture")) // 替换详情页
		case "9":
			autoPicture.SaveForWebDetailsPage() // 导出web格式详情页
		case "-":
			break OuterLoop
		default:
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入...\n", tools.ColourString(layoutType, ctc.ForegroundGreen))
			}
		}
	}
}
