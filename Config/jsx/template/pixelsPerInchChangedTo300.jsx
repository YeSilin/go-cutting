﻿// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 设置首选项新文档预设单位是像素
app.preferences.rulerUnits = Units.PIXELS;

// 源文件
{{.}} // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

// 这是源文件索引
var src = 0;

// 获取当前脚本所在路径
const scriptPath = (new File($.fileName)).parent;
// 为源文件加上完整路径
for (i = 0; i < srcArray.length; i++) {
    srcArray[i] = scriptPath + "/../Picture/" + srcArray[i];

    // 为指定路径文件定义变量
    var fileRef = new File(srcArray[i]);
    if (fileRef.exists) {    // 如果图像存在
        app.open(fileRef);   // 打开文档

        // 清理元数据
        deleteDocumentAncestorsMetadata();

        // 修改图像大小和分辨率并且重新采样为自动，resizeImage( 宽度, 高度, 分辨率, 重新采样, 减少杂色 ) 
        app.activeDocument.resizeImage(1500, 1500, 300, ResampleMethod.AUTOMATIC);

        // 避免弹窗使用另存文件，定义一个变量[exportOptionsSave]，用来表示导出文档为jpeg格式的设置属性。
        var exportOptionsSave = new JPEGSaveOptions();

        // 嵌入彩色配置文件
        exportOptionsSave.embedColorProfile = true;

        // 设置杂边为无
        exportOptionsSave.matte = MatteType.NONE;

        //设置导出文档时，图片的压缩质量。数字范围为1至12。
        exportOptionsSave.quality = 12;

        // 保存为基线已优化
        exportOptionsSave.formatOptions = FormatOptions.OPTIMIZEDBASELINE;

        // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )，调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档
        app.activeDocument.saveAs(fileRef, exportOptionsSave, true, Extension.LOWERCASE);

        // 关闭文档 不存储改变
        app.activeDocument.close(SaveOptions.DONOTSAVECHANGES);
    }
}