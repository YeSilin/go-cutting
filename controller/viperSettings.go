// 这是一个名为毒蛇的配置
package controller

import (
	"fmt"
	"github.com/spf13/viper"
)

// 初始化
func InitSetting() {
	// 初始化的 Viper实例
	//v := viper.New()

	// 设置默认值
	viper.SetDefault("memory", false)              // 记忆框架
	viper.SetDefault("openPs", true)               // 自动新建文档
	viper.SetDefault("blackEdge", true)            // 黑边
	viper.SetDefault("prefix", "")                 // 前缀
	viper.SetDefault("reserve", 5)                 // 画布预留
	viper.SetDefault("gui", true)                  // 自动开启gui
	viper.SetDefault("cipherList", false)          // 自动开启暗号列表
	viper.SetDefault("picture", "config\\picture") // 正斜杠会出错
	viper.SetDefault("automaticDeletion", false)   // 自动主图时删除来源
	viper.SetDefault("darkTheme", true)            // 默认黑色主题

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
