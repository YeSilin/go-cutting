package model

// 这里放一些Photoshop保存文档的脚本
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// SaveAsJPEG 储存为jpeg格式的初始化 暗号-10的实现
func SaveAsJPEG() {
	const script = `// 储存为JPEG格式并清理与更新链接的智能对象

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
    // 更新智能对象很有可能出错
    try {
        var idplacedLayerUpdateAllModified = stringIDToTypeID("placedLayerUpdateAllModified");
        executeAction(idplacedLayerUpdateAllModified, undefined, DialogModes.NO);
    } catch (error) {
        // 忽略错误
    }
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
    var TmpFile = new File("~/Desktop/GoCutting/" + name[0] + " 副本");

    // 保存文件类型
    // var saveType = new Array("JPEG Files: *.jpg", "PNG Files: *.png");

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档
    app.activeDocument.saveAs(TmpFile.saveDlg("储存为", "JPEG Files: *.jpg"), exportOptionsSave, true, Extension.LOWERCASE);
}


// 判断是否有打开的文件
if (!documents.length) {
    //alert("没有打开的文档，请打开一个文档来运行此脚本！");
    // return;
} else {
    // 如果出错就返回最开始
    try {
        // 生成历史记录
        app.activeDocument.suspendHistory("储存为", "optimized()");
    } catch (error) {
        // 忽略错误
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("data/jsx/saveAsJPEG.jsx", script)
}

// SaveAllJPEG 将所有打开的文档储存为jpeg格式的初始化 暗号-11的实现
func SaveAllJPEG() {
	const script = `// 用于快捷另存全部打开的文件

// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript")
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData)
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors")
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}

// 更新所有链接的智能对象
function updateAllModified() {
    //更新智能对象很有可能出错
    try {
        var idplacedLayerUpdateAllModified = stringIDToTypeID("placedLayerUpdateAllModified");
        executeAction(idplacedLayerUpdateAllModified, undefined, DialogModes.NO);
    } catch (error) {
        // 忽略错误
    }
}

