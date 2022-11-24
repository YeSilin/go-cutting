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
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算
	width = width - 10 + reserve
	height = height - 10 + reserve

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_常规座屏_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true) // 创建ps文档

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter2 对左右镂空进行处理
func FramePresenter2(widthStr, heightStr, leftHollowStr, rightHollowStr, hingeStr string) (width, height float64, frameType string) {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	leftHollow, _ := strconv.ParseFloat(leftHollowStr, 64)
	rightHollow, _ := strconv.ParseFloat(rightHollowStr, 64)
	hinge, _ := strconv.ParseFloat(hingeStr, 64)

	// 进行框架公式计算
	if hinge == 0 {
		width = width - 10 + reserve
		if leftHollow > 0 {
			width -= leftHollow + 5 // 如果有左镂空的话
		}
		if rightHollow > 0 {
			width -= rightHollow + 5 // 如果有右镂空的话
		}
	} else {
		width = width - (leftHollow + rightHollow) - 10 + reserve
	}
	height = height - 10 + reserve

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

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}

// FramePresenter8to1  对卷帘座屏进行处理
func FramePresenter8to1(widthStr, heightStr string) (width, height float64) {
	// 定义一个预留尺寸
	var reserve = viper.GetFloat64("reserve")

	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算
	width = width - 20             // 由于卷帘座屏左右两边的画布没有被嵌套，因此不需要计算预留
	height = height - 35 + reserve // 上下镂空各-15，横梁再-5

	// 为当前框架指定名字
	frameName := fmt.Sprintf("%s_卷帘座屏_%.0fx%.0f", tools.NowTime(), width, height)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, true) // 创建ps文档

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()
	return
}

// FramePresenter8to2  对拉布座屏进行处理
func FramePresenter8to2(widthStr, heightStr, thicknessStr string) (width, height float64) {
	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)
	thickness, _ := strconv.ParseFloat(thicknessStr, 64)

	// 进行框架公式计算
	saveWidth := width + thickness*2 + 2   // 保存时的宽度
	saveHeight := height + thickness*2 + 2 // 保存时的高度

	// 为当前框架指定名字，此框架特殊，保存时进行框架计算
	frameName := fmt.Sprintf("%s_拉布座屏_%.0fx%.0f", tools.NowTime(), saveWidth, saveHeight)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, false)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave8to2(frameName, thickness)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(saveWidth, saveHeight)

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

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave8to3(frameName, width, height, count)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width+8, height+8)

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

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSaveDef(frameName)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width, height)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	return
}
