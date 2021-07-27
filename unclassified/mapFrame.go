package unclassified

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strconv"
)

// 贴图小座屏
func MapFrame1() {
	for {
		tools.ChineseTitle("当前框架常规座屏贴图 ", 74) // 请注意切图的工厂与框架的选择

		widthStr := input.InputCanvasSize("\n:: 请输入常规座屏的宽：", 6)
		if widthStr == "-" {
			break
		}
		heightStr := input.InputCanvasSize("\n:: 请输入常规座屏的高：", 6)
		if heightStr == "-" {
			break
		}
		reserveStr := input.InputCanvasSize("\n:: 请输入要减去的单边框大小(一般为40)：", 0)
		if reserveStr == "-" {
			break
		}

		// 字符串转 int
		width, _ := strconv.Atoi(widthStr)
		height, _ := strconv.Atoi(heightStr)
		reserve, _ := strconv.Atoi(reserveStr)

		// 去掉边框
		width -= reserve * 2
		height -= reserve * 2

		fmt.Printf("\n:: 常规座屏：宽 %d pixels，高 %d pixels", width, height)

		generate.MaxCanvas(float64(width)/10, float64(height)/10)

		generate.NewDocumentForMap(width, height, "常规座屏贴图") // 生成创建ps文档脚本

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

// 贴图折屏
func MapFrame6() {
	for {
		tools.ChineseTitle("当前框架各种折屏贴图", 74) // 请注意切图的工厂与框架的选择

		widthStr := input.InputCanvasSize("\n:: 请输入折屏单扇的宽：", 6)
		if widthStr == "-" {
			break
		}
		heightStr := input.InputCanvasSize("\n:: 请输入折屏单扇的高：", 6)
		if heightStr == "-" {
			break
		}
		upperHollowOutStr := input.InputCanvasSize("\n:: 请输入上镂空的大小：", 0)
		if heightStr == "-" {
			break
		}

		downHollowOutStr := input.InputCanvasSize("\n:: 请输入下镂空的大小：", 0)
		if heightStr == "-" {

			break
		}

		numberStr := input.InputCanvasSize("\n:: 请输入共拥有几扇：", 1)
		if numberStr == "-" {
			break
		}

		reserveStr := input.InputCanvasSize("\n:: 请输入要减去的单边框大小(一般为40)：", 0)
		if reserveStr == "-" {
			break
		}

		// 字符串转 int
		width, _ := strconv.Atoi(widthStr)
		height, _ := strconv.Atoi(heightStr)
		upperHollowOut, _ := strconv.Atoi(upperHollowOutStr)
		downHollowOut, _ := strconv.Atoi(downHollowOutStr)
		number, _ := strconv.Atoi(numberStr)
		reserve, _ := strconv.Atoi(reserveStr)

		// 计算边框
		width -= reserve * 2         // 单扇的宽
		totalWidth := width * number // 总宽
		height -= reserve * 2        // 单扇的高

		if upperHollowOut > 0 { // 如果有上镂空的话
			height -= upperHollowOut + reserve
		}
		if downHollowOut > 0 { // 如果有下镂空的话
			height -= downHollowOut + reserve
		}

		fmt.Printf("\n:: 常规折屏：宽 %d pixels，高 %d pixels", totalWidth, height)
		generate.MaxCanvas(float64(width)/10, float64(height)/10)

		//获取当前时间，进行格式化 2006-01-02 15:04:05
		now := tools.NowTime()

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_折屏贴图_%dx%d", now, totalWidth, height)

		// 定义单片名字
		singleName := fmt.Sprintf("%s_折屏贴图", now)

		generate.NewDocumentForMap(totalWidth, height, frameName)               // 生成创建ps文档脚本
		generate.Line3DMapJs6(width, number)                                    // 生成专属参考线
		go generate.TailorForMap6(width, height, number, frameName, singleName) // 生成暗号【-1】可以用的另存脚本
		if viper.GetBool("openPs") {                                            // 是否自动新建ps文档
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

// 贴图多座屏
func MapFrame7() {

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
	parseIntSlice := func(s []string, reserve int) (f []int) {
		// 分配下内存
		f = make([]int, len(s))

		// 将字符串转int
		for i := 0; i < len(s); i++ {
			size, _ := strconv.Atoi(s[i])
			// 计算一下实际切图尺寸，赋值到切片
			f[i] = size - reserve*2
		}
		return
	}

	// 计算最大的临时函数
	maxSize := func(s []int) (max int) {
		max = s[0]
		for i := 1; i < len(s); i++ {
			if max < s[i] {
				max = s[i]
			}
		}
		return
	}

	//计算最小的临时函数
	minSize := func(s []int) (min int) {
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
		tools.ChineseTitle("当前框架多座屏贴图", 74) // 请注意切图的工厂与框架的选择
		numberStr := input.InputCanvasSize("\n:: 请输入拥有几个座屏：", 0)
		// 一开始就返回直接退出函数
		if numberStr == "-" || numberStr == "--" {
			tools.CallClear() // 清屏
			return
		}

		reserveStr := input.InputCanvasSize("\n:: 请输入要减去的单边框大小(一般为40)：", 0)
		if reserveStr == "-" || numberStr == "--" {
			return
		}

		// 字符串转 int---------------------
		// 得到具体要切几个座屏
		number, _ := strconv.Atoi(numberStr)
		// 减去的边框大小
		reserve, _ := strconv.Atoi(reserveStr)

		// 替换宽度和高度文案
		inputWidth := replaceText(number, "\n:: 请输入第%s个座屏的宽：")
		inputHeight := replaceText(number, "\n:: 请输入第%s个座屏的高：")

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

		// 将字符串转int
		widthSlice := parseIntSlice(widthStrSlice, reserve)
		heightSlice := parseIntSlice(heightStrSlice, reserve)

		downHollowOut, _ := strconv.Atoi(downHollowOutStr)

		if downHollowOut > 0 { // 如果有下镂空的话
			// 顺序遍历
			for i := 0; i < len(heightSlice); i++ {
				heightSlice[i] = heightSlice[i] - (downHollowOut + 5)
			}
		}

		// 遍历出总宽
		var widthSum int
		for i := 0; i < len(widthSlice); i++ {
			widthSum += widthSlice[i]
		}

		// 遍历出最大的宽度
		widthMax := maxSize(widthSlice)

		// 遍历出最大的高度
		heightMax := maxSize(heightSlice)

		// 遍历出最小的高度
		heightMin := minSize(heightSlice)

		color.Yellow.Printf("\n:: 多座屏：总宽 %d pixels，高 %d pixels", widthSum, heightMax)

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_%s座屏贴图_%dx%d", tools.NowTime(), tools.Transfer(len(widthSlice)), widthSum, heightMax)
		generate.NewDocumentForMap(widthSum, heightMax, frameName) // 生成创建ps文档脚本
		generate.Line3DMapJs7(widthSlice, heightSlice, heightMax, heightMin) // 生成参考线和遮罩层
		// TODO: 代修改
		go generate.TailorForMap7(widthSlice, heightSlice, heightMax, frameName) // 生成暗号【-1】可以用的另存脚本// 生成参考线与遮罩层
		generate.MaxCanvas(float64(widthMax)/10, float64(heightMax)/10) // 最大画布判断

		isOpenPs() // 是否打开自动新建文档

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}
