package cli

import (
	"fmt"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
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

// 读取终端输入，返回字符串数字，对暗号进行拦截并执行，对非数字进行拦截重新输入
func inputPro(prompt string) string {
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

		var temp string
		_, _ = fmt.Scanln(&temp)

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

		return temp
	}
}
