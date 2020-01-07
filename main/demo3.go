package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func init() {
	//// 初始化的 Viper实例
	//v := viper.New()



}

// 声明一个设置的结构体
type Settings struct {
	Memory    bool    `yaml:"memory"`
	OpenPs    bool    `yaml:"openPs"`
	BlackEdge bool    `yaml:"blackEdge"`
	Prefix    string  `yaml:"prefix"`
	Reserve   float64 `yaml:"reserve"`
}

func main3() {


	// 设置默认值
	viper.SetDefault("memory", false)
	viper.SetDefault("openPs", true)
	viper.SetDefault("blackEdge", true)
	viper.SetDefault("prefix", "")
	viper.SetDefault("reserve", 5)

	//  设置配置文件名，不带后缀
	viper.SetConfigName("settings")
	// 第一个搜索路径
	viper.AddConfigPath("./Config")
	//设置配置文件类型
	viper.SetConfigType("yaml")

	// 如果找不到，就保存当前默认配置文件
	_ = viper.SafeWriteConfig()

	// 搜索路径，并读取配置数据
	err := viper.ReadInConfig()
	if err != nil {

		fmt.Println(err)
		viper.WriteConfig()
	}

	fmt.Println(viper.GetInt("reserve"))

	//也可以直接反序列化为Struct
	//var settings Settings
	//if err := v.Unmarshal(&settings); err != nil {
	//	fmt.Printf("viper.Unmarshal err: %s", err)
	//}





	fmt.Println(viper.GetInt("reserve"))
	//viper.WatchConfig()
	//
	//
	for   {
		time.Sleep(time.Second)
		fmt.Println(viper.GetInt("reserve"))
	}


//	v.Set("reserve" , 10)
//	fmt.Println(v.GetInt("reserve"))
//	//err = v.WriteConfig()
//fmt.Println(err)
	//time.Sleep(time.Second*5)
}
