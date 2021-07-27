package view

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/unclassified"
	"github.com/yesilin/go-cutting/tools"
)

// 上下座屏的选择
func (v *CliView) upperAndLowerScreenChoice(){
OuterLoop:
	for {
		tools.EnglishTitle("Upper And Lower Screen Choice", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]上下镂空.                  [2]上下画布.                  [3]暂未开发`
		fmt.Println(text)

		v.key, v.info = input.InputMenuSelection("\n:: 请选择上方的边框类型：")
		tools.CallClear() // 清屏
		switch v.key {
		case "1":
			unclassified.OldFrame4to1()   // 上下镂空
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "2":
			unclassified.OldFrame4to2()   // 上下画布
			if !viper.GetBool("memory") { // 是否记忆框架
				return
			}
		case "-":
			break OuterLoop
		default:
			if len(v.info) != 0 {
				fmt.Println(v.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的边框类型，请重新输入...\n", tools.ColourString(v.key, ctc.ForegroundGreen))
			}
		}
	}

}



// 旧厂框架的选择
func (v *CliView) oldFrameChoice() {
OuterLoop:
	for {
		tools.EnglishTitle("Cutting", 74)
		text := `
:: 下方所有框架的切图单位均是厘米，支持使用小数点来表示毫米，但是意义不大！

   [1]常规座屏.                  [2]左右镂空.                  [3]左右画布.

   [4]上下座屏.                  [5]顶天立地.                  [6]各种折屏.

   [7]多个座屏.                  [8]卷帘座屏.                  [9]不扣补切.`
		fmt.Println(text)

		v.key, v.info = input.InputMenuSelection("\n:: 请选择上方的边框类型：")
		tools.CallClear() // 清屏
		switch v.key {
		case "1":
			unclassified.OldFrame1() // 小座屏
		case "2":
			unclassified.OldFrame2() // 左右镂空
		case "3":
			unclassified.OldFrame3() // 左右画布
		case "4":
			v.upperAndLowerScreenChoice() // 上下座屏框架选择
		case "5":
			unclassified.OldFrame5() // 顶天立地
		case "6":
			unclassified.OldFrame6() // 常规折屏
		case "7":
			unclassified.OldFrame7() // 多座屏
		case "8":
			unclassified.OldFrame8() // 卷帘座屏
		case "9":
			unclassified.OldFrame9() // 补切画布
		case "-":
			break OuterLoop
		default:
			if len(v.info) != 0 {
				fmt.Println(v.info)
			} else {
				fmt.Printf("\n:: 输入的 [%s] 不是已知的边框类型，请重新输入...\n", tools.ColourString(v.key, ctc.ForegroundGreen))
			}
		}
	}
}