// 打开资源管理器文件夹窗口
function openFolder(path) {
    // execute在操作系统中打开文件或运行程序
    Folder(path).execute();
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

    // 得到保存位置的数组，索引0为路径，索引1为后缀
    var pathArr = pathSplit(userSavePath)
    // 如何用户没有指定后缀，就统一改成副本这个后缀
    if (pathArr[1] == "默认名") {
        pathArr[1] = "副本"
    }

    // 得到最终会保存出来的文件名
    var fileNameArr = new Array();
    for (var i = 0; i < app.documents.length; i++) {
        // 获取指定文档的文件名并分割
        var nameArr = app.documents[i].name.split(".");

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
    app.doForcedProgress("正在另存全部... ", "saveAll(fileNameArr)");
    // 执行完打开文件夹
    openFolder(pathArr[0]);
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
	tools.CreateFile("data/jsx/saveAllJPEG.jsx", script)
}

// SaveAndCloseAllDocuments 保存并关闭全部文档的初始化 暗号-12的实现
func SaveAndCloseAllDocuments() {
	const script = `// 保存并关闭全部文档
function saveAndCloseAllDocuments() {
    // 得到要保存关闭的文档的数量，主要是documents关闭后会直接刷新
    const count = documents.length;
 
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
    app.doForcedProgress("正在保存并关闭全部文件... ", "saveAndCloseAllDocuments()");
}

// 主函数
function main() {
    if (!documents.length) {
        // alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 运行进度条
    progressBar();
}

// 运行
main();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("data/jsx/saveAndCloseAllDocuments.jsx", script)
}

// SaveForWeb 导出web格式脚本的初始化 暗号-98的实现 副作用分辨率会被强制修改为72ppi 并且无法选择保存路径
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
	tmpl, err := template.ParseFiles("data/jsx/template/saveForWeb.gohtml")
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("data/jsx/saveForWeb.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 保存路径定义
	savePath := fmt.Sprintf("%s/主图/", originalPath)

	// 利用给定数据渲染模板，并将结果写入f
	_ = tmpl.Execute(f, savePath)
}

// LoadSaveScript 根据当前文档名选择正确的快捷裁剪脚本
func LoadSaveScript() {
	const script = `// 载入一个调用针对当前文档的专属另存脚本
function loadSaveScript() {
    // 判断是否有打开的文件
    if (!documents.length) {
        // alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 获取当前脚本所在路径
    const scriptPath = (new File($.fileName)).parent;

    // 获取当前文档名字
    var docName = app.activeDocument.name;
    // 如果有文件后缀名就去掉
    const index = docName.lastIndexOf(".")
    if (index != -1) {
        docName = docName.substring(0, index);
    }

    // 要运行的脚本路径
    const runScript = scriptPath + "/temp/tailor_" + docName + ".jsx";

    // 获得脚本对象
    var fileRef = new File(runScript);

    // 如果脚本存在
    if (fileRef.exists) {
        // 运行脚本
        app.load(fileRef);
        return;
    }

    // 前面没有直接返回说明专属导出脚本不存在
    var answer = confirm("未找到当前文档专属脚本，是否调用默认脚本？", false, "储存为 JPG 格式副本")

    // 如果确定运行默认脚本
    if (answer) {
        app.load(new File(scriptPath + "/frameSaveDef.jsx"));
    }
}

// 载入专属另存脚本
loadSaveScript();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("data/jsx/loadSaveScript.jsx", script)
}

// FrameSaveDef 生成大部分框架的自动裁剪，例如左右镂空，小座屏等
func FrameSaveDef(frameName string) {
	const script = `// 定义一个函数用来设置黑边
function addBlackEdge() {
    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width.value + 0.1;
    var height = app.activeDocument.height.value + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(width, height, AnchorPosition.MIDDLECENTER);
    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    const layer = function () {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 清理元数据
function deleteDocumentAncestorsMetadata() {
    const clear = function () {
        // 清理元数据四步骤
        if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
        var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
        // Begone foul Document Ancestors!
        xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
        app.activeDocument.xmpMetadata.rawData = xmp.serialize();
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("清理元数据", "clear()");
}


// 储存为 JPEG 格式副本
function saveAsJPEG() {
    var name = app.activeDocument.name

    // 获取当前文档的文件名
    var TmpFile = new File("~/Desktop/GoCutting/" + name);
    // 获取用户自定义储存位置
    TmpFile = TmpFile.saveDlg("储存副本", "JPEG Files: *.jpg")

    // 如果用户取消储存就直接返回
    if (TmpFile == null) {
        return false;
    }

    //定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();
    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;
    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;
    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;
    //设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    app.activeDocument.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE);

    return true;
}


// 默认的框架储存脚本
function frameSaveDef() {
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();

    // 添加黑边
    addBlackEdge();

    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK)
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";

    // 保存图片
    IsSaved = saveAsJPEG()
}


// 主函数
function main() {
    // 判断是否有打开的文件
    if (!documents.length) {
        // alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }
    // 清理元数据
    deleteDocumentAncestorsMetadata()
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;

    // 保存开始的历史记录状态
    var saveState = app.activeDocument.activeHistoryState;
    // 生成历史记录
    app.activeDocument.suspendHistory("储存副本", "frameSaveDef()");
    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = saveState;
    // 保存图片，同时成功就加个提示
    if (IsSaved) {
        createLayer()
    }
}


// 是否已保存
var IsSaved = false;
main();`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		BlackEdge bool // 是否自动黑边

	}{viper.GetBool("blackEdge")}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 生成通用的文件名字
	fileName := "data/jsx/frameSaveDef.jsx"
	// 框架名不是空，就生成专属裁剪脚本名字
	if frameName != "" {
		fileName = fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName)
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

// FrameSave3 左右画布专属保存
func FrameSave3(width, height, hollow float64, frameName string) {
	const script = `// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 定义一个函数用来设置黑边
function addEdge() {
    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width + 0.1;
    var height = app.activeDocument.height + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    function layer() {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }

    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 合并图层用于提速
function speedUp() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();

    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK)
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";

    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true
}


//  返回上一级函数
function returnUpper() {
    // 为历史定义的变量
    var idslct = charIDToTypeID("slct");
    var idnull = charIDToTypeID("null");
    var idHstS = charIDToTypeID("HstS");
    var idOrdn = charIDToTypeID("Ordn");
    var idPrvs = charIDToTypeID("Prvs");

    // 返回上一级历史
    var ref = new ActionReference();
    var desc = new ActionDescriptor();
    ref.putEnumerated(idHstS, idOrdn, idPrvs);
    desc.putReference(idnull, ref);
    executeAction(idslct, desc, DialogModes.NO);
}


//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName);
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("优化另存为",["不要带扩展名:*","默认保存为 JPG 文件:*"]))
}


// 用于裁剪另存
function optimized(x1, y1, x2, y2, savePath) {
    var bounds0 = [x1, y1, x2, y2];
    document.crop(bounds0, 0);

    // 添加黑边
    addEdge();

    // 裁剪之后进行保存的位置和你想要的默认名称
    var TmpFile = new File(savePath);

    //saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档。
    document.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE);
}

// 主函数
function main() {
    // 得到用户想要保存的位置
    var userSavePath = getPathName("请输入订单编号");
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return
    }

    // 得到最终会保存出来的文件名
    var tempNameArr = new Array()
    tempNameArr[0] = userSavePath + "_左右画布_左_" + hollow + "x" + height;
    tempNameArr[1] = userSavePath + "_左右画布_中_" + width + "x" + height;
    tempNameArr[2] = userSavePath + "_左右画布_右_" + hollow + "x" + height;

    // 遍历全部可能覆盖的文件名
    for (var i = 0; i < tempNameArr.length; i++){
        // 避免覆盖已保存的文件
        if (new File(tempNameArr[i] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return
        }
    }

    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("性能加速", "speedUp()");

    // 清理元数据
    deleteDocumentAncestorsMetadata();


    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("优化另存为", "optimized(0,0,hollow,height,tempNameArr[0])");
    // 返回上一级历史
    returnUpper();


    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("优化另存为", "optimized(hollow,0,hollow+width,height,tempNameArr[1])");
    // 返回上一级历史
    returnUpper();


    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("优化另存为", "optimized(hollow+width,0,hollow+width+hollow,height,tempNameArr[2])");

    // 全部保存成功
    isSave = true;
}


// 判断是否有打开的文件
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    // 中间大两边小的镂空均是
    var hollow = {{.Hollow}};   // 这里传golang变量哦！！！！！！！！！！！！！！
    // 中间大两边小的中间宽
    var width = {{.Width}};    // 这里传golang变量哦！！！！！！！！！！！！！！
    // 中间大两边小的高
    var height = {{.Height}};    // 这里传golang变量哦！！！！！！！！！！！！！！

    // 是否保存成功
    var isSave = false;

    // 保存开始状态
    var savedState = app.activeDocument.activeHistoryState;

    // 为当前文档定义变量
    var document = app.activeDocument;

    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    // 设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // 执行主函数
    main();

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 保存成功就加个提示
    if (isSave) {
        createLayer();
    }
}`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width  float64
		Height float64
		Hollow float64 // 中间大两边小的镂空均是
	}{width, height, hollow}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName))
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

