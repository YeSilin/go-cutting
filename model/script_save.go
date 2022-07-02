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

//SaveAsJPEG 储存为jpeg格式的初始化 暗号-10的实现
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
	tools.CreateFile("config/jsx/saveAsJPEG.jsx", script)
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
	tools.CreateFile("config/jsx/saveAllJPEG.jsx", script)
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
	tools.CreateFile("config/jsx/saveAndCloseAllDocuments.jsx", script)
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
    const nowName = app.activeDocument.name;

    // 要运行的脚本路径
    const runScript = scriptPath + "/temp/tailor_" + nowName + ".jsx";

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
	tools.CreateFile("config/jsx/loadSaveScript.jsx", script)
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
    // 图层数量必须大于1才能合并
    if (app.activeDocument.layers.length > 1) {
        // 合并全部可见图层
        app.activeDocument.mergeVisibleLayers();
    }

    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true

    // 添加黑边
    if (blackEdge) {
        addBlackEdge();
    }

    // 保存图片，同时成功就加个提示
    if (saveAsJPEG()) {
        createLayer()
    }
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
}


// 是否自动黑边
const blackEdge = {{.BlackEdge}};   // 这里传golang变量哦！！！！！！！！！！！！！！！！！！！！！！！！！！！！
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
	fileName := "config/jsx/frameSaveDef.jsx"
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

// FrameSave8to2 拉布座屏专属保存
func FrameSave8to2(frameName string) {
	jsx := `
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

// 用来保存的函数
function saveJPEG() {
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

    // 获取当前文档的文件名
    var name = app.activeDocument.name
    var TmpFile1 = new File("~/Desktop/GoCutting/" + name);

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档
    app.activeDocument.saveAs(TmpFile1.saveDlg("优化另存为", "JPEG Files: *.jpg"), exportOptionsSave, true, Extension.LOWERCASE);
}


// 全部整合在一起
function optimized() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;

    // 新建一个空白图层用于合并
    app.activeDocument.artLayers.add();
    // 合并全部可见图层
    app.activeDocument.mergeVisibleLayers();
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true

    // 复制图层
    app.activeDocument.activeLayer.duplicate()
    app.activeDocument.activeLayer.duplicate()
    app.activeDocument.activeLayer.duplicate()
    app.activeDocument.activeLayer.duplicate()
    
    // 扩大画布
    const width = app.activeDocument.width
    const height = app.activeDocument.height
    app.activeDocument.resizeCanvas(width + 8, height + 8, AnchorPosition.MIDDLECENTER);
    
    // 垂直翻转
    app.activeDocument.artLayers[0].resize(undefined, -100); //will flip layer vertically
    // 向上移动图层
    app.activeDocument.artLayers[0].translate(0, -height);
    
    // 垂直翻转
    app.activeDocument.artLayers[1].resize(undefined, -100); //will flip layer vertically
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
    
    // 合并全部可见图层
    app.activeDocument.mergeVisibleLayers();


    if (BlackEdge) {
        // 添加黑边
        addEdge();
    }

    // 清理元数据
    deleteDocumentAncestorsMetadata()

    // 保存JPEG
    saveJPEG()
}

// 判断是否有打开的文件
if (!documents.length) {
    // return;
} else {
    // 是否自动黑边
    var BlackEdge = {{.BlackEdge}};   // 这里传golang变量哦！！！！！！！！！！！！！！！！！！！！！！！！！！！！

    // 是否保存成功
    var isSave = false;

    // 保存开始的历史记录状态
    var savedState = app.activeDocument.activeHistoryState;

    // 如果出错就返回最开始
    try {
        // 生成历史记录
        app.activeDocument.suspendHistory("优化另存", "optimized()");
        isSave = true;
    } catch (error) {
        // 忽略错误
    }
    // 当你完成了你正在做的任何事情，返回这个状态
    app.activeDocument.activeHistoryState = savedState;

    // 保存成功就加个提示
    if (isSave) {
        createLayer()
    }
}
`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		BlackEdge bool // 是否自动黑边

	}{viper.GetBool("blackEdge")}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(jsx)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 生成的文件名字
	fileName := fmt.Sprintf("config/jsx/temp/tailor_%s.jsx", frameName)

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
