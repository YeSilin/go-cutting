package generate

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//生成清除元数据第二版js，让文件跟小巧，带进度条
func ClearMetadata() {
	const script = `// 清除元数据的主要函数
function deleteDocumentAncestorsMetadata() {
    if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
    // 如果当前是主文档就更新进度条
    if (app.activeDocument == mainDocument) {
        // 更新进度条
        updateProgress(1, maxBar);
    }

    var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);
    // 如果当前是主文档就更新进度条
    if (app.activeDocument == mainDocument) {
        // 更新进度条
        updateProgress(2, maxBar);
    }

    // Begone foul Document Ancestors!
    xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");
    // 如果当前是主文档就更新进度条
    if (app.activeDocument == mainDocument) {
        // 更新进度条
        updateProgress(3, maxBar);
    }

    app.activeDocument.xmpMetadata.rawData = xmp.serialize();
    // 如果当前是主文档就更新进度条
    if (app.activeDocument == mainDocument) {
        // 更新进度条
        updateProgress(4, maxBar);
    }

    // 递归调用函数，清除当前文档中的所有智能对象元数据
    clearDocumentAncestorsForAllLayers(app.activeDocument);

    if (app.activeDocument !== mainDocument) {
        // 不是主文档，那么就是智能对象，关闭并保存它
        app.activeDocument.close(SaveOptions.SAVECHANGES);

    } else {
        // 主文档就不自动保存了
        //app.activeDocument.save();
    }
}

// 清除文档中所有智能对象的元数据
function clearDocumentAncestorsForAllLayers(doc) {
    try {
        // 如果当前文档未定义就返回
        if (doc == undefined) {
            return;
        }

        for (var i = 0; i < doc.layers.length; i++) {
            var curLayer = doc.layers[i];


            // 如果选中的图层类型不是普通图层，即代表是图层组
            if (curLayer.typename != "ArtLayer") {
                clearDocumentAncestorsForAllLayers(curLayer);
                continue;
            }

            // 如果当前图层是智能对象
            if (curLayer.kind == "LayerKind.SMARTOBJECT") {

                // 如果当前是主文档就更新进度条
                if (app.activeDocument == mainDocument) {
                    // 更新进度条
                    updateProgress(4 + i, maxBar);
                }

                app.activeDocument.activeLayer = curLayer;

                // 打开这个智能对象
                var idplacedLayerEditContents = stringIDToTypeID("placedLayerEditContents");
                var actionDescriptor = new ActionDescriptor();
                executeAction(idplacedLayerEditContents, actionDescriptor, DialogModes.NO);

                updateProgress(0, maxBar);
                // 这行代码看不懂，大概意思是，所打开的智能对象当前选中的图层，如果是被打开的智能对象本身，那么就不清理
                if (app.activeDocument.activeLayer == curLayer) {
                    continue;
                }

                // 递归调用清除元数据
                deleteDocumentAncestorsMetadata();
            }
        }
    } catch (e) {
        // 清除失败就不弹窗了
        //alert("Layer clean fail.name="+doc+";e="+e)
    }
}


// 生成进度条函数
function progressBar() {
    // 进度条调用清除元数据函数
    app.doForcedProgress("正在清除元数据... ", "deleteDocumentAncestorsMetadata()")  // 添加进度条
}


function main() {
    whatApp = String(app.name);  //String version of the app name
    if (whatApp.search("Photoshop") > 0) {  //Check for photoshop specifically, or this will cause errors

        //  函数scrubs从文件中提取文档祖先
        if (!documents.length) {
            alert("没有打开的文档，请打开一个文档来运行此脚本！");
            return;
        }

        // 定义主文档，因为要打开很多智能对象
        mainDocument = app.activeDocument;

        // 定义一个变量表示总进度与当前进度
        maxBar = app.activeDocument.layers.length + 4;


        // 生成历史，历史调用进度条
        app.activeDocument.suspendHistory("清除元数据", "progressBar()");  // 生成历史记录
    }
}

// 定义主文档，因为要打开很多智能对象
var mainDocument;

// 定义一个变量表示总进度与当前进度
var maxBar;

// 运行此脚本
main();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/clearMetadata.jsx", script)
}

//生成清除元数据js，不清理智能对象，让文件跟小巧 无弹窗版本！
func ClearMetadataNoPopUp() {
	const script = `// 清除元数据无弹窗版，并且不会清理智能对象
function deleteDocumentAncestorsMetadata() {
    whatApp = String(app.name);  //String version of the app name
    if (whatApp.search("Photoshop") > 0) {  //Check for photoshop specifically, or this will cause errors

        // Function Scrubs Document Ancestors from Files
        if (!documents.length) {
            // alert("没有打开的文档，请打开一个文档来运行此脚本！")
            return;
        }

        if (ExternalObject.AdobeXMPScript == undefined) ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");

        var xmp = new XMPMeta(activeDocument.xmpMetadata.rawData);

        // Begone foul Document Ancestors!
        xmp.deleteProperty(XMPConst.NS_PHOTOSHOP, "DocumentAncestors");

        app.activeDocument.xmpMetadata.rawData = xmp.serialize();
    }
}

// Now run the function to remove the document ancestors
deleteDocumentAncestorsMetadata();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/clearMetadataNoPopUp.jsx", script)
}