// FrameSave4to2  生成上下画布专属保存
func FrameSave4to2(width, upperCanvas, middleCanvas, downCanvas float64, frameType, frameName string) {
	const script = `// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 定义一个函数用来设置黑边
function addEdge() {
    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width + 0.1;
    var height = app.activeDocument.height + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    function layer() {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }

    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 合并图层用于提速
function speedUp() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();
    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK)
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true
}


//  返回上一级函数
function returnUpper() {
    // 为历史定义的变量
    var idslct = charIDToTypeID("slct");
    var idnull = charIDToTypeID("null");
    var idHstS = charIDToTypeID("HstS");
    var idOrdn = charIDToTypeID("Ordn");
    var idPrvs = charIDToTypeID("Prvs");

    // 返回上一级历史
    var ref = new ActionReference();
    var desc = new ActionDescriptor();
    ref.putEnumerated(idHstS, idOrdn, idPrvs);
    desc.putReference(idnull, ref);
    executeAction(idslct, desc, DialogModes.NO);
}


//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName);
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("优化另存为", ["不要带扩展名:*", "默认保存为 JPG 文件:*"]))
}


// 用于裁剪另存
function optimized(x1, y1, x2, y2, savePath) {
    var bounds0 = [x1, y1, x2, y2];
    document.crop(bounds0, 0);

    // 添加黑边
    addEdge();

    // 裁剪之后进行保存的位置和你想要的默认名称
    var TmpFile = new File(savePath);

    //saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档。
    document.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE);
}

// 主函数
function main() {
    // 得到用户想要保存的位置
    var userSavePath = getPathName("请输入订单编号");
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return
    }

    // 得到最终会保存出来的文件名
    var tempNameDic = new Array()

    // 如果有上画布就追加
    if (upperCanvas > 0) {
        tempNameDic["upper"] = userSavePath + "_" + frameName + "_上_" + width + "x" + upperCanvas
    }

    // 中间这个画布肯定存在，直接添加
    tempNameDic["middle"] = userSavePath + "_" + frameName + "_中_" + width + "x" + middleCanvas;

    // 如果有下画布就追加
    if (downCanvas > 0) {
        tempNameDic["down"] = userSavePath + "_" + frameName + "_下_" + width + "x" + downCanvas
    }


    // 遍历全部可能覆盖的文件名
    for (var key in tempNameDic) { // 输出字典元素，如果字典的key是数字，输出时会自动按序输出
        if (new File(tempNameDic[key] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return
        }
    }


    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("性能加速", "speedUp()");

    // 清理元数据
    deleteDocumentAncestorsMetadata();


    // 为上画布裁剪
    if (upperCanvas > 0) {
        app.activeDocument.suspendHistory("合并所有操作为一条历史", "optimized(0, 0, width, upperCanvas, tempNameDic[\"upper\"])")
        // 返回上一级历史
        returnUpper();
    }


    // 为中间画布调用裁剪函数
    app.activeDocument.suspendHistory("合并所有操作为一条历史", "optimized(0, upperCanvas, width, upperCanvas + middleCanvas, tempNameDic[\"middle\"])")
    // 返回上一级历史
    returnUpper();


    // 为下画布裁剪
    if (downCanvas > 0) {
        app.activeDocument.suspendHistory("合并所有操作为一条历史", "optimized(0, upperCanvas + middleCanvas, width, upperCanvas + middleCanvas + downCanvas, tempNameDic[\"down\"])")
        // 返回上一级历史
        returnUpper();
    }


    // 全部保存成功
    isSave = true;
}


// 判断是否有打开的文件
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {


    // 确定是 上画布 上下画布 下画布
    var frameName = "{{.FrameName}}" // 这里传golang变量哦！！！！！！！！！！！！！！

    // 上下画布的总宽
    var width = {{.Width}};    // 这里传golang变量哦！！！！！！！！！！！！！！
    
    // 上画布高度
    var upperCanvas = {{.UpperCanvas}};   // 这里传golang变量哦！！！！！！！！！！！！！！

    // 中画布高度
    var middleCanvas = {{.MiddleCanvas}}  // 这里传golang变量哦！！！！！！！！！！！！！！

    // 下画布高度
    var downCanvas = {{.DownCanvas}}    // 这里传golang变量哦！！！！！！！！！！！！！！

    // 是否保存成功
    var isSave = false;

    // 保存开始状态
    var savedState = app.activeDocument.activeHistoryState;

    // 为当前文档定义变量
    var document = app.activeDocument;

    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    // 设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // 执行主函数
    main();

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 保存成功就加个提示
    if (isSave) {
        createLayer();
    }
}`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		FrameName    string  // 确定是 上画布 上下画布 下画布
		Width        float64 // 总宽
		UpperCanvas  float64 // 上高
		MiddleCanvas float64 // 中高
		DownCanvas   float64 // 下高
	}{frameType, width, upperCanvas, middleCanvas, downCanvas}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName))
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

