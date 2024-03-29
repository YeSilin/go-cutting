package model

// 生成专属参考线的脚本
import (
	"github.com/sirupsen/logrus"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"text/template"
)

// FrameGuide3 生成左右画布专属参考线
func FrameGuide3(width, hollow float64) {
	const script = `
// ----------------------------------------------------------------------------------------------------------------------------
// 定义一个函数用于新建参考线
function addLine() {
    activeDocument.guides.add(Direction.VERTICAL, UnitValue("{{printf "%.2f" .Line1}}cm"));
    activeDocument.guides.add(Direction.VERTICAL, UnitValue("{{printf "%.2f" .Line2}}cm"));
}

// 生成历史记录
app.activeDocument.suspendHistory("左右画布参考线！", "addLine()");`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Line1 float64
		Line2 float64
	}{hollow, hollow + width}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// FrameGuide4to2 生成上下画布专属参考线
func FrameGuide4to2(upHeight, middleHeight float64) {
	const script = `
// ----------------------------------------------------------------------------------------------------------------------------
// 定义一个函数用于新建参考线
function addLine() {
    activeDocument.guides.add(Direction.HORIZONTAL, UnitValue("{{printf "%.2f" .Line1}}cm"));
    activeDocument.guides.add(Direction.HORIZONTAL, UnitValue("{{printf "%.2f" .Line2}}cm"));
}

// 生成历史记录
app.activeDocument.suspendHistory("上下画布参考线！", "addLine()");`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Line1 float64
		Line2 float64
	}{upHeight, upHeight + middleHeight}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// FrameGuide6 生成折屏专属参考线
func FrameGuide6(width, number float64) {
	const script = `
// 新建折屏参考线函数
function addFoldingScreenGuide() {
    const width = {{.Width}};  // 这里传golang参数！！！！！！！！！！！！！！！！！！！！！！！！！！
    const number = {{.Number}};  // 这里传golang参数！！！！！！！！！！！！！！！！！！！！！！！！！！

    const line = function (width, number) {
        for (var i = 1; i < number; i++) {
            // 添加垂直参考线
            app.activeDocument.guides.add(Direction.VERTICAL, UnitValue(i * width + "cm"));
        }
    }
    // 调用并生成历史记录
    app.activeDocument.suspendHistory("新建折屏参考线", "line(width, number)");
}

// 调用新建折屏参考线
addFoldingScreenGuide();
`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		Width  float64
		Number int
	}{width, int(number)}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// FrameGuide7 生成多座屏专属参考线
func FrameGuide7(widthSlice, heightSlice []float64, maxHeight, minHeight float64) {
	const script = `// ----------------------------------------------------------------------------------------------------------------------------
// 定义一个函数用来新建透明图层
function mask(layerName) {
    // 新建一个图层
    app.activeDocument.artLayers.add().name = layerName;

    // 不透明度改30%
    app.activeDocument.artLayers[0].opacity = 30;
}

// 定义一个函数用来设置区域填充颜色
function makeSelection(x, y, sw, sh) {
    // 设置选区
    app.activeDocument.selection.select([[x, y], [x, y + sh], [x + sw, y + sh], [x + sw, y]]);

    // 生成一个随机色
    var color = new SolidColor();
    color.rgb.red = Math.random() * 255;
    color.rgb.green = Math.random() * 255;
    color.rgb.blue = Math.random() * 255;

    // 填充选区
    app.activeDocument.selection.fill(color);
    // 取消选区
    app.activeDocument.selection.deselect();
}


// 将所有遮罩图层的操作归纳到此函数
function addMaskLayer(layerName, widthArray, heightArray, equal) {
    // 如果相等就不制作遮罩图层
    if (equal) {
        return
    }

    // 先新建一个用于填充的空图层
    mask(layerName);

    // 先提前定义一个左上角 x 轴位置
    var x = 0
    for (var i = 0; i < heightArray.length; i++) {
        // 遍历到一样高的座屏就不绘制
        if (heightArray[i] == maxHeight) {

            // 就算不绘制也要更新 x 轴位置
            x += widthArray[i] * 39.37
            continue
        }
        //alert (heightArray[i])
        // 开始填充遮罩，这里乘39.37是因为要把厘米转成像素
        makeSelection(x, 0, widthArray[i] * 39.37, (maxHeight - heightArray[i]) * 39.37)
        x += widthArray[i] * 39.37
    }
}


// 定义一个函数用于新建宽度参考线
function addWidthLine(widthArray) {
    var singleFan = 0
    for (var i = 0; i < widthArray.length - 1; i++) {
        // 定义单扇宽度
        singleFan += widthArray[i];
        activeDocument.guides.add(Direction.VERTICAL, UnitValue(singleFan + "cm"));
    }
}


// 定义一个函数用于新建高度参考线
function addHeightLine(heightArray, maxHeight, equal) {
    // 如果相等就不制作高度参考线
    if (equal) {
        return
    }
    for (var i = 0; i < heightArray.length; i++) {
        // 遍历到一样高的座屏就不绘制
        if (maxHeight == heightArray[i]) {
            continue
        }

        // 定义单扇多出的度
        var singleFan = maxHeight - heightArray[i];

        activeDocument.guides.add(Direction.HORIZONTAL, UnitValue(singleFan + "cm"));
    }
}


// 定义框架名字
var frameName = "{{.ScreenName}}座屏" // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

// 定义图层名字
var layerName = frameName + "遮罩图层"

// 用数组定义多座屏的宽和高
var widthArray = {{.WidthSliceJS}};   // 这里传golang排版好的字符串哦！！！！！！！！！！！
var heightArray = {{.HeightSliceJS}};  // 这里传golang排版好的字符串哦！！！！！！！！！！！

// 定义最高的座屏和最矮的座屏是否相等
var equal = {{.Equal}};  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

// 定义最高座屏的高度
var maxHeight = {{.MaxHeight}}; // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！


// 生成遮罩层并写入历史
app.activeDocument.suspendHistory(layerName, "addMaskLayer(layerName,widthArray,heightArray,equal)");

// 生成宽度参考线并写入历史
app.activeDocument.suspendHistory(frameName + "宽度参考线", "addWidthLine(widthArray)");

// 生成高度参考线并写入历史
app.activeDocument.suspendHistory(frameName + "高度参考线", "addHeightLine(heightArray,maxHeight,equal)");

// 当前文档只有1个背景图层时不激活其他图层
if (app.activeDocument.layers.length != 1){
    // 设置激活的图层
    app.activeDocument.activeLayer = app.activeDocument.layers[1]
}`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		WidthSliceJS  string
		HeightSliceJS string
		MaxHeight     float64 // 最大的高
		ScreenName    string  // 是几座屏
		Equal         bool    //最高和最矮座屏是否相等
	}{tools.Float64SliceToJsArray(widthSlice), tools.Float64SliceToJsArray(heightSlice), maxHeight, tools.Transfer(len(widthSlice)), maxHeight == minHeight}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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

