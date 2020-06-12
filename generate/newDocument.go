package generate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

//生成用来新建ps文档js
//@param width 传入宽度
//@param height 传入高度
//@param wordLine 是否创建文字不要超过的参考线
func NewDocument(width, height float64, frameName string, wordLine bool) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width     float64
		Height    float64
		FrameName string // 新文档名
		WordLine  bool   // 文字参考线
	}{width, height, frameName, wordLine}

	// 为模板自定义加法函数
	//add := func(left float64, right float64) float64 {
	//	return left + right
	//}

	// 为模板自定义减法函数
	sub := func(Minuend float64, Reduction float64) float64 {
		return Minuend - Reduction
	}

	// 采用链式操作在Parse解析之前调用Funcs添加自定义的kua函数
	// 这边有个地方值得注意，template.New()函数中参数名字要和ParseFiles（）
	// 函数的文件名要相同，要不然就会报错："" is an incomplete template
	tmpl, err := template.New("newDocument.gohtml").Funcs(template.FuncMap{"sub": sub}).ParseFiles("config/jsx/template/newDocument.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	//// 解析指定文件生成模板对象
	//tmpl, err = template.ParseFiles("config/jsx/template/newDocument.gohtml")
	//if err != nil {
	//	fmt.Println("template.ParseFiles err: ", err)
	//	return
	//}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("config/jsx/newDocument.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}

	// 关闭文件
	f.Close()
}

// 生成用来新建ps文档3d作图js
// @param width 传入宽度
// @param height 传入高度
func NewDocumentForMap(width, height int, frameName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width     int
		Height    int
		FrameName string // 新文档名
	}{width, height, frameName}

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/newDocumentForMap.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("config/jsx/newDocument.jsx")
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

// 新建临时效果图文档
func NewTempDocumentJs() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now := time.Now().Format("0102150405")

	var jsx = strings.Builder{}

	jsx.WriteString("app.preferences.rulerUnits = Units.PIXELS;\n")
	jsx.WriteString("var width = 5000;\n")
	jsx.WriteString("var height = 3500;\n")
	jsx.WriteString("var resolution = 72;\n")
	jsx.WriteString(fmt.Sprintf("var docName = \"%s_临时效果图\";   // 这里传golang 时间参数哦！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！\n", now))
	jsx.WriteString("var mode = NewDocumentMode.RGB;\n")
	jsx.WriteString("var initialFill = DocumentFill.WHITE;\n")
	jsx.WriteString("var pixelAspectRatio = 1;\n")
	jsx.WriteString("app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio);\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 为主文档定义变量\n")
	jsx.WriteString("var masterDocument = app.activeDocument;\n")
	jsx.WriteString("\n")
	jsx.WriteString(" // 获取当前脚本所在路径\n")
	jsx.WriteString("var scriptPath = (new File($.fileName)).parent;\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 为指定路径文件定义变量\n")
	jsx.WriteString("var fileRef = new File(scriptPath+\"/../PSD/ColorCard.psd\");\n")
	jsx.WriteString("if(fileRef.exists){    // 如果图像存在\n")
	jsx.WriteString("    app.open(fileRef);   // 打开文档\n")
	jsx.WriteString("    selectAllLayers();   // 选择全部图层\n")
	jsx.WriteString("    app.activeDocument.activeLayer.duplicate (masterDocument); // 复制到主文档\n")
	jsx.WriteString("    app.activeDocument.close(SaveOptions.DONOTSAVECHANGES); // 关闭临时文档而不保存更改\n")
	jsx.WriteString("    //app.activeDocument = app.documents.getByName(\"临时效果图\");    //  切换窗口到这里\n")
	jsx.WriteString("   app.activeDocument.activeLayer = app.activeDocument.layers[1];  // 选择背景图层\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 选择全部图层，大神代码\n")
	jsx.WriteString("function selectAllLayers() {\n")
	jsx.WriteString("    var desc29 = new ActionDescriptor();\n")
	jsx.WriteString("    var ref23 = new ActionReference();\n")
	jsx.WriteString("    ref23.putEnumerated(charIDToTypeID('Lyr '), charIDToTypeID('Ordn'), charIDToTypeID('Trgt'));\n")
	jsx.WriteString("    desc29.putReference(charIDToTypeID('null'), ref23);\n")
	jsx.WriteString("    executeAction(stringIDToTypeID('selectAllLayers'), desc29, DialogModes.NO);\n")
	jsx.WriteString("}\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/newTempDocument.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/newTempDocument.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}
