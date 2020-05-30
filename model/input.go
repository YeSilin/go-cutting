// 对输入的字符进行判断
package model

import (
	"bufio"
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/clib"
	"github.com/yesilin/go-cutting/code"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 运行暗号
func runCode(num string) (ok bool, info string) {
	// 开始指定功能
	switch num {
	case "-1":
		code.StartCode1()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在调用快捷裁剪..."
	case "-2":
		code.StartCode2()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在重建新文档..."
	case "-3":
		code.StartCode3() // 深度清除源数据
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在深度清理PSD..."
	case "-4":
		code.StartCode4() // 工作目录
	case "-5":
		return true, "\n:: 检测到输入的内容为隐藏暗号，但是此暗号未指定功能..."
	case "-6":
		code.StartCode6() // 简单清除元数据
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在快速清理PSD..."
	case "-7":
		code.StartCode7() // 为当前文档添加黑边
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在为当前文档添加黑边..."
	case "-8":
		tools.CallClear() // 清屏
		return true, ""
	case "-9":
		code.StartCode9() // 打开历史记录
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在打开切图历史..."
	case "-10":
		code.StartCode10() // 快捷另存为jpg
		return true, ""
	case "-97":
		code.StartCode97()
		return true, ""
	case "-98":
		code.StartCode98()
		return true, ""
	case "-99":
		code.StartCode99()
		return true, ""
	}
	return false, ""
}

// 从终端读取一行输入
func readTerminalInput(tips string) (input string) {
	// 用户输入提示，获取键盘输入
	fmt.Print(tips)

	// 把终端当文件读取
	inputReader := bufio.NewReader(os.Stdin)
	// 读到换行才算结束
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	// 删除换行符号
	input = strings.ReplaceAll(input, "\n", "")

	// 删除前后端所有空白
	input = strings.TrimSpace(input)
	//fmt.Printf("%#v",input)
	return
}

// 回到一开始光标的位置重新输入
func refreshRow(x, y int) {
	// 重新指定xy位置，覆盖之前的输入
	clib.GotoPostion(x-1, y)
	// 打一些空格覆盖之前的内容
	fmt.Print("                                                                                           ")

	// 重新指定xy位置，覆盖之前的输入
	clib.GotoPostion(x-1, y-1)
}

// 判断是否为数字，并根据指定值提供指定的全局功能，是否输入的是画布，是否支持 cls 清屏
func Input(tips string, canvasMode, cls bool) (num, info string) {
	for {
		// 获取还没开始打印提示信息之前的光标位置
		x, y := clib.WhereXY()

		// 获取用户输入
		num = readTerminalInput(tips)

		// 如果输入空内容
		if num == "" {
			if cls {
				// 返回一个清屏命令
				return "cls", ""
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 二级返回
		if num == "--" {

			return "--", ""
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			//fmt.Printf("%#v", num)
			return "-", ""
		}

		// 如果包含中文就转拼音
		if tools.IncludeChinese2(num) {
			if cls {
				// 返回拼音
				return strings.Join(tools.ToPinyin(num), " "), ""
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果是暗号
		if ok, info := runCode(num); ok {
			if cls {

				// 返回一个清屏命令
				return "cls", info
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果小数点多于一个就循环
		if strings.Count(num, ".") > 1 {
			if cls {
				// 返回一个清屏命令
				return "cls", ""
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果不是整数或浮点数就循环
		integer, _ := regexp.MatchString(`^(\-|\+)?\d+(\.\d+)?$`, num)
		if !integer {
			if cls {
				// 返回一个清屏命令
				return "cls", ""
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果输入的数字小于6就循环
		tempNum, _ := strconv.ParseFloat(num, 64)
		if tempNum < 6 && canvasMode {
			// 查找当前提示信息中是否已包含 插入的提示信息
			if !strings.Contains(tips, "（尺寸不可小于6厘米）") {
				// 没有就插入
				tips = tools.StrRightInsert(tips, tools.ColourString("（尺寸不可小于6厘米）", ctc.ForegroundRed), 3)
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}
		return num, ""
	}
}

// 支持暗号的获取用户输入的数字
func InputMenuSelection(tips string) (num, info string) {
	for {
		// 获取用户输入
		num = readTerminalInput(tips)

		// 如果输入空内容
		if num == "" {
			return num, "\n:: 输入的内容为空，请重新输入，若输入中文可快捷转换为拼音..."
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			//fmt.Printf("%#v", num)
			return "-", ""
		}

		// 如果包含中文就转拼音
		if tools.IncludeChinese2(num) {
			return num, fmt.Sprintf("\n:: 检测到输入的 [%s] 为中文，转换成拼音为 [%s]",
				tools.ColourString(num, ctc.ForegroundGreen),
				tools.ColourString(strings.Join(tools.ToPinyin(num), " "), ctc.ForegroundGreen))
		}

		// 如果是暗号
		if ok, info := runCode(num); ok {
			return num, info
		}

		return num, ""
	}
}
