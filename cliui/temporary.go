package cliui

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
)

// 临时框架的选择
func temporaryChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Temporary", 74)

		// 提示标题
		tips := `
:: 提供简单框架的自动效果图一键生成，由于效果图框架复杂暂时只提供如下功能！

   [1]新建效果图      [2]置入小座屏      [3]置入单折屏      [4]置入单镂空`
		fmt.Println(tips)
		frameType, info := model.InputMenuSelection("\n:: 请选择上方的功能类型：")
		tools.CallClear() // 清屏
		switch frameType {
		case "1":
			generate.NewTempDocumentJs() // 新建临时文档
		case "2":
			model.TempFame1() // 小座屏
		case "3":
			model.TempFame2() // 折屏
		case "4":
			model.TempFame3() // 镂空
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
			if len(info) != 0 {
				fmt.Println(info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的框架类型，请重新输入...\n", tools.ColourString(frameType, ctc.ForegroundGreen))
			}
		}
	}
}
