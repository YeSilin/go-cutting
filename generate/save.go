package generate

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// 暗号-98的实现
func SaveForWeb(originalPath string) {
	// 返回绝对路径
	originalPath, err := filepath.Abs(originalPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	// 全部换成正斜杠
	originalPath = strings.Replace(originalPath, "\\", "/", -1)
	// 修改成js脚本可以看懂的路径
	originalPath = "/" + strings.Replace(originalPath, ":", "", 1)

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("config/jsx/template/saveForWeb.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("config/jsx/saveForWeb.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 保存路径定义
	savePath := fmt.Sprintf("%s/主图/DP.jpg", originalPath)

	// 利用给定数据渲染模板，并将结果写入f
	_ = tmpl.Execute(f, savePath)
}

//如果画布的高和宽同时大于148则提示
func MaxCanvas(width, height float64) {
	// 默认不写入
	write := false

	if width > 180 && height > 180 {
		// 这其实是红色
		color.LightBlue.Println("（已超不透最大180cm。）")
		write = true
	} else if width > 148 && height > 148 {
		color.LightBlue.Println("（已超半透最大148cm。）")
		write = true
	} else {
		fmt.Println()
		return
	}

	// 如果要写入
	if write {
		//最大画布字体图层提示
		var jsx = strings.Builder{}

		jsx.WriteString("// 无聊加了个画布大小判断\n")
		jsx.WriteString(fmt.Sprintf("if ((%f>180) && (%f>180)) {\n", width, height))
		jsx.WriteString("     // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超不透最大180cm\", \"maxCanvas(\\\"注意：已超不透最大180cm。\\\",  \\\"9d2e2d\\\")\");\n")
		jsx.WriteString(fmt.Sprintf("} else if ((%f>148) && (%f>148)) {\n", width, height))
		jsx.WriteString("      // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超半透最大148cm\", \"maxCanvas(\\\"注意：已超半透最大148cm。\\\",  \\\"77bb11\\\")\");\n")
		jsx.WriteString("}\n")
		jsx.WriteString("\n")
		jsx.WriteString("\n")
		jsx.WriteString("function maxCanvas(text, rgbValue){\n")
		jsx.WriteString("    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。\n")
		jsx.WriteString("    var newLayer = app.activeDocument.artLayers.add();\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。\n")
		jsx.WriteString("    newLayer.kind = LayerKind.TEXT;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //把图层 newLayer 的文本内容类型变为”文本框“。\n")
		jsx.WriteString("    newLayer.textItem.kind = TextType.PARAGRAPHTEXT;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置图层 newLayer 的文本框宽度与高度。\n")
		jsx.WriteString("    newLayer.textItem.width = app.activeDocument.width*0.8;\n")
		jsx.WriteString("    newLayer.textItem.height = app.activeDocument.width*0.1;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素。\n")
		jsx.WriteString("    //newLayer.textItem.position= [UnitValue(\"50px\"), UnitValue(\"100px\")]\n")
		jsx.WriteString("    newLayer.textItem.position= [UnitValue(app.activeDocument.width*0.1), UnitValue((app.activeDocument.height*0.5)-(app.activeDocument.width*0.025))];\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本字体大小为“40 点”。\n")
		jsx.WriteString("    newLayer.textItem.size = UnitValue(app.activeDocument.width*0.05);\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本内容。\n")
		jsx.WriteString("    newLayer.textItem.contents= text;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本框对齐方式为居中对齐。\n")
		jsx.WriteString("    newLayer.textItem.justification = Justification.CENTER;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //创建一个色彩变量 c   ，颜色为 #77bb11。\n")
		jsx.WriteString("    var c = new SolidColor();\n")
		jsx.WriteString("    c.rgb.hexValue = rgbValue;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本颜色为 c。\n")
		jsx.WriteString("    newLayer.textItem.color = c;\n")
		jsx.WriteString("}\n")

		// 转成字符串格式
		jsxStr := jsx.String()
		// 追加写入
		tools.WriteFile("config/jsx/newDocument.jsx", jsxStr)
	}
}

//生成自带清除元数据的另存
func SaveAsJPEG() {
	const script = `// 用于快捷另存为jpg

// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}

// 更新所有链接的智能对象
function updateAllModified() {
    var idplacedLayerUpdateAllModified = stringIDToTypeID( "placedLayerUpdateAllModified" );
    executeAction( idplacedLayerUpdateAllModified, undefined, DialogModes.NO );
}

// 全部整合在一起
function optimized() {
    // 清理元数据
    deleteDocumentAncestorsMetadata()

	// 更新所有链接的智能对象
	updateAllModified() 

    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    // 设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

	// 获取当前文档的文件名并分割
    var name = app.activeDocument.name.split(".")
    var TmpFile = new File("~/Desktop/GoCutting/" + name[0]+" 副本");

    // 保存文件类型
    // var saveType = new Array("JPEG Files: *.jpg", "PNG Files: *.png");

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档
    app.activeDocument.saveAs(TmpFile.saveDlg("另存为", "JPEG Files: *.jpg"), exportOptionsSave, true, Extension.LOWERCASE);
}


// 判断是否有打开的文件
if (!documents.length) {
    //alert("没有打开的文档，请打开一个文档来运行此脚本！");
    // return;
} else {
    // 如果出错就返回最开始
    try {
        // 生成历史记录
        app.activeDocument.suspendHistory("另存为", "optimized()");
    } catch (error) {
        // 忽略错误
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/saveAsJPEG.jsx", script)
}

//生成自带清除元数据的另存全部打开的文件
func SaveAllJPEG() {
	const script = `// 用于快捷另存全部打开的文件

// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript")
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData)
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors")
    app.activeDocument.xmpMetadata.rawData = xmp.serialize()
}

// 更新所有链接的智能对象
function updateAllModified() {
    var idplacedLayerUpdateAllModified = stringIDToTypeID("placedLayerUpdateAllModified")
    executeAction(idplacedLayerUpdateAllModified, undefined, DialogModes.NO)
}


//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName)
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("优化另存为", ["不要带扩展名:*", "默认保存为 JPG 文件:*"]))
}


