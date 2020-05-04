// 专门用来做临时效果图的
package model

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"strconv"
	"strings"
)

// 临时框架的选择
func Choice() {
	for {
		EnglishTitle("Temporary", 74)

		// 提示标题
		tips := `
:: 提供简单框架的自动效果图一键生成，由于效果图框架复杂暂时只提供如下功能！

   [1]新建效果图      [2]置入小座屏      [3]置入单折屏      [4]置入单镂空`
		fmt.Println(tips)
		frameType := Input("\n:: 请选择上方的功能类型：", false)

		switch frameType {
		case "1":
			generate.NewTempDocumentJs() // 新建临时文档

		case "2":
			tempFame1() // 小座屏
		case "3":
			tempFame2() // 折屏
		case "4":
			tempFame3() // 镂空
		case "5":
			fmt.Println("未开发") // 顶天立地
		case "6":
			fmt.Println("未开发") // 顶天立地
		case "7":
			fmt.Println("未开发") // 多座屏
		case "8":
			fmt.Println("未开发") // 补切画布
		case "-":
			goto FLAG
		default:
			continue
		}
	}
FLAG:
}

// 临时效果图小座屏框架细分
func tempFame1() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]80-180\t[2]100-180\t[3]120-180\t[4]自定义尺寸")
		frameType := Input("\n:: 请选择上方的边框尺寸：", false)

		switch frameType {
		case "1":
			generate.SelectionTempFrameJS("Frame01", 0)
		case "2":
			generate.SelectionTempFrameJS("Frame01", 1)
		case "3":
			generate.SelectionTempFrameJS("Frame01", 2)
		case "4":
			tempFame1To4() // 小座屏自定义框架
		case "-":
			goto FLAG
		default:
			continue
		}
	}
FLAG:
}

func tempFame1To4() {
	for {
		widthStr := Input("\n:: 请输入小座屏的宽：", true)
		if widthStr == "-" {
			break
		}

		heightStr := Input("\n:: 请输入小座屏的高：", true)
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

// 临时效果图折屏框架细分
func tempFame2() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]45-180\t[2]50-190\t[3]60-190\t[4]60-200")
		frameType := Input("\n:: 请选择上方的边框尺寸：", false)

		switch frameType {
		case "1":
			generate.SelectionTempFrameJS("Frame02", 0)
		case "2":
			generate.SelectionTempFrameJS("Frame02", 1)
		case "3":
			generate.SelectionTempFrameJS("Frame02", 2)
		case "4":
			generate.SelectionTempFrameJS("Frame02", 3)
		case "-":
			goto FLAG
		default:
			continue
		}
	}
FLAG:
}

// 临时效果图折屏框架细分
func tempFame3() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]回字形\t[2]竖条形\t[3]功能待定\t[4]功能待定")
		frameType := Input("\n:: 请选择上方的镂空类型：", false)

		switch frameType {
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
			goto FLAG
		default:
			continue
		}
	}
FLAG:
}
