package generate

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"path/filepath"
	"strings"
)

// 生成通用套主图js脚本，生成并运行。目前已被内部算法取代
func UniversalMasterGraph(count int) {
	var jsx = strings.Builder{}

	jsx.WriteString("// 嵌入智能对象\n")
	jsx.WriteString("function embedSmartObjects(path) {\n")
	jsx.WriteString("    var idPlc = charIDToTypeID(\"Plc \");\n")
	jsx.WriteString("    var desc122 = new ActionDescriptor();\n")
	jsx.WriteString("    var idnull = charIDToTypeID(\"null\");\n")
	jsx.WriteString("    desc122.putPath(idnull, new File(path));\n")
	jsx.WriteString("    executeAction(idPlc, desc122, DialogModes.NO);\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 保存为web格式jpeg\n")
	jsx.WriteString("function saveForWebJPEG(path) {\n")
	jsx.WriteString("    var fileObj = new File(path); // 文件保存路径\n")
	jsx.WriteString("    var jpgOpt = new ExportOptionsSaveForWeb();\n")
	jsx.WriteString("    jpgOpt.format = SaveDocumentType.JPEG;  // 保存为jpg\n")
	jsx.WriteString("    jpgOpt.includeProfile = false;  //装入颜色配置文件\n")
	jsx.WriteString("    jpgOpt.interlaced = false;  // 交错\n")
	jsx.WriteString("    jpgOpt.optimized = true;  //最优化\n")
	jsx.WriteString("    jpgOpt.blur = 0;    // 默认 0.0 不模糊。\n")
	jsx.WriteString("    jpgOpt.matteColor = new RGBColor(); // 把杂边颜色染成白色\n")
	jsx.WriteString("    jpgOpt.matteColor.red = 255;\n")
	jsx.WriteString("    jpgOpt.matteColor.green = 255;\n")
	jsx.WriteString("    jpgOpt.matteColor.blue = 255;\n")
	jsx.WriteString("    jpgOpt.quality = 100;  // 品质   100是最高画质\n")
	jsx.WriteString("    activeDocument.exportDocument(fileObj, ExportType.SAVEFORWEB, jpgOpt);\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("function main() {\n")
	jsx.WriteString(fmt.Sprintf("    var count = %d; // 这里传golang参数哦！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！\n", count))
	jsx.WriteString("\n")
	jsx.WriteString("    // 获取当前脚本所在路径\n")
	jsx.WriteString("    var scriptPath = (new File($.fileName)).parent;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    // 定义文件变量\n")
	jsx.WriteString("    var fileRef = new File(scriptPath + \"/../PSD/UniversalMasterGraph.psd\");\n")
	jsx.WriteString("    if (fileRef.exists) {    // 如果图像存在\n")
	jsx.WriteString("        app.open(fileRef);   // 打开文档\n")
	jsx.WriteString("\n")
	jsx.WriteString("        //  根据数量循环导出\n")
	jsx.WriteString("        for (i = 1; i <= count; i++) {\n")
	jsx.WriteString("            embedSmartObjects(scriptPath + \"/../Picture/\" + i + \".jpg\"); // 嵌入智能对象\n")
	jsx.WriteString("            saveForWebJPEG(scriptPath + \"/../Picture/主图/\" + i + \".jpg\");  // 保存路径\n")
	jsx.WriteString("        }\n")
	jsx.WriteString("        app.activeDocument.close(SaveOptions.DONOTSAVECHANGES); // 关闭文档而不保存更改\n")
	jsx.WriteString("    }\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 运行\n")
	jsx.WriteString("main();\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/UniversalMasterGraph.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "Config/jsx/UniversalMasterGraph.jsx"
	exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 生成通用套详情页js脚本，生成并运行。
func ReplaceDetailsPage() {
	var jsx = strings.Builder{}

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 {
		fmt.Println("\n【提示】脚本注入失败，因为 Picture 文件夹下没有 jpg 格式图片！")
		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
		return
	}

	// 获取白底图
	minImage, isMinImage := tools.MinWhiteBackground("Config\\Picture\\*.jpg")

	result := []string{}
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
	for i := 0; i < len(jpgSlice); i++ {
		// 去掉Config\Picture\`
		jpgSlice[i] = strings.TrimPrefix(jpgSlice[i], `Config\Picture\`)
	}

	jsx.WriteString("// 替换智能对象\r\n")
	jsx.WriteString("function replaceContents(newFile, theSO) {\r\n")
	jsx.WriteString("    app.activeDocument.activeLayer = theSO;\r\n")
	jsx.WriteString("    // =======================================================\r\n")
	jsx.WriteString("    var idplacedLayerReplaceContents = stringIDToTypeID(\"placedLayerReplaceContents\");\r\n")
	jsx.WriteString("    var desc3 = new ActionDescriptor();\r\n")
	jsx.WriteString("    var idnull = charIDToTypeID(\"null\");\r\n")
	jsx.WriteString("    desc3.putPath(idnull, new File(newFile));\r\n")
	jsx.WriteString("    var idPgNm = charIDToTypeID(\"PgNm\");\r\n")
	jsx.WriteString("    desc3.putInteger(idPgNm, 1);\r\n")
	jsx.WriteString("    executeAction(idplacedLayerReplaceContents, desc3, DialogModes.NO);\r\n")
	jsx.WriteString("    return app.activeDocument.activeLayer\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("// 递归函数\r\n")
	jsx.WriteString("function replaceSmartObjects(doc) {\r\n")
	jsx.WriteString("    for (var i = 0; i < doc.layers.length; i++) {\r\n")
	jsx.WriteString("        var curLayer = doc.layers[i];\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("        // 如果选中的图层类型不是普通图层，即代表是图层组\r\n")
	jsx.WriteString("        if (curLayer.typename != \"ArtLayer\") {\r\n")
	jsx.WriteString("            replaceSmartObjects(curLayer);\r\n")
	jsx.WriteString("            continue;\r\n")
	jsx.WriteString("        }\r\n")
	jsx.WriteString("        // 如果当前图层是智能对象\r\n")
	jsx.WriteString("        if (curLayer.kind == \"LayerKind.SMARTOBJECT\") {\r\n")
	jsx.WriteString("            // 判断字符串是否以 dp 开头 upper and lower case\r\n")
	jsx.WriteString("            var upperCase = curLayer.name.indexOf(\"dp\"); // 小写\r\n")
	jsx.WriteString("            var upperCase2 = curLayer.name.indexOf(\"dP\"); // 小写\r\n")
	jsx.WriteString("             var lowerCase = curLayer.name.indexOf(\"DP\");  // 大写\r\n")
	jsx.WriteString("             var lowerCase2 = curLayer.name.indexOf(\"Dp\");  // 大写\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("            if (upperCase == 0 || lowerCase == 0 || upperCase2 == 0 || lowerCase2 == 0) { // 表示curLayer.name是以dp开头；\r\n")
	jsx.WriteString("                // 替换智能对象\r\n")
	jsx.WriteString("                replaceContents(new File(srcArray[src]), curLayer);\r\n")
	jsx.WriteString("                curLayer.name = \"DP 此图层由GoCutting替换\"\r\n")
	jsx.WriteString("                src++; // 索引加1\r\n")
	jsx.WriteString("                // 如果索引大于源文件数量就不再替换\r\n")
	jsx.WriteString("                if (src == srcArray.length) {\r\n")
	jsx.WriteString("                    return\r\n")
	jsx.WriteString("                }\r\n")
	jsx.WriteString("            }\r\n")
	jsx.WriteString("        }\r\n")
	jsx.WriteString("    }\r\n")
	jsx.WriteString("}\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("//  函数scrubs从文件中提取文档祖先\r\n")
	jsx.WriteString("if (!documents.length) {\r\n")
	jsx.WriteString("    alert(\"没有打开的文档，请打开一个文档来运行此脚本！\");\r\n")
	jsx.WriteString("} else {\r\n")
	jsx.WriteString("    try {\r\n")
	jsx.WriteString("        var doc = app.activeDocument;\r\n")
	jsx.WriteString("        // 源文件\r\n")
	jsx.WriteString(fmt.Sprintf("        var srcArray = %s; //  这里传golang排版好的字符串哦！！！！！！！！！！！\r\n", tools.StrToJsArray(jpgSlice)))
	jsx.WriteString("        // 这是源文件索引\r\n")
	jsx.WriteString("        var src = 0;\r\n")
	jsx.WriteString("\r\n")
	jsx.WriteString("        // 获取当前脚本所在路径\r\n")
	jsx.WriteString("        var scriptPath = (new File($.fileName)).parent;\r\n")
	jsx.WriteString("        // 为源文件加上完整路径\r\n")
	jsx.WriteString("        for (i = 0; i < srcArray.length; i++) {\r\n")
	jsx.WriteString("            srcArray[i] = scriptPath + \"/../Picture/\" + srcArray[i]\r\n")
	jsx.WriteString("        }\r\n")
	jsx.WriteString("        // 运行图层替换\r\n")
	jsx.WriteString("        replaceSmartObjects(doc);\r\n")
	jsx.WriteString("    } catch (error) {\r\n")
	jsx.WriteString("        //发生错误执行的代码\r\n")
	jsx.WriteString("    }\r\n")
	jsx.WriteString("}")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("Config/JSX/ReplaceDetailsPage.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "Config/jsx/ReplaceDetailsPage.jsx"
	exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	fmt.Println("\n【提示】脚本注入成功，正在自动替换详情页中名字以【dp】开头的智能对象图层！")
}


// 生成 修改全部jpg为 300ppi
func AllResolution300() {
	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(".\\Config\\Picture\\*.jpg")
	// 如果png和jpg都小于一张就不执行
	if len(jpgSlice) < 1 {
		fmt.Println("\n【提示】脚本注入失败，因为 Picture 文件夹下没有 jpg 格式图片！")
		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
		return
	}

	for i := 0; i < len(jpgSlice); i++ {
		// 去掉Config\Picture\`
		jpgSlice[i] = strings.TrimPrefix(jpgSlice[i], `Config\Picture\`)
	}


}