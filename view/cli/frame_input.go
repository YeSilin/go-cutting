package cli

// 快捷切图的具体框架参数获取
import (
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
)

// 小座屏 扣掉两个边框-5-5 然后再加回来5厘米  可以理解为扣5
func frame1() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架常规座屏", 74)         // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入常规座屏的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入常规座屏的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter1(widthStr, heightStr)

		// 输出提示
		color.Yellow.Printf("\n:: 常规座屏：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 卷帘座屏  宽-20   高上下各镂空各-15 上横梁-5    高预留5
func frame8to1() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架卷帘座屏", 74)         // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入卷帘座屏的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入卷帘座屏的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter8to1(widthStr, heightStr)

		// 输出提示
		color.Yellow.Printf("\n:: 卷帘座屏：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//  拉布座屏框架，需要一个宽和高
func frame8to2() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架拉布座屏", 74)         // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入拉布座屏的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入拉布座屏的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter8to2(widthStr, heightStr)

		// 输出提示
		color.Yellow.Printf("\n:: 拉布座屏：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

//  拉布折屏框架，需要：宽，高，片数
func frame8to3() {
	
}
