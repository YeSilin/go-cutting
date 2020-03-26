package model

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model/quickCipher"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strconv"
	"time"
)

// 旧厂框架的选择
func OldFrameChoice() {
	for {
		EnglishTitle("Cutting", 79)
		text := `
【切图】[1]常规座屏                 [2]左右镂空                 [3]左右画布

【切图】[4]上下镂空                 [5]顶天立地                 [6]各种折屏

【切图】[7]多个座屏                 [8]卷帘座屏                 [9]不扣补切`
		fmt.Println(text)

		frameType := Input("\n【切图】请选择上方的边框类型：", false)

		switch frameType {
		case "1":
			tools.CallClear() // 清屏
			OldFrame1()       // 小座屏
		case "2":
			tools.CallClear() // 清屏
			OldFrame2()       // 左右镂空
		case "3":
			tools.CallClear() // 清屏
			OldFrame3()       // 左右画布
		case "4":
			tools.CallClear() // 清屏
			OldFrame4()       // 上下镂空
		case "5":
			tools.CallClear() // 清屏
			OldFrame5()       // 顶天立地
		case "6":
			tools.CallClear() // 清屏
			OldFrame6()       // 常规折屏
		case "7":
			tools.CallClear() // 清屏
			OldFrame7()       // 多座屏
		case "8":
			tools.CallClear() // 清屏
			OldFrame8()       // 卷帘座屏
		case "9":
			tools.CallClear() // 清屏
			OldFrame9()       // 补切画布
		case "-", "--":
			goto FLAG
		default:
			tools.CallClear() // 清屏
			fmt.Printf("\n【错误】输入的 [%s] 不是已知的边框类型，请重新输入！\n", ColourString(frameType, ctc.ForegroundGreen))
			continue
		}
	}
FLAG:
}

// 返回当前时间
func nowTime() (now string) {
	// 获取当前时间，进行格式化 2006-01-02 15:04:05
	return time.Now().Format("060102150405")
}

