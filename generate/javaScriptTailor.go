package generate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"text/template"
)

// 根据当前文档名选择正确的快捷裁剪脚本
func SelectTailor() {
	const script = `// 声明：这是一个调用针对当前文档的自动裁剪脚本

// 判断是否有打开的文件
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
    // return;
} else {
    // 获取当前脚本所在路径
    var scriptPath = (new File($.fileName)).parent;

    // 获取当前文档名字
    var nowName = app.activeDocument.name;

    // 要运行的脚本路径
    var runScript = scriptPath + "/temp/tailor_" + nowName + ".jsx";

    // 获得脚本对象
    var fileRef = new File(runScript);
    if (fileRef.exists) {    // 如果脚本存在
        app.load(fileRef);   // 运行脚本
    } else {// 不存在就运行默认裁剪
        alert("未找到当前文档定制版【-1】脚本，已自动调用默认版脚本！");
        app.load(new File(scriptPath + "/GeneralCutting.jsx"));
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/selectTailor.jsx", script)
}

// 生成大部分框架的自动裁剪，例如左右镂空，小座屏等
func GeneralCutting(frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		BlackEdge bool // 是否自动黑边
	}{viper.GetBool("blackEdge")}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/generalCutting.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 生成的文件名字
	fileName := "config/jsx/generalCutting.jsx"
	// 框架名不是空，就生成专属裁剪脚本名字
	if frameName != "" {
		fileName = fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName)
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fileName)
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

// 生成中间大两边小的自动裁剪js
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

//生成折屏的自动裁剪js
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

// 生成多座屏的自动裁剪js
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
	}{tools.ToJsArray(widthSlice), tools.ToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice)), viper.GetBool("blackEdge")}

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

//生成贴图折屏的自动裁剪js
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
	tmpl, err := template.ParseFiles("config/jsx/template/foldingScreensForMap.gohtml")
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
