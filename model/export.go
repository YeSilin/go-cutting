package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// SaveForWebInit 导出web格式脚本的初始化 暗号-98的实现 副作用分辨率会被强制修改为72ppi 并且无法选择保存路径
func SaveForWebInit(originalPath string) {
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

// SaveForWeb 导出web格式的调用
func SaveForWeb() {
	go func() {
		// 自动套图工作路径
		picturePath := viper.GetString("picture")

		// 创建套图文件夹
		_ = tools.CreateMkdirAll(fmt.Sprintf("%s/主图", picturePath))

		// 创建一个协程使用cmd来运行脚本
		dataPath := "config/jsx/saveForWeb.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

		//time.Sleep(time.Second) // 停一秒
		//
		//// 如果存在images就打开
		//if ok, _ := tools.IsPathExists(fmt.Sprintf("%s/主图/images", picturePath)); ok {
		//	exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图\\images", picturePath)).Run()
		//} else {
		//	exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图", picturePath)).Run()
		//}
	}()
}

//SaveAsJPEGInit 储存为jpeg格式的初始化 暗号-10的实现
func SaveAsJPEGInit() {
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

// SaveAsJPEG 储存为jpeg格式的调用
func SaveAsJPEG() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/SaveAsJPEGInit.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// SaveAllJPEGInit 将所有打开的文档储存为jpeg格式的初始化 暗号-11的实现
func SaveAllJPEGInit() {
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

// SaveAllJPEG 将所有打开的文档储存为jpeg格式的调用
func SaveAllJPEG() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/saveAllJPEG.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}
