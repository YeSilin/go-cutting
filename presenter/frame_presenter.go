package presenter

// 负责运算从视图收集到的参数
import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
)

// FramePresenter1 对常规座屏进行处理
func FramePresenter1(widthStr, heightStr string) (width, height float64) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算
	width = width - border*2 + reserve
	height = height - border*2 + reserve

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_常规座屏_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true) // 创建ps文档

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter2 对左右镂空进行处理
func FramePresenter2(widthStr, heightStr, leftHollowStr, rightHollowStr, hingeStr string) (width, height float64, frameType string) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	leftHollow, _ := strconv.ParseFloat(leftHollowStr, 64)
	rightHollow, _ := strconv.ParseFloat(rightHollowStr, 64)
	hinge, _ := strconv.ParseFloat(hingeStr, 64)

	// 进行框架公式计算
	if hinge == 0 {
		width = width - border*2 + reserve
		if leftHollow > 0 {
			width -= leftHollow + border // 如果有左镂空的话
		}
		if rightHollow > 0 {
			width -= rightHollow + border // 如果有右镂空的话
		}
	} else {
		width = width - (leftHollow + rightHollow) - border*2 + reserve
	}
	height = height - border*2 + reserve

	// 求出框架类型
	if leftHollow > 0 {
		frameType += "左"
	}
	if rightHollow > 0 {
		frameType += "右"
	}
	if frameType == "" {
		frameType += "无"
	}
	frameType += "镂空"

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), frameType, width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter3 对左右画布进行处理
func FramePresenter3(widthStr, heightStr, hollowStr, hingeStr string) (width, height, hollow float64) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	hollow, _ = strconv.ParseFloat(hollowStr, 64)
	hinge, _ := strconv.ParseFloat(hingeStr, 64)

	// 进行框架公式计算
	if hinge == 0 {
		width = width - hollow*2 - border*4 + reserve
		hollow += reserve
	} else {
		width = width - hollow*2 - hinge*border + reserve
		hollow = hollow - border*2 + reserve
	}
	totalWidth := width + hollow*2
	height = height - border*2 + reserve

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_左右画布_%.0fx%.0f", tools.NowTime(), totalWidth, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(totalWidth, height, frameName, false)

	// 追加专属的切图参考线
	model.FrameGuide3(width, hollow)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave3(width, height, hollow, frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter4to1 对上下镂空进行处理
func FramePresenter4to1(widthStr, heightStr, upHollowStr, downHollowStr, hingeStr string) (width, height float64, frameType string) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	upHollow, _ := strconv.ParseFloat(upHollowStr, 64)
	downHollow, _ := strconv.ParseFloat(downHollowStr, 64)
	hinge, _ := strconv.ParseFloat(hingeStr, 64)

	// 进行框架公式计算
	if hinge == 0 {
		height = height - border*2 + reserve
		if upHollow > 0 {
			height -= upHollow + border // 如果有上镂空的话
		}
		if downHollow > 0 {
			height -= downHollow + border // 如果有下镂空的话
		}
	} else {
		height = height - (upHollow + downHollow) - border*2 + reserve
	}
	width = width - border*2 + reserve

	// 求出框架类型
	if upHollow > 0 {
		frameType += "上"
	}
	if downHollow > 0 {
		frameType += "下"
	}
	if frameType == "" {
		frameType += "无"
	}
	frameType += "镂空"

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), frameType, width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter4to2 对上下画布进行处理
func FramePresenter4to2(widthStr, heightStr, upHollowStr, downHollowStr, hingeStr string) (width, height, upHollow, downHollow float64, frameType string) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	upHollow, _ = strconv.ParseFloat(upHollowStr, 64)
	downHollow, _ = strconv.ParseFloat(downHollowStr, 64)
	hinge, _ := strconv.ParseFloat(hingeStr, 64)

	// 进行框架公式计算
	width = width - border*2 + reserve
	if hinge == 0 {
		height = height - (upHollow + downHollow) - border*4 + reserve
		if upHollow > 0 { // 如果有上画布
			upHollow += reserve
		}
		if downHollow > 0 { // 如果有下画布
			downHollow += reserve
		}
	} else { // 有合页的话就完全按照多座屏的方式
		height = height - (upHollow + downHollow) - hinge*border + reserve
		if upHollow > 0 { // 如果有上画布
			upHollow = upHollow - border*2 + reserve
		}
		if downHollow > 0 { // 如果有下画布
			downHollow = downHollow - border*2 + reserve
		}
	}
	totalHeight := upHollow + downHollow + height

	// 求出框架类型
	if upHollow > 0 {
		frameType += "上"
	}
	if downHollow > 0 {
		frameType += "下"
	}
	if frameType == "" {
		frameType += "无"
	}
	frameType += "画布"

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_%s_%.0fx%.0f", tools.NowTime(), frameType, width, totalHeight)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, totalHeight, frameName, false)

	// 追加专属的切图参考线
	model.FrameGuide4to2(upHollow, height)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave4to2(width, upHollow, height, downHollow, frameType, frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter5 顶天立地进行处理
func FramePresenter5(widthStr, heightStr, upHollowStr, downHollowStr, numberStr string) (width, height float64) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	upHollow, _ := strconv.ParseFloat(upHollowStr, 64)
	downHollow, _ := strconv.ParseFloat(downHollowStr, 64)
	number, _ := strconv.ParseFloat(numberStr, 64)

	// 进行框架公式计算
	width = width - border*2 + reserve
	height = height - upHollow - downHollow - border*2 - number*border + reserve

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_顶天立地_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter6 传统折屏进行处理
func FramePresenter6(widthStr, heightStr, upHollowStr, downHollowStr, numberStr string) (totalWidth, height float64, frameType string) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ := strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	upHollow, _ := strconv.ParseFloat(upHollowStr, 64)
	downHollow, _ := strconv.ParseFloat(downHollowStr, 64)
	number, _ := strconv.ParseFloat(numberStr, 64)

	// 进行框架公式计算
	width = width - border*2 + reserve         // 单扇的宽
	totalWidth = width * number                // 总宽
	height = height - border*2 + (reserve + 1) // 单扇的高  折屏高要多预留1厘米

	// 求出框架类型
	switch {
	case upHollow > 0 && downHollow > 0: // 如果有上镂空 下镂空的话
		height = height - upHollow - downHollow - border*2
		frameType = "上下镂空"
	case upHollow > 0: // 如果有上镂空的话
		height = height - upHollow - border
		frameType = "上镂空"
	case downHollow > 0: // 如果有下镂空的话
		height = height - downHollow - border
		frameType = "下镂空"
	default:
		frameType = "常规"
	}

	// 为当前框架指定名字，获取当前时间，进行格式化 2006-01-02 15:04:05
	frameName := fmt.Sprintf("%s_%s折屏_%.0fx%.0f", tools.NowTime(), frameType, totalWidth, height)

	// 定义单片名字
	singleName := fmt.Sprintf("%s折屏", frameType)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(totalWidth, height, frameName, false)

	// 追加专属的切图参考线
	model.FrameGuide6(width, number)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave6(width, height, number, frameName, singleName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()
	return
}

// FramePresenter7 多个座屏进行处理
func FramePresenter7(widthStrSlice, heightStrSlice []string, upHollowStr, downHollowStr string) (totalWidth, maxHeight float64) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 将全部字符串切片转浮点数切片的函数，顺便计算一下实际切图尺寸
	parseFloatSlice := func(s []string) (f []float64) {
		// 分配下内存
		f = make([]float64, len(s))

		// 将字符串转浮点数
		for i := 0; i < len(s); i++ {
			size, _ := strconv.ParseFloat(s[i], 64)
			// 顺便计算一下实际切图尺寸，赋值到切片
			f[i] = size - border*2 + reserve
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

	// 强制类型转换成浮点数
	widthSlice := parseFloatSlice(widthStrSlice)
	heightSlice := parseFloatSlice(heightStrSlice)
	upHollow, _ := strconv.ParseFloat(upHollowStr, 64)
	downHollow, _ := strconv.ParseFloat(downHollowStr, 64)

	// 进行框架公式计算
	if upHollow > 0 { // 如果有上镂空的话
		// 顺序遍历
		for i := 0; i < len(heightSlice); i++ {
			heightSlice[i] = heightSlice[i] - upHollow - border
		}
	}
	if downHollow > 0 { // 如果有下镂空的话
		// 顺序遍历
		for i := 0; i < len(heightSlice); i++ {
			heightSlice[i] = heightSlice[i] - downHollow - border
		}
	}

	// 遍历出总宽
	for i := 0; i < len(widthSlice); i++ {
		totalWidth += widthSlice[i]
	}

	// 遍历出最大的宽度
	maxWidth := maxSize(widthSlice)

	// 遍历出最大的高度
	maxHeight = maxSize(heightSlice)

	// 遍历出最小的高度
	minHeight := minSize(heightSlice)

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_%s座屏_%.0fx%.0f", tools.NowTime(), tools.Transfer(len(widthSlice)), totalWidth, maxHeight)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(totalWidth, maxHeight, frameName, false)

	// 追加专属的切图参考线与遮罩层
	model.FrameGuide7(widthSlice, heightSlice, maxHeight, minHeight)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(maxWidth, maxHeight)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave7(widthSlice, heightSlice, maxHeight, frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter8to1  对卷帘座屏进行处理
func FramePresenter8to1(widthStr, heightStr string) (width, height float64) {
	// 定义预留尺寸和传统边框宽度
	var reserve = viper.GetFloat64("reserve")
	var border = viper.GetFloat64("border")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算
	width = width - 20                      // 由于卷帘座屏左右两边的画布没有被嵌套，因此不需要计算预留
	height = height - 30 - border + reserve // 上下镂空各-15，横梁再-5

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_卷帘座屏_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true) // 创建ps文档

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()
	return
}

// FramePresenter8to2  对拉布座屏进行处理
func FramePresenter8to2(widthStr, heightStr, thicknessStr string) (width, height float64) {
	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	thickness, _ := strconv.ParseFloat(thicknessStr, 64) // 边框厚度

	// 进行框架公式计算
	saveWidth := width + thickness*2 + 2   // 保存时的宽度
	saveHeight := height + thickness*2 + 2 // 保存时的高度

	// 为当前框架指定名字，此框架特殊，保存时进行框架计算
	frameName := fmt.Sprintf("%s_拉布座屏_%.0fx%.0f", tools.NowTime(), saveWidth, saveHeight)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, false)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(saveWidth, saveHeight)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave8to2(frameName, thickness)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	// 此框架是保存时才将画布变大
	width = saveWidth
	height = saveHeight
	return
}

// FramePresenter8to3 对拉布折屏进行处理
func FramePresenter8to3(widthStr, heightStr, countStr string) (totalWidth, height float64) {
	// 强制类型转换成浮点数
	width, _ := strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	count, _ := strconv.ParseFloat(countStr, 64)
	// 算出总宽
	totalWidth = width * count

	// 进行框架公式计算

	// 为当前框架指定名字，此框架特殊，保存时进行框架计算
	frameName := fmt.Sprintf("%s_拉布折屏_%.0fx%.0f", tools.NowTime(), totalWidth, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(totalWidth, height, frameName, false)

	// 生成专属的切图参考线
	model.FrameGuide6(width, count)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width+8, height+8)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave8to3(frameName, width, height, count)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter9 对补切画布进行处理
func FramePresenter9(widthStr, heightStr string) (width, height float64) {
	// 定义一个预留尺寸

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_补切画布_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true) // 创建ps文档

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}