// FrameSave6  生成传统折屏专属保存
func FrameSave6(width, height, number float64, frameName, singleName string) {
	const script = `// 各种折屏的快捷裁剪实现

//var str = "js实现用{two}自符串替换占位符{two} {three}  {one} ".format({one: "I",two: "LOVE",three: "YOU"});
//var str2 = "js实现用{1}自符串替换占位符{1} {2}  {0} ".format("I","LOVE","YOU");
String.prototype.format = function () {
    if (arguments.length == 0) return this;
    var param = arguments[0];
    var s = this;
    if (typeof (param) == 'object') {
        for (var key in param)
            s = s.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
        return s;
    } else {
        for (var i = 0; i < arguments.length; i++)
            s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
        return s;
    }
};


// 定义一个函数用来设置黑边
function addEdge() {
    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width + 0.1;
    var height = app.activeDocument.height + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    function layer() {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }

    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 合并图层用于提速
function speedUp() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();
    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK);
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true
}

//  返回上一级函数
function returnUpper() {
    // 为历史定义的变量
    var idslct = charIDToTypeID("slct");
    var idnull = charIDToTypeID("null");
    var idHstS = charIDToTypeID("HstS");
    var idOrdn = charIDToTypeID("Ordn");
    var idPrvs = charIDToTypeID("Prvs");

    // 返回上一级历史
    var ref = new ActionReference();
    var desc = new ActionDescriptor();
    ref.putEnumerated(idHstS, idOrdn, idPrvs);
    desc.putReference(idnull, ref);
    executeAction(idslct, desc, DialogModes.NO);
}


//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName);
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("优化另存为",["不要带扩展名:*","默认保存为 JPG 文件:*"]))
}


// 这是负责裁剪另存的函数
function optimized(x1, y1, x2, y2, savePath) {
    var bounds = [x1, y1, x2, y2];
    app.activeDocument.crop(bounds, 0);

    // 添加黑边
    addEdge();

    // 裁剪之后进行保存的位置和你想要的默认名称
    var TmpFile = new File(savePath);

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档。
    app.activeDocument.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE);
}


// 主函数
function main() {
    // 得到用户想要保存的位置
    var userSavePath = getPathName("请输入订单编号");
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return
    }

    // 得到最终会保存出来的文件名
    var fileNameArr = new Array();
    for (var i = 0; i < num; i++) {
        fileNameArr[i] = "{0}_{1}_{2}-{3}_".format(userSavePath, singleName, i + 1, num) + width + "x" + height;
    }

    // 遍历全部可能覆盖的文件名
    for (var i = 0; i < fileNameArr.length; i++) {
        // 避免覆盖已保存的文件
        if (new File(fileNameArr[i] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return
        }
    }

    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("性能加速", "speedUp()");

    // 清除元数据
    deleteDocumentAncestorsMetadata();

    // 开始自动保存文件
    for (var i = 0; i < num; i++) {
        // 生成历史记录并调用函数
        app.activeDocument.suspendHistory("优化另存为", "optimized(width*i, 0, width*(i+1), height, fileNameArr[i])");

        // 返回上一级历史
        returnUpper();
    }

    // 全部保存成功
    isSave = true;
}


// 判断是否有打开的文件
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    //设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // 是否保存成功
    var isSave = false;

    // 保存活动历史记录状态
    var savedState = app.activeDocument.activeHistoryState;

    // 定义单片名字
    var singleName = "{{.SingleName}}";  // 这里传golang变量哦！！！！！！！！！！！
    // 定义折屏单扇的宽和高
    var width = {{.Width}};  // 这里传golang变量哦！！！！！！！！！！！
    var height = {{.Height}};  // 这里传golang变量哦！！！！！！！！！！！
    // 定义一个变量表示几扇
    var num = {{.Number}};  // 这里传golang变量哦！！！！！！！！！！！


    // 执行主函数
    main();

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 保存成功就加个提示
    if (isSave) {
        createLayer();
    }
}`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width      float64
		Height     float64
		Number     int    // 几片折屏
		SingleName string // 单片名字
	}{width, height, int(number), singleName}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName))
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