// 另存为web格式png
function saveAsWebPNG(savePath) {
    // 定义文件保存位置
    var savePath = new File(savePath + ".png")
    var pngOpts = new ExportOptionsSaveForWeb()
    pngOpts.format = SaveDocumentType.PNG   // 保存为png
    pngOpts.transparency = true  // 透明度
    pngOpts.interlaced = false  // 交错
    pngOpts.blur = 0    // 默认 0.0 不模糊。
    pngOpts.PNG8 = false   // 为真保存为 PNG8 ，否则保存为 PNG24
    activeDocument.exportDocument(savePath, ExportType.SAVEFORWEB, pngOpts)
}

//  负责另存文件的函数
function saveFileAs(savePath) {
    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions()

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE

    // 设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true

    // 实例化文件
    var TmpFile = new File(savePath)

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档。
    app.activeDocument.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE)
}


// 把前面的操作整合成一条历史记录
function start(userSavePath) {
    // 新建一个空白图层用于更新，不然更新会报错
    app.activeDocument.artLayers.add()

    // 更新所有链接的智能对象
    updateAllModified()

    // 删除刚刚新建的空白图层
    app.activeDocument.activeLayer.remove()

    // 另存文件
    saveFileAs(userSavePath)

    // 如果是白底图再存一份png
    if (userSavePath.indexOf("白底图") != -1) {
        saveAsWebPNG(userSavePath)
    }
}


// 得到路径前面全部和最后一个斜杠后的名字
function pathSplit(path) {
    var arr = path.split("/")

    // 得到最后一个元素
    var after = arr[arr.length - 1]

    // 删除最后一个元素
    path = path.replace(after, "")

    return [path, after]
}


// 把循环保存归成函数，用于添加进度条
function saveAll(fileNameArr) {
    // 开始循环保存
    for (var i = 0; i < app.documents.length; i++) {
        app.activeDocument = app.documents[i]

        // 清理元数据
        deleteDocumentAncestorsMetadata()

        // 生成历史记录并调用函数
        app.activeDocument.suspendHistory("更新所有修改的智能对象（储存为）", "start(fileNameArr[i])");
    }
}


// 主函数
function main() {
    // 得到用户想要保存的位置
    var userSavePath = getPathName("默认名")
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return
    }


    var pathArr = pathSplit(userSavePath)
    // 如何用户没有指定后缀，就统一改成副本这个后缀
    if (pathArr[1] == "默认名") {
        pathArr[1] = "副本"
    }

    // 得到最终会保存出来的文件名
    var fileNameArr = new Array();
    for (var i = 0; i < app.documents.length; i++) {
        // 获取指定文档的文件名并分割
        var nameArr = app.documents[i].name.split(".")

        // 保存文件名为 路径 + 当前文档名字 + 自定义后缀
        fileNameArr[i] = pathArr[0] + nameArr[0] + " " + pathArr[1]
    }

    // 遍历全部可能覆盖的文件名
    for (var i = 0; i < fileNameArr.length; i++) {
        // 避免覆盖已保存的文件
        if (new File(fileNameArr[i] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return
        }

        // 包含白底图的文件名，此文件名带完整路径
        if (fileNameArr[i].indexOf("白底图") != -1) {
            // 避免覆盖已保存的白底图
            if (new File(fileNameArr[i] + ".png").exists) {
                alert("输入的编号重复，已自动取消操作！");
                return
            }
        }
    }


    // 生成进度条调并调用函数
    app.doForcedProgress("正在另存全部... ", "saveAll(fileNameArr)")
}


// 判断是否有打开的文件
if (!documents.length) {
    //alert("没有打开的文档，请打开一个文档来运行此脚本！");
    // return;
} else {
    // 如果出错就返回最开始
    try {
        main()
    } catch (error) {
        // 忽略错误
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/saveAllJPEG.jsx", script)
}

// 生成保存并关闭全部文档的脚本
func SaveAndCloseAllDocuments() {
	const script = `// 保存并关闭所有文件
function saveAndClose() {
    // 得到要保存关闭的文档的数量，主要是documents关闭后会直接刷新
    const count = documents.length
 
    // 循环保存所有
    for (var i = 0; i < count; i++) {

        try {
            // 最后保存并关闭
            app.activeDocument.close(SaveOptions.SAVECHANGES);
        } catch (error) {
            // 忽略错误
        }
    }

}

// 添加进度条
function progressBar() {
    app.doForcedProgress("正在保存并关闭全部文件... ", "saveAndClose()")
}

// 主函数
function main() {
    if (!documents.length) {
        //alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 运行进度条
    progressBar()
}

// 运行
main()`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/saveAndCloseAllDocuments.jsx", script)
}
