package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"strings"
)

//生成清除元数据第三版js，让文件跟小巧，带进度条
func ClearMetadata() {
	const script = `// 清除元数据的主要函数
function deleteDocumentAncestorsMetadata() {
    // 清理垃圾四步骤
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
}


// 数组去重，双层循环，外层循环元素，内层循环时比较值。值相同则删去这个值。
function unique(arr) {
    for (var i = 0; i < arr.length; i++) {
        for (var j = i + 1; j < arr.length; j++) {
            if (arr[i] == arr[j]) {         //第一个等同于第二个，splice方法删除第二个
                arr.splice(j, 1);
                j--;
            }
        }
    }
    return arr;
}

// 打开当前选中的智能对象
function openSmartObject() {
    // 打开这个智能对象
    var idplacedLayerEditContents = stringIDToTypeID("placedLayerEditContents");
    var actionDescriptor = new ActionDescriptor();
    executeAction(idplacedLayerEditContents, actionDescriptor, DialogModes.NO);
}


// 为了避免重复清理打开，一次性打开全部智能对象
function openAllSmartObject(doc) {
    try {
// 如果当前文档未定义就返回
        if (doc == undefined) {
            return;
        }

        // 遍历当前文档的全部图层
        for (var i = 0; i < doc.layers.length; i++) {
            var curLayer = doc.layers[i];

            // 如果当前图层类型不是普通图层，即代表是图层组
            if (curLayer.typename != "ArtLayer") {
                // 那么继续打开
                openAllSmartObject(curLayer);
                continue;
            }

            // 如果当前图层是智能对象，并且名字不叫矢量智能对象
            if (curLayer.kind == "LayerKind.SMARTOBJECT" && curLayer.name != "矢量智能对象") {


                // 激活图层
                app.activeDocument.activeLayer = curLayer;

                // 打开之前先定义一下当前文档
                var curDoc = app.activeDocument

                // 先确认下当前打开了多少文档
                var openDocCount = documents.length

                // 打开智能对象
                openSmartObject()

                // 如果没有打开成功说明可能是矢量智能对象
                if (openDocCount == documents.length) {
                    continue
                }

                // 打开后追加保存到已打开的智能列表
                openDocumentList.push(app.activeDocument)

                // 那么继续打开全部智能对象
                openAllSmartObject(app.activeDocument);

                // 打开后回到之前的文档
                app.activeDocument = curDoc
            }
        }
    } catch (e) {
        // 清除失败就不弹窗了
    }
}


// 删除全部文档元数据 
function deleteAllDocumentAncestorsMetadata() {
    // 先清理主文档元数据
    deleteDocumentAncestorsMetadata()

    // 打开全部智能对象
    openAllSmartObject(mainDocument)

    // 数组去重
    openDocumentList = unique(openDocumentList)

    // 开始循环清理
    for (var i = 0; i < openDocumentList.length; i++) {
        // 先激活文档
        app.activeDocument = openDocumentList[i]
        // 然后清理数据
        deleteDocumentAncestorsMetadata()
        // 最后保存并关闭
        app.activeDocument.close(SaveOptions.SAVECHANGES);
    }
}


// 生成进度条函数
function progressBar() {
    // 进度条调用清除元数据函数
    app.doForcedProgress("正在清除元数据... ", "deleteAllDocumentAncestorsMetadata()")
}


function main() {
    whatApp = String(app.name);  //String version of the app name
    if (whatApp.search("Photoshop") > 0) {  //Check for photoshop specifically, or this will cause errors

        //  函数scrubs从文件中提取文档祖先
        if (!documents.length) {
            //alert("没有打开的文档，请打开一个文档来运行此脚本！");
            return;
        }

        // 定义主文档
        mainDocument = app.activeDocument;

        // 生成历史，历史调用进度条
        app.activeDocument.suspendHistory("清除元数据", "progressBar()");  // 生成历史记录
    }
}

// 声明主文档，因为要打开很多智能对象
var mainDocument;

// 保存因为是智能对象而打开的文档列表
var openDocumentList = new Array();

// 运行此脚本
main();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/clearMetadata.jsx", script)
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
