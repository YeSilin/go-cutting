package model

// 这里放一些创建Photoshop文档的脚本
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

// NewDocument 生成用来新建ps文档jsx；wordLine 是否创建文字不要超过的参考线
func NewDocument(width, height float64, frameName string, exceeded bool) {
	jsx := `// 新建文档函数
function newDocument(width, height, docName){
	// 设置首选项新文档预设单位是厘米，PIXELS是像素
	app.preferences.rulerUnits = Units.CM;
	// 新文档的分辨率
	const resolution = 100;
	// 新文档的颜色模式
	const mode = NewDocumentMode.CMYK;
	// 新文档的默认背景填充颜色
	const initialFill = DocumentFill.WHITE;
	// 新文档的像素比率
	const pixelAspectRatio = 1;
	// 设置颜色位数为8位
	const bitsPerChannel = BitsPerChannelType.EIGHT;
	// 设置颜色配置文件为日本常规用途3
	const colorProfileName = "Japan Color 2011 Coated";
	// 将设置好的参数放在[add]方法里面
	app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio, bitsPerChannel, colorProfileName);
	// 切换锁定参考线
	app.runMenuItem(stringIDToTypeID('toggleLockGuides'));
}

// 新文档的宽度
const width = {{printf "%.2f" .Width}}; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 新文档的高度
const height = {{printf "%.2f" .Height}}; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 新文档的名称
const docName = "{{.FrameName}}";
// 执行新建文档
newDocument(width,height,docName)


{{/*只有类似座屏那种单张的才需要参考线*/}}
{{if .Exceeded}} // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 定义一个函数用于新建参考线
function addLine() {
    // 添加文字垂直参考线
    activeDocument.guides.add(Direction.VERTICAL, UnitValue("5cm"));
    activeDocument.guides.add(Direction.VERTICAL, UnitValue("{{printf "%.2f" (sub .Width 5)}}cm")); // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
    // 添加文字水平参考线
    activeDocument.guides.add(Direction.HORIZONTAL, UnitValue("5cm"));
    activeDocument.guides.add(Direction.HORIZONTAL, UnitValue("{{printf "%.2f" (sub .Height 5)}}cm")); // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
}

// 生成历史记录
app.activeDocument.suspendHistory("建议：字不要在此参考线外！", "addLine()");
{{end}}
`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width     float64
		Height    float64
		FrameName string // 新文档名
		Exceeded  bool   // 文字超出提醒的参考线
	}{width, height, frameName, exceeded}

	// 为模板自定义加法函数
	//add := func(left float64, right float64) float64 {
	//	return left + right
	//}

	// 为模板自定义减法函数
	sub := func(Minuend float64, Reduction float64) float64 {
		return Minuend - Reduction
	}

	// 采用链式操作在Parse解析之前调用 Funcs 添加自定义的sub函数
	// 这边有个地方值得注意，template.New()函数中参数名字要和ParseFiles（）
	// 函数的文件名要相同，要不然就会报错："" is an incomplete template
	tmpl, err := template.New("newDocument.gohtml").Funcs(template.FuncMap{"sub": sub}).Parse(jsx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("resources/jsx/newDocument.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		fmt.Println(err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil { // 如果有错误，打印错误，同时返回
		fmt.Println(err)
		return
	}

	// 关闭文件
	f.Close()
}

// IsMaxCanvasExceeded 是否超过最大画布尺寸
func IsMaxCanvasExceeded(width, height float64) bool {
	// 如果宽或高小于限定，那就直接返回
	if width < 150 || height < 150 {
		return false
	}

	var jsx = strings.Builder{}
	jsx.WriteString(`
// 创建一个文字提示层函数
function promptLayer(text, rgbValue){
    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。
    var newLayer = app.activeDocument.artLayers.add();

    // 把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。
    newLayer.kind = LayerKind.TEXT;

    // 把图层 newLayer 的文本内容类型变为”文本框“。
    newLayer.textItem.kind = TextType.PARAGRAPHTEXT;

    // 设置图层 newLayer 的文本框宽度与高度。
    newLayer.textItem.width = app.activeDocument.width*0.8;
    newLayer.textItem.height = app.activeDocument.width*0.1;

    // 设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素，例子 newLayer.textItem.position= [UnitValue("50px"), UnitValue("100px")]
    newLayer.textItem.position= [UnitValue(app.activeDocument.width*0.1), UnitValue((app.activeDocument.height*0.5)-(app.activeDocument.width*0.025))];

    // 设置 newLayer 的文本字体大小为“40 点”。
    newLayer.textItem.size = UnitValue(app.activeDocument.width*0.05);

    // 设置 newLayer 的文本内容。
    newLayer.textItem.contents= text;

    // 设置 newLayer 的文本框对齐方式为居中对齐。
    newLayer.textItem.justification = Justification.CENTER;

    // 创建一个色彩变量 c   ，颜色为 #77bb11。
    var c = new SolidColor();
    c.rgb.hexValue = rgbValue;

    // 设置 newLayer 的文本颜色为 c。
    newLayer.textItem.color = c;
}
`)

	// 剩下的都大于150了，于是先从最过分的尺寸判断
	if width > 200 && height > 200 {
		jsx.WriteString(`
app.activeDocument.suspendHistory("注意：已超半透最大200cm", "promptLayer(\"注意：已超半透最大200cm。\",  \"f1362c\")");
`)
	} else {
		jsx.WriteString(`
app.activeDocument.suspendHistory("注意：已超不透最大150cm", "promptLayer(\"注意：已超不透最大150cm。\",  \"77bb11\")");
`)
	}

	// 追加写入
	tools.WriteFile("resources/jsx/newDocument.jsx", jsx.String())
	return true
}