// 为当前文档添加黑边
func BlackEdge() {
	const script = `// 为当前文档添加黑边
function BlackEdge() {
    // 判断是否有打开的文件
    if (!documents.length) {
        alert("没有打开的文档，请打开一个文档来运行此脚本！");
        return;
    }

    // 设置首选项新文档预设单位是厘米，PIXELS是像素
    app.preferences.rulerUnits = Units.CM;

    // 保存当前背景颜色
    var nowColor = app.backgroundColor;

    // 定义一个对象颜色是黑色
    var black = new SolidColor();
    black.rgb.hexValue = "000000";
    app.backgroundColor = black;

    // 生成历史记录
    app.activeDocument.suspendHistory("向四周添加0.5厘米黑边！", "addEdge()");

    // 恢复之前的背景颜色
    app.backgroundColor = nowColor;
}

// 此函数用于生成历史记录
function addEdge() {
    // 新建一个空白图层用于合并
    app.activeDocument.artLayers.add();
    // 合并全部可见图层
    app.activeDocument.mergeVisibleLayers();
    // 转为背景图层不然添加黑边会无效
    app.activeDocument.activeLayer.isBackgroundLayer = true;

    // 获取当前文档的高度与宽度
    var width = app.activeDocument.width + 0.5;
    var height = app.activeDocument.height + 0.5;

    // 重设画布大小
    app.activeDocument.resizeCanvas(UnitValue(width), UnitValue(height), AnchorPosition.MIDDLECENTER);
}

// 生成黑边
BlackEdge();`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/blackEdge.jsx", script)
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

// 暗号-98的实现
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
	savePath := fmt.Sprintf("%s/主图/dp.jpg", originalPath)

	// 利用给定数据渲染模板，并将结果写入f
	_ = tmpl.Execute(f, savePath)
}

//如果画布的高和宽同时大于148则提示
func MaxCanvas(width, height float64) {
	// 默认不写入
	write := false

	if width > 180 && height > 180 {
		// 这其实是红色
		color.LightBlue.Println("（已超不透最大180cm。）")
		write = true
	} else if width > 148 && height > 148 {
		color.LightBlue.Println("（已超半透最大148cm。）")
		write = true
	} else {
		fmt.Println()
		return
	}

	// 如果要写入
	if write {
		//最大画布字体图层提示
		var jsx = strings.Builder{}

		jsx.WriteString("// 无聊加了个画布大小判断\n")
		jsx.WriteString(fmt.Sprintf("if ((%f>180) && (%f>180)) {\n", width, height))
		jsx.WriteString("     // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超不透最大180cm\", \"maxCanvas(\\\"注意：已超不透最大180cm。\\\",  \\\"9d2e2d\\\")\");\n")
		jsx.WriteString(fmt.Sprintf("} else if ((%f>148) && (%f>148)) {\n", width, height))
		jsx.WriteString("      // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超半透最大148cm\", \"maxCanvas(\\\"注意：已超半透最大148cm。\\\",  \\\"77bb11\\\")\");\n")
		jsx.WriteString("}\n")
		jsx.WriteString("\n")
		jsx.WriteString("\n")
		jsx.WriteString("function maxCanvas(text, rgbValue){\n")
		jsx.WriteString("    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。\n")
		jsx.WriteString("    var newLayer = app.activeDocument.artLayers.add();\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。\n")
		jsx.WriteString("    newLayer.kind = LayerKind.TEXT;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //把图层 newLayer 的文本内容类型变为”文本框“。\n")
		jsx.WriteString("    newLayer.textItem.kind = TextType.PARAGRAPHTEXT;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置图层 newLayer 的文本框宽度与高度。\n")
		jsx.WriteString("    newLayer.textItem.width = app.activeDocument.width*0.8;\n")
		jsx.WriteString("    newLayer.textItem.height = app.activeDocument.width*0.1;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素。\n")
		jsx.WriteString("    //newLayer.textItem.position= [UnitValue(\"50px\"), UnitValue(\"100px\")]\n")
		jsx.WriteString("    newLayer.textItem.position= [UnitValue(app.activeDocument.width*0.1), UnitValue((app.activeDocument.height*0.5)-(app.activeDocument.width*0.025))];\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本字体大小为“40 点”。\n")
		jsx.WriteString("    newLayer.textItem.size = UnitValue(app.activeDocument.width*0.05);\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本内容。\n")
		jsx.WriteString("    newLayer.textItem.contents= text;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本框对齐方式为居中对齐。\n")
		jsx.WriteString("    newLayer.textItem.justification = Justification.CENTER;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //创建一个色彩变量 c   ，颜色为 #77bb11。\n")
		jsx.WriteString("    var c = new SolidColor();\n")
		jsx.WriteString("    c.rgb.hexValue = rgbValue;\n")
		jsx.WriteString("\n")
		jsx.WriteString("    //设置 newLayer 的文本颜色为 c。\n")
		jsx.WriteString("    newLayer.textItem.color = c;\n")
		jsx.WriteString("}\n")

		// 转成字符串格式
		jsxStr := jsx.String()
		// 追加写入
		tools.WriteFile("config/jsx/newDocument.jsx", jsxStr)
	}
}

