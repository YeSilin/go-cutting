package model

// 生成专属参考线的脚本
import (
	"github.com/sirupsen/logrus"
	"os"
	"text/template"
)

// FrameGuide6 生成折屏参考线
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
	f, err := os.OpenFile("resources/jsx/newDocument.jsx", os.O_APPEND, 0644)
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