// FrameGuide9to2 生成补切圆形专属参考线圆形辅助
func FrameGuide9to2() {
	const script = `
/* accepts parameters
 * h  Object = {h:x, s:y, v:z}
 * OR 
 * h, s, v
*/
function HSVtoRGB(h, s, v) {
    var r, g, b, i, f, p, q, t;
    if (arguments.length === 1) {
        s = h.s, v = h.v, h = h.h;
    }
    i = Math.floor(h * 6);
    f = h * 6 - i;
    p = v * (1 - s);
    q = v * (1 - f * s);
    t = v * (1 - (1 - f) * s);
    switch (i % 6) {
        case 0: r = v, g = t, b = p; break;
        case 1: r = q, g = v, b = p; break;
        case 2: r = p, g = v, b = t; break;
        case 3: r = p, g = q, b = v; break;
        case 4: r = t, g = p, b = v; break;
        case 5: r = v, g = p, b = q; break;
    }
    return {
        r: Math.round(r * 255),
        g: Math.round(g * 255),
        b: Math.round(b * 255)
    };
}


// 创建椭圆形状
function createEllipse(theBounds, theR, theG, theB) {
    var idpixelsUnit = stringIDToTypeID("pixelsUnit");
    var idmake = stringIDToTypeID("make");
    var desc15 = new ActionDescriptor();
    var idnull = stringIDToTypeID("null");
    var ref4 = new ActionReference();
    var idcontentLayer = stringIDToTypeID("contentLayer");
    ref4.putClass(idcontentLayer);
    desc15.putReference(idnull, ref4);
    var idusing = stringIDToTypeID("using");
    var desc16 = new ActionDescriptor();
    var idtype = stringIDToTypeID("type");
    var desc17 = new ActionDescriptor();
    var idcolor = stringIDToTypeID("color");
    var desc18 = new ActionDescriptor();
    var idred = stringIDToTypeID("red");
    desc18.putDouble(idred, theR);
    var idgrain = stringIDToTypeID("grain");
    desc18.putDouble(idgrain, theG);
    var idblue = stringIDToTypeID("blue");
    desc18.putDouble(idblue, theB);
    var idRGBColor = stringIDToTypeID("RGBColor");
    desc17.putObject(idcolor, idRGBColor, desc18);
    var idsolidColorLayer = stringIDToTypeID("solidColorLayer");
    desc16.putObject(idtype, idsolidColorLayer, desc17);
    var idshape = stringIDToTypeID("shape");
    var desc19 = new ActionDescriptor();
    var idunitValueQuadVersion = stringIDToTypeID("unitValueQuadVersion");
    desc19.putInteger(idunitValueQuadVersion, 1);
    var idtop = stringIDToTypeID("top");
    desc19.putUnitDouble(idtop, idpixelsUnit, theBounds[1]);
    var idleft = stringIDToTypeID("left");
    desc19.putUnitDouble(idleft, idpixelsUnit, theBounds[0]);
    var idbottom = stringIDToTypeID("bottom");
    desc19.putUnitDouble(idbottom, idpixelsUnit, theBounds[3]);
    var idright = stringIDToTypeID("right");
    desc19.putUnitDouble(idright, idpixelsUnit, theBounds[2]);
    var idellipse = stringIDToTypeID("ellipse");
    desc16.putObject(idshape, idellipse, desc19);
    var idcontentLayer = stringIDToTypeID("contentLayer");
    desc15.putObject(idusing, idcontentLayer, desc16);
    var idlayerID = stringIDToTypeID("layerID");
    desc15.putInteger(idlayerID, 3);
    executeAction(idmake, desc15, DialogModes.NO);
};


// 创建圆形参考
function circularReference() {
    var originalRulerUnits = app.preferences.rulerUnits;
    app.preferences.rulerUnits = Units.PIXELS;
	// HSV 转 RGB
    var rgbColor = HSVtoRGB(Math.random(), 20 / 100, 85 / 100);
    // 创建一个随机颜色的椭圆
    createEllipse([0, 0, activeDocument.width, activeDocument.width], rgbColor.r, rgbColor.g, rgbColor.b);
    app.preferences.rulerUnits = originalRulerUnits;
}

// 生成历史记录
app.activeDocument.suspendHistory("圆形参考", "circularReference()");
`

	// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
	info := struct {
		diameter float64
	}{}

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 打开要追加数据的文件
	f, err := os.OpenFile("data/jsx/newDocument.jsx", os.O_APPEND, 0644)
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