//生成自带清除元数据的另存
func SaveAsJPEG() {
	const script = `// 用于快捷另存为jpg

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
    var idplacedLayerUpdateAllModified = stringIDToTypeID( "placedLayerUpdateAllModified" );
    executeAction( idplacedLayerUpdateAllModified, undefined, DialogModes.NO );
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
    var TmpFile = new File("~/Desktop/GoCutting/" + name[0]+" 副本");

    // 保存文件类型
    // var saveType = new Array("JPEG Files: *.jpg", "PNG Files: *.png");

    // saveAs( 文件, 选项, 作为副本, 扩展名大小写 )
    //调用[document]的[saveAs]另存方法，使用上面设置的各种参数，将当前文档导出并转换为JPEG格式的文档
    app.activeDocument.saveAs(TmpFile.saveDlg("另存为", "JPEG Files: *.jpg"), exportOptionsSave, true, Extension.LOWERCASE);
}


// 判断是否有打开的文件
if (!documents.length) {
    //alert("没有打开的文档，请打开一个文档来运行此脚本！");
    // return;
} else {
    // 如果出错就返回最开始
    try {
        // 生成历史记录
        app.activeDocument.suspendHistory("另存为", "optimized()");
    } catch (error) {
        // 忽略错误
    }
}`

	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("config/jsx/saveAsJPEG.jsx", script)
}

// 生成详情页替换智能对象的脚本
func ReplaceDetailsPage(originalPath string) {
	// 返回绝对路径
	originalPath, err := filepath.Abs(originalPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	// 全部换成正斜杠
	originalPath = strings.Replace(originalPath, "\\", "/", -1)

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(fmt.Sprintf("%s/*.jpg", originalPath))
	// 如果jpg小于一张就不执行
	if len(jpgSlice) < 1 {
		fmt.Println("\n【提示】脚本注入失败，因为 Picture 文件夹下没有 jpg 格式图片！")
		// 打开套图文件夹
		go exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", viper.GetString("picture"))).Run()
		return
	}

	go func() {
		// 获取白底图
		minImage, isMinImage := tools.MinWhiteBackground(fmt.Sprintf("%s/*.jpg", originalPath))
		// 去掉白底图
		var result []string
		if isMinImage {
			// 不是白底图的都添加到result切片
			for i := 0; i < len(jpgSlice); i++ {
				if minImage != jpgSlice[i] {
					result = append(result, jpgSlice[i])
				}
			}
			// 更新最新切片
			jpgSlice = result
		}

		// 修改成js脚本可以看懂的路径
		for i := 0; i < len(jpgSlice); i++ {
			jpgSlice[i] = strings.Replace(jpgSlice[i], "\\", "/", -1)
			jpgSlice[i] = strings.Replace(jpgSlice[i], ":", "", 1)
			jpgSlice[i] = "/" + jpgSlice[i]
		}

		// 解析指定文件生成模板对象
		tmpl, err := template.ParseFiles("config/jsx/template/replaceDetailsPage.gohtml")
		if err != nil {
			logrus.Error(err)
			return
		}

		// 创建文件，返回两个值，一是创建的文件，二是错误信息
		f, err := os.Create("config/jsx/replaceDetailsPage.jsx")
		if err != nil { // 如果有错误，打印错误，同时返回
			logrus.Error(err)
			return
		}

		// 利用给定数据渲染模板，并将结果写入f
		_ = tmpl.Execute(f, tools.StrToJsArray("srcArray", jpgSlice))

		// 关闭文件
		f.Close()

		// 创建一个协程使用cmd来运行脚本
		dataPath := "Config/jsx/ReplaceDetailsPage.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
	}()

	fmt.Println("\n【提示】脚本注入成功，正在自动替换详情页中名字以【dp】开头的智能对象图层！")
}