// NewDocumentForMap 生成用来新建ps文档3d作图的jsx
func NewDocumentForMap(width, height int, frameName string) {
	const script = `// 新建文档函数
function newDocument(width, height, docName){
	// 设置首选项新文档预设单位是厘米，PIXELS是像素
	app.preferences.rulerUnits = Units.PIXELS;
	// 新文档的分辨率
	const resolution = 72;
	// 新文档的颜色模式
	const mode = NewDocumentMode.RGB;
	// 新文档的默认背景填充颜色
	const initialFill = DocumentFill.WHITE;
	// 新文档的像素比率
	const pixelAspectRatio = 1;
	// 设置颜色位数为8位
	const bitsPerChannel = BitsPerChannelType.EIGHT;
	// 设置颜色配置文件为日本常规用途3
	const colorProfileName = "sRGB IEC61966-2.1";
	// 将设置好的参数放在[add]方法里面
	app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio, bitsPerChannel, colorProfileName);
}

// 新文档的宽度
const width = {{printf "%d" .Width}}; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 新文档的高度
const height = {{printf "%d" .Height}}; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 新文档的名称
const docName = "{{.FrameName}}"; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！
// 执行新建文档
newDocument(width,height,docName);`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width     int
		Height    int
		FrameName string // 新文档名
	}{width, height, frameName}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("resources/jsx/newDocument.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}
}

// NewDocumentForTemp 新建临时效果图文档
func NewDocumentForTemp() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now := time.Now().Format("0102150405")

	var jsx = strings.Builder{}

	jsx.WriteString("app.preferences.rulerUnits = Units.PIXELS;\n")
	jsx.WriteString("var width = 5000;\n")
	jsx.WriteString("var height = 3500;\n")
	jsx.WriteString("var resolution = 72;\n")
	jsx.WriteString(fmt.Sprintf("var docName = \"%s_临时效果图\";   // 这里传golang 时间参数哦！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！\n", now))
	jsx.WriteString("var mode = NewDocumentMode.RGB;\n")
	jsx.WriteString("var initialFill = DocumentFill.WHITE;\n")
	jsx.WriteString("var pixelAspectRatio = 1;\n")
	jsx.WriteString("app.documents.add(width, height, resolution, docName, mode, initialFill, pixelAspectRatio);\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 为主文档定义变量\n")
	jsx.WriteString("var masterDocument = app.activeDocument;\n")
	jsx.WriteString("\n")
	jsx.WriteString(" // 获取当前脚本所在路径\n")
	jsx.WriteString("var scriptPath = (new File($.fileName)).parent;\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 为指定路径文件定义变量\n")
	jsx.WriteString("var fileRef = new File(scriptPath+\"/../PSD/ColorCard.psd\");\n")
	jsx.WriteString("if(fileRef.exists){    // 如果图像存在\n")
	jsx.WriteString("    app.open(fileRef);   // 打开文档\n")
	jsx.WriteString("    selectAllLayers();   // 选择全部图层\n")
	jsx.WriteString("    app.activeDocument.activeLayer.duplicate (masterDocument); // 复制到主文档\n")
	jsx.WriteString("    app.activeDocument.close(SaveOptions.DONOTSAVECHANGES); // 关闭临时文档而不保存更改\n")
	jsx.WriteString("    //app.activeDocument = app.documents.getByName(\"临时效果图\");    //  切换窗口到这里\n")
	jsx.WriteString("   app.activeDocument.activeLayer = app.activeDocument.layers[1];  // 选择背景图层\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("// 选择全部图层，大神代码\n")
	jsx.WriteString("function selectAllLayers() {\n")
	jsx.WriteString("    var desc29 = new ActionDescriptor();\n")
	jsx.WriteString("    var ref23 = new ActionReference();\n")
	jsx.WriteString("    ref23.putEnumerated(charIDToTypeID('Lyr '), charIDToTypeID('Ordn'), charIDToTypeID('Trgt'));\n")
	jsx.WriteString("    desc29.putReference(charIDToTypeID('null'), ref23);\n")
	jsx.WriteString("    executeAction(stringIDToTypeID('selectAllLayers'), desc29, DialogModes.NO);\n")
	jsx.WriteString("}\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 71.0 更新 先强制生成的文本写覆盖入目标文件
	tools.CreateFile("resources/jsx/newTempDocument.jsx", jsxStr)

	// 创建一个协程使用cmd来运行脚本
	dataPath := "resources/jsx/newTempDocument.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}
