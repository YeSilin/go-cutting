// 生成进度条函数
function progressBar() {
    // 进度条调用清除元数据函数
    app.doForcedProgress("正在替换智能对象... ", "replaceSmartObjects(doc)")  // 添加进度条
}


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

// 判断当前图层名字是DE开头还是DP开头
function judgeDeOrDp(layerName) {
    // 得到图层名字后，统一转成大写
    var upperCase = layerName.toUpperCase();

    // 再判断字符串是否以 DE 或DP 开头
    var isDE = upperCase.indexOf("DE")
    if (isDE == 0) { // 表示t图层名字是以DE开头； 
        return "DE"
    }

    var isDP = upperCase.indexOf("DP")
    if (isDP == 0) { // 表示t图层名字是以DP开头；
        return "DP"
    }
    return "NO"
}

// 替换图层时智能切换资源，选择优先消耗的资源
function resourcesSwitchDE(curLayer) {
    // 替换之前看下资源是否已经用完，相等表示用完了
    if (de == deArray.length) {
        // 替换智能对象，因为用完了所以使用dp补
        replaceContents(new File(dpArray[dp]), curLayer);
        curLayer.name = "DE 细节图层(不够用替补)";
        // 表示有dp图层
        isDPDE = true;
        dp++; // 索引加1
        countBar++; // 进度条加1
        // 更新进度条
        updateProgress(countBar, maxBar);
        return
    }

    // 替换智能对象
    replaceContents(new File(deArray[de]), curLayer);
    curLayer.name = "DE 细节图层";
    // 表示有dp图层
    isDPDE = true;
    de++; // 索引加1
    countBar++; // 进度条加1
    // 更新进度条
    updateProgress(countBar, maxBar);
}

// 替换图层时智能切换资源，选择优先消耗的资源
function resourcesSwitchDP(curLayer) {
    // 替换之前看下资源是否已经用完，相等表示用完了
    if (dp == dpArray.length) {
        // 替换智能对象
        replaceContents(new File(deArray[de]), curLayer);
        curLayer.name = "DP 普通图层(不够用替补)";
        // 表示有dp图层
        isDPDE = true;
        de++; // 索引加1
        countBar++; // 进度条加1
        // 更新进度条
        updateProgress(countBar, maxBar);
        return
    }

    replaceContents(new File(dpArray[dp]), curLayer);
    curLayer.name = "DP 普通图层";
    // 表示有dp图层
    isDPDE = true;
    dp++; // 索引加1
    countBar++; // 进度条加1
    // 更新进度条
    updateProgress(countBar, maxBar);
}


// 递归函数
function replaceSmartObjects(doc) {
    // 这里倒序，因为最上层的图层是0
    for (var i = doc.layers.length - 1; i >= 0; i--) {
        // 如果索引大于源文件数量就不再替换
        if (countBar == maxBar) {
            return
        }

        var curLayer = doc.layers[i];

        // 如果选中的图层类型不是普通图层，即代表是图层组
        if (curLayer.typename != "ArtLayer") {
            replaceSmartObjects(curLayer);
            continue;
        }
        // 如果当前图层是智能对象
        if (curLayer.kind == "LayerKind.SMARTOBJECT") {
            // 得到前面的头
            var frontName = judgeDeOrDp(curLayer.name)

            // 没有指定的头
            if (frontName == "NO") {
                continue
            }
            // 表示是以DE开头；
            if (frontName == "DE") {
                // 替换的时候智能切换
                resourcesSwitchDE(curLayer)
            }
            // 表示是以DP开头；
            if (frontName == "DP") {
                // 替换的时候智能切换
                resourcesSwitchDP(curLayer)
            }

        }
    }
}


// 没有打开的文档直接略过
if (!documents.length) {
    // alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    try {
        var doc = app.activeDocument;

        // dp源文件
        var dpArray = ["/E/淘宝美工/套图汇总/5.jpg"];  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

        // de源文件
        var deArray = ["/E/淘宝美工/套图汇总/de4.jpg"];  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

        // 这是源文件索引
        var dp = 0;
        var de = 0;

        // 进度条最大进度
        const maxBar = dpArray.length + deArray.length;
        // 进度条计数
        var countBar = 0

        // 默认当做没有dp de智能对象
        var isDPDE = false;

        // 运行带进度条的图层替换
        progressBar();

        if (!isDPDE) {
            alert("此文档没有以 DP 或 DE 开头命名的智能对象图层！");
        }
    } catch (error) {
        //发生错误执行的代码
    }
}