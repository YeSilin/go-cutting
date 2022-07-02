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
func FramePresenter8to2(widthStr, heightStr string) (width, height float64) {
	// 强制类型转换成浮点数
	width, _ = strconv.ParseFloat(widthStr, 64)
	height, _ = strconv.ParseFloat(heightStr, 64)

	// 进行框架公式计算

	// 为当前框架指定名字，此框架特殊，保存时进行框架计算
	frameName := fmt.Sprintf("%s_拉布座屏_%.0fx%.0f", tools.NowTime(), width+8, height+8)

	// 生成创建Photoshop新文档脚本
	model.NewDocument(width, height, frameName, false)

	// 生成暗号【-1】可以用的另存脚本
	go model.FrameSave8to2(frameName)

	// 追加最大画布判断
	model.IsMaxCanvasExceeded(width+8, height+8)

	// 是否打开自动新建文档
	model.RunAutoCreateDocuments()

	// 此框架是保存时才将画布变大
	width += 8
	height += 8
	return
}
