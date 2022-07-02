package generate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"text/template"
)

//Tailor3 生成中间大两边小的自动裁剪js
// @param width 传入中间宽度
// @param height 传入高度
// @param hollowOut 传入镂空
func Tailor3(width, height, hollowOut float64, frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width     float64
		Height    float64
		HollowOut float64 // 中间大两边小的镂空均是
		BlackEdge bool    // 是否自动黑边
	}{width, height, hollowOut, viper.GetBool("blackEdge")}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/leftAndRightCanvas.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}

//Tailor4to2  生成上下画布的自动裁剪JavaScript
func Tailor4to2(width, upperCanvas, middleCanvas, downCanvas float64, shortName, frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		FrameName    string  // 确定是 上画布 上下画布 下画布
		Width        float64 // 总宽
		UpperCanvas  float64 // 上高
		MiddleCanvas float64 // 中高
		DownCanvas   float64 // 下高
		BlackEdge    bool    // 是否自动黑边
	}{shortName, width, upperCanvas, middleCanvas, downCanvas, viper.GetBool("blackEdge")}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/upAndDownCanvas.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}

//Tailor6 生成折屏的自动裁剪js
//@param width 传入单扇宽度
//@param height 传入高度
//@param number 传入扇数
func Tailor6(width, height, number float64, frameName, singleName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width      float64
		Height     float64
		Number     int    // 几片折屏
		SingleName string // 单片名字
		BlackEdge  bool   // 是否自动黑边
	}{width, height, int(number), singleName, viper.GetBool("blackEdge")}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/foldingScreens.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}

}

//Tailor7 生成多座屏的自动裁剪js
// @param width 传入中间宽度
// @param height 传入高度
// @param hollowOut 传入镂空
func Tailor7(widthSlice, heightSlice []float64, heightMax float64, frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		HeightMax     float64 // 最大的高
		ScreenName    string  // 是几座屏
		BlackEdge     bool    // 是否自动黑边
	}{tools.Float64SliceToJsArray(widthSlice), tools.Float64SliceToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice)), viper.GetBool("blackEdge")}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/multiScreen.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}

//TailorForMap6 生成贴图折屏的自动裁剪js
//@param width 传入单扇宽度
//@param height 传入高度
//@param number 传入扇数
func TailorForMap6(width, height, number int, frameName, singleName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width      int // 像素没有小数点
		Height     int
		Number     int    // 几片折屏
		SingleName string // 单片名字
	}{width, height, number, singleName}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/map/foldingScreensCutting.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}

//TailorForMap7 生成多座屏贴图的自动裁剪
func TailorForMap7(widthSlice, heightSlice []int, heightMax int, frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		HeightMax     int    // 最大的高
		ScreenName    string // 是几座屏
	}{tools.IntSliceToJsArray(widthSlice), tools.IntSliceToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice))}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/map/multiScreenCutting.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName))
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}
