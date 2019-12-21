// 清理元数据
function deleteDocumentAncestorsMetadata() {
    // 清理元数据四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 另存为web
function saveAsWeb() {
    // 更新进度条
    updateProgress(1, 4);
    // 清理元数据
    deleteDocumentAncestorsMetadata();

    // 更新进度条
    updateProgress(2, 4);
    // 获取当前脚本所在路径
    var scriptPath = (new File($.fileName)).parent;

    // 定义文件保存位置
    var savePath = new File(scriptPath + "/../Picture/主图/dp.jpg");

    var jpgOpt = new ExportOptionsSaveForWeb();
    jpgOpt.format = SaveDocumentType.JPEG;  // 保存为jpg
    jpgOpt.includeProfile = false;  //装入颜色配置文件
    jpgOpt.interlaced = false;  // 交错
    jpgOpt.optimized = true;  //最优化
    jpgOpt.blur = 0;    // 默认 0.0 不模糊。
    jpgOpt.matteColor = new RGBColor(); // 把杂边颜色染成白色
    jpgOpt.matteColor.red = 255;
    jpgOpt.matteColor.green = 255;
    jpgOpt.matteColor.blue = 255;
    jpgOpt.quality = 100;  // 品质   100是最高画质

    // 更新进度条
    updateProgress(3, 4);
    activeDocument.exportDocument(savePath, ExportType.SAVEFORWEB,);

    // 更新进度条
    updateProgress(4, 4);
}


function main() {
    // 判断是否有打开的文件
    if (!documents.length) {
        alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 进度条调用另存函数
    app.doForcedProgress("正在导出文件... ", "saveAsWeb()")  // 添加进度条


}

// 主函数
main();