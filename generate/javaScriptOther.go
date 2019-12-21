package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strings"
)

//生成清除元数据第二版js，让文件跟小巧，带进度条
func ClearMetadata() {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("// 清除元数据的主要函数\n")
	jsx.WriteString("function deleteDocumentAncestorsMetadata() {\n")
	jsx.WriteString("    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject(\"lib:AdobeXMPScript\");\n")
	jsx.WriteString("    // 如果当前是主文档就更新进度条\n")
	jsx.WriteString("    if (app.activeDocument == mainDocument) {\n")
	jsx.WriteString("        // 更新进度条\n")
	jsx.WriteString("        updateProgress(1, maxBar);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var xmp = new XMPMeta( activeDocument.xmpMetadata.rawData);\n")
	jsx.WriteString("    // 如果当前是主文档就更新进度条\n")
	jsx.WriteString("    if (app.activeDocument == mainDocument) {\n")
	jsx.WriteString("        // 更新进度条\n")
	jsx.WriteString("        updateProgress(2, maxBar);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // Begone foul Document Ancestors!\n")
	jsx.WriteString("    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, \"DocumentAncestors\");\n")
	jsx.WriteString("    // 如果当前是主文档就更新进度条\n")
	jsx.WriteString("    if (app.activeDocument == mainDocument) {\n")
	jsx.WriteString("        // 更新进度条\n")
	jsx.WriteString("        updateProgress(3, maxBar);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    app.activeDocument.xmpMetadata.rawData = xmp.serialize();\n")
	jsx.WriteString("    // 如果当前是主文档就更新进度条\n")
	jsx.WriteString("    if (app.activeDocument == mainDocument) {\n")
	jsx.WriteString("        // 更新进度条\n")
	jsx.WriteString("        updateProgress(4, maxBar);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 递归调用函数，清除当前文档中的所有智能对象元数据\n")
	jsx.WriteString("    clearDocumentAncestorsForAllLayers(app.activeDocument);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    if (app.activeDocument !== mainDocument) {\n")
	jsx.WriteString("        // 不是主文档，那么就是智能对象，关闭并保存它\n")
	jsx.WriteString("        app.activeDocument.close(SaveOptions.SAVECHANGES);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    }else{\n")
	jsx.WriteString("        // 主文档就不自动保存了\n")
	jsx.WriteString("        //app.activeDocument.save();\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 清除文档中所有智能对象的元数据\n")
	jsx.WriteString("function clearDocumentAncestorsForAllLayers(doc) {\n")
	jsx.WriteString("    try {\n")
	jsx.WriteString("        // 如果当前文档未定义就返回\n")
	jsx.WriteString("        if (doc == undefined) {\n")
	jsx.WriteString("            return;\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("\n")
	jsx.WriteString("        for (var i = 0; i < doc.layers.length; i++) {\n")
	jsx.WriteString("            var curLayer = doc.layers[i];\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("            // 如果选中的图层类型不是普通图层，即代表是图层组\n")
	jsx.WriteString("            if (curLayer.typename != \"ArtLayer\") {\n")
	jsx.WriteString("                clearDocumentAncestorsForAllLayers(curLayer);\n")
	jsx.WriteString("                continue;\n")
	jsx.WriteString("            }\n")
	jsx.WriteString("\n")
	jsx.WriteString("            // 如果当前图层是智能对象\n")
	jsx.WriteString("            if (curLayer.kind == \"LayerKind.SMARTOBJECT\") {\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("                // 如果当前是主文档就更新进度条\n")
	jsx.WriteString("                if (app.activeDocument == mainDocument) {\n")
	jsx.WriteString("                    // 更新进度条\n")
	jsx.WriteString("                    updateProgress(4+i, maxBar);\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("                }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("                app.activeDocument.activeLayer = curLayer;\n")
	jsx.WriteString("\n")
	jsx.WriteString("                // 打开这个智能对象\n")
	jsx.WriteString("                var idplacedLayerEditContents = stringIDToTypeID(\"placedLayerEditContents\");\n")
	jsx.WriteString("                var actionDescriptor = new ActionDescriptor();\n")
	jsx.WriteString("                executeAction(idplacedLayerEditContents, actionDescriptor, DialogModes.NO);\n")
	jsx.WriteString("\n")
	jsx.WriteString("                updateProgress( 0, maxBar);\n")
	jsx.WriteString("                // 这行代码看不懂，大概意思是，所打开的智能对象当前选中的图层，如果是被打开的智能对象本身，那么就不清理\n")
	jsx.WriteString("                if(app.activeDocument.activeLayer == curLayer){\n")
	jsx.WriteString("                    continue;\n")
	jsx.WriteString("                }\n")
	jsx.WriteString("\n")
	jsx.WriteString("                // 递归调用清除元数据\n")
	jsx.WriteString("                deleteDocumentAncestorsMetadata();\n")
	jsx.WriteString("\n")
	jsx.WriteString("            }\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("    } catch (e) {\n")
	jsx.WriteString("        // 清除失败就不弹窗了\n")
	jsx.WriteString("        //alert(\"Layer clean fail.name=\"+doc+\";e=\"+e)\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 生成进度条函数\n")
	jsx.WriteString("function progressBar() {\n")
	jsx.WriteString("    // 进度条调用清除元数据函数\n")
	jsx.WriteString("    app.doForcedProgress(\"正在清除源数据... \",\"deleteDocumentAncestorsMetadata()\")  // 添加进度条\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("function main() {\n")
	jsx.WriteString("    whatApp = String(app.name);  //String version of the app name\n")
	jsx.WriteString("    if(whatApp.search(\"Photoshop\") > 0) {  //Check for photoshop specifically, or this will cause errors\n")
	jsx.WriteString("\n")
	jsx.WriteString("        //  函数scrubs从文件中提取文档祖先\n")
	jsx.WriteString("        if (!documents.length) {\n")
	jsx.WriteString("            alert(\"没有打开的文档，请打开一个文档来运行此脚本！\");\n")
	jsx.WriteString("            return;\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 定义主文档，因为要打开很多智能对象\n")
	jsx.WriteString("        mainDocument = app.activeDocument;\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 定义一个变量表示总进度与当前进度\n")
	jsx.WriteString("        maxBar = app.activeDocument.layers.length + 4;\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 生成历史，历史调用进度条\n")
	jsx.WriteString("        app.activeDocument.suspendHistory(\"清除源数据\", \"progressBar()\");  // 生成历史记录\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 定义主文档，因为要打开很多智能对象\n")
	jsx.WriteString("var mainDocument;\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 定义一个变量表示总进度与当前进度\n")
	jsx.WriteString("var maxBar;\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 运行此脚本\n")
	jsx.WriteString("main();\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/ClearMetadataJS.jsx", jsxStr)
}

//生成清除元数据js，让文件跟小巧 无弹窗版本！
func ClearMetadataNoPopUp() {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("function deleteDocumentAncestorsMetadata() {\n")
	jsx.WriteString("    whatApp = String(app.name);  //String version of the app name\n")
	jsx.WriteString("    if(whatApp.search(\"Photoshop\") > 0) {  //Check for photoshop specifically, or this will cause errors\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Function Scrubs Document Ancestors from Files\n")
	jsx.WriteString("        if(!documents.length) {\n")
	jsx.WriteString("            // alert(\"没有打开的文档，请打开一个文档来运行此脚本！\")\n")
	jsx.WriteString("            return;\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("\n")
	jsx.WriteString("        if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject(\"lib:AdobeXMPScript\");\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var xmp = new XMPMeta( activeDocument.xmpMetadata.rawData);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Begone foul Document Ancestors!\n")
	jsx.WriteString("        xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, \"DocumentAncestors\");\n")
	jsx.WriteString("\n")
	jsx.WriteString("        app.activeDocument.xmpMetadata.rawData = xmp.serialize();\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// Now run the function to remove the document ancestors\n")
	jsx.WriteString("deleteDocumentAncestorsMetadata();\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/ClearMetadataNoPopUpJS.jsx", jsxStr)
}

// 为当前文档添加黑边
func BlackEdge() {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("function BlackEdge() {\n")
	jsx.WriteString("    // 判断是否有打开的文件\n")
	jsx.WriteString("    if(!documents.length) {\n")
	jsx.WriteString("        alert(\"没有打开的文档，请打开一个文档来运行此脚本！\");\n")
	jsx.WriteString("       return;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 设置首选项新文档预设单位是厘米，PIXELS是像素\n")
	jsx.WriteString("    app.preferences.rulerUnits = Units.CM;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 保存当前背景颜色\n")
	jsx.WriteString("    var nowColor = app.backgroundColor;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 定义一个对象颜色是黑色\n")
	jsx.WriteString("    var black = new SolidColor();\n")
	jsx.WriteString("    black.rgb.hexValue = \"000000\";\n")
	jsx.WriteString("    app.backgroundColor = black;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 生成历史记录\n")
	jsx.WriteString("    app.activeDocument.suspendHistory(\"向四周添加0.5厘米黑边！\", \"addEdge()\");\n")
	jsx.WriteString("\n")
	jsx.WriteString("     // 恢复之前的背景颜色\n")
	jsx.WriteString("    app.backgroundColor = nowColor;\n")
	jsx.WriteString(" }\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 此函数用于生成历史记录\n")
	jsx.WriteString("function addEdge(){\n")
	jsx.WriteString("        // 新建一个空白图层用于合并\n")
	jsx.WriteString("        app.activeDocument.artLayers.add();\n")
	jsx.WriteString("        // 合并全部可见图层\n")
	jsx.WriteString("        app.activeDocument.mergeVisibleLayers();\n")
	jsx.WriteString("        // 转为背景图层不然添加黑边会无效\n")
	jsx.WriteString("        app.activeDocument.activeLayer.isBackgroundLayer = true\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 获取当前文档的高度与宽度\n")
	jsx.WriteString("        var width = app.activeDocument.width + 0.5;\n")
	jsx.WriteString("        var height = app.activeDocument.height + 0.5;\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 重设画布大小\n")
	jsx.WriteString("        app.activeDocument.resizeCanvas(UnitValue (width), UnitValue (height), AnchorPosition.MIDDLECENTER);\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 生成黑边\n")
	jsx.WriteString("BlackEdge()\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/BlackEdgeJS.jsx", jsxStr)
}

// 生成临时效果图选择框架代码
func SelectionTempFrameJS(frame string, layer int) {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("// 为了代码可读性\n")
	jsx.WriteString("function cTID(s){return charIDToTypeID(s)}\n")
	jsx.WriteString("function sTID(s){return stringIDToTypeID(s)}\n")
	jsx.WriteString("// =============================\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 选定组\n")
	jsx.WriteString("function groupSelected(name) {\n")
	jsx.WriteString("   var m_Dsc01 = new ActionDescriptor();\n")
	jsx.WriteString("   var m_Ref01 = new ActionReference();\n")
	jsx.WriteString("   m_Ref01.putClass( sTID( \"layerSection\" ) );\n")
	jsx.WriteString("   m_Dsc01.putReference(  cTID( \"null\" ), m_Ref01 );\n")
	jsx.WriteString("   var m_Ref02 = new ActionReference();\n")
	jsx.WriteString("   m_Ref02.putEnumerated( cTID( \"Lyr \" ), cTID( \"Ordn\" ), cTID( \"Trgt\" ) );\n")
	jsx.WriteString("   m_Dsc01.putReference( cTID( \"From\" ), m_Ref02 );\n")
	jsx.WriteString("   var m_Dsc02 = new ActionDescriptor();\n")
	jsx.WriteString("   m_Dsc02.putString( cTID( \"Nm  \" ), name);\n")
	jsx.WriteString("   m_Dsc01.putObject( cTID( \"Usng\" ), sTID( \"layerSection\" ), m_Dsc02 );\n")
	jsx.WriteString("   executeAction( cTID( \"Mk  \" ), m_Dsc01, DialogModes.NO );\n")
	jsx.WriteString("\n")
	jsx.WriteString("   return activeDocument.activeLayer;\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 取消分组\n")
	jsx.WriteString("function ungroup() {\n")
	jsx.WriteString("   var m_Dsc01 = new ActionDescriptor();\n")
	jsx.WriteString("   var m_Ref01 = new ActionReference();\n")
	jsx.WriteString("   m_Ref01.putEnumerated( cTID( \"Lyr \" ), cTID( \"Ordn\" ), cTID( \"Trgt\" ) );\n")
	jsx.WriteString("   m_Dsc01.putReference( cTID( \"null\" ), m_Ref01 );\n")
	jsx.WriteString("\n")
	jsx.WriteString("   try {\n")
	jsx.WriteString("      executeAction( sTID( \"ungroupLayersEvent\" ), m_Dsc01, DialogModes.NO );\n")
	jsx.WriteString("   } catch(e) {}\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 有图层蒙版\n")
	jsx.WriteString("function hasLayerMask() {\n")
	jsx.WriteString("   var m_Ref01 = new ActionReference();\n")
	jsx.WriteString("   m_Ref01.putEnumerated( sTID( \"layer\" ), cTID( \"Ordn\" ), cTID( \"Trgt\" ));\n")
	jsx.WriteString("   var m_Dsc01= executeActionGet( m_Ref01 );\n")
	jsx.WriteString("   return m_Dsc01.hasKey(cTID('Usrs'));\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 折叠图层组的函数\n")
	jsx.WriteString("function closeGroup(layerSet) {\n")
	jsx.WriteString("   var m_Name = layerSet.name;\n")
	jsx.WriteString("   var m_Opacity = layerSet.opacity;\n")
	jsx.WriteString("   var m_BlendMode = layerSet.blendMode;\n")
	jsx.WriteString("   var m_LinkedLayers = layerSet.linkedLayers;\n")
	jsx.WriteString("\n")
	jsx.WriteString("   var m_bHasMask = hasLayerMask();\n")
	jsx.WriteString("   if(m_bHasMask) loadSelectionOfMask();\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("   if(layerSet.layers.length <= 1) {\n")
	jsx.WriteString("      addLayer();\n")
	jsx.WriteString("      var m_Tmp = activeDocument.activeLayer;\n")
	jsx.WriteString("      m_Tmp.name = \"dummy - feel free to remove me\";\n")
	jsx.WriteString("      activeDocument.activeLayer = layerSet;\n")
	jsx.WriteString("      ungroup();\n")
	jsx.WriteString("      addToSelection(\"dummy - feel free to remove me\");\n")
	jsx.WriteString("      groupSelected(m_Name);\n")
	jsx.WriteString("\n")
	jsx.WriteString("   } else {\n")
	jsx.WriteString("      activeDocument.activeLayer = layerSet;\n")
	jsx.WriteString("      ungroup();\n")
	jsx.WriteString("      groupSelected(m_Name);\n")
	jsx.WriteString("   }\n")
	jsx.WriteString("\n")
	jsx.WriteString("   var m_Closed = activeDocument.activeLayer;\n")
	jsx.WriteString("   m_Closed.opacity = m_Opacity;\n")
	jsx.WriteString("   m_Closed.blendMode = m_BlendMode;\n")
	jsx.WriteString("\n")
	jsx.WriteString("   for(x in m_LinkedLayers) {\n")
	jsx.WriteString("      if(m_LinkedLayers[x].typename == \"LayerSet\")\n")
	jsx.WriteString("         activeDocument.activeLayer.link(m_LinkedLayers[x]);\n")
	jsx.WriteString("   }\n")
	jsx.WriteString("\n")
	jsx.WriteString("   if(m_bHasMask) maskFromSelection();\n")
	jsx.WriteString("\n")
	jsx.WriteString("   return m_Closed;\n")
	jsx.WriteString("}\n")
	jsx.WriteString("//////////////////////////////////////////////////////////////// 上面那一串全部为了折叠图层组 /////////////////////////////////////////////////////////////////////////////////////////\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 选择临时效果图模板图层\n")
	jsx.WriteString("function selectionTempFrame(layerIndex){\n")
	jsx.WriteString("    // 检查ps和检查是否有打开的文档\n")
	jsx.WriteString("    whatApp = String(app.name);  //String version of the app name\n")
	jsx.WriteString("    if(whatApp.search(\"Photoshop\") > 0) {  //Check for photoshop specifically, or this will cause errors\n")
	jsx.WriteString("        // Function Scrubs Document Ancestors from Files\n")
	jsx.WriteString("        if(!documents.length) {\n")
	jsx.WriteString("            alert(\"没有打开的文档，请先使用软件新建专属文档！\")\n")
	jsx.WriteString("            return;\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("        // 为主文档定义变量\n")
	jsx.WriteString("        var masterDocument = app.activeDocument;\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 获取当前脚本所在路径\n")
	jsx.WriteString("        var scriptPath = (new File($.fileName)).parent\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // 定义文件变量\n")
	jsx.WriteString(fmt.Sprintf("        var fileRef = new File(scriptPath+\"/../PSD/%s.psd\");  // 这里传当前路径哦！！！！！！！！！！！！！！！！！！\n", frame))
	jsx.WriteString("        if(fileRef.exists){    // 如果图像存在\n")
	jsx.WriteString("            app.open(fileRef);   // 打开文档\n")
	jsx.WriteString("            app.activeDocument.activeLayer = app.activeDocument.layers[layerIndex];  // 选择图层\n")
	jsx.WriteString("            app.activeDocument.activeLayer.duplicate (masterDocument); // 复制到主文档\n")
	jsx.WriteString("            app.activeDocument.close(SaveOptions.DONOTSAVECHANGES); // 关闭临时文档而不保存更改\n")
	jsx.WriteString("            app.activeDocument.activeLayer.visible = true;  //  切换所选/活动层的可见性\n")
	jsx.WriteString("            closeGroup(activeDocument.activeLayer);  // 折叠当前激活的图层组\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString(fmt.Sprintf("selectionTempFrame(%d) //这里传golang变量哦！！！！！！！！！！！！！！\n", layer))

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/SelectionTempFrame.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "Config/JSX/SelectionTempFrame.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 将矩形选框转换为标记测量标志
func SizeMarks() {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("/**\n")
	jsx.WriteString(" * Size Marks 1.3\n")
	jsx.WriteString(" *\n")
	jsx.WriteString(" * Copyright (c) 2014 Roman Shamin https://github.com/romashamin\n")
	jsx.WriteString(" * and licenced under the MIT licence. All rights not explicitly\n")
	jsx.WriteString(" * granted in the MIT license are reserved. See the included\n")
	jsx.WriteString(" * LICENSE file for more details.\n")
	jsx.WriteString(" *\n")
	jsx.WriteString(" * https://github.com/romashamin\n")
	jsx.WriteString(" * https://twitter.com/romanshamin\n")
	jsx.WriteString(" *\n")
	jsx.WriteString(" * Converts rectangular selection to labeled measurement mark.\n")
	jsx.WriteString(" * Landscape selection → horizontal mark. Portrait or square\n")
	jsx.WriteString(" * selection → vertical mark.\n")
	jsx.WriteString(" *将矩形选择转换为标记的测量标记\n")
	jsx.WriteString(" */\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("var doc = null,\n")
	jsx.WriteString("    docIsExist = false,\n")
	jsx.WriteString("    selBounds = null,\n")
	jsx.WriteString("    selIsExist = false;\n")
	jsx.WriteString("\n")
	jsx.WriteString("var store = {\n")
	jsx.WriteString("    activeLayer: null,\n")
	jsx.WriteString("    rulerUnits: app.preferences.rulerUnits,\n")
	jsx.WriteString("    typeUnits: app.preferences.typeUnits,\n")
	jsx.WriteString("    font: null\n")
	jsx.WriteString("};\n")
	jsx.WriteString("\n")
	jsx.WriteString("app.preferences.rulerUnits = Units.PIXELS;\n")
	jsx.WriteString("app.preferences.typeUnits = TypeUnits.POINTS;\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("try {\n")
	jsx.WriteString("    doc = app.activeDocument;\n")
	jsx.WriteString("    docIsExist = true;\n")
	jsx.WriteString("} catch (e) {\n")
	jsx.WriteString("    alert('Size Marks Script: no document\\n' +\n")
	jsx.WriteString("          'Use File → New... to create one');\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("if (docIsExist) {\n")
	jsx.WriteString("    try {\n")
	jsx.WriteString("        selBounds = doc.selection.bounds;\n")
	jsx.WriteString("        selIsExist = true;\n")
	jsx.WriteString("    } catch (e) {\n")
	jsx.WriteString("        alert('Size Marks Script: no selection\\n' +\n")
	jsx.WriteString("              'Use Rectangular Marquee Tool (M) to create one');\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("if (docIsExist && selIsExist) {\n")
	jsx.WriteString("    doc.suspendHistory(\"Add Size Mark\", \"makeSizeMark()\");\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("function makeSizeMark() {\n")
	jsx.WriteString("  try {\n")
	jsx.WriteString("    var halfMark = 3,\n")
	jsx.WriteString("        txtMargin = 5,\n")
	jsx.WriteString("        baseRes = 72,\n")
	jsx.WriteString("        decimPlaces = 1,\n")
	jsx.WriteString("        layerOpacity = 65,\n")
	jsx.WriteString("        docRes = doc.resolution,\n")
	jsx.WriteString("        scaleRatio = docRes / baseRes,\n")
	jsx.WriteString("        scale = setScaleF(scaleRatio),\n")
	jsx.WriteString("        realUnits = 'px',\n")
	jsx.WriteString("        scaledUnits = 'pt',\n")
	jsx.WriteString("        charThinSpace = '\\u200A'; /* Thin space: \\u2009, hair space: \\u200A */\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var selX1 = selBounds[0].value,\n")
	jsx.WriteString("        selX2 = selBounds[2].value - 1,\n")
	jsx.WriteString("        selY1 = selBounds[1].value,\n")
	jsx.WriteString("        selY2 = selBounds[3].value - 1;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var selWidth = selX2 - selX1,\n")
	jsx.WriteString("        selHeight = selY2 - selY1;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var val = 0,\n")
	jsx.WriteString("        txtLayerPos = [0, 0],\n")
	jsx.WriteString("        layerNamePrefix = 'MSRMNT',\n")
	jsx.WriteString("        txtJ11n = Justification.LEFT;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    store.activeLayer = doc.activeLayer;\n")
	jsx.WriteString("    doc.selection.deselect();\n")
	jsx.WriteString("    var markLayer = doc.artLayers.add();\n")
	jsx.WriteString("\n")
	jsx.WriteString("    setPenToolSize(1.0);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    if (selWidth > selHeight) {\n")
	jsx.WriteString("        // Draw Main Line\n")
	jsx.WriteString("        drawLine([selX1, selY1], [selX2, selY1]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Draw Edge Marks\n")
	jsx.WriteString("        drawLine([selX1, selY1 - halfMark], [selX1, selY1 + halfMark]);\n")
	jsx.WriteString("        drawLine([selX2, selY1 - halfMark], [selX2, selY1 + halfMark]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Set some values for text layer\n")
	jsx.WriteString("        layerNamePrefix = 'W';\n")
	jsx.WriteString("        val = selWidth + 1;\n")
	jsx.WriteString("        txtLayerPos = [selX1 + val / 2, selY1 - txtMargin];\n")
	jsx.WriteString("        txtJ11n = Justification.CENTER;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    } else {\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Draw Main Line\n")
	jsx.WriteString("        drawLine([selX1, selY1], [selX1, selY2]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Draw Edge Marks\n")
	jsx.WriteString("        drawLine([selX1 - halfMark, selY1], [selX1 + halfMark, selY1]);\n")
	jsx.WriteString("        drawLine([selX1 - halfMark, selY2], [selX1 + halfMark, selY2]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        // Set some values for text layer\n")
	jsx.WriteString("        layerNamePrefix = 'H';\n")
	jsx.WriteString("        val = selHeight + 1;\n")
	jsx.WriteString("        txtLayerPos = [selX1 + txtMargin, selY1 + val / 2 + 4];\n")
	jsx.WriteString("        txtJ11n = Justification.LEFT;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    markLayer.opacity = 85;\n")
	jsx.WriteString("    markLayer.move(store.activeLayer, ElementPlacement.PLACEBEFORE);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // Draw label\n")
	jsx.WriteString("    disableArtboardAutoNest();\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var txtLayer = makeTextLayer(),\n")
	jsx.WriteString("        txtLayerItem = txtLayer.textItem;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    store.font = txtLayerItem.font;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    txtLayerItem.font = 'ArialMT';\n")
	jsx.WriteString("    txtLayerItem.autoKerning = AutoKernType.OPTICAL;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    txtLayer.translate(txtLayerPos[0], txtLayerPos[1]);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    txtLayerItem.justification = txtJ11n;\n")
	jsx.WriteString("    txtLayerItem.color = app.foregroundColor;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var label = '';\n")
	jsx.WriteString("\n")
	jsx.WriteString("    if (baseRes !== docRes) {\n")
	jsx.WriteString("        label = formatValueWithUnits(scale(val).toFixed(decimPlaces),\n")
	jsx.WriteString("                scaledUnits,\n")
	jsx.WriteString("                charThinSpace) +\n")
	jsx.WriteString("            ' / ';\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("    label += formatValueWithUnits(val, realUnits, charThinSpace);\n")
	jsx.WriteString("    txtLayerItem.contents = label;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // Finish\n")
	jsx.WriteString("    txtLayer.rasterize(RasterizeType.TEXTCONTENTS);\n")
	jsx.WriteString("    txtLayer.move(markLayer, ElementPlacement.PLACEBEFORE);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    var finalLayer = txtLayer.merge();\n")
	jsx.WriteString("    finalLayer.name = layerNamePrefix + ' ' + label;\n")
	jsx.WriteString("    finalLayer.opacity = layerOpacity;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    app.preferences.rulerUnits = store.rulerUnits;\n")
	jsx.WriteString("    app.preferences.typeUnits = store.typeUnits;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    enableArtboardAutoNest();\n")
	jsx.WriteString("\n")
	jsx.WriteString("    pickTool('marqueeRectTool');\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // HELPERS\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function makePoint(pnt) {\n")
	jsx.WriteString("\n")
	jsx.WriteString("        for (var i = 0; i < pnt.length; i++) {\n")
	jsx.WriteString("            pnt[i] = scale(pnt[i]);\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var point = new PathPointInfo();\n")
	jsx.WriteString("\n")
	jsx.WriteString("        point.anchor = pnt;\n")
	jsx.WriteString("        point.leftDirection = pnt;\n")
	jsx.WriteString("        point.rightDirection = pnt;\n")
	jsx.WriteString("        point.kind = PointKind.CORNERPOINT;\n")
	jsx.WriteString("\n")
	jsx.WriteString("        return point;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function setScaleF(ratio) {\n")
	jsx.WriteString("        return function (value) {\n")
	jsx.WriteString("            return value / ratio;\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function formatValueWithUnits(v, u, space) {\n")
	jsx.WriteString("        space = space || '';\n")
	jsx.WriteString("        return '' + v + space + u;\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function drawLine(start, stop) {\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var startPoint = makePoint(start),\n")
	jsx.WriteString("            stopPoint = makePoint(stop);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var spi = new SubPathInfo();\n")
	jsx.WriteString("        spi.closed = false;\n")
	jsx.WriteString("        spi.operation = ShapeOperation.SHAPEXOR;\n")
	jsx.WriteString("        spi.entireSubPath = [startPoint, stopPoint];\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var uniqueName = 'Line ' + Date.now();\n")
	jsx.WriteString("        var line = doc.pathItems.add(uniqueName, [spi]);\n")
	jsx.WriteString("        line.strokePath(ToolType.PENCIL);\n")
	jsx.WriteString("        line.remove();\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function pickTool(toolName) {\n")
	jsx.WriteString("        var idslct = charIDToTypeID('slct');\n")
	jsx.WriteString("        var desc4 = new ActionDescriptor();\n")
	jsx.WriteString("        var idnull = charIDToTypeID('null');\n")
	jsx.WriteString("        var ref2 = new ActionReference();\n")
	jsx.WriteString("        var idmarqueeRectTool = stringIDToTypeID(toolName);\n")
	jsx.WriteString("        ref2.putClass(idmarqueeRectTool);\n")
	jsx.WriteString("        desc4.putReference(idnull, ref2);\n")
	jsx.WriteString("        executeAction(idslct, desc4, DialogModes.NO);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function makeTextLayer() {\n")
	jsx.WriteString("        var desc = new ActionDescriptor();\n")
	jsx.WriteString("        var ref = new ActionReference();\n")
	jsx.WriteString("        ref.putClass(app.charIDToTypeID('TxLr'));\n")
	jsx.WriteString("        desc.putReference(app.charIDToTypeID('null'), ref);\n")
	jsx.WriteString("        var desc2 = new ActionDescriptor();\n")
	jsx.WriteString("        desc2.putString(app.charIDToTypeID('Txt '), \"text\");\n")
	jsx.WriteString("        var list2 = new ActionList();\n")
	jsx.WriteString("        desc2.putList(app.charIDToTypeID('Txtt'), list2);\n")
	jsx.WriteString("        desc.putObject(app.charIDToTypeID('Usng'), app.charIDToTypeID('TxLr'), desc2);\n")
	jsx.WriteString("        executeAction(app.charIDToTypeID('Mk  '), desc, DialogModes.NO);\n")
	jsx.WriteString("        return doc.activeLayer\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function disableArtboardAutoNest() {\n")
	jsx.WriteString("        var ideditArtboardEvent = stringIDToTypeID( \"editArtboardEvent\" );\n")
	jsx.WriteString("        var desc3 = new ActionDescriptor();\n")
	jsx.WriteString("        var idnull = charIDToTypeID( \"null\" );\n")
	jsx.WriteString("            var ref2 = new ActionReference();\n")
	jsx.WriteString("            var idLyr = charIDToTypeID( \"Lyr \" );\n")
	jsx.WriteString("            var idOrdn = charIDToTypeID( \"Ordn\" );\n")
	jsx.WriteString("            var idTrgt = charIDToTypeID( \"Trgt\" );\n")
	jsx.WriteString("            ref2.putEnumerated( idLyr, idOrdn, idTrgt );\n")
	jsx.WriteString("        desc3.putReference( idnull, ref2 );\n")
	jsx.WriteString("        var idautoNestEnabled = stringIDToTypeID( \"autoNestEnabled\" );\n")
	jsx.WriteString("        desc3.putBoolean( idautoNestEnabled, false );\n")
	jsx.WriteString("        executeAction( ideditArtboardEvent, desc3, DialogModes.NO );\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    function enableArtboardAutoNest() {\n")
	jsx.WriteString("        var ideditArtboardEvent = stringIDToTypeID( \"editArtboardEvent\" );\n")
	jsx.WriteString("        var desc3 = new ActionDescriptor();\n")
	jsx.WriteString("        var idnull = charIDToTypeID( \"null\" );\n")
	jsx.WriteString("            var ref2 = new ActionReference();\n")
	jsx.WriteString("            var idLyr = charIDToTypeID( \"Lyr \" );\n")
	jsx.WriteString("            var idOrdn = charIDToTypeID( \"Ordn\" );\n")
	jsx.WriteString("            var idTrgt = charIDToTypeID( \"Trgt\" );\n")
	jsx.WriteString("            ref2.putEnumerated( idLyr, idOrdn, idTrgt );\n")
	jsx.WriteString("        desc3.putReference( idnull, ref2 );\n")
	jsx.WriteString("        var idautoNestEnabled = stringIDToTypeID( \"autoNestEnabled\" );\n")
	jsx.WriteString("        desc3.putBoolean( idautoNestEnabled, true );\n")
	jsx.WriteString("        executeAction( ideditArtboardEvent, desc3, DialogModes.NO );\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("    /**\n")
	jsx.WriteString("     * Source: https://forums.adobe.com/thread/962285?start=0&tstart=0\n")
	jsx.WriteString("     * Comment for Feb 16, 2012 7:18 AM\n")
	jsx.WriteString("     */\n")
	jsx.WriteString("    function setPenToolSize(dblSize) {\n")
	jsx.WriteString("        var idslct = charIDToTypeID('slct');\n")
	jsx.WriteString("        var desc3 = new ActionDescriptor();\n")
	jsx.WriteString("        var idnull = charIDToTypeID('null');\n")
	jsx.WriteString("        var ref2 = new ActionReference();\n")
	jsx.WriteString("        var idPcTl = charIDToTypeID('PcTl');\n")
	jsx.WriteString("        ref2.putClass(idPcTl);\n")
	jsx.WriteString("        desc3.putReference(idnull, ref2);\n")
	jsx.WriteString("        executeAction(idslct, desc3, DialogModes.NO);\n")
	jsx.WriteString("\n")
	jsx.WriteString("        var idsetd = charIDToTypeID('setd');\n")
	jsx.WriteString("        var desc2 = new ActionDescriptor();\n")
	jsx.WriteString("        var ref1 = new ActionReference();\n")
	jsx.WriteString("        var idBrsh = charIDToTypeID('Brsh');\n")
	jsx.WriteString("        var idOrdn = charIDToTypeID('Ordn');\n")
	jsx.WriteString("        var idTrgt = charIDToTypeID('Trgt');\n")
	jsx.WriteString("        ref1.putEnumerated(idBrsh, idOrdn, idTrgt);\n")
	jsx.WriteString("        desc2.putReference(idnull, ref1);\n")
	jsx.WriteString("        var idT = charIDToTypeID('T   ');\n")
	jsx.WriteString("        var idmasterDiameter = stringIDToTypeID('masterDiameter');\n")
	jsx.WriteString("        var idPxl = charIDToTypeID('#Pxl');\n")
	jsx.WriteString("        desc3.putUnitDouble(idmasterDiameter, idPxl, dblSize);\n")
	jsx.WriteString("        desc2.putObject(idT, idBrsh, desc3);\n")
	jsx.WriteString("        executeAction(idsetd, desc2, DialogModes.NO);\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("  } catch (e) {\n")
	jsx.WriteString("    alert(e.line + '\\n' + e)\n")
	jsx.WriteString("  }\n")
	jsx.WriteString("}\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/SizeMarks.jsx", jsxStr)
}

// 暗号-98的实现
func SaveForWeb() {
	// 使用最高效的字符串拼接
	var jsx = strings.Builder{}

	jsx.WriteString("// 清理元数据\r\n")
	jsx.WriteString("function deleteDocumentAncestorsMetadata() {\r\n")
	jsx.WriteString("    // 清理元数据四步骤\r\n")
	jsx.WriteString("    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject(\"lib:AdobeXMPScript\");\r\n")
	jsx.WriteString("    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);\r\n")
	jsx.WriteString("    // Begone foul Document Ancestors!\r\n")
	jsx.WriteString("    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, \"DocumentAncestors\");\r\n")
	jsx.WriteString("    app.activeDocument.xmpMetadata.rawData = xmp.serialize();\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 另存为web\r\n")
	jsx.WriteString("function saveAsWeb() {\r\n")
	jsx.WriteString("    // 更新进度条\r\n")
	jsx.WriteString("    updateProgress(1, 4);\r\n")
	jsx.WriteString("    // 清理元数据\r\n")
	jsx.WriteString("    deleteDocumentAncestorsMetadata();\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 更新进度条\r\n")
	jsx.WriteString("    updateProgress(2, 4);\r\n")
	jsx.WriteString("    // 获取当前脚本所在路径\r\n")
	jsx.WriteString("    var scriptPath = (new File($.fileName)).parent;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 定义文件保存位置\r\n")
	jsx.WriteString("    var savePath = new File(scriptPath + \"/../Picture/主图/dp.jpg\");\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    var jpgOpt = new ExportOptionsSaveForWeb();\r\n")
	jsx.WriteString("    jpgOpt.format = SaveDocumentType.JPEG;  // 保存为jpg\r\n")
	jsx.WriteString("    jpgOpt.includeProfile = false;  //装入颜色配置文件\r\n")
	jsx.WriteString("    jpgOpt.interlaced = false;  // 交错\r\n")
	jsx.WriteString("    jpgOpt.optimized = true;  //最优化\r\n")
	jsx.WriteString("    jpgOpt.blur = 0;    // 默认 0.0 不模糊。\r\n")
	jsx.WriteString("    jpgOpt.matteColor = new RGBColor(); // 把杂边颜色染成白色\r\n")
	jsx.WriteString("    jpgOpt.matteColor.red = 255;\r\n")
	jsx.WriteString("    jpgOpt.matteColor.green = 255;\r\n")
	jsx.WriteString("    jpgOpt.matteColor.blue = 255;\r\n")
	jsx.WriteString("    jpgOpt.quality = 100;  // 品质   100是最高画质\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 更新进度条\r\n")
	jsx.WriteString("    updateProgress(3, 4);\r\n")
	jsx.WriteString("    activeDocument.exportDocument(savePath, ExportType.SAVEFORWEB,);\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 更新进度条\r\n")
	jsx.WriteString("    updateProgress(4, 4);\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("function main() {\r\n")
	jsx.WriteString("    // 判断是否有打开的文件\r\n")
	jsx.WriteString("    if (!documents.length) {\r\n")
	jsx.WriteString("        alert(\"没有打开的文档，请打开一个文档来运行此脚本！\");\r\n")
	jsx.WriteString("        return;\r\n")
	jsx.WriteString("    }\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("    // 进度条调用另存函数\r\n")
	jsx.WriteString("    app.doForcedProgress(\"正在导出文件... \", \"saveAsWeb()\")  // 添加进度条\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 主函数\r\n")
	jsx.WriteString("main();")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/SaveForWeb.jsx", jsxStr)
}
