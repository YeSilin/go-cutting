package generate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"strings"
	"text/template"
)

//生成中间大两边小参考线js
//@param width 传入宽度
//@param hollowOut: 传入镂空
func LineJs3(width, hollowOut float64) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Line1 float64
		Line2 float64
	}{hollowOut, hollowOut + width}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/leftAndRightCanvasReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("config/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// 上下画布参考线
func LineJs4to2(upperHeight, middleHeight float64) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Line1 float64
		Line2 float64
	}{upperHeight, upperHeight + middleHeight}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/upAndDownCanvasReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("config/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

//生成折屏参考线js
//@param width: 传入宽
//@param number: 传入扇数
func LineJs6(width, number float64) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width  float64
		Number int
	}{width, int(number)}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/foldingScreensReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("config/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// 生成多座屏参考线js
func LineJs7(widthSlice, heightSlice []float64, heightMax, heightMin float64) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		HeightMax     float64 // 最大的高
		ScreenName    string  // 是几座屏
		Equal         bool    //最高和最矮座屏是否相等
	}{tools.Float64SliceToJsArray(widthSlice), tools.Float64SliceToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice)), heightMax == heightMin}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/multiScreenReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("config/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

	//// 如果高度最高和最低相等，那么就添加注释，就是没有必要添加遮罩层
	//var notes = ""
	//if heightMax == heightMin {
	//	notes = "//"
	//}
	//
	//var jsx = strings.Builder{}
	//
	//// 生成遮罩图层加宽度参考线
	//jsx.WriteString("\n")
	//jsx.WriteString("// 定义一个函数用来新建透明图层\n")
	//jsx.WriteString("function mask() {\n")
	//jsx.WriteString("    // 新建一个图层\n")
	//jsx.WriteString(fmt.Sprintf("    app.activeDocument.artLayers.add().name = \"%s座屏遮罩区域\";\n", tools.Transfer(len(widthSlice))))
	//jsx.WriteString("\n")
	//jsx.WriteString("    // 不透明度改30%\n")
	//jsx.WriteString("    app.activeDocument.artLayers[0].opacity=30;\n")
	//jsx.WriteString("}\n")
	//jsx.WriteString("\n")
	//jsx.WriteString("// 定义一个函数用来设置区域填充颜色\n")
	//jsx.WriteString("function makeSelection(x,y,sw,sh){\n")
	//jsx.WriteString("    // 设置选区\n")
	//jsx.WriteString("    app.activeDocument.selection.select([ [x,y], [x,y+sh], [x+sw,y+sh], [x+sw,y] ]);\n")
	//jsx.WriteString("\n")
	//jsx.WriteString("    // 生成一个随机色\n")
	//jsx.WriteString("    var color = new SolidColor();\n")
	//jsx.WriteString("    color.rgb.red = Math.random() * 255;\n")
	//jsx.WriteString("    color.rgb.green = Math.random() * 255;\n")
	//jsx.WriteString("    color.rgb.blue = Math.random() * 255;\n")
	//jsx.WriteString("\n")
	//jsx.WriteString("    // 填充选区\n")
	//jsx.WriteString("    app.activeDocument.selection.fill(color);\n")
	//jsx.WriteString("    // 取消选区\n")
	//jsx.WriteString("    app.activeDocument.selection.deselect();\n")
	//jsx.WriteString("}\n")
	//jsx.WriteString("\n")
	//jsx.WriteString("// 生成历史记录\n")
	//jsx.WriteString(fmt.Sprintf("%sapp.activeDocument.suspendHistory(\"%s座屏遮罩图层\", \"mask()\");\n", notes, tools.Transfer(len(widthSlice))))
	//
	//for i := 0; i < len(heightSlice); i++ { // 遍历每一片座屏的高
	//
	//	if heightSlice[i] < heightMax { // 如果其中某些比较矮
	//		var x = 0.0              // x 默认是0
	//		for j := 0; j < i; j++ { // 如果编号不是第一片
	//			x += widthSlice[j]
	//		}
	//		jsx.WriteString("// 遮罩层写入历史\n")
	//		jsx.WriteString(fmt.Sprintf("app.activeDocument.suspendHistory(\"第%s个座屏遮罩区域\", \"makeSelection(%f, 0, %f, %f)\");\n", tools.Transfer(i+1), x*39.37, widthSlice[i]*39.37, (heightMax-heightSlice[i])*39.37))
	//	}
	//}
	//
	//// 这里开始是宽度参考线
	//var notes2 = ""
	//if len(heightSlice) == 1 {
	//	notes2 = "//"
	//}
	//
	//jsx.WriteString("\n")
	//jsx.WriteString(fmt.Sprintf("%sapp.activeDocument.suspendHistory(\"%s座屏宽度参考线\", \"addLine()\");  // 生成历史记录\n", notes2, tools.Transfer(len(widthSlice))))
	//jsx.WriteString("function addLine(){   // 定义一个函数用于新建参考线\n")
	//
	//var lineSum = 0.0 // 一开始参考线是0
	//for i := 0; i < len(widthSlice)-1; i++ {
	//	lineSum += widthSlice[i]
	//	jsx.WriteString(fmt.Sprintf("	activeDocument.guides.add (Direction.VERTICAL,UnitValue(\"%fcm\"));\n", lineSum))
	//}
	//jsx.WriteString("}\n")
	//
	//// 除了最高的那片，其他都生成高度参考线
	//for i := 0; i < len(heightSlice); i++ {
	//	if heightSlice[i] < heightMax {
	//		jsx.WriteString("\n")
	//		jsx.WriteString(fmt.Sprintf("app.activeDocument.suspendHistory(\"第%s个座屏高度参考线\", \"addLine%d()\");  // 生成历史记录\n", tools.Transfer(i+1), i))
	//		jsx.WriteString(fmt.Sprintf("function addLine%d(){   // 定义一个函数用于新建参考线\n", i))
	//		jsx.WriteString(fmt.Sprintf("    activeDocument.guides.add (Direction.HORIZONTAL,UnitValue(\"%fcm\"));\n", heightMax-heightSlice[i]))
	//		jsx.WriteString("}\n")
	//	}
	//}
	//
	//jsx.WriteString("// 设置激活的图层\n")
	//jsx.WriteString(fmt.Sprintf("%sapp.activeDocument.activeLayer = app.activeDocument.layers[1]\n", notes))
	//
	//// 转成字符串格式
	//jsxStr := jsx.String()
	//// 追加写入
	//tools.WriteFile("config/jsx/newDocument.jsx", jsxStr)
}

