package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strings"
	"time"
)

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
	tools.CreateFile("resources/jsx/newTempDocument.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "resources/jsx/newTempDocument.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}
