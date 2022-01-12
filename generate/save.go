package generate

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/yesilin/go-cutting/tools"
	"strings"
)

// MaxCanvas 如果画布的高和宽同时大于180则提示
func MaxCanvas(width, height float64) {
	// 默认不写入
	var write bool

	if width > 200 && height > 200 {
		// 这其实是红色
		color.LightBlue.Println("（已超半透最大200cm。）")
		write = true
	} else if width > 180 && height > 180 {
		color.LightBlue.Println("（已超不透最大180cm。）")
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
		jsx.WriteString(fmt.Sprintf("if ((%f>200) && (%f>200)) {\n", width, height))
		jsx.WriteString("     // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超半透最大200cm\", \"maxCanvas(\\\"注意：已超半透最大200cm。\\\",  \\\"9d2e2d\\\")\");\n")
		jsx.WriteString(fmt.Sprintf("} else if ((%f>180) && (%f>180)) {\n", width, height))
		jsx.WriteString("      // 生成历史记录\n")
		jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超不透最大180cm\", \"maxCanvas(\\\"注意：已超不透最大180cm。\\\",  \\\"77bb11\\\")\");\n")
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

