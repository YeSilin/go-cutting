package presenter

// 限制切图软件使用

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"strings"
)

// Limit 初始化的 Viper配置实例 限制使用
var Limit *viper.Viper

// 内部做好初始化无需而外调用
func init() {
	Limit = viper.New()
	// 指定配置文件类型
	Limit.SetConfigType("yaml")
	err := Limit.ReadConfig(strings.NewReader(tools.GetYNote("aa95080948125d854a5ca245b778ee22")))
	if err != nil {
		fmt.Println("获取授权文件出错，请关闭重试！")
	}
}

// GetLicense 获取许可证
func GetLicense(version string) bool {
	// 先获取许可证列表
	ls := Limit.GetStringSlice("version")

	for i := 0; i < len(ls); i++ {
		// 在许可版本内允许运行
		if ls[i] == version {
			return true
		}
	}
	return false
}
