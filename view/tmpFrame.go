package view

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
)

// 小座屏自定义框架
func (this *CliView) tempFame1To4() {
	for {
		tools.EnglishTitle("Custom size", 74)

		widthStr := input.InputCanvasSize("\n:: 请输入小座屏的宽：", 6)
		if widthStr == "-" {
			break
		}

		heightStr := input.InputCanvasSize("\n:: 请输入小座屏的高：", 6)
		if heightStr == "-" {
			break
		}

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)

		color.Green.Printf("\n:: 小座屏：宽 %.2f cm，高 %.2f cm", width, height)

		generate.TempFrame1JS(width*10, height*10) // 生成小座屏效果图框架

		generate.MaxCanvas(width-5, height-5) // 最大画布判断

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 临时效果图小座屏框架细分
func (this *CliView) tempFame1() {
OuterLoop:
	for {
		tools.EnglishTitle("Size selection", 74)
		// 提示标题
		tips := `
:: 以下为常见小座屏边框尺寸也可自定义，置入成功的前提是已新建好效果图背景！

   [1]80-180         [2]100-180         [3]120-180         [4]自定义尺寸.`
		fmt.Println(tips)

		this.key, this.info = input.InputMenuSelection("\n:: 请选择上方的边框尺寸：")
		tools.CallClear() // 清屏

		switch this.key {
		case "1":
			generate.SelectionTempFrameJS("Frame01", 0)
		case "2":
			generate.SelectionTempFrameJS("Frame01", 1)
		case "3":
			generate.SelectionTempFrameJS("Frame01", 2)
		case "4":
			this.tempFame1To4() // 小座屏自定义框架
		case "-":
			break OuterLoop
		default:
			if len(this.info) != 0 {
				fmt.Println(this.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的框架类型，请重新输入...\n", tools.ColourString(this.key, ctc.ForegroundGreen))
			}
		}
	}
}

// 临时效果图折屏框架细分
func (this *CliView) tempFame2() {
OuterLoop:
	for {
		tools.EnglishTitle("Size selection", 74)
		// 提示标题
		tips := `
:: 以下为常见的折屏边框尺寸不可自定义，置入成功的前提是已新建好效果图背景！

   [1]45-180           [2]50-190           [3]60-190           [4]60-200`
		fmt.Println(tips)

		this.key, this.info = input.InputMenuSelection("\n:: 请选择上方的边框尺寸：")
		tools.CallClear() // 清屏

		switch this.key {
		case "1":
			generate.SelectionTempFrameJS("Frame02", 0)
		case "2":
			generate.SelectionTempFrameJS("Frame02", 1)
		case "3":
			generate.SelectionTempFrameJS("Frame02", 2)
		case "4":
			generate.SelectionTempFrameJS("Frame02", 3)
		case "-":
			break OuterLoop
		default:
			if len(this.info) != 0 {
				fmt.Println(this.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的框架类型，请重新输入...\n", tools.ColourString(this.key, ctc.ForegroundGreen))
			}
		}
	}
}

// 镂空框架细分
func (this *CliView) tempFame3() {
OuterLoop:
	for {
		tools.EnglishTitle("Hollow type selection", 74)
		// 提示标题
		tips := `
:: 以下为常见的单侧镂空框架请按需选择，置入成功的前提是已新建好效果图背景！

   [1]回字镂空         [2]竖条镂空         [3]功能待定         [4]功能待定`
		fmt.Println(tips)

		this.key, this.info = input.InputMenuSelection("\n:: 请选择上方的镂空类型：")
		tools.CallClear() // 清屏

		switch this.key {
		case "1":
			generate.SelectionTempFrameJS("HollowFrame", 0)
		case "2":
			generate.SelectionTempFrameJS("HollowFrame", 1)
		case "3":
			//generate.SelectionTempFrameJS("Frame02", 2)
			fmt.Println("未开发")
		case "4":
			//generate.SelectionTempFrameJS("Frame02", 3)
			fmt.Println("未开发")
		case "-":
			break OuterLoop
		default:
			if len(this.info) != 0 {
				fmt.Println(this.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的镂空类型，请重新输入...\n", tools.ColourString(this.key, ctc.ForegroundGreen))
			}
		}
	}
}

// 临时框架的选择
func (this *CliView) temporaryChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Temporary renderings", 74)

		// 提示标题
		tips := `
:: 提供简单框架的自动效果图一键生成，由于效果图框架复杂暂时只提供如下功能！

   [1]新建背景         [2]常规座屏.        [3]单扇折屏.        [4]单侧镂空.`
		fmt.Println(tips)
		this.key, this.info = input.InputMenuSelection("\n:: 请选择上方的功能类型：")
		tools.CallClear() // 清屏
		switch this.key {
		case "1":
			generate.NewTempDocumentJs() // 新建临时文档
		case "2":
			this.tempFame1() // 小座屏
		case "3":
			this.tempFame2() // 折屏
		case "4":
			this.tempFame3() // 镂空
		case "5":
			fmt.Println("未开发") // 顶天立地
		case "6":
			fmt.Println("未开发") // 顶天立地
		case "7":
			fmt.Println("未开发") // 多座屏
		case "8":
			fmt.Println("未开发") // 补切画布
		case "-":
			break OuterLoop
		default:
			if len(this.info) != 0 {
				fmt.Println(this.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的框架类型，请重新输入...\n", tools.ColourString(this.key, ctc.ForegroundGreen))
			}
		}
	}
}
