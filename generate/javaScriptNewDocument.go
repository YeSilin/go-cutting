package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strings"
	"time"
)

//生成用来新建ps文档js
//@param width 传入宽度
//@param height 传入高度
//@param wordLine 是否创建文字不要超过的参考线
func NewDocumentJS(width, height float64, frameName string, wordLine bool) {
	var jsx = strings.Builder{}

	jsx.WriteString("app.preferences.rulerUnits = Units.CM;\n")
	jsx.WriteString(fmt.Sprintf("var width = %.2f;\n", width))
	jsx.WriteString(fmt.Sprintf("var height = %.2f;\n", height))
	jsx.WriteString("var resolution = 100;\n")
	jsx.WriteString(fmt.Sprintf("var docName = \"%s\";\n", frameName))
	jsx.WriteString("var mode = NewDocumentMode.CMYK;\n")
	jsx.WriteString("var initialFill = DocumentFill.WHITE;\n")
	jsx.WriteString("var pixelAspectRatio = 1;\n")
	jsx.WriteString("app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio);\n")

	if wordLine {
		jsx.WriteString("app.activeDocument.suspendHistory(\"建议：字不要在此参考线外！\", \"addLine()\");  // 生成历史记录\n")
		jsx.WriteString("function addLine(){   // 定义一个函数用于新建参考线\n")
		jsx.WriteString("    // 添加文字垂直参考线\n")
		jsx.WriteString("    activeDocument.guides.add (Direction.VERTICAL,UnitValue(\"5cm\"));\n")
		jsx.WriteString(fmt.Sprintf("    activeDocument.guides.add (Direction.VERTICAL,UnitValue(\"%.2fcm\"));\n", width-5))
		jsx.WriteString("    // 添加文字水平参考线\n")
		jsx.WriteString("    activeDocument.guides.add (Direction.HORIZONTAL,UnitValue(\"5cm\"));\n")
		jsx.WriteString(fmt.Sprintf("    activeDocument.guides.add (Direction.HORIZONTAL,UnitValue(\"%.2fcm\"));\n", height-5))
		jsx.WriteString("}\n")
	}

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/NewDocumentJS.jsx", jsxStr)
}

// 生成用来新建ps文档3d作图js
// @param width 传入宽度
// @param height 传入高度
func NewDocument3DMapJS(width, height float64, frameName string) {
	var jsx = strings.Builder{}

	jsx.WriteString("#target photoshop\n")
	jsx.WriteString("// 设置首选项新文档预设单位是厘米，PIXELS是像素\n")
	jsx.WriteString("app.preferences.rulerUnits = Units.PIXELS;\n")
	jsx.WriteString(fmt.Sprintf("var width = %f;\n", width*10))
	jsx.WriteString(fmt.Sprintf("var height = %f;\n", height*10))
	jsx.WriteString("// 定义一个变量[resolution]，表示新文档的分辨率。\n")
	jsx.WriteString("var resolution = 72;\n")
	jsx.WriteString(fmt.Sprintf("var docName = \"%s\";\n", frameName))
	jsx.WriteString("//定义一个变量[mode]，表示新文档的颜色模式。\n")
	jsx.WriteString("var mode = NewDocumentMode.RGB;\n")
	jsx.WriteString("var initialFill = DocumentFill.WHITE;\n")
	jsx.WriteString("var pixelAspectRatio = 1;\n")
	jsx.WriteString("app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio);\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/NewDocumentJS.jsx", jsxStr)
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
	tools.CreateFile("Config/JSX/NewTempDocument.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "Config/jsx/NewTempDocument.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}
