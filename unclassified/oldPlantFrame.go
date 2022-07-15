package unclassified

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strconv"
)

// 是否打开自动新建文档
func isOpenPs() {
	if viper.GetBool("openPs") { // 是否自动新建ps文档
		// 创建一个协程使用cmd来运行脚本
		dataPath := "resources/jsx/newDocument.jsx"
		cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
		go cmd.Run()
	}
}

// OldFrame2 旧厂左右镂空
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米  旧厂的边框实际厚度是5厘米
func OldFrame2() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [5]string{"\n:: 请输入左右镂空的总宽：", "\n:: 请输入左右镂空的总高：", "\n:: 请输入左镂空的大小：",
		"\n:: 请输入右镂空的大小：", "\n:: 请输入合页数量（若订单无备注请输入“0”）："}

	// 保存尺寸的切片
	saveSizeStr := [5]string{}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架左右镂空", 74) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {

			// 只有前两个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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

		color.Yellow.Printf("\n:: %s：宽 %.2f cm，高 %.2f cm", tempName, width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("%s：宽 %.2f cm，高 %.2f cm\n", tempName, width, height)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), tempName, width, height)

		model.NewDocument(width, height, frameName, true) // 创建ps文档
		go model.FrameSaveDef(frameName)                  // 生成暗号【-1】可以用的另存脚本
		model.IsMaxCanvasExceeded(width, height)          // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// OldFrame3 旧厂中间大两边小
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米
func OldFrame3() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [4]string{"\n:: 请输入左右画布的总宽：", "\n:: 请输入左右画布的总高：",
		"\n:: 请输入单边画布的大小：", "\n:: 请输入合页数量（若订单无备注请输入“0”）："}

	// 保存尺寸的切片
	saveSizeStr := [4]string{}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架左右画布", 74) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前3个需要开启画布模式
			if i < 3 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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

		color.Yellow.Printf("\n:: 左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm", width, hollowOut, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm\n", width, hollowOut, height)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_左右画布_%.0fx%.0f", tools.NowTime(), totalWidth, height)

		model.NewDocument(totalWidth, height, frameName, false)  // 创建ps文档
		generate.LineJs3(width, hollowOut)                       // 生成专属参考线
		go generate.Tailor3(width, height, hollowOut, frameName) // 生成暗号【-1】可以用的另存脚本
		model.IsMaxCanvasExceeded(width, height)                 // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂上下镂空
//先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个横边 空出的中间画面加5厘米
func OldFrame4to1() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [4]string{"\n:: 请输入上下镂空的总宽：", "\n:: 请输入上下镂空的总高：",
		"\n:: 请输入上镂空的大小：", "\n:: 请输入下镂空的大小："}

	// 保存尺寸的切片
	saveSizeStr := [4]string{}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架上下镂空", 74) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前2个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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

		color.Yellow.Printf("\n:: %s：宽 %.2f cm，高 %.2f cm", tempName, width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("%s：宽 %.2f cm，高 %.2f cm\n", tempName, width, height)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), tempName, width, height)

		model.NewDocument(width, height, frameName, true) // 创建ps文档
		go model.FrameSaveDef(frameName)                  // 生成暗号【-1】可以用的另存脚本
		model.IsMaxCanvasExceeded(width, height)          // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 上下座屏