// FrameSave7 生成多座屏专属保存
func FrameSave7(widthSlice, heightSlice []float64, heightMax float64, frameName string) {
	const script = `//var str = "js实现用{two}自符串替换占位符{two} {three}  {one} ".format({one: "I",two: "LOVE",three: "YOU"});
//var str2 = "js实现用{1}自符串替换占位符{1} {2}  {0} ".format("I","LOVE","YOU");
String.prototype.format = function () {
    if (arguments.length == 0) return this;
    var param = arguments[0];
    var s = this;
    if (typeof (param) == 'object') {
        for (var key in param)
            s = s.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
        return s;
    } else {
        for (var i = 0; i < arguments.length; i++)
            s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
        return s;
    }
};


// 定义一个函数用来设置黑边
function addEdge() {
    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width + 0.1;
    var height = app.activeDocument.height + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    function layer() {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }

    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 合并图层用于提速
function speedUp() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();
    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK);
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true;
}

//  返回上一级函数
function returnUpper() {
    // 为历史定义的变量
    var idslct = charIDToTypeID("slct");
    var idnull = charIDToTypeID("null");
    var idHstS = charIDToTypeID("HstS");
    var idOrdn = charIDToTypeID("Ordn");
    var idPrvs = charIDToTypeID("Prvs");

    // 返回上一级历史
    var ref = new ActionReference();
    var desc = new ActionDescriptor();
    ref.putEnumerated(idHstS, idOrdn, idPrvs);
    desc.putReference(idnull, ref);
    executeAction(idslct, desc, DialogModes.NO);
}

//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName);
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("优化另存为",["不要带扩展名:*","默认保存为 JPG 文件:*"]))
}


// 这是负责裁剪另存的函数
function optimized(x1, y1, x2, y2, savePath) {
    var bounds = [x1, y1, x2, y2];
    app.activeDocument.crop(bounds, 0);

    // 添加黑边
    addEdge();

    // 裁剪之后进行保存的位置和你想要的默认名称
    var TmpFile = new File(savePath);

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档。
    app.activeDocument.saveAs(TmpFile, exportOptionsSave, true, Extension.LOWERCASE);
}

// 主函数
function main() {
    // 得到用户想要保存的位置
    var userSavePath = getPathName("请输入订单编号");
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return
    }

    // 得到最终会保存出来的文件名
    var fileNameArr = new Array();
    for (var i = 0; i < widthArray.length; i++) {
        fileNameArr[i] = "{0}_{1}_{2}-{3}_".format(userSavePath, screenName, i + 1, widthArray.length) + widthArray[i] + "x" + heightArray[i];
    }

    // 遍历全部可能覆盖的文件名
    for (var i = 0; i < fileNameArr.length; i++) {
        // 避免覆盖已保存的文件
        if (new File(fileNameArr[i] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return
        }
    }

    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("性能加速", "speedUp()");

    // 清除元数据
    deleteDocumentAncestorsMetadata();


    // 开始自动保存文件
    for (i = 0; i < widthArray.length; i++) {
        // 求出x1 坐标的值
        var x1 = 0;
        for (j = 0; j < i; j++) {
            x1 += widthArray[j]
        }

        // 求出y1 坐标的值
        var y1 = heightMax - heightArray[i];

        // 求出x2 坐标的值
        var x2 = 0;
        for (k = 0; k < i + 1; k++) {
            x2 += widthArray[k]
        }

        // 求出y2 坐标的值
        var y2 = heightMax;

        // 生成历史记录并调用函数
        app.activeDocument.suspendHistory("优化另存", "optimized(x1, y1, x2, y2, fileNameArr[i])");

        // 返回上一级历史
        returnUpper();
    }

    // 全部保存成功
    isSave = true;
}


// 判断是否有打开的文件
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    //设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // 是否保存成功
    var isSave = false;

    // 保存活动历史记录状态
    var savedState = app.activeDocument.activeHistoryState;

    // 用数组定义多座屏的宽和高
    var widthArray = {{.WidthSliceJS}};   // 这里传golang排版好的字符串哦！！！！！！！！！！！
    var heightArray = {{.HeightSliceJS}};  // 这里传golang排版好的字符串哦！！！！！！！！！！！
    // 最高的高度哦
    var heightMax = {{.HeightMax}};        // 这里传golang变量哦！！！！！！！！！！！
    var screenName = "{{.ScreenName}}座屏"    // 这里传golang变量哦！！！！！！！！！！！

    // 执行主函数
    main();

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 保存成功就加个提示
    if (isSave) {
        createLayer();
    }
}`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		HeightMax     float64 // 最大的高
		ScreenName    string  // 是几座屏
	}{tools.Float64SliceToJsArray(widthSlice), tools.Float64SliceToJsArray(heightSlice), heightMax, tools.Transfer(len(widthSlice))}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create(fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName))
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

