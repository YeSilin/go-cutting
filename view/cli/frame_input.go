package cli

// 快捷切图的具体框架参数获取
import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"strconv"
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
		color.Yellow.Printf("\n:: %s折屏：总宽 %.2f cm，高 %.2f cm\n", frameType, totalWidth, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 旧厂多座屏
func frame7() {
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

	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架多座屏", 74) // 请注意切图的工厂与框架的选择

		numberStr := inputPro("\n:: 请输入拥有几个座屏：", 1)
		// 输入返回当然要返回啦
		if numberStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 字符串转换成int，得到具体要切几个座屏
		number, _ := strconv.Atoi(numberStr)

		// 第几座屏的文案由阿拉伯数字改成纯中文
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
				saveSizeStr[i] = inputPro(inputPrompt[i], 6)
			} else {
				saveSizeStr[i] = inputPro(inputPrompt[i], 0)
			}
			// 输入返回当然要返回啦
			if saveSizeStr[i] == "-" {
				tools.CallClear() // 清屏
				return
			}
		}

		// 开始解码得到的值
		widthStrSlice, heightStrSlice, upHollowStr, downHollowStr := deSliceStacking(saveSizeStr)

		// 处理框架生成脚本
		totalWidth, maxHeight := presenter.FramePresenter7(widthStrSlice, heightStrSlice, upHollowStr, downHollowStr)

		// 输出提示
		color.Yellow.Printf("\n:: %s座屏：总宽 %.2f cm，最高 %.2f cm\n", tools.Transfer(number), totalWidth, maxHeight)

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

// 不扣补切，不扣任何大小
func frame9to1() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架不扣补切", 74) // 请注意切图的工厂与框架的选择
		fmt.Println("\n:: 此框架主要用来补切画布，不减去任何边框尺寸，适合不想手动新建画布时使用！")

		widthStr := inputPro("\n:: 请输入不扣补切的宽：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if widthStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		heightStr := inputPro("\n:: 请输入不扣补切的高：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if heightStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		width, height := presenter.FramePresenter9to1(widthStr, heightStr)

		// 输出提示
		color.Yellow.Printf("\n:: 不扣补切：宽 %.2f cm，高 %.2f cm\n", width, height)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 圆形补切，扣2
func frame9to2() {
	// 由配置决定是否循环使用此框架
	for {
		tools.ChineseTitle("当前框架圆形补切", 74) // 请注意切图的工厂与框架的选择

		diameterStr := inputPro("\n:: 请输入圆形补切的直径：", 6) // 获取键盘输入
		// 输入返回当然要返回啦
		if diameterStr == "-" {
			tools.CallClear() // 清屏
			return
		}

		// 处理框架生成脚本
		diameter := presenter.FramePresenter9to2(diameterStr)

		// 输出提示
		color.Yellow.Printf("\n:: 圆形补切：直径 %.2f cm\n", diameter)

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}
