package cliui

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
)

// 旧厂框架的选择
func oldFrameChoice() {
	for {
		tools.EnglishTitle("Cutting", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]常规座屏                   [2]左右镂空                   [3]左右画布

   [4]上下镂空                   [5]顶天立地                   [6]各种折屏

   [7]多个座屏                   [8]卷帘座屏                   [9]不扣补切`
		fmt.Println(text)

		frameType , info:= model.Input("\n:: 请选择上方的边框类型：", false,true)
		tools.CallClear() // 清屏
		switch frameType {
		case "1":
			model.OldFrame1()       // 小座屏
		case "2":
			model.OldFrame2()       // 左右镂空
		case "3":
			model.OldFrame3()       // 左右画布
		case "4":
			model.OldFrame4()       // 上下镂空
		case "5":
			model.OldFrame5()       // 顶天立地
		case "6":
			model.OldFrame6()       // 常规折屏
		case "7":
			model.OldFrame7()       // 多座屏
		case "8":
			model.OldFrame8()       // 卷帘座屏
		case "9":
			model.OldFrame9()       // 补切画布
		case "-", "--":
			goto FLAG
		case "cls":
			// 收到清屏命令
			if len(info) != 0 {
				fmt.Println(info)
			}
			continue
		default:
			fmt.Printf("\n:: 输入的 [%s] 不是已知的边框类型，请重新输入！\n", tools.ColourString(frameType, ctc.ForegroundGreen))
			continue
		}
	}
FLAG:
}



