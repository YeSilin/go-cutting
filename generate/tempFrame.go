package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strings"
)


// 生成小座屏临时框架
func TempFrame1JS(width, height float64) {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("/* Code by  Mike Hale http://www.ps-scripts.com/bb/viewtopic.php?f=14&t=1802&start=15\n")
	jsx.WriteString("with small modification by Vladimir Carrer\n")
	jsx.WriteString("*/\n")
	jsx.WriteString("// 又是一个大神代码，用于创建矢量矩形\n")
	jsx.WriteString("function DrawShape() {\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var doc = app.activeDocument;\n")
	jsx.WriteString("    var y = arguments.length;\n")
	jsx.WriteString("    var i = 0;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var lineArray = [];\n")
	jsx.WriteString("    for (i = 0; i < y; i++) {\n")
	jsx.WriteString("        lineArray[i] = new PathPointInfo;\n")
	jsx.WriteString("        lineArray[i].kind = PointKind.CORNERPOINT;\n")
	jsx.WriteString("        lineArray[i].anchor = arguments[i];\n")
	jsx.WriteString("        lineArray[i].leftDirection = lineArray[i].anchor;\n")
	jsx.WriteString("        lineArray[i].rightDirection = lineArray[i].anchor;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var lineSubPathArray = new SubPathInfo();\n")
	jsx.WriteString("    lineSubPathArray.closed = true;\n")
	jsx.WriteString("    lineSubPathArray.operation = ShapeOperation.SHAPEADD;\n")
	jsx.WriteString("    lineSubPathArray.entireSubPath = lineArray;\n")
	jsx.WriteString("    var myPathItem = doc.pathItems.add(\"myPath\", [lineSubPathArray]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var desc88 = new ActionDescriptor();\n")
	jsx.WriteString("    var ref60 = new ActionReference();\n")
	jsx.WriteString("    ref60.putClass(stringIDToTypeID(\"contentLayer\"));\n")
	jsx.WriteString("    desc88.putReference(charIDToTypeID(\"null\"), ref60);\n")
	jsx.WriteString("    var desc89 = new ActionDescriptor();\n")
	jsx.WriteString("    var desc90 = new ActionDescriptor();\n")
	jsx.WriteString("    var desc91 = new ActionDescriptor();\n")
	jsx.WriteString("    desc91.putDouble(charIDToTypeID(\"Rd  \"), 200); // R\n")
	jsx.WriteString("    desc91.putDouble(charIDToTypeID(\"Grn \"), 200); // G\n")
	jsx.WriteString("    desc91.putDouble(charIDToTypeID(\"Bl  \"), 200); // B\n")
	jsx.WriteString("    var id481 = charIDToTypeID(\"RGBC\");\n")
	jsx.WriteString("    desc90.putObject(charIDToTypeID(\"Clr \"), id481, desc91);\n")
	jsx.WriteString("    desc89.putObject(charIDToTypeID(\"Type\"), stringIDToTypeID(\"solidColorLayer\"), desc90);\n")
	jsx.WriteString("    desc88.putObject(charIDToTypeID(\"Usng\"), stringIDToTypeID(\"contentLayer\"), desc89);\n")
	jsx.WriteString("    executeAction(charIDToTypeID(\"Mk  \"), desc88, DialogModes.NO);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    myPathItem.remove();\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// X,Y\n")
	jsx.WriteString("// Put the coordinates in clockwise order\n")
	jsx.WriteString("//DrawShape([100, 100], [100, 200], [200, 200], [200, 100]);\n")
	jsx.WriteString("//DrawShape([512, 128], [600, 256], [684, 320], [600, 386], [686, 514], [512, 450],[340,512],[428,386],[340,320],[428,256]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 创建矩形  左上角xy，宽高，透明度，是否使用前景色填充\n")
	jsx.WriteString("function createRectangle(x, y, w, h, transparent, modifyColor,name) {\n")
	jsx.WriteString("    // 这个函数是顺时针坐标定位\n")
	jsx.WriteString("    DrawShape([x, y], [x+w, y], [x+w, y+h], [x, y+h]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("   // 不透明度改80%\n")
	jsx.WriteString("    app.activeDocument.activeLayer.opacity=transparent;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    app.activeDocument.activeLayer.name=name\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 是否使用前景色\n")
	jsx.WriteString("    if (modifyColor) {\n")
	jsx.WriteString("        // 用前景色填充\n")
	jsx.WriteString("        var idFl = charIDToTypeID( \"Fl  \" );\n")
	jsx.WriteString("        var desc76 = new ActionDescriptor();\n")
	jsx.WriteString("        var idUsng = charIDToTypeID( \"Usng\" );\n")
	jsx.WriteString("        var idFlCn = charIDToTypeID( \"FlCn\" );\n")
	jsx.WriteString("        var idFrgC = charIDToTypeID( \"FrgC\" );\n")
	jsx.WriteString("        desc76.putEnumerated( idUsng, idFlCn, idFrgC );\n")
	jsx.WriteString("        executeAction( idFl, desc76, DialogModes.NO );\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 创建小座屏\n")
	jsx.WriteString("function createBaseScreen(inputWidth, inputHeight){\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 左上角的坐标需要使用parseInt转成数字\n")
	jsx.WriteString("    var x = parseInt(app.activeDocument.width/2-inputWidth/2)\n")
	jsx.WriteString("    var y = parseInt(app.activeDocument.height/2-inputHeight/2)\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 新建图层组\n")
	jsx.WriteString("    var newLayerSetRef =app.activeDocument.layerSets.add()\n")
	jsx.WriteString("    // 图层组设置名字\n")
	jsx.WriteString("    newLayerSetRef.name=\"layerset-\"+inputWidth/10+\"-\"+inputHeight/10\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 右底座\n")
	jsx.WriteString("    createRectangle(x+inputWidth+4,y+inputHeight-360,40,400,100,true,\"右底座\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 左底座\n")
	jsx.WriteString("    createRectangle(x-44,y+inputHeight-360,40,400,100,true,\"左底座\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 创建画布\n")
	jsx.WriteString("    createRectangle(x,y,inputWidth,inputHeight,80,false,\"画布\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 左边框\n")
	jsx.WriteString("    createRectangle(x,y,40,inputHeight,100,true,\"左\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 右边框\n")
	jsx.WriteString("    createRectangle(x+inputWidth-40,y,40,inputHeight,100,true,\"右\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 上边框\n")
	jsx.WriteString("    createRectangle(x,y,inputWidth,40,100,true,\"上\")\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 下边框\n")
	jsx.WriteString("    createRectangle(x,y+inputHeight-40,inputWidth,40,100,true,\"下\")\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 主函数\n")
	jsx.WriteString("function main(){\n")
	jsx.WriteString("    if(!documents.length) {\n")
	jsx.WriteString("        alert(\"没有打开的文档，请先使用软件新建专属文档！\")\n")
	jsx.WriteString("        return;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 打开前景色的拾色器\n")
	jsx.WriteString("    // 返回值是布尔类型，如果用户没有选择颜色返回假\n")
	jsx.WriteString("    if(!showColorPicker()) {\n")
	jsx.WriteString("        return;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 用户输入的宽和高\n")
	jsx.WriteString(fmt.Sprintf("    var inputWidth = %.2f     // 这里传golang变量哦！！！！！！！！！！！！！！！！！！！！\n", width))
	jsx.WriteString(fmt.Sprintf("    var inputHeight = %.2f    // 这里传golang变量哦！！！！！！！！！！！！！！！！！！！！\n", height))
	jsx.WriteString("\n")
	jsx.WriteString("    // 生成历史记录\n")
	jsx.WriteString("    app.activeDocument.suspendHistory(\"生成小座屏框架！\", \"createBaseScreen(inputWidth, inputHeight)\");\n")
	jsx.WriteString("\n")
	jsx.WriteString(" }\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 调用\n")
	jsx.WriteString("main ()\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/GenerateTempFrame.jsx", jsxStr)

	// 同时运行运行脚本
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/GenerateTempFrame.jsx"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}