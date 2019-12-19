package globa

// 定义全局结构体，用于储存配置信息
type SetInformation struct {
	Memory    bool
	OpenPs    bool
	BlackEdge bool
	Reserve   float64
}

// 定义全局默认配置参数
var DefaultSetting SetInformation = SetInformation{false, true, true, 5}

// 预定义全局解码结果  获取当前设置
var NowSetting SetInformation