// FrameSave8to2 拉布座屏专属保存
func FrameSave8to2(frameName string, thickness float64) {
	const script = `// 定义一个函数用来设置黑边
function addBlackEdge() {
    // 保存当前背景颜色
    const nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";

    // 设置背景颜色
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    const width = app.activeDocument.width.value + 0.1;
    const height = app.activeDocument.height.value + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(width, height, AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 清理元数据
function deleteDocumentAncestorsMetadata() {
    const clear = function () {
        // 清理元数据四步骤
        if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
        var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
        // Begone foul Document Ancestors!
        xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
        app.activeDocument.xmpMetadata.rawData = xmp.serialize();
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("清理元数据", "clear()");
}


// 创建一个文字提示层
function promptLayer(text) {
    // 设置坐标底部居中
    const x = app.activeDocument.width.value / 2;
    const y = app.activeDocument.height.value - 1;
    // 添加一个颜色采样器
    const pointSample1 = app.activeDocument.colorSamplers.add([x - 1, y - 1]);
    const pointSample2 = app.activeDocument.colorSamplers.add([x + 1, y - 1]);
    // 求出平均值备用
    const average = (pointSample1.color.cmyk.black + pointSample2.color.cmyk.black) / 2
    // 删除全部颜色取样器
    app.activeDocument.colorSamplers.removeAll();
    // 不在循环内直接关闭信息面板
    app.runMenuItem(stringIDToTypeID("closeInfoPanel"));
    // 创建一个色彩变量 c
    var c = new SolidColor();
    // 如果吸取的颜色K小于40说明偏白，那么字就改成黑色
    if (average < 40) {
        c.rgb.hexValue = "000000";
    } else {
        c.rgb.hexValue = "ffffff";
    }

    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。
    var newLayer = app.activeDocument.artLayers.add();
    // 把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。
    newLayer.kind = LayerKind.TEXT;
    // 设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素，例子 newLayer.textItem.position= [UnitValue("50px"), UnitValue("100px")]
    newLayer.textItem.position = [x, y];
    // 设置 newLayer 的文本字体大小为“40 点”。
    newLayer.textItem.size = UnitValue("2cm");
    // 设置 newLayer 的文本内容。
    newLayer.textItem.contents = text;
    // 设置 newLayer 的文本框对齐方式为居中对齐。
    newLayer.textItem.justification = Justification.CENTER;
    // 设置 newLayer 的文本颜色为 c。
    newLayer.textItem.color = c;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    const layer = function () {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 让用户确定保存的文件位置，返回文件对象
function setSavePath() {
    var docName = app.activeDocument.name;

    // 如果有后缀
    var index = docName.lastIndexOf(".")
    if (index != -1) {
        // 就去掉后缀名
        docName = docName.substring(0, index);
    }

    // 自己特意加的后缀可以取代类型选择的大写
    var TmpFile = new File("~/Desktop/GoCutting/" + docName + ".jpg");
    // 获取用户自定义储存位置
    TmpFile = TmpFile.saveDlg("储存副本", "JPEG:*.JPG;*.JPEG;*.JPE, 不要修改保存类型:*.*");

    return TmpFile;
}


// 获取文件名，返回字符串
function getFileName(fileObject) {
    // 先进行URL解码
    var fileName = decodeURI(fileObject.name);
    // 去掉后缀名
    fileName = fileName.substring(0, fileName.lastIndexOf("."));

    return fileName;
}


// 用来保存的函数
function saveJPEG(fileObject) {
    //定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();

    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;

    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;

    //设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;

    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    app.activeDocument.saveAs(fileObject, exportOptionsSave, true, Extension.LOWERCASE);
}


// 全部整合在一起
function frameSave(fileObject) {
    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();
	// 改变当前文档的色彩模型为 CMYK
	app.activeDocument.changeMode(ChangeMode.CMYK);
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";

    // 复制图层
    app.activeDocument.activeLayer.duplicate();
    app.activeDocument.activeLayer.duplicate();
    app.activeDocument.activeLayer.duplicate();
    app.activeDocument.activeLayer.duplicate();

    // 扩大画布
    const width = app.activeDocument.width.value;
    const height = app.activeDocument.height.value;
    // 叠加画布预留重设大小
    app.activeDocument.resizeCanvas(width + Reserve, height + Reserve, AnchorPosition.MIDDLECENTER);

    // 垂直翻转
    app.activeDocument.artLayers[0].resize(undefined, -100);
    // 向上移动图层
    app.activeDocument.artLayers[0].translate(0, -height);

    // 垂直翻转
    app.activeDocument.artLayers[1].resize(undefined, -100);
    // 向下移动图层
    app.activeDocument.artLayers[1].translate(0, height);

    // 水平翻转
    app.activeDocument.artLayers[2].resize(-100, undefined);
    // 向左移动图层
    app.activeDocument.artLayers[2].translate(-width, 0);

    // 水平翻转
    app.activeDocument.artLayers[3].resize(-100, undefined);
    // 向左移动图层
    app.activeDocument.artLayers[3].translate(width, 0);

    // 获取用户设定的文件名
    //const userName = getFileName(fileObject);
    // 按工厂要求添加提示
    //promptLayer(userName);

    // 拼合活动文档的所有图层并扔掉隐藏的图层
    app.activeDocument.flatten();

    // 添加黑边
    addBlackEdge();

    // 最后另存为JPEG
    saveJPEG(fileObject);
}


// 主函数
function main() {
    // 判断是否有打开的文件
    if (!documents.length) {
        // alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 清理元数据
    deleteDocumentAncestorsMetadata();

    // 确定用户的保存位置
    var filePath = setSavePath();
    // 用户没有选择位置直接返回
    if (filePath == null) {
        return;
    }

    // 保存开始的历史记录状态
    var saveState = app.activeDocument.activeHistoryState;

    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("储存副本", "frameSave(filePath)");

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = saveState;

    // 保存成功就加个提示
    createLayer();
}

// 设置拉布屏风边框厚度，提前算出叠加厚度后需要预留的画布
const Reserve = {{.Thickness}} * 2 + 2  // 这里传golang变量哦！！！！！！！！！！！
main();`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Thickness float64 // 拉布屏风的边框厚度
	}{thickness}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 生成的文件名字
	fileName := fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName)

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

