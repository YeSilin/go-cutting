package unclassified

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"strings"
	"text/template"
)

// 生成折屏3d贴图参考线js
// @param width: 传入宽
// @param number: 传入扇数
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
	tools.WriteFile("data/jsx/newDocument.jsx", jsxStr)
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
	tmpl, err := template.ParseFiles("data/jsx/template/map/multiScreenReferenceLine.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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
