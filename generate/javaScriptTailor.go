package generate

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"strings"
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
        app.load(new File(scriptPath + "/Tailor.jsx"));
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/selectTailor.jsx", script)
}

// 生成大部分框架的自动裁剪，例如左右镂空，小座屏等
func Tailor(frameName string) {
	var jsx = strings.Builder{}

	jsx.WriteString("// 定义一个函数用来设置黑边\r\n")
	jsx.WriteString("function addEdge() {\r\n")
	jsx.WriteString("    // 保存当前背景颜色\r\n")
	jsx.WriteString("    var nowColor = app.backgroundColor;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 定义一个对象颜色是黑色\r\n")
	jsx.WriteString("    var black = new SolidColor();\r\n")
	jsx.WriteString("    black.rgb.hexValue = \"000000\";\r\n")
	jsx.WriteString("    app.backgroundColor = black;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 获取当前文档的高度与宽度\r\n")
	jsx.WriteString("    var width = app.activeDocument.width + 0.5;\r\n")
	jsx.WriteString("    var height = app.activeDocument.height + 0.5;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 重设画布大小\r\n")
	jsx.WriteString("    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 恢复之前的背景颜色\r\n")
	jsx.WriteString("    app.backgroundColor = nowColor;\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 清理元数据\r\n")
	jsx.WriteString("function deleteDocumentAncestorsMetadata() {\r\n")
	jsx.WriteString("    // 清理元数据四步骤\r\n")
	jsx.WriteString("    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject(\"lib:AdobeXMPScript\");\r\n")
	jsx.WriteString("    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);\r\n")
	jsx.WriteString("    // Begone foul Document Ancestors!\r\n")
	jsx.WriteString("    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, \"DocumentAncestors\");\r\n")
	jsx.WriteString("    app.activeDocument.xmpMetadata.rawData = xmp.serialize();\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 全部整合在一起\r\n")
	jsx.WriteString("function optimized() {\r\n")
	jsx.WriteString("    // 设置首选项新文档预设单位是厘米，PIXELS是像素\r\n")
	jsx.WriteString("    app.preferences.rulerUnits = Units.CM;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 新建一个空白图层用于合并\r\n")
	jsx.WriteString("    app.activeDocument.artLayers.add();\r\n")
	jsx.WriteString("    // 合并全部可见图层\r\n")
	jsx.WriteString("    app.activeDocument.mergeVisibleLayers();\r\n")
	jsx.WriteString("    // 转为背景图层不然添加黑边会无效\r\n")
	jsx.WriteString("    app.activeDocument.activeLayer.isBackgroundLayer = true\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    if (BlackEdge) {\r\n")
	jsx.WriteString("        // 添加黑边\r\n")
	jsx.WriteString("        addEdge();\r\n")
	jsx.WriteString("    }\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 清理元数据\r\n")
	jsx.WriteString("    deleteDocumentAncestorsMetadata()\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    //定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。\r\n")
	jsx.WriteString("    var exportOptionsSave = new JPEGSaveOptions();\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 嵌入彩色配置文件\r\n")
	jsx.WriteString("    exportOptionsSave.embedColorProfile = true;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 设置杂边为无\r\n")
	jsx.WriteString("    exportOptionsSave.matte = MatteType.NONE;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    //设置导出文档时，图片的压缩质量。数字范围为1至12。\r\n")
	jsx.WriteString("    exportOptionsSave.quality = 12;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 保存为基线已优化\r\n")
	jsx.WriteString("    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 获取当前文档的文件名\r\n")
	jsx.WriteString("    var name = app.activeDocument.name\r\n")
	jsx.WriteString("    var TmpFile1 = new File(\"~/Desktop/GoCutting/\" + name);\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )\r\n")
	jsx.WriteString("    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档\r\n")
	jsx.WriteString("    app.activeDocument.saveAs(TmpFile1.saveDlg(\"优化另存为\", \"JPEG Files: *.jpg\"), exportOptionsSave, true, Extension.LOWERCASE);\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 判断是否有打开的文件\r\n")
	jsx.WriteString("if (!documents.length) {\r\n")
	jsx.WriteString("    alert(\"没有打开的文档，请打开一个文档来运行此脚本！\");\r\n")
	jsx.WriteString("    // return;\r\n")
	jsx.WriteString("} else {\r\n")
	jsx.WriteString("    // 是否自动黑边\r\n")
	jsx.WriteString(fmt.Sprintf("    var BlackEdge = %t;   // 这里传golang变量哦！！！！！！！！！！！！！！\n", viper.GetBool("blackEdge")))
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 保存开始的历史记录状态\r\n")
	jsx.WriteString("    var savedState = app.activeDocument.activeHistoryState;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 如果出错就返回最开始\r\n")
	jsx.WriteString("    try {\r\n")
	jsx.WriteString("        // 生成历史记录\r\n")
	jsx.WriteString("        app.activeDocument.suspendHistory(\"优化另存\", \"optimized()\");\r\n")
	jsx.WriteString("    } catch (error) {\r\n")
	jsx.WriteString("        // 忽略错误\r\n")
	jsx.WriteString("    }\r\n")
	jsx.WriteString("    // 当你完成了你正在做的任何事情，返回这个状态\r\n")
	jsx.WriteString("    app.activeDocument.activeHistoryState = savedState;\r\n")
	jsx.WriteString("}")

	// 转成字符串格式
	jsxStr := jsx.String()

	// 框架名不是空，就生成专属裁剪脚本
	if frameName != "" {
		tools.CreateFile(fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName), jsxStr)
	} else {
		tools.CreateFile("config/jsx/tailor.jsx", jsxStr)
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
func TailorForMap6(width, height, number float64, frameName, singleName string) {
	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width      int // 像素没有小数点
		Height     int
		Number     int    // 几片折屏
		SingleName string // 单片名字
	}{int(width * 10), int(height * 10), int(number), singleName}

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
