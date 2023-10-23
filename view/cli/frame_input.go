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

// 旧厂中间大两边小 先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个竖边 空出的中间画面加5厘米
func frame3() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架左右画布", 74)          // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入左右画布的总宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入左右画布的总高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		hollowStr := inputPro("\n:: 请输入单边画布的大小：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if hollowStr == "-" {
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
		width, height, hollow := presenter.FramePresenter3(widthStr, heightStr, hollowStr, hingeStr)

		// 输出提示
		color.Yellow.Printf("\n:: 左右画布：中间 %.2f cm，两边各 %.2f cm，高 %.2f cm\n", width, hollow, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 旧厂上下镂空 先扣镂空尺寸 先扣两个镂空的大小  再扣掉 几个边框5 两镂空就有4个横边 空出的中间画面加5厘米
func frame4to1() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架上下镂空", 74)          // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入上下镂空的总宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入上下镂空的总高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		upHollowStr := inputPro("\n:: 请输入上镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if upHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		downHollowStr := inputPro("\n:: 请输入下镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if downHollowStr == "-" {
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
		width, height, frameType := presenter.FramePresenter4to1(widthStr, heightStr, upHollowStr, downHollowStr, hingeStr)

		// 输出提示
		color.Yellow.Printf("\n:: %s：宽 %.2f cm，高 %.2f cm\n", frameType, width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}

	}

}

// 上下画布 例如 总高：180 上下分别：30+5  中间为：100+5，也就是说边框5都在中间画布扣了
func frame4to2() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架上下画布", 74) // 请注意切图的工厂与框架的选择

		widthStr := inputPro("\n:: 请输入上下画布的总宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入上下画布的总高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		upHollowStr := inputPro("\n:: 请输入上画布的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if upHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		downHollowStr := inputPro("\n:: 请输入下画布的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if downHollowStr == "-" {
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
		width, height, upHollow, downHollow, frameType := presenter.FramePresenter4to2(widthStr, heightStr, upHollowStr, downHollowStr, hingeStr)

		// 输出提示
		color.Yellow.Printf("\n:: %s：宽 %.2f cm，上高 %.2f cm，中高 %.2f cm，下高 %.2f cm\n", frameType, width, upHollow, height, downHollow)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 旧厂顶天立地 扣掉镂空部分 再扣5
func frame5() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架顶天立地", 74)          // 请注意切图的工厂与框架的选择
		widthStr := inputPro("\n:: 请输入顶天立地的总宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入顶天立地的总高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		upHollowStr := inputPro("\n:: 请输入上镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if upHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		downHollowStr := inputPro("\n:: 请输入下镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if downHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		numberStr := inputPro("\n:: 请输入拥有几个贴地或贴顶横杆：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if numberStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter5(widthStr, heightStr, upHollowStr, downHollowStr, numberStr)

		// 输出提示
		color.Yellow.Printf("\n:: 顶天立地：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 旧厂常规折屏
func frame6() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架各种折屏", 74) // 请注意切图的工厂与框架的选择

		widthStr := inputPro("\n:: 请输入折屏单扇的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入折屏单扇的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		upHollowStr := inputPro("\n:: 请输入上镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if upHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		downHollowStr := inputPro("\n:: 请输入下镂空的大小：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if downHollowStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		numberStr := inputPro("\n:: 请输入共拥有几扇：", 0) // 获取键盘输入
		// 输入返回当然要返回啦
		if numberStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		totalWidth, height, frameType := presenter.FramePresenter6(widthStr, heightStr, upHollowStr, downHollowStr, numberStr)

		// 输出提示
		color.Yellow.Printf("\n:: %s折屏：总宽 %.2f cm，高 %.2f cm", frameType, totalWidth, height)

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

		thicknessStr := inputPro("\n:: 请输入边框的厚度（默认3）：", 3) // 获取键盘输入
		// 输入返回当然要返回啦
		if thicknessStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter8to2(widthStr, heightStr, thicknessStr)

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
