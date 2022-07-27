package cli

// 快捷切图的具体框架参数获取
import (
	"fmt"
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

// 左右镂空 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米  旧厂的边框厚度是5厘米
func frame2() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架左右镂空", 74)          // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入左右镂空的总宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入左右镂空的总高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		leftHollowStr := inputPro("\n:: 请输入左镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if leftHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		rightHollowStr := inputPro("\n:: 请输入右镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if rightHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		hingeStr := inputPro("\n:: 请输入合页数量（若订单无备注请输入“0”）：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if hingeStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height, frameType := presenter.FramePresenter2(widthStr, heightStr, leftHollowStr, rightHollowStr, hingeStr)

		// 输出提示
		color.Yellow.Printf("\n:: %s：宽 %.2f cm，高 %.2f cm\n", frameType, width, height)

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

// 拉布座屏框架，需要一个宽和高
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

// 拉布折屏框架，需要：宽，高，片数
func frame8to3() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架拉布折屏", 74)         // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入拉布折屏的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入拉布折屏的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		countStr := inputPro("\n:: 请输入共拥有几扇：", 1) // 获取键盘输入
		// 输入返回当然要返回啦
		if countStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		totalWidth, height := presenter.FramePresenter8to3(widthStr, heightStr, countStr)

		// 输出提示
		color.Yellow.Printf("\n:: 拉布折屏：宽 %.2f cm，高 %.2f cm\n", totalWidth, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 补切画布不扣任何大小
func frame9() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架补切画布", 74) // 请注意切图的工厂与框架的选择
		fmt.Println("\n:: 此框架主要用来补切画布，不减去任何边框尺寸，适合不想手动新建画布时使用！")

		widthStr := inputPro("\n:: 请输入补切画布的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入补切画布的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter9(widthStr, heightStr)

		// 输出提示
		color.Yellow.Printf("\n:: 补切画布：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}
