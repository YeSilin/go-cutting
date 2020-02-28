// 替换智能对象
function replaceContents(newFile, theSO) {
    app.activeDocument.activeLayer = theSO;
    // =======================================================
    var idplacedLayerReplaceContents = stringIDToTypeID("placedLayerReplaceContents");
    var desc3 = new ActionDescriptor();
    var idnull = charIDToTypeID("null");
    desc3.putPath(idnull, new File(newFile));
    var idPgNm = charIDToTypeID("PgNm");
    desc3.putInteger(idPgNm, 1);
    executeAction(idplacedLayerReplaceContents, desc3, DialogModes.NO);
    return app.activeDocument.activeLayer
}


// 递归函数
function replaceSmartObjects(doc) {
    for (var i = 0; i < doc.layers.length; i++) {
        var curLayer = doc.layers[i];

        // 如果选中的图层类型不是普通图层，即代表是图层组
        if (curLayer.typename != "ArtLayer") {
            replaceSmartObjects(curLayer);
            continue;
        }
        // 如果当前图层是智能对象
        if (curLayer.kind == "LayerKind.SMARTOBJECT") {
            // 判断字符串是否以 dp 开头 upper and lower case
            var upperCase = curLayer.name.indexOf("dp"); // 小写
            var upperCase2 = curLayer.name.indexOf("dP"); // 小写
             var lowerCase = curLayer.name.indexOf("DP");  // 大写
             var lowerCase2 = curLayer.name.indexOf("Dp");  // 大写

            if (upperCase == 0 || lowerCase == 0 || upperCase2 == 0 || lowerCase2 == 0) { // 表示curLayer.name是以dp开头；
                // 替换智能对象
                replaceContents(new File(srcArray[src]), curLayer);
                curLayer.name = "DP 此图层由GoCutting替换"
                src++; // 索引加1
                // 如果索引大于源文件数量就不再替换
                if (src == srcArray.length) {
                    return
                }
            }
        }
    }
}


//  函数scrubs从文件中提取文档祖先
if (!documents.length) {
    alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    try {
        var doc = app.activeDocument;

        // 源文件
        {{.}}  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

        // 这是源文件索引
        var src = 0;

        // 获取当前脚本所在路径
        var scriptPath = (new File($.fileName)).parent;
        // 为源文件加上完整路径
        for (i = 0; i < srcArray.length; i++) {
            srcArray[i] = scriptPath + "/../Picture/" + srcArray[i]
        }
        // 运行图层替换
        replaceSmartObjects(doc);
    } catch (error) {
        //发生错误执行的代码
    }
}