// Package input 对输入的字符进行判断
package input

import (
	"bufio"
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/presenter"
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
		presenter.Command1()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在调用快捷裁剪..."
	case "-2":
		presenter.Command2()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在重建新文档..."
	case "-3":
		presenter.Command3() // 深度清除源数据
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在深度清理PSD..."
	case "-4":
		return false, ""
	case "-5":
		presenter.Command5() // 复制并关闭其他文档
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在复制并关闭其他文档..."
	case "-6":
		presenter.Command6() // 简单清除元数据
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在快速清理PSD..."
	case "-7":
		presenter.Command7() // 为当前文档添加黑边
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在为当前文档添加黑边..."
	case "-8":
		tools.CallClear() // 清屏
		return true, ""
	case "-9":
		presenter.Command9() // 打开历史记录
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在打开切图历史..."
	case "-10":
		presenter.Command10() // 快捷另存为jpg
		return true, ""
	case "-11":
		presenter.Command11() // 快捷另存全部打开的文件
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在另存全部打开的文件..."
	case "-12":
		presenter.Command12() // 快捷保存并关闭全部文档
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在保存并关闭全部文档..."
	case "-41":
		presenter.Command41() // 快捷保存并关闭全部文档
		return true, "\n:: 检测到输入的内容为隐藏暗号，已打开套图文件夹..."
	case "-42":
		presenter.Command42()
		return true, "\n:: 检测到输入的内容为隐藏暗号，已执行随机重命名..."
	case "-48":
		presenter.Command48()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在替换详情页DP智能对象..."
	case "-49":
		presenter.Command49()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在导出为Web所用格式..."
	case "-99":
		presenter.Command99()
		return true, "\n:: 检测到输入的内容为隐藏暗号，正在后台激活Win10系统..."
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

// InputMenuSelection 支持暗号的获取用户输入的内容
func InputMenuSelection(tips string) (num, info string) {
	for {
		// 获取用户输入
		num = readTerminalInput(tips)

		// 如果输入空内容
		if num == "" {
			return num, "\n:: 输入的内容为空，请重新输入..."
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			//fmt.Printf("%#v", num)
			return "-", ""
		}

		//// 如果包含中文就转拼音
		//if tools.IncludeChinese2(num) {
		//	return num, fmt.Sprintf("\n:: 检测到输入的 [%s] 为中文，转换成拼音为 [%s]",
		//		tools.ColourString(num, ctc.ForegroundGreen),
		//		tools.ColourString(strings.Join(tools.ToPinyin(num), " "), ctc.ForegroundGreen))
		//}

		// 如果是暗号
		if ok, info := runCode(num); ok {
			return num, info
		}

		return num, ""
	}
}

// InputCanvasSize 只返回属于数字的字符串，并且支持暗号，用作画布输入
func InputCanvasSize(tips string, minimum int) (num string) {
	// 回到一开始光标的位置重新输入
	refreshRow := func(x, y int) {
		/// 重新指定xy位置，覆盖之前的输入
		tools.GotoPostion(0, y-1)

		// 打一些空格覆盖之前的内容
		fmt.Print("                                                                                           ")

		// 重新指定xy位置，覆盖之前的输入
		tools.GotoPostion(x, y-2)
	}
	for {
		// 用户输入提示，获取键盘输入
		fmt.Print(tips)

		// 获取光标位置
		x, y := tools.WhereXY()

		// 换行也是一次输入结束
		_, err := fmt.Scanln(&num)
		// 如果输入空内容就当 0 处理，前提是运行的最小值小于或等于0，
		if err != nil && minimum <= 0 {

			// 重新指定xy位置
			tools.GotoPostion(x-1, y-1)
			fmt.Println("0")
			return "0"
		}
		// 有错误，但又小于最小值，忽略重来
		if err != nil {
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 二级返回
		if num == "--" {
			// 覆盖之前的信息
			refreshRow(x, y)
			// 覆盖之前的信息
			refreshRow(x, y-2)
			return "--"
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			return "-"
		}

		// 如果是暗号
		if ok, _ := runCode(num); ok {
			// 覆盖之前的信息
			refreshRow(x, y)

			continue
		}

		// 如果小数点多于一个就循环
		if strings.Count(num, ".") > 1 {
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果不是整数或浮点数就循环
		integer, _ := regexp.MatchString(`^(\-|\+)?\d+(\.\d+)?$`, num)
		if !integer {
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}

		// 如果输入的数字小于6就循环，其实是小于最小值
		tempNum, _ := strconv.ParseFloat(num, 64)
		if tempNum < float64(minimum) {
			// 查找当前提示信息中是否已包含 插入的提示信息
			if !strings.Contains(tips, "（尺寸不可小于6厘米）") {
				// 没有就插入
				tips = tools.StrRightInsert(tips, tools.ColourString("（尺寸不可小于6厘米）", ctc.ForegroundRed), 3)
			}
			// 覆盖之前的信息
			refreshRow(x, y)
			continue
		}
		return
	}
}