//旧厂小座屏
//边框是5  扣掉两个边框5+5 然后再加回来5厘米  可以理解为扣5*/
func OldFrame1() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [2]string{"\n【切图】请输入常规座屏的宽：", "\n【切图】请输入常规座屏的高："}

	// 保存尺寸的切片
	saveSizeStr := [2]string{}

	// 循环使用此框架
	for {
		ChineseTitle("当前框架常规座屏", 79) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			saveSizeStr[i] = Input(inputPrompt[i], true)

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				return
			}

			// 第一次就输入返回就退出此框架
			if i == 0 && saveSizeStr[i] == "--" {
				return
			}

			// 退回上一级输入
			if saveSizeStr[i] == "--" {
				i -= 2
			}
		}

		//存储未计算时的历史记录
		history := fmt.Sprintf("常规座屏的宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("常规座屏的高：%s\n", saveSizeStr[1])

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)

		// 进行框架公式计算
		width = width - 10 + reserve
		height = height - 10 + reserve

		// 输出提示
		color.Yellow.Printf("\n【切图】常规座屏：宽 %.2f cm，高 %.2f cm", width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("常规座屏：宽 %.2f cm，高 %.2f cm\n", width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_常规座屏_%.0fx%.0f", nowTime(), width, height)

		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}
		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂左右镂空
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米  旧厂的边框实际厚度是5厘米
func OldFrame2() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [5]string{"\n【切图】请输入左右镂空的总宽：", "\n【切图】请输入左右镂空的总高：", "\n【切图】请输入左镂空的大小：",
		"\n【切图】请输入右镂空的大小：", "\n【切图】请输入合页数量（若订单无备注请输入“0”）："}

	// 保存尺寸的切片
	saveSizeStr := [5]string{}

	// 循环使用此框架
	for {
		ChineseTitle("当前框架左右镂空", 79) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {

			// 只有前两个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = Input(inputPrompt[i], true)
			} else {
				saveSizeStr[i] = Input(inputPrompt[i], false)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				return
			}

			// 第一次就输入返回就退出此框架
			if i == 0 && saveSizeStr[i] == "--" {
				return
			}

			// 退回上一级输入
			if saveSizeStr[i] == "--" {
				i -= 2
			}
		}

		//存储未计算时的历史记录
		history := fmt.Sprintf("左右镂空的总宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("左右镂空的总高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("左镂空的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("右镂空的大小：%s\n", saveSizeStr[3])
		history += fmt.Sprintf("合页数量：%s\n", saveSizeStr[4])

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		leftHollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		rightHollowOut, _ := strconv.ParseFloat(saveSizeStr[3], 64)
		hinges, _ := strconv.ParseFloat(saveSizeStr[4], 64)

		// 声明临时框架名字
		var tempName = "左右镂空"

		// 镂空判断
		if leftHollowOut > 0 && rightHollowOut == 0 {
			tempName = "左镂空"
		}
		if leftHollowOut == 0 && rightHollowOut > 0 {
			tempName = "右镂空"
		}

		// 进行框架公式计算
		if hinges == 0 {
			width = width - 10 + reserve
			if leftHollowOut > 0 {
				width -= leftHollowOut + 5 // 如果有左镂空的话
			}
			if rightHollowOut > 0 {
				width -= rightHollowOut + 5 // 如果有右镂空的话
			}
		} else {
			width = width - (leftHollowOut + rightHollowOut) - 10 + reserve
		}
		height = height - 10 + reserve

		color.Yellow.Printf("\n【切图】%s：宽 %.2f cm，高 %.2f cm", tempName, width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("%s：宽 %.2f cm，高 %.2f cm\n", tempName, width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", nowTime(), tempName, width, height)

		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂中间大两边小
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米
func OldFrame3() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [4]string{"\n【切图】请输入左右画布的总宽：", "\n【切图】请输入左右画布的总高：",
		"\n【切图】请输入单边画布的大小：", "\n【切图】请输入合页数量（若订单无备注请输入“0”）："}

	// 保存尺寸的切片
	saveSizeStr := [4]string{}

	// 循环使用此框架
	for {
		ChineseTitle("当前框架左右画布", 79) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前3个需要开启画布模式
			if i < 3 {
				saveSizeStr[i] = Input(inputPrompt[i], true)
			} else {
				saveSizeStr[i] = Input(inputPrompt[i], false)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				return
			}

			// 第一次就输入返回就退出此框架
			if i == 0 && saveSizeStr[i] == "--" {
				return
			}

			// 退回上一级输入
			if saveSizeStr[i] == "--" {
				i -= 2
			}
		}

		//存储未计算时的历史记录
		history := fmt.Sprintf("左右画布的总宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("左右画布的总高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("单边画布的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("合页数量：%s\n", saveSizeStr[3])

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		hollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		hinges, _ := strconv.ParseFloat(saveSizeStr[3], 64)

		if hinges == 0 {
			width = width - hollowOut*2 - 4*5 + reserve
			hollowOut += reserve
		} else {
			width = width - hollowOut*2 - hinges*5 + reserve
			hollowOut = hollowOut - 10 + reserve
		}

		totalWidth := width + hollowOut*2
		height = height - 10 + reserve

		color.Yellow.Printf("\n【切图】左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm", width, hollowOut, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm\n", width, hollowOut, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_左右画布_%.0fx%.0f", nowTime(), totalWidth, height)

		generate.NewDocument(totalWidth, height, frameName, false) // 创建ps文档
		generate.LineJs3(width, hollowOut)                           // 生成专属参考线
		go generate.Tailor3(width, height, hollowOut, frameName)     // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                               // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂上下镂空
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个横边 空出的中间画面加5厘米
func OldFrame4() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [4]string{"\n【切图】请输入上下镂空的总宽：", "\n【切图】请输入上下镂空的总高：",
		"\n【切图】请输入上镂空的大小：", "\n【切图】请输入下镂空的大小："}

	// 保存尺寸的切片
	saveSizeStr := [4]string{}

	// 循环使用此框架
	for {
		ChineseTitle("当前框架上下镂空", 79) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前2个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = Input(inputPrompt[i], true)
			} else {
				saveSizeStr[i] = Input(inputPrompt[i], false)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				return
			}

			// 第一次就输入返回就退出此框架
			if i == 0 && saveSizeStr[i] == "--" {
				return
			}

			// 退回上一级输入
			if saveSizeStr[i] == "--" {
				i -= 2
			}
		}

		//存储未计算时的历史记录
		history := fmt.Sprintf("上下镂空的总宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("上下镂空的总高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("上镂空的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("下镂空的大小：%s\n", saveSizeStr[3])

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		upperHollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		downHollowOut, _ := strconv.ParseFloat(saveSizeStr[3], 64)

		// 声明临时框架名字
		var tempName = "上下镂空"

		// 镂空判断
		if upperHollowOut > 0 && downHollowOut == 0 {
			tempName = "上镂空"
		}
		if upperHollowOut == 0 && downHollowOut > 0 {
			tempName = "下镂空"
		}

		// 进行框架公式计算
		width = width - 10 + reserve
		height = height - 10 + reserve
		if upperHollowOut > 0 {
			height -= upperHollowOut + 5 // 如果有上镂空的话
		}
		if downHollowOut > 0 {
			height -= downHollowOut + 5 // 如果有下镂空的话
		}

		color.Yellow.Printf("\n【切图】%s：宽 %.2f cm，高 %.2f cm", tempName, width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("%s：宽 %.2f cm，高 %.2f cm\n", tempName, width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", nowTime(), tempName, width, height)

		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂顶天立地
//扣掉镂空部分 再扣5
func OldFrame5() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	for {
	FLAG1:
		ChineseTitle("当前框架顶天立地", 79) // 请注意切图的工厂与框架的选择
		widthStr := Input("\n【切图】请输入顶天立地的总宽：", true)
		if widthStr == "-" || widthStr == "--" {
			break
		}
	FLAG2:
		heightStr := Input("\n【切图】请输入顶天立地的总高：", true)
		if heightStr == "-" {
			break
		}
		// 返回上一次输入
		if heightStr == "--" {
			goto FLAG1
		}
	FLAG3:
		upperHollowOutStr := Input("\n【切图】请输入上镂空的大小：", false)
		if upperHollowOutStr == "-" {
			break
		}
		// 返回上一次输入
		if upperHollowOutStr == "--" {
			goto FLAG2
		}
	FLAG4:
		downHollowOutStr := Input("\n【切图】请输入下镂空的大小：", false)
		if downHollowOutStr == "-" {
			break
		}
		// 返回上一次输入
		if downHollowOutStr == "--" {
			goto FLAG3
		}

		numberStr := Input("\n【切图】请输入拥有几个贴地或贴顶横杆：", false)
		if numberStr == "-" {
			break
		}
		// 返回上一次输入
		if numberStr == "--" {
			goto FLAG4
		}
		//存储未计算时的历史记录
		var history = fmt.Sprintf("顶天立地的总宽：%s\n", widthStr)
		history += fmt.Sprintf("顶天立地的总高：%s\n", heightStr)
		history += fmt.Sprintf("上镂空的大小：%s\n", upperHollowOutStr)
		history += fmt.Sprintf("下镂空的大小：%s\n", downHollowOutStr)
		history += fmt.Sprintf("拥有几个贴地或贴顶横杆：%s\n", numberStr)

		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)
		upperHollowOut, _ := strconv.ParseFloat(upperHollowOutStr, 64)
		downHollowOut, _ := strconv.ParseFloat(downHollowOutStr, 64)
		number, _ := strconv.ParseFloat(numberStr, 64)

		width = width - 10 + reserve
		height = height - upperHollowOut - downHollowOut - 10 - number*5 + reserve

		color.Yellow.Printf("\n【切图】顶天立地：宽 %.2f cm，高 %.2f cm", width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("顶天立地：宽 %.2f cm，高 %.2f cm\n", width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_顶天立地_%.0fx%.0f", nowTime(), width, height)

		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 旧厂常规折屏
//扣5
func OldFrame6() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	for {
	FLAG1:
		ChineseTitle("当前框架常规折屏", 79) // 请注意切图的工厂与框架的选择
		widthStr := Input("\n【切图】请输入折屏单扇的宽：", true)
		if widthStr == "-" || widthStr == "--" {
			break
		}
	FLAG2:
		heightStr := Input("\n【切图】请输入折屏单扇的高：", true)
		if heightStr == "-" {
			break
		}
		// 返回上一次输入
		if heightStr == "--" {
			goto FLAG1
		}
	FLAG3:
		upperHollowOutStr := Input("\n【切图】请输入上镂空的大小：", false)
		if upperHollowOutStr == "-" {
			break
		}
		// 返回上一次输入
		if upperHollowOutStr == "--" {
			goto FLAG2
		}
	FLAG4:
		downHollowOutStr := Input("\n【切图】请输入下镂空的大小：", false)
		if downHollowOutStr == "-" {
			break
		}
		// 返回上一次输入
		if downHollowOutStr == "--" {
			goto FLAG3
		}

		numberStr := Input("\n【切图】请输入共拥有几扇：", false)
		if numberStr == "-" {
			break
		}
		// 返回上一次输入
		if numberStr == "--" {
			goto FLAG4
		}

		//存储未计算时的历史记录
		var history = fmt.Sprintf("折屏单扇的宽：%s\n", widthStr)
		history += fmt.Sprintf("折屏单扇的高：%s\n", heightStr)
		history += fmt.Sprintf("上镂空的大小：%s\n", upperHollowOutStr)
		history += fmt.Sprintf("下镂空的大小：%s\n", downHollowOutStr)
		history += fmt.Sprintf("共拥有几扇：%s\n", numberStr)

		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)
		number, _ := strconv.ParseFloat(numberStr, 64)
		upperHollowOut, _ := strconv.ParseFloat(upperHollowOutStr, 64)
		downHollowOut, _ := strconv.ParseFloat(downHollowOutStr, 64)

		width = width - 10 + reserve         // 单扇的宽
		totalWidth := width * number         // 总宽
		height = height - 10 + (reserve + 1) // 单扇的高  折屏高要多预留1厘米

		// 声明临时框架名字
		var tempName string
		// 判断镂空
		switch {
		case upperHollowOut > 0 && downHollowOut > 0: // 如果有上镂空 下镂空的话
			height -= upperHollowOut + 5
			height -= downHollowOut + 5
			tempName = "上下镂空"
		case upperHollowOut > 0: // 如果有上镂空的话
			height -= upperHollowOut + 5
			tempName = "上镂空"
		case downHollowOut > 0: // 如果有下镂空的话
			height -= downHollowOut + 5
			tempName = "下镂空"
		default:
			tempName = "常规"
		}

		color.Yellow.Printf("\n【切图】%s折屏：总宽 %.2f cm，高 %.2f cm", tempName, totalWidth, height)
		//存储已计算的历史记录
		history += fmt.Sprintf("%s折屏：总宽 %.2f cm，高 %.2f cm\n", tempName, totalWidth, height)
		go quickCipher.History(history) // 写入历史

		//获取当前时间，进行格式化 2006-01-02 15:04:05
		now := nowTime()

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s折屏_%.0fx%.0f", now, tempName, totalWidth, height)
		// 定义单片名字
		singleName := fmt.Sprintf("%s折屏", tempName)

		generate.NewDocument(totalWidth, height, frameName, false)      // 创建ps文档
		generate.LineJs6(width, number)                                   // 生成专属参考线
		go generate.Tailor6(width, height, number, frameName, singleName) // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                                    // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂多座屏
func OldFrame7() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	for {
	UP1: // 主要给宽返回
		ChineseTitle("当前框架多座屏", 79) // 请注意切图的工厂与框架的选择
		numberStr := Input("\n【切图】请输入拥有几个座屏：", false)
		if numberStr == "-" {
			break
		}

		// 字符串转换成int64后再转int
		number64, _ := strconv.ParseInt(numberStr, 10, 64)
		number := int(number64)

		// 定义一个宽度切片，长度为number
		widthSlice := make([]float64, number)
		// 定义一个高度切片，长度为number
		heightSlice := make([]float64, number)

		// 定义一个宽度切片，长度为number
		widthStrSlice := make([]string, number)
		// 定义一个高度切片，长度为number
		heightStrSlice := make([]string, number)

		// 定义镂空是否返回
		var lr = false
	UP4: // 主要给镂空返回

		for i := 0; i < number; i++ {
			// 提前声明临时宽
			var widthTemp string

			// 如果镂空要返回
			if lr {
				i = number - 1
				lr = false // 用完返回要复位
				goto UP3
			}

		UP2: // 主要给高返回
			// 可以用 Sprintf 来将格式化后的字符串赋值给一个变量
			widthTemp = Input(fmt.Sprintf("\n【切图】请输入第%s个座屏的宽：", tools.Transfer(i+1)), true)
			if widthTemp == "-" {
				goto FLAG // 跳转到函数结束
			}
			if widthTemp == "--" {
				if i > 0 {
					i--
					goto UP3
				}
				goto UP1
			}

			// 赋值到切片
			widthStrSlice[i] = widthTemp
		UP3: // 主要给宽返回
			// 开始接收高度至列表
			heightTemp := Input(fmt.Sprintf("\n【切图】请输入第%s个座屏的高：", tools.Transfer(i+1)), true)
			if heightTemp == "-" {
				goto FLAG // 跳转到函数结束
			}
			if heightTemp == "--" {
				goto UP2
			}

			// 赋值到切片
			heightStrSlice[i] = heightTemp
		}

		downHollowOutStr := Input("\n【切图】每个座屏的下镂空均是：", false)
		if downHollowOutStr == "-" {
			break
		}
		if downHollowOutStr == "--" {
			lr = true // 更新镂空要返回
			goto UP4
		}

		// 将字符串转浮点数
		for i := 0; i < len(widthStrSlice); i++ {
			width, _ := strconv.ParseFloat(widthStrSlice[i], 64)
			// 计算一下实际切图尺寸，赋值到切片
			widthSlice[i] = width - 10 + reserve
		}
		for i := 0; i < len(heightStrSlice); i++ {
			height, _ := strconv.ParseFloat(heightStrSlice[i], 64)
			// 计算一下实际切图尺寸，赋值到切片
			heightSlice[i] = height - 10 + reserve
		}

		downHollowOut, _ := strconv.ParseFloat(downHollowOutStr, 64)

		if downHollowOut > 0 { // 如果有下镂空的话
			// 顺序遍历
			for i := 0; i < len(heightSlice); i++ {
				heightSlice[i] = heightSlice[i] - (downHollowOut + 5)
			}

		}

		// 遍历出总宽
		var widthSum float64
		for i := 0; i < len(widthSlice); i++ {
			widthSum += widthSlice[i]
		}

		// 遍历出最大的宽度
		var widthMax = widthSlice[0]
		for i := 1; i < len(widthSlice); i++ {
			if widthMax < widthSlice[i] {
				widthMax = widthSlice[i]
			}
		}

		// 遍历出最大的高度
		var heightMax = heightSlice[0]
		for i := 1; i < len(heightSlice); i++ {
			if heightMax < heightSlice[i] {
				heightMax = heightSlice[i]
			}
		}

		// 遍历出最小的高度
		var heightMin = heightSlice[0]
		for i := 1; i < len(heightSlice); i++ {
			if heightMin > heightSlice[i] {
				heightMin = heightSlice[i]
			}
		}

		// 定义历史变量
		var history string
		//存储未计算时的历史记录
		for i := 0; i < len(widthStrSlice); i++ {
			history += fmt.Sprintf("请输入第%s个座屏的宽：%s\n", tools.Transfer(i+1), widthStrSlice[i])
			history += fmt.Sprintf("请输入第%s个座屏的高：%s\n", tools.Transfer(i+1), heightStrSlice[i])
		}
		history += fmt.Sprintf("每个座屏的下镂空均是：%s\n", downHollowOutStr)
		//存储已计算的历史记录
		history += fmt.Sprintf("多座屏：总宽 %.2f cm，高 %.2f cm\n", widthSum, heightMax)
		go quickCipher.History(history) // 写入历史

		color.Yellow.Printf("\n【切图】多座屏：总宽 %.2f cm，高 %.2f cm", widthSum, heightMax)

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s座屏_%.0fx%.0f", nowTime(), tools.Transfer(len(widthSlice)), widthSum, heightMax)

		generate.NewDocument(widthSum, heightMax, frameName, false) // 创建ps文档
		generate.LineJs7(widthSlice, heightSlice, heightMax, heightMin)
		go generate.Tailor7(widthSlice, heightSlice, heightMax, frameName) // 生成暗号【-1】可以用的另存脚本// 生成参考线与遮罩层
		generate.MaxCanvas(widthMax, heightMax)                               // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}

	// 标签
FLAG:
}

// 卷帘座屏
// 宽-20   高上下各镂空15 长度要预留10
func OldFrame8() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	for {
	FLAG:
		ChineseTitle("当前框架卷帘座屏", 79) // 请注意切图的工厂与框架的选择
		widthStr := Input("\n【切图】请输入卷帘座屏的宽：", true)
		if widthStr == "-" || widthStr == "--" {
			break
		}

		heightStr := Input("\n【切图】请输入卷帘座屏的高：", true)
		if heightStr == "-" {
			break
		}
		// 返回上一次输入
		if heightStr == "--" {
			goto FLAG
		}

		//存储未计算时的历史记录
		var history = fmt.Sprintf("卷帘座屏的宽：%s\n", widthStr)
		history += fmt.Sprintf("卷帘座屏的高：%s\n", heightStr)

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)

		// 由于卷帘座屏左右两边的画布没有被嵌套，因此不需要计算预留
		width = width - 20
		height = height - 40 + reserve + 5 // 卷帘预留要而外 +5

		color.Yellow.Printf("\n【切图】卷帘座屏：宽 %.2f cm，高 %.2f cm", width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("卷帘座屏：宽 %.2f cm，高 %.2f cm\n", width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_卷帘座屏_%.0fx%.0f", nowTime(), width, height)

		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}
		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 补切画布
func OldFrame9() {
	for {
	FLAG:
		ChineseTitle("当前框架补切画布", 79) // 请注意切图的工厂与框架的选择
		fmt.Println("\n【补切】主要用来补切画布，不减去任何边框尺寸，适合不想手动新建画布时使用！")
		widthStr := Input("\n【补切】请输入补切画布的宽：", false)
		if widthStr == "-" || widthStr == "--" {
			break
		}

		heightStr := Input("\n【补切】请输入补切画布的高：", false)
		if heightStr == "-" {
			break
		}
		// 返回上一次输入
		if heightStr == "--" {
			goto FLAG
		}

		//存储未计算时的历史记录
		var history = fmt.Sprintf("补切画布的宽：%s\n", widthStr)
		history += fmt.Sprintf("补切画布的高：%s\n", heightStr)

		// 强制转换成浮点数
		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)

		color.Yellow.Printf("\n【补切】补切画布的切图：宽为 %.2f cm，高为 %.2f cm", width, height)
		//存储已计算的历史记录
		history += fmt.Sprintf("补切画布的切图：宽为 %.2f cm，高为 %.2f cm", width, height)
		go quickCipher.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_补切画布_%.0fx%.0f", nowTime(), width, height)

		go generate.Tailor(frameName)                         // 生成暗号【-1】可以用的另存脚本
		generate.NewDocument(width, height, frameName, true) // 创建ps文档
		generate.MaxCanvas(width, height)                         // 最大画布判断

		if viper.GetBool("openPs") { // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}
