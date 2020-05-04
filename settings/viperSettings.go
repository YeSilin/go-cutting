// 这是一个名为毒蛇的配置
package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
)

// 初始化
func init() {
	// 初始化的 Viper实例
	//v := viper.New()

	// 设置默认值
	viper.SetDefault("memory", false)
	viper.SetDefault("openPs", true)
	viper.SetDefault("blackEdge", true)
	viper.SetDefault("prefix", "")
	viper.SetDefault("reserve", 5)
	viper.SetDefault("cipherList", false)  // 自动开启暗号列表
	viper.SetDefault("picture", "config\\picture") // 正斜杠会出错
	viper.SetDefault("automaticDeletion", false)   // 自动主图时删除来源

	//  设置配置文件名，不带后缀
	viper.SetConfigName("settings")

	// 第一个搜索路径
	viper.AddConfigPath("./config/")

	//设置配置文件类型
	viper.SetConfigType("yaml")

	// 安全保存配置文件，如果没有配置文件就保存当前配置
	_ = viper.SafeWriteConfig()

	// 搜索路径，并读取配置数据
	err := viper.ReadInConfig()
	if err != nil {
		// 如果读取失败，就保存当前默认配置文件
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("viper.WriteConfig err: ", err)
			return
		}
	}

	// 监视配置文件，重新读取配置数据
	viper.WatchConfig()

	// 显示更新信息，不稳
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("【提示】配置文件已更新，来自：", e.Name)
	//})
}

// 当前状态
func current() {
	model.EnglishTitle("Settings", 74)
	fmt.Println("\n:: 这里提供简单的设置端口，如果大家有其他需要实现的功能设置可以在群里反馈！")

	var memoryStr string
	switch viper.GetBool("memory") {
	case true:
		memoryStr = "已启用"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		memoryStr = "已关闭"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		memoryStr = "参数错误"
		memoryStr = model.ColourString(memoryStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	var openPsStr string
	switch viper.GetBool("openPs") {
	case true:
		openPsStr = "已启用"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		openPsStr = "已关闭"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		openPsStr = "参数错误"
		openPsStr = model.ColourString(openPsStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 自动黑边状态
	var blackEdgeStr string
	switch viper.GetBool("blackEdge") {
	case true:
		blackEdgeStr = "已启用"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		blackEdgeStr = "已关闭"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		blackEdgeStr = "参数错误"
		blackEdgeStr = model.ColourString(blackEdgeStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 自定前缀状态
	var prefixStr string
	if viper.GetString("prefix") != "" {
		prefixStr = "已定义"
		prefixStr = model.ColourString(prefixStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	} else {
		prefixStr = "未定义"
		prefixStr = model.ColourString(prefixStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 带颜色的预留画布提示
	var reserveStr = fmt.Sprintf("%.2fcm", viper.GetFloat64("reserve"))
	reserveStr = model.ColourString(reserveStr, ctc.ForegroundGreen) // 设置带颜色的字符串

	// 自动开启暗号列表
	var cipherListStr string
	switch viper.GetBool("cipherList") {
	case true:
		cipherListStr = "自启动"
		cipherListStr = model.ColourString(cipherListStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		cipherListStr = "已关闭"
		cipherListStr = model.ColourString(cipherListStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		cipherListStr = "参数错误"
		cipherListStr = model.ColourString(cipherListStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 修改套图文件夹位置
	var pictureStr string
	if viper.GetString("picture") != "config\\picture" {
		pictureStr = "已修改"
		pictureStr = model.ColourString(pictureStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	} else {
		pictureStr = "默认值"
		pictureStr = model.ColourString(pictureStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	// 主图自动删除来源
	var automaticDeletionStr string
	switch viper.GetBool("automaticDeletion") {
	case true:
		automaticDeletionStr = "已启用"
		automaticDeletionStr = model.ColourString(automaticDeletionStr, ctc.ForegroundGreen) // 设置带颜色的字符串
	case false:
		automaticDeletionStr = "已关闭"
		automaticDeletionStr = model.ColourString(automaticDeletionStr, ctc.ForegroundBright) // 设置带颜色的字符串
	default:
		automaticDeletionStr = "参数错误"
		automaticDeletionStr = model.ColourString(automaticDeletionStr, ctc.ForegroundBright) // 设置带颜色的字符串
	}

	fmt.Printf("\n   [1]记忆框架：%s       [2]自动新建：%s       [3]自动黑边：%s\n", memoryStr, openPsStr, blackEdgeStr)
	fmt.Printf("\n   [4]自定前缀：%s       [5]切布预留：%s       [6]暗号列表：%s\n", prefixStr, reserveStr,cipherListStr)
	fmt.Printf("\n   [7]套图位置：%s       [8]主图自删：%s       [9]全部恢复默认设置\n", pictureStr, automaticDeletionStr)
}