// FrameSave8to3 拉布折屏的专属保存
func FrameSave8to3(frameName string, width, height, count float64) {
	const script = `//var str = "js实现用{two}自符串替换占位符{two} {three}  {one} ".format({one: "I",two: "LOVE",three: "YOU"});
//var str2 = "js实现用{1}自符串替换占位符{1} {2}  {0} ".format("I","LOVE","YOU");
String.prototype.format = function () {
    if (arguments.length == 0) return this;
    var param = arguments[0];
    var s = this;
    if (typeof (param) == 'object') {
        for (var key in param)
            s = s.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
        return s;
    } else {
        for (var i = 0; i < arguments.length; i++)
            s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
        return s;
    }
};


// 用于添加黑边的函数
function addBlackEdge() {
    // 保存当前背景颜色
    const nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";

    // 设置背景颜色
    app.backgroundColor = black;

    // 获取当前文档的高度与宽度
    const width = app.activeDocument.width.value + 0.1;
    const height = app.activeDocument.height.value + 0.1;

    // 重设画布大小
    app.activeDocument.resizeCanvas(width, height, AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}


// 创建一个文字提示层
function promptLayer(text) {
    // 设置坐标底部居中
    const x = app.activeDocument.width.value / 2;
    const y = 3;
    // 添加一个颜色采样器
    const pointSample1 = app.activeDocument.colorSamplers.add([x - 1, y - 1]);
    const pointSample2 = app.activeDocument.colorSamplers.add([x, y - 1]);
    const pointSample3 = app.activeDocument.colorSamplers.add([x + 1, y - 1]);
    // 求出平均值备用
    const average = (pointSample1.color.cmyk.black + pointSample2.color.cmyk.black + pointSample3.color.cmyk.black) / 3
    // 删除全部颜色取样器
    app.activeDocument.colorSamplers.removeAll();
    // 创建一个色彩变量 c
    var c = new SolidColor();
    // 如果吸取的颜色K小于40说明偏白，那么字就改成黑色
    if (average < 40) {
        c.rgb.hexValue = "000000";
    } else {
        c.rgb.hexValue = "ffffff";
    }

    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。
    var newLayer = app.activeDocument.artLayers.add();
    // 把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。
    newLayer.kind = LayerKind.TEXT;
    // 设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素，例子 newLayer.textItem.position= [UnitValue("50px"), UnitValue("100px")]
    newLayer.textItem.position = [x, y];
    // 设置 newLayer 的文本字体大小为“40 点”。
    newLayer.textItem.size = UnitValue("2cm");
    // 设置 newLayer 的文本内容。
    newLayer.textItem.contents = text;
    // 设置 newLayer 的文本框对齐方式为居中对齐。
    newLayer.textItem.justification = Justification.CENTER;
    // 设置 newLayer 的文本颜色为 c。
    newLayer.textItem.color = c;
}


// 创建一个透明图层
function createLayer() {
    // 新建一个图层
    const layer = function () {
        app.activeDocument.artLayers.add().name = "注意：已快捷裁剪成功！";
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("注意：已快捷裁剪成功！", "layer()");
}


// 清理元数据
function deleteDocumentAncestorsMetadata() {
    const clear = function () {
        // 清理元数据四步骤
        if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
        var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
        // Begone foul Document Ancestors!
        xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
        app.activeDocument.xmpMetadata.rawData = xmp.serialize();
    }
    // 生成历史记录
    app.activeDocument.suspendHistory("清理元数据", "clear()");
}


//  获取用户想要保存的位置加文件名
function getPathName(saveName) {
    // 裁剪之后进行保存的位置和你想要的默认名称
    var tempFile = new File("~/Desktop/GoCutting/" + saveName);
    // 返回带路径的名字，注意要先数字解码
    return decodeURI(tempFile.saveDlg("储存副本", ["不要带扩展名:*", "默认保存为 JPG 文件:*"]))
}


// 全部整合在一起
function frameSave(fileNameArr) {
    // 定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
    var exportOptionsSave = new JPEGSaveOptions();
    // 嵌入彩色配置文件
    exportOptionsSave.embedColorProfile = true;
    // 设置杂边为无
    exportOptionsSave.matte = MatteType.NONE;
    //设置导出文档时，图片的压缩质量。数字范围为1至12。
    exportOptionsSave.quality = 12;
    // 保存为基线已优化
    exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;
    // 拼合图像
    app.activeDocument.flatten();
    // 改变当前文档的色彩模型为 CMYK
    app.activeDocument.changeMode(ChangeMode.CMYK);
	// 指定配置文件的色彩配置文件名
	app.activeDocument.colorProfileName = "Japan Color 2011 Coated";
    // 保存活动历史记录状态
    var work = app.activeDocument.activeHistoryState;

    // 循环保存每一片
    for (var i = 0; i < Count; i++) {
        // 根据左上右下裁剪且边距是0
        app.activeDocument.crop([Width * i, 0, Width * (i + 1), Height], 0);

        // 复制图层
        app.activeDocument.activeLayer.duplicate();
        app.activeDocument.activeLayer.duplicate();
        app.activeDocument.activeLayer.duplicate();
        app.activeDocument.activeLayer.duplicate();

        // 扩大画布，循环里不能使用常量
        var currentWidth = app.activeDocument.width.value;
        var currentHeight = app.activeDocument.height.value;
        app.activeDocument.resizeCanvas(currentWidth + 8, currentHeight + 8, AnchorPosition.MIDDLECENTER);

        // 垂直翻转
        app.activeDocument.artLayers[0].resize(undefined, -100);
        // 向上移动图层
        app.activeDocument.artLayers[0].translate(0, -currentHeight);

        // 垂直翻转
        app.activeDocument.artLayers[1].resize(undefined, -100);
        // 向下移动图层
        app.activeDocument.artLayers[1].translate(0, currentHeight);

        // 水平翻转
        app.activeDocument.artLayers[2].resize(-100, undefined);
        // 向左移动图层
        app.activeDocument.artLayers[2].translate(-currentWidth, 0);

        // 水平翻转
        app.activeDocument.artLayers[3].resize(-100, undefined);
        // 向左移动图层
        app.activeDocument.artLayers[3].translate(currentWidth, 0);

        // 按工厂要求添加提示会自动打开信息面板
        promptLayer((i + 1) + "/" + Count);

        // 拼合活动文档的所有图层并扔掉隐藏的图层
        app.activeDocument.flatten();

        // 添加黑边
        addBlackEdge();

        // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
        app.activeDocument.saveAs(new File(fileNameArr[i]), exportOptionsSave, true, Extension.LOWERCASE);

        // 当你完成了你正在做的任何事情，返回这个状态
        app.activeDocument.activeHistoryState = work;
    }
    // 循环结束后关闭信息面板
    app.runMenuItem(stringIDToTypeID("closeInfoPanel"));
}


// 主函数
function main() {
    // 判断是否有打开的文件
    if (!documents.length) {
        return;
    }

    // 得到用户想要保存的位置
    var userSavePath = getPathName("请输入订单编号");
    // 没有得到路径就返回
    if (userSavePath == "null") {
        return;
    }

    // 得到最终会保存出来的文件名
    var fileNameArr = new Array();
    for (var i = 0; i < Count; i++) {
        // 这里额外加8是因为保存时还要重设画布大小
        fileNameArr[i] = "{0}_拉布折屏_{1}_{2}-{3}".format(userSavePath, (Width + 8) + "x" + (Height + 8), i + 1, Count);
       
        // 避免覆盖已保存的文件
        if (new File(fileNameArr[i] + ".jpg").exists) {
            alert("输入的编号重复，已自动取消操作！");
            return;
        }
    }

    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;
    // 清除元数据
    deleteDocumentAncestorsMetadata();

    // 保存活动历史记录状态
    const savedState = app.activeDocument.activeHistoryState;

    // 调用保存函数，这里如果调用了历史函数内就不能再次调用历史
    frameSave(fileNameArr);

    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 全部保存成功加个提示
    createLayer();
}


// 定义折屏单扇的宽和高
const Width = {{.Width}};  // 这里传golang变量哦！！！！！！！！！！！
const Height = {{.Height}};  // 这里传golang变量哦！！！！！！！！！！！
// 定义一个变量表示几扇
const Count = {{.Count}};  // 这里传golang变量哦！！！！！！！！！！！
// 执行主函数
main();`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width  float64
		Height float64
		Count  int // 几片折屏
	}{width, height, int(count)}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 生成的文件名字
	fileName := fmt.Sprintf("data/jsx/temp/tailor_%s.jsx", frameName)

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