// 上下画布 一般没有合页，上下画布的两边是不扣的 例如 总高：180 上下分别：30+5  中间为：100+5，也就是说边框5都中中间画布扣了
func OldFrame4to2() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 初始化输入提示的切片
	inputPrompt := [4]string{"\n:: 请输入上下画布的总宽：", "\n:: 请输入上下画布的总高：",
		"\n:: 请输入上画布的大小：", "\n:: 请输入下画布的大小："}

	// 保存尺寸的切片
	saveSizeStr := [4]string{}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架上下画布", 74) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前2个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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
		history := fmt.Sprintf("上下画布的总宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("上下画布的总高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("上画布的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("下画布的大小：%s\n", saveSizeStr[3])

		// 强制类型转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		upperHollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		downHollowOut, _ := strconv.ParseFloat(saveSizeStr[3], 64)

		// 声明临时框架名字
		var tempName = "上下画布"

		// 镂空判断
		if upperHollowOut > 0 && downHollowOut == 0 {
			tempName = "上画布"
		}
		if upperHollowOut == 0 && downHollowOut > 0 {
			tempName = "下画布"
		}

		// 进行框架公式计算
		width = width - 10 + reserve
		height = height - 10 + reserve
		if upperHollowOut > 0 {
			height -= upperHollowOut + 5 // 如果有上镂空的话
			// 上下画布要有画布预留
			upperHollowOut += reserve
		}
		if downHollowOut > 0 {
			height -= downHollowOut + 5 // 如果有下镂空的话
			// 上下画布要有画布预留
			downHollowOut += reserve
		}

		// 总高度
		totalHeight := upperHollowOut + downHollowOut + height

		//color.Yellow.Printf("\n:: 左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm", width, hollowOut, height)
		color.Yellow.Printf("\n:: %s：宽 %.2f cm，上高 %.2f cm，中高 %.2f cm，下高 %.2f cm", tempName, width, upperHollowOut, height, downHollowOut)

		//存储已计算的历史记录
		history += fmt.Sprintf("%s：宽 %.2f cm，上高 %.2f cm，中高 %.2f cm，下高 %.2f cm\n", tempName, width, upperHollowOut, height, downHollowOut)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), tempName, width, totalHeight)

		model.NewDocument(width, totalHeight, frameName, false) // 创建ps文档

		// 生成专属参考线
		generate.LineJs4to2(upperHollowOut, height)

		// 生成暗号【-1】可以用的另存脚本
		go generate.Tailor4to2(width, upperHollowOut, height, downHollowOut, tempName, frameName)

		model.IsMaxCanvasExceeded(width, height) // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

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

	// 初始化输入提示的切片
	inputPrompt := [5]string{"\n:: 请输入顶天立地的总宽：", "\n:: 请输入顶天立地的总高：",
		"\n:: 请输入上镂空的大小：", "\n:: 请输入下镂空的大小：", "\n:: 请输入拥有几个贴地或贴顶横杆："}

	// 保存尺寸的数组
	saveSizeStr := [5]string{}

	for {
		tools.ChineseTitle("当前框架顶天立地", 74) // 请注意切图的工厂与框架的选择

		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前2个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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
		var history = fmt.Sprintf("顶天立地的总宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("顶天立地的总高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("上镂空的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("下镂空的大小：%s\n", saveSizeStr[3])
		history += fmt.Sprintf("拥有几个贴地或贴顶横杆：%s\n", saveSizeStr[4])

		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		upperHollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		downHollowOut, _ := strconv.ParseFloat(saveSizeStr[3], 64)
		number, _ := strconv.ParseFloat(saveSizeStr[4], 64)

		width = width - 10 + reserve
		height = height - upperHollowOut - downHollowOut - 10 - number*5 + reserve

		color.Yellow.Printf("\n:: 顶天立地：宽 %.2f cm，高 %.2f cm", width, height)

		//存储已计算的历史记录
		history += fmt.Sprintf("顶天立地：宽 %.2f cm，高 %.2f cm\n", width, height)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_顶天立地_%.0fx%.0f", tools.NowTime(), width, height)

		model.NewDocument(width, height, frameName, true) // 创建ps文档
		go model.FrameSaveDef(frameName)                  // 生成暗号【-1】可以用的另存脚本
		model.IsMaxCanvasExceeded(width, height)          // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

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

	// 初始化输入提示的切片
	inputPrompt := [5]string{"\n:: 请输入折屏单扇的宽：", "\n:: 请输入折屏单扇的高：",
		"\n:: 请输入上镂空的大小：", "\n:: 请输入下镂空的大小：", "\n:: 请输入共拥有几扇："}

	// 保存尺寸的数组
	saveSizeStr := [5]string{}

	for {
		tools.ChineseTitle("当前框架各种折屏", 74) // 请注意切图的工厂与框架的选择
		for i := 0; i < len(saveSizeStr); i++ {
			// 只有前2个需要开启画布模式
			if i < 2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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
		var history = fmt.Sprintf("折屏单扇的宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("折屏单扇的高：%s\n", saveSizeStr[1])
		history += fmt.Sprintf("上镂空的大小：%s\n", saveSizeStr[2])
		history += fmt.Sprintf("下镂空的大小：%s\n", saveSizeStr[3])
		history += fmt.Sprintf("共拥有几扇：%s\n", saveSizeStr[4])

		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)
		upperHollowOut, _ := strconv.ParseFloat(saveSizeStr[2], 64)
		downHollowOut, _ := strconv.ParseFloat(saveSizeStr[3], 64)
		number, _ := strconv.ParseFloat(saveSizeStr[4], 64)

		width = width - 10 + reserve // 单扇的宽
		totalWidth := width * number // 总宽

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

		color.Yellow.Printf("\n:: %s折屏：总宽 %.2f cm，高 %.2f cm", tempName, totalWidth, height)
		//存储已计算的历史记录
		history += fmt.Sprintf("%s折屏：总宽 %.2f cm，高 %.2f cm\n", tempName, totalWidth, height)
		go presenter.History(history) // 写入历史

		//获取当前时间，进行格式化 2006-01-02 15:04:05
		now := tools.NowTime()

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s折屏_%.0fx%.0f", now, tempName, totalWidth, height)
		// 定义单片名字
		singleName := fmt.Sprintf("%s折屏", tempName)

		model.NewDocument(totalWidth, height, frameName, false)           // 创建ps文档
		generate.LineJs6(width, number)                                   // 生成专属参考线
		go generate.Tailor6(width, height, number, frameName, singleName) // 生成暗号【-1】可以用的另存脚本
		model.IsMaxCanvasExceeded(width, height)                          // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//旧厂多座屏
func OldFrame7() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 替换文本的临时函数，第几座屏改成中文数字
	replaceText := func(n int, format string) []string {
		s := make([]string, n)
		for i := 0; i < n; i++ {
			s[i] = fmt.Sprintf(format, tools.Transfer(i+1))
		}
		return s
	}

	// 切片叠加的编码临时函数，宽和高交叉叠加，上下镂空直接追加
	enSliceStacking := func(w, h []string, s ...string) (ret []string) {
		sum := len(w) + len(h) + len(s)

		// 为结果分配内存
		ret = make([]string, sum)

		// 最后镂空
		//ret[sum] = s

		// 计算使用次数
		countW, countH := 0, 0
		// 先把宽和高互相叠加进切片
		for i := 0; i < len(w)+len(h); i++ {
			if i%2 == 0 { // 偶数给宽，0也是偶数
				ret[i] = w[countW]
				countW++
			} else {
				ret[i] = h[countH]
				countH++
			}
		}

		// 得到开始追加的索引
		index := len(w) + len(h)
		// 最后把上下镂空追加到切片
		for i := 0; i < len(s); i++ {
			ret[index] = s[i]
			index++
		}

		return
	}

	// 切片叠加的解码临时函数
	deSliceStacking := func(ret []string) (w, h []string, up, down string) {
		// 剪切上下镂空后的切片长度
		sum := len(ret) - 2

		// 为宽和高的切片分配一次内存
		w = make([]string, sum/2)
		h = make([]string, sum/2)

		// 计算宽和高的使用次数
		countW, countH := 0, 0
		// 宽和高的切片先求出
		for i := 0; i < sum; i++ {
			if i%2 == 0 { // 偶数给宽
				w[countW] = ret[i]
				countW++
			} else {
				h[countH] = ret[i]
				countH++
			}
		}

		// 得到上下镂空的值
		up = ret[sum]
		down = ret[sum+1]
		return
	}

	// 将全部字符串切片转浮点数切片的临时函数，计算一下实际切图尺寸
	parseFloatSlice := func(s []string) (f []float64) {
		// 分配下内存
		f = make([]float64, len(s))

		// 将字符串转浮点数
		for i := 0; i < len(s); i++ {
			size, _ := strconv.ParseFloat(s[i], 64)
			// 计算一下实际切图尺寸，赋值到切片
			f[i] = size - 10 + reserve
		}
		return
	}

	// 计算最大的临时函数
	maxSize := func(s []float64) (max float64) {
		max = s[0]
		for i := 1; i < len(s); i++ {
			if max < s[i] {
				max = s[i]
			}
		}
		return
	}

	// 计算最小的临时函数
	minSize := func(s []float64) (min float64) {
		min = s[0]
		for i := 1; i < len(s); i++ {
			if min > s[i] {
				min = s[i]
			}
		}
		return
	}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架多座屏", 74) // 请注意切图的工厂与框架的选择
		numberStr := input.InputCanvasSize("\n:: 请输入拥有几个座屏：", 1)
		// 一开始就返回直接退出函数
		if numberStr == "-" || numberStr == "--" {
			tools.CallClear() // 清屏
			return
		}
		// 字符串转换成int64后再转int
		number64, _ := strconv.ParseInt(numberStr, 10, 64)
		// 得到具体要切几个座屏
		number := int(number64)

		// 替换宽度和高度文案
		inputWidth := replaceText(number, "\n:: 请输入第%s个座屏的宽：")
		inputHeight := replaceText(number, "\n:: 请输入第%s个座屏的高：")

		// 初始化输入提示的切片汇总
		inputPrompt := enSliceStacking(inputWidth, inputHeight, "\n:: 每个座屏的上镂空均是：", "\n:: 每个座屏的下镂空均是：")
		// 保存尺寸的切片
		saveSizeStr := make([]string, len(inputPrompt))

		// 循环输入尺寸信息
		for i := 0; i < len(saveSizeStr); i++ {
			// 除了最后两个都需要开启画布模式
			if i < len(saveSizeStr)-2 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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

		// 开始解码得到的值
		widthStrSlice, heightStrSlice, upperHollowOutStr, downHollowOutStr := deSliceStacking(saveSizeStr)

		// 将字符串转浮点数
		widthSlice := parseFloatSlice(widthStrSlice)
		heightSlice := parseFloatSlice(heightStrSlice)
		upperHollowOut, _ := strconv.ParseFloat(upperHollowOutStr, 64)
		downHollowOut, _ := strconv.ParseFloat(downHollowOutStr, 64)

		// 如果有上镂空的话
		if upperHollowOut > 0 {
			// 顺序遍历
			for i := 0; i < len(heightSlice); i++ {
				heightSlice[i] -= upperHollowOut + 5
			}
		}

		// 如果有下镂空的话
		if downHollowOut > 0 {
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
		widthMax := maxSize(widthSlice)

		// 遍历出最大的高度
		heightMax := maxSize(heightSlice)

		// 遍历出最小的高度
		heightMin := minSize(heightSlice)

		// 定义历史变量
		var history string
		//存储未计算时的历史记录
		for i := 0; i < len(widthStrSlice); i++ {
			history += fmt.Sprintf("请输入第%s个座屏的宽：%s\n", tools.Transfer(i+1), widthStrSlice[i])
			history += fmt.Sprintf("请输入第%s个座屏的高：%s\n", tools.Transfer(i+1), heightStrSlice[i])
		}
		history += fmt.Sprintf("每个座屏的上镂空均是：%s\n", upperHollowOutStr)
		history += fmt.Sprintf("每个座屏的下镂空均是：%s\n", downHollowOutStr)
		//存储已计算的历史记录
		history += fmt.Sprintf("多座屏：总宽 %.2f cm，高 %.2f cm\n", widthSum, heightMax)
		go presenter.History(history) // 写入历史

		color.Yellow.Printf("\n:: 多座屏：总宽 %.2f cm，高 %.2f cm", widthSum, heightMax)

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s座屏_%.0fx%.0f", tools.NowTime(), tools.Transfer(len(widthSlice)), widthSum, heightMax)

		model.NewDocument(widthSum, heightMax, frameName, false) // 创建ps文档
		generate.LineJs7(widthSlice, heightSlice, heightMax, heightMin)
		go generate.Tailor7(widthSlice, heightSlice, heightMax, frameName) // 生成暗号【-1】可以用的另存脚本// 生成参考线与遮罩层
		model.IsMaxCanvasExceeded(widthMax, heightMax)                     // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

func OldFrame7bk() {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 替换文本的临时函数，第几座屏改成中文数字
	replaceText := func(n int, format string) []string {
		s := make([]string, n)
		for i := 0; i < n; i++ {
			s[i] = fmt.Sprintf(format, tools.Transfer(i+1))
		}
		return s
	}

	// 切片叠加的编码临时函数，宽和高交叉叠加，下镂空直接追加
	enSliceStacking := func(w, h []string, s string) (ret []string) {
		sum := len(w) + len(h)

		// 为结果分配内存
		ret = make([]string, sum+1)

		// 最后镂空
		ret[sum] = s

		// 计算宽和高的使用次数
		countW, countH := 0, 0
		for i := 0; i < sum; i++ {
			if i%2 == 0 { // 偶数给宽
				ret[i] = w[countW]
				countW++
			} else {
				ret[i] = h[countH]
				countH++
			}
		}
		return
	}

	// 切片叠加的解码临时函数
	deSliceStacking := func(ret []string) (w, h []string, s string) {
		sum := len(ret) - 1
		// 最后一个是下镂空
		s = ret[sum]
		// 计算宽和高的使用次数
		countW, countH := 0, 0
		for i := 0; i < sum; i++ {
			if i%2 == 0 { // 偶数给宽
				w = append(w, ret[i])
				countW++
			} else {
				h = append(h, ret[i])
				countH++
			}
		}
		return
	}

	// 将全部字符串切片转浮点数切片的临时函数，计算一下实际切图尺寸
	parseFloatSlice := func(s []string) (f []float64) {
		// 分配下内存
		f = make([]float64, len(s))

		// 将字符串转浮点数
		for i := 0; i < len(s); i++ {
			size, _ := strconv.ParseFloat(s[i], 64)
			// 计算一下实际切图尺寸，赋值到切片
			f[i] = size - 10 + reserve
		}
		return
	}

	// 计算最大的临时函数
	maxSize := func(s []float64) (max float64) {
		max = s[0]
		for i := 1; i < len(s); i++ {
			if max < s[i] {
				max = s[i]
			}
		}
		return
	}

	// 计算最小的临时函数
	minSize := func(s []float64) (min float64) {
		min = s[0]
		for i := 1; i < len(s); i++ {
			if min > s[i] {
				min = s[i]
			}
		}
		return
	}

	// 循环使用此框架
	for {
		tools.ChineseTitle("当前框架多座屏", 74) // 请注意切图的工厂与框架的选择
		numberStr := input.InputCanvasSize("\n:: 请输入拥有几个座屏：", 0)
		// 一开始就返回直接退出函数
		if numberStr == "-" || numberStr == "--" {
			tools.CallClear() // 清屏
			return
		}
		// 字符串转换成int64后再转int
		number64, _ := strconv.ParseInt(numberStr, 10, 64)
		// 得到具体要切几个座屏
		number := int(number64)

		// 替换宽度和高度文案
		inputWidth := replaceText(number, "\n:: 请输入第%s个座屏的宽：")
		inputHeight := replaceText(number, "\n:: 请输入第%s个座屏的高：")

		//fmt.Println(inputWidth)
		//fmt.Println(inputHeight)

		// 初始化输入提示的切片汇总
		inputPrompt := enSliceStacking(inputWidth, inputHeight, "\n:: 每个座屏的下镂空均是：")
		// 保存尺寸的切片
		saveSizeStr := make([]string, len(inputPrompt))

		// 循环输入尺寸信息
		for i := 0; i < len(saveSizeStr); i++ {
			// 除了最后一个都需要开启画布模式
			if i != len(saveSizeStr)-1 {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 0)
			}

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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

		// 开始解码得到的值
		widthStrSlice, heightStrSlice, downHollowOutStr := deSliceStacking(saveSizeStr)

		// 将字符串转浮点数
		widthSlice := parseFloatSlice(widthStrSlice)
		heightSlice := parseFloatSlice(heightStrSlice)

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
		widthMax := maxSize(widthSlice)

		// 遍历出最大的高度
		heightMax := maxSize(heightSlice)

		// 遍历出最小的高度
		heightMin := minSize(heightSlice)

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
		go presenter.History(history) // 写入历史

		color.Yellow.Printf("\n:: 多座屏：总宽 %.2f cm，高 %.2f cm", widthSum, heightMax)

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s座屏_%.0fx%.0f", tools.NowTime(), tools.Transfer(len(widthSlice)), widthSum, heightMax)

		model.NewDocument(widthSum, heightMax, frameName, false) // 创建ps文档
		generate.LineJs7(widthSlice, heightSlice, heightMax, heightMin)
		go generate.Tailor7(widthSlice, heightSlice, heightMax, frameName) // 生成暗号【-1】可以用的另存脚本// 生成参考线与遮罩层
		model.IsMaxCanvasExceeded(widthMax, heightMax)                     // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 补切画布
func OldFrame9() {
	// 初始化输入提示的切片
	inputPrompt := [2]string{"\n:: 请输入补切画布的宽：", "\n:: 请输入补切画布的高："}

	// 保存尺寸的切片
	saveSizeStr := [2]string{}

	for {
		tools.ChineseTitle("当前框架补切画布", 74) // 请注意切图的工厂与框架的选择
		fmt.Println("\n【补切】主要用来补切画布，不减去任何边框尺寸，适合不想手动新建画布时使用！")

		for i := 0; i < len(saveSizeStr); i++ {
			saveSizeStr[i] = input.InputCanvasSize(inputPrompt[i], 6)

			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
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
		var history = fmt.Sprintf("补切画布的宽：%s\n", saveSizeStr[0])
		history += fmt.Sprintf("补切画布的高：%s\n", saveSizeStr[1])

		// 强制转换成浮点数
		width, _ := strconv.ParseFloat(saveSizeStr[0], 64)
		height, _ := strconv.ParseFloat(saveSizeStr[1], 64)

		color.Yellow.Printf("\n【补切】补切画布的切图：宽为 %.2f cm，高为 %.2f cm", width, height)
		//存储已计算的历史记录
		history += fmt.Sprintf("补切画布的切图：宽为 %.2f cm，高为 %.2f cm", width, height)
		go presenter.History(history) // 写入历史

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_补切画布_%.0fx%.0f", tools.NowTime(), width, height)

		go model.FrameSaveDef(frameName)                  // 生成暗号【-1】可以用的另存脚本
		model.NewDocument(width, height, frameName, true) // 创建ps文档
		model.IsMaxCanvasExceeded(width, height)          // 最大画布判断
		isOpenPs()                                        // 是否打开自动新建文档
		if !viper.GetBool("memory") {                     // 是否记忆框架
			break
		}
	}
}
