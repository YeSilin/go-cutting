package model

import "github.com/yesilin/go-cutting/tools"

// AddBlackEdge 暗号-7，为当前文档添加黑边
func AddBlackEdge() {
	const script = `// 对当前文档添加黑边，缺点是会合并全部图层
function addBlackEdge() {
    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;

    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "d5d5d5";
    
    // 设置背景颜色为新颜色
    app.backgroundColor = black;

    // 新建一个空白图层用于合并
    app.activeDocument.artLayers.add();
    
    // 合并全部可见图层
    app.activeDocument.mergeVisibleLayers();
    
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width.value + 0.1
    var height = app.activeDocument.height.value + 0.1
    
    // 重设画布大小
    app.activeDocument.resizeCanvas(width, height, AnchorPosition.MIDDLECENTER);

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}

// 判断是否有打开的文件
if (!documents.length) {
    // alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    // 生成历史记录
    app.activeDocument.suspendHistory("向四周添加0.1厘米黑边！", "addBlackEdge()");
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/addBlackEdge.jsx", script)
}

// ClearMetadataStd 生成清除元数据标准版，不清理智能对象，让文件跟小巧 无弹窗
func ClearMetadataStd() {
	const script = `// 清除元数据无弹窗版，并且不会清理智能对象
function deleteDocumentAncestorsMetadata() {
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
     // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


//  专门检查photoshop，否则会导致错误
if(String(app.name).search("Photoshop") > 0) {
    if (!documents.length) {
    // alert("没有打开的文档，请打开一个文档来运行此脚本！")
    } else {
        // 生成历史记录
        app.activeDocument.suspendHistory("清除元数据", "deleteDocumentAncestorsMetadata()");
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/clearMetadataStd.jsx", script)
}

// CopyAndCloseOtherDocuments 生成复制并关闭其他文档脚本
func CopyAndCloseOtherDocuments() {
	const script = `// For code readability. 图层操作要用到的函数
function cTID(s) {
    return charIDToTypeID(s)
}

function sTID(s) {
    return stringIDToTypeID(s)
}

// =============================  

// 将选中的图层编组  
function groupSelected(name) {
    var m_Dsc01 = new ActionDescriptor();
    var m_Ref01 = new ActionReference();
    m_Ref01.putClass(sTID("layerSection"));
    m_Dsc01.putReference(cTID("null"), m_Ref01);
    var m_Ref02 = new ActionReference();
    m_Ref02.putEnumerated(cTID("Lyr "), cTID("Ordn"), cTID("Trgt"));
    m_Dsc01.putReference(cTID("From"), m_Ref02);
    var m_Dsc02 = new ActionDescriptor();
    m_Dsc02.putString(cTID("Nm  "), name);
    m_Dsc01.putObject(cTID("Usng"), sTID("layerSection"), m_Dsc02);
    executeAction(cTID("Mk  "), m_Dsc01, DialogModes.NO);

    return activeDocument.activeLayer;
}


// 解锁背景图层
function unlockBackgroundLayer() {
    var idsetd = charIDToTypeID("setd");
    var desc8 = new ActionDescriptor();
    var idnull = charIDToTypeID("null");
    var ref2 = new ActionReference();
    var idLyr = charIDToTypeID("Lyr ");
    var idBckg = charIDToTypeID("Bckg");
    ref2.putProperty(idLyr, idBckg);
    desc8.putReference(idnull, ref2);
    var idT = charIDToTypeID("T   ");
    var desc9 = new ActionDescriptor();
    var idLyr = charIDToTypeID("Lyr ");
    desc8.putObject(idT, idLyr, desc9);
    executeAction(idsetd, desc8, DialogModes.NO);
}

// 选择全部图层，但不包括背景图层
function selectAllLayers() {
    var desc29 = new ActionDescriptor();
    var ref23 = new ActionReference();
    ref23.putEnumerated(charIDToTypeID('Lyr '), charIDToTypeID('Ordn'), charIDToTypeID('Trgt'));
    desc29.putReference(charIDToTypeID('null'), ref23);
    executeAction(stringIDToTypeID('selectAllLayers'), desc29, DialogModes.NO);
}


// 复制所有图层到指定文档
function copyAllLayers(srcDoc, dstDoc) {
    // 先激活文档
    app.activeDocument = srcDoc

    // 如果图层只有一个就不用解锁
    if (srcDoc.layers.length == 1) {
        // 直接复制到主文档
        srcDoc.activeLayer.duplicate(dstDoc);
        return
    }


    // 解锁背景图层
    unlockBackgroundLayer()

    // 选择全部图层
    selectAllLayers()

    // 只有图层数大于1的才打包
    if (srcDoc.layers.length > 1) {
        // 将选中的图层编组 
        groupSelected(srcDoc.name)
    }

    // 复制到主文档
    srcDoc.activeLayer.duplicate(dstDoc);
}


//  复制并关闭其他文档
function copyAndCloseOtherDocuments() {
    // 把所有要复制的文档保存到组
    var documentsArr = new Array

    // 得到要复制的文档，主要是documents关闭后会直接刷新，所以存自定义的数组里
    for (var i = 0; i < documents.length; i++) {
        // 如果是自己就不复制
        if (documents[i] == masterDocument) {
            continue
        }
        // 追加到数组
        documentsArr.push(documents[i])
    }

    // 循环关闭所有
    for (var i = 0; i < documentsArr.length; i++) {

        // 复制全部图层到指定文档
        copyAllLayers(documentsArr[i], masterDocument)

        // 关闭文档而不保存更改
        documentsArr[i].close(SaveOptions.DONOTSAVECHANGES);
    }
}


// 添加进度条
function progressBar() {
    app.doForcedProgress("正在复制并关闭其他文档... ", "copyAndCloseOtherDocuments()")
}

// 添加历史记录
function history() {
    // 生成历史记录并调用函数
    app.activeDocument.suspendHistory("复制并关闭其他文档", "progressBar()");
}


// 主函数
function main() {
    if (!documents.length) {
        //alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }
    // 主文档等于当前激活的文档
    masterDocument = app.activeDocument

    // 运行历史记录
    history()
}


// 先申明主文档
var masterDocument

// 运行
main()`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/copyAndCloseOtherDocuments.jsx", script)
}
