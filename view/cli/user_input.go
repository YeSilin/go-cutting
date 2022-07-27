package cli

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
	"strings"
)

// 读取终端输入，返回字符串
func inputString(prompt string) string {
	// 先打印提示
	fmt.Print(prompt)

	var temp string
	_, _ = fmt.Scanln(&temp)
	return temp
}

// 读取终端输入，返回浮点数，不是浮点数重新输入
func inputFloat(prompt string) float64 {
	// 获取光标位置
	x, y := tools.WhereXY()

	for {
		// 一开始就定位重新输入，如果放在输入完会有bug
		// 重新指定xy位置
		tools.GotoPostion(x, y-1)
		// 打一些空格覆盖之前的内容，只清空两行
		fmt.Println("                                                                                           ")
		fmt.Print("                                                                                           ")
		// 重新指定xy位置
		tools.GotoPostion(x, y-1)

		// 先打印提示
		fmt.Print(prompt)

		var temp float64
		_, err := fmt.Scanln(&temp)
		if err != nil {
			continue
		}
		return temp
	}
}

// 读取终端输入，返回字符串数字，对暗号进行拦截并执行，对非数字进行拦截重新输入，least 限制至少是多大值
func inputPro(prompt string, least int) string {
	// 获取光标位置
	x, y := tools.WhereXY()
	// 用来储存新的提示
	var newPrompt string

	for {
		// 一开始就定位重新输入，如果放在输入完会有bug
		// 重新指定xy位置
		tools.GotoPostion(x, y-1)
		// 打一些空格覆盖之前的内容，只清空两行
		fmt.Println("                                                                                             ")
		fmt.Print("                                                                                             ")
		// 重新指定xy位置
		tools.GotoPostion(x, y-1)

		// 先打印提示，如果没有新提示就打印最开始的

		if newPrompt != "" {
			fmt.Print(newPrompt)
		} else {
			fmt.Print(prompt)
		}

		var temp string
		_, err := fmt.Scanln(&temp)
		// 说明是未输入直接回车，如果输入空内容就当 0 处理
		if err != nil {
			temp = "0"
		}

		// 如果要0直接返回，那必须满足不可小于最小值
		if temp == "0" && least <= 0 {
			// 重新指定xy位置
			tools.GotoPostion(x, y-1)
			fmt.Println(prompt + "0") // 重新打印一遍
			return "0"
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串；如果减号出现在最后一位，就提前返回
		if lastIndex := strings.LastIndex(temp, "-"); lastIndex == len(temp)-1 {
			return "-"
		}

		// 如果是暗号就执行
		if ok, _ := presenter.SelectCommand(temp); ok {
			continue
		}

		// 如果非数字就重新输入
		if !tools.IsNumber(temp) {
			continue
		}

		// 如果输入的数字小于至少值就循环
		if num, _ := strconv.ParseFloat(temp, 64); num < float64(least) {
			newPrompt = fmt.Sprintf("（输入值不可小于%d）", least)
			newPrompt = tools.StrRightInsert(prompt, tools.ColourString(newPrompt, ctc.ForegroundRed), 3)
			continue
		}

		return temp
	}
}
