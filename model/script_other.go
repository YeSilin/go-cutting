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

// ClearMetadata 生成清除元数据第四版js，让文件更小巧，带进度条
func ClearMetadata() {
	const script = `// 方法用来判断当前字符串是否是以另外一个给定的子字符串“结尾”的
if (!String.prototype.endsWith) {
    String.prototype.endsWith = function (search, this_len) {
        if (this_len === undefined || this_len > this.length) {
            this_len = this.length;
        }
        return this.substring(this_len - search.length, this_len) === search;
    };
}

// 方法返回在数组中可以找到一个给定元素的第一个索引，如果不存在，则返回-1
// Production steps of ECMA-262, Edition 5, 15.4.4.14
// Reference: http://es5.github.io/#x15.4.4.14
if (!Array.prototype.indexOf) {
    Array.prototype.indexOf = function (searchElement, fromIndex) {

        var k;

        // 1. Let O be the result of calling ToObject passing
        //    the this value as the argument.
        if (this == null) {
            throw new TypeError('"this" is null or not defined');
        }

        var O = Object(this);

        // 2. Let lenValue be the result of calling the Get
        //    internal method of O with the argument "length".
        // 3. Let len be ToUint32(lenValue).
        var len = O.length >>> 0;

        // 4. If len is 0, return -1.
        if (len === 0) {
            return -1;
        }

        // 5. If argument fromIndex was passed let n be
        //    ToInteger(fromIndex); else let n be 0.
        var n = +fromIndex || 0;

        if (Math.abs(n) === Infinity) {
            n = 0;
        }

        // 6. If n >= len, return -1.
        if (n >= len) {
            return -1;
        }

        // 7. If n >= 0, then Let k be n.
        // 8. Else, n<0, Let k be len - abs(n).
        //    If k is less than 0, then let k be 0.
        k = Math.max(n >= 0 ? n : len - Math.abs(n), 0);

        // 9. Repeat, while k < len
        while (k < len) {
            // a. Let Pk be ToString(k).
            //   This is implicit for LHS operands of the in operator
            // b. Let kPresent be the result of calling the
            //    HasProperty internal method of O with argument Pk.
            //   This step can be combined with c
            // c. If kPresent is true, then
            //    i.  Let elementK be the result of calling the Get
            //        internal method of O with the argument ToString(k).
            //   ii.  Let same be the result of applying the
            //        Strict Equality Comparison Algorithm to
            //        searchElement and elementK.
            //  iii.  If same is true, return k.
            if (k in O && O[k] === searchElement) {
                return k;
            }
            k++;
        }
        return -1;
    };
}


// 清除元数据的主要函数
function deleteDocumentAncestorsMetadata() {
    // 清理垃圾四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 获得活动智能对象图层的文件名
function getSmartObjectName() {
    // 不打开智能对象的情况下获取其文件名，不是智能对象会报错
    var ref = new ActionReference();
    ref.putProperty(stringIDToTypeID("property"), stringIDToTypeID("smartObject"));
    ref.putEnumerated(charIDToTypeID("Lyr "), charIDToTypeID("Ordn"), charIDToTypeID("Trgt"));
    var layerDesc = executeActionGet(ref);
    var soDesc = layerDesc.getObjectValue(stringIDToTypeID('smartObject'));
    var theName = soDesc.getString(stringIDToTypeID("fileReference"));
    return theName;
}


// 打开全部智能对象的递归函数
function openAllSmartObject(doc) {
    // 打开的是文档不是图层组时才清理
    if (doc.typename == "Document") {
        // 清理元数据
        deleteDocumentAncestorsMetadata();
    }

    // 遍历全部图层
    for (var i = 0; i < doc.layers.length; i++) {
        var theLayer = doc.layers[i];
        // 如果这个图层类型不是普通图层，即代表是图层组
        if (theLayer.typename != "ArtLayer") {
            // 那么递归打开
            openAllSmartObject(theLayer);
            continue;
        }

        // 如果这个图层不是智能对象
        if (theLayer.kind != "LayerKind.SMARTOBJECT") {
            continue;
        }

        // 设置成活动图层
        app.activeDocument.activeLayer = theLayer;
        // 获取智能对象文件名
        var theName = getSmartObjectName();

        //  转成小写好查询扩展名
        var lowerName = theName.toLowerCase()
        //  如果是 pdf 就不处理
        if (lowerName.endsWith('.pdf')) {
            continue;
        }
        //  如果是 svg 就不处理
        if (lowerName.endsWith('.svg')) {
            continue;
        }
        //  如果是 ai 就不处理
        if (lowerName.endsWith('.ai')) {
            continue;
        }

        // 如果从没打开过这个智能对象
        if (smartObjectName.indexOf(theName) == -1) {
            // 不管能不能成功打开先记录下来
            smartObjectName.push(theName);

            // 打开当前活动图层的智能对象
            app.runMenuItem(stringIDToTypeID('placedLayerEditContents'));

            // 继续递归打开当前文档
            openAllSmartObject(app.activeDocument);
        }
    }

    // 如果是图层组不需要关闭文档只需退出函数
    if (doc.typename == "LayerSet") {
        return;
    }

    // 全部图层遍历完就关闭这个文档
    if (doc != mainDocument) {
        // 最后保存并关闭
        app.activeDocument.close(SaveOptions.SAVECHANGES);
    }
}


// 生成进度条函数
function progressBar() {
    // 进度条调用 打开所有智能对象并清理
    app.doForcedProgress("正在清除元数据... ", "openAllSmartObject(app.activeDocument)")
}

// 主函数
function main() {
    // 不是 Photoshop 就返回 或者 没有打开的文档就返回
    if (app.name.search("Photoshop") == -1 || !documents.length) {
        return;
    }
    // 定义主文档
    mainDocument = app.activeDocument;
    // 生成历史，历史调用进度条
    app.activeDocument.suspendHistory("清除元数据", "progressBar()");
}


// 声明主文档，因为要打开很多智能对象
var mainDocument;
// 智能对象名称数组，清理过的存里面
var smartObjectName = [];
main();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/clearMetadata.jsx", script)
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
