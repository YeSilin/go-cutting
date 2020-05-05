// 对输入的字符进行判断
package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/clib"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/tools"
	"regexp"
	"strconv"
	"strings"
)

// 判断是否为数字，并根据指定值提供指定的全局功能
func Input(text string, canvasMode bool) (num string) {
	for {
		// 用户输入提示，获取键盘输入
		fmt.Print(text)

		// 获取光标位置
		x, y := clib.WhereXY()

		_, err := fmt.Scanln(&num) // Scanln 换行也是一次输入结束
		if err != nil {
			if !canvasMode { // 不是画布模式就输出0
				// 重新指定xy位置
				clib.GotoPostion(x-1, y-1)
				fmt.Println("0")
				return "0"
			} else {
				// 重新指定xy位置
				clib.GotoPostion(0, y-2)
				continue
			}
		}

		// 二级返回
		if num == "--" {
			return "--"
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			return "-"
		}

		// 开始指定功能
		switch num {
		case "":
			continue
		case "-1":
			StartCode1()
			continue
		case "-2":
			StartCode2()
			continue
		case "-3":
			StartCode3() // 深度清除源数据
			continue
		case "-4":
			StartCode4() // 工作目录
		case "-5":
		case "-6":
			StartCode6() // 简单清除元数据
			continue
		case "-7":
			StartCode7() // 为当前文档添加黑边
			continue
		case "-8":
			tools.CallClear() // 清屏
			continue
		case "-9":
			StartCode9() // 打开历史记录
			continue
		case "-10":
			StartCode10() // 快捷另存为jpg
			continue
		case "-11":
			OldFrame1()
			continue
		case "-12":
			OldFrame2()
			continue
		case "-13":
			OldFrame3()
			continue
		case "-14":
			OldFrame4()
			continue
		case "-15":
			OldFrame5()
			continue
		case "-16":
			OldFrame6()
			continue
		case "-17":
			OldFrame7()
			continue
		case "-18":
			OldFrame8()
			continue
		case "-19":
			OldFrame9()
			continue
		case "-97":
			generate.ReplaceDetailsPage(viper.GetString("picture")) // 替换详情页
			continue
		case "-98":
			StartCode98()
			continue
		case "-99":
			StartCode99()
			continue
		}

		//strings.Replace(num,".","",0)  // 删掉全部小数点
		// 如果小数点多于一个就循环
		if strings.Count(num, ".") > 1 {
			continue
		}

		// 如果包含中文就转拼音
		if IncludeChinese2(num){
			return strings.Join(ToPinyin(num)," ")
		}


		// 如果不是整数或浮点数就循环
		integer, _ := regexp.MatchString(`^(\-|\+)?\d+(\.\d+)?$`, num)
		if !integer {
			continue
		}

		// 如果输入的数字小于6就循环
		tempNum, _ := strconv.ParseFloat(num, 64)
		if tempNum < 6 && canvasMode {
			// 查找当前提示信息中是否已包含 插入的提示信息
			if !strings.Contains(text, "（尺寸不可小于6厘米）") {
				// 没有就插入
				text = tools.StrRightInsert(text, ColourString("（尺寸不可小于6厘米）", ctc.ForegroundRed), 3)
			}
			continue
		}
		return num
	}
}
