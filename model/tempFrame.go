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



// 临时效果图小座屏框架细分
func TempFame1() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]80-180\t[2]100-180\t[3]120-180\t[4]自定义尺寸")
		frameType, info := Input("\n:: 请选择上方的边框尺寸：", false,true)

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
		case "cls":
			// 收到清屏命令
			if len(info) != 0 {
				fmt.Println(info)
			}
			continue
		default:
			continue
		}
	}
FLAG:
}

func tempFame1To4() {
	for {
		widthStr, _ := Input("\n:: 请输入小座屏的宽：", true,false)
		if widthStr == "-" {
			break
		}

		heightStr , _:= Input("\n:: 请输入小座屏的高：", true,false)
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
func TempFame2() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]45-180\t[2]50-190\t[3]60-190\t[4]60-200")
		frameType , info:= Input("\n:: 请选择上方的边框尺寸：", false,true)

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
		case "cls":
			// 收到清屏命令
			if len(info) != 0 {
				fmt.Println(info)
			}
			continue
		default:
			continue
		}
	}
FLAG:
}

// 临时效果图折屏框架细分
func TempFame3() {
	for {
		fmt.Println("\n" + strings.Repeat("-", 37) + " Size " + strings.Repeat("-", 36))
		//tools.PrintLine(2)
		fmt.Println("\n:: [1]回字形\t[2]竖条形\t[3]功能待定\t[4]功能待定")
		frameType , info:= Input("\n:: 请选择上方的镂空类型：", false,true)

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
		case "cls":
			// 收到清屏命令
			if len(info) != 0 {
				fmt.Println(info)
			}
			continue
		default:
			continue
		}
	}
FLAG:
}
