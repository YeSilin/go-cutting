package cliui

import (
	"fmt"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
)

// 贴图框架的选择
func mapFrameChoice() {
	for {
		tools.EnglishTitle("3ds Max map frame", 74)
		text := `
:: 下方所有框架的切图单位均是像素，小数点位之后值均会被舍弃，并非四舍五入！

   [1]常规座屏贴图             [2]左右镂空贴图             [3]左右画布贴图

   [4]上下镂空贴图             [5]顶天立地贴图             [6]各种折屏贴图

   [7]多个座屏贴图             [8]卷帘座屏贴图             [9]不扣补切贴图`
		fmt.Println(text)

		frameType , info:= model.Input("\n:: 请选择上方的边框类型：", false,true)
		tools.CallClear() // 清屏
		switch frameType {
		case "1":
			model.MapFrame1() // 小座屏
		case "2":
			fmt.Println("未开发") // 左右镂空
		case "3":
			fmt.Println("未开发") // 中间大两边小
		case "4":
			fmt.Println("未开发") // 上下镂空
		case "5":
			fmt.Println("未开发") // 顶天立地
		case "6":
			model.MapFrame6() // 常规折屏
		case "7":
			fmt.Println("未开发") // 多座屏
		case "8":
			fmt.Println("未开发") // 补切画布
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