//生成折屏3d贴图参考线js
//@param width: 传入宽
//@param number: 传入扇数
func Line3DMapJs6(width, number int) {
	var line = strings.Builder{}

	var i = 1
	for ; i < number; i++ {
		w := width * i
		// 追加参考线js代码
		line.WriteString(fmt.Sprintf("\nactiveDocument.guides.add (Direction.VERTICAL,UnitValue(\"%dPIXELS\"));", w))
	}
	// 转成字符串格式
	lineStr := line.String()

	// 收集所有生成的js
	var jsx = strings.Builder{}

	jsx.WriteString("app.activeDocument.suspendHistory(\"折屏参考线！\", \"addLine()\");  // 生成历史记录\n")
	jsx.WriteString("function addLine(){   // 定义一个函数用于新建参考线\n")
	jsx.WriteString(fmt.Sprintf("    %s\n", lineStr))
	jsx.WriteString("}\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 追加写入
	tools.WriteFile("config/jsx/newDocument.jsx", jsxStr)
}

// 生成多座屏贴图参考线
func Line3DMapJs7(widthSlice, heightSlice []int, heightMax, heightMin int) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		HeightMax     int    // 最大的高
		ScreenName    string // 是几座屏
		Equal         bool   //最高和最矮座屏是否相等
	}{tools.IntSliceToJsArray(widthSlice), tools.IntSliceToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice)), heightMax == heightMin}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/map/multiScreenReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("config/jsx/newDocument.jsx", os.O_APPEND, 0644)
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
