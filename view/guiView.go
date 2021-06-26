package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/yesilin/go-cutting/signal"
	"github.com/yesilin/go-cutting/tools"
)


// 绑定资源到此变量
var logo = &fyne.StaticResource{
	StaticName:    "logo7-32.png",
	StaticContent: tools.ReadFileEasy("config/static/img/logo7-32.png"),
}

//Logo 返回实现资源接口的logo图标
func Logo() fyne.Resource {
	return logo
}





//Start 开始运行窗口
func Start(){
	// New返回一个具有默认驱动程序且没有唯一ID的新应用程序实例
	app := app.New()

	// 主题默认设置成深色
	app.Settings().SetTheme(theme.DarkTheme())

	// 为应用程序创建新窗口。第一个打开的窗口被认为是“主窗口”，当它关闭时应用程序将退出。
	windows := app.NewWindow("GoCutting")

	// 修改窗口大小
	windows.Resize(fyne.NewSize(250, 80))

	// 设置一个默认图标
	windows.SetIcon(Logo())

	//实例一个标签对象，用于文本提示
	tips := widget.NewLabelWithStyle("Quick operation code!~", fyne.TextAlignCenter, fyne.TextStyle{})

	//实例一个容器对象，此容器每行可以放两个画布对象
	code := container.New(layout.NewGridLayout(2),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-1]", func() {
			signal.ExecuteSignal1()
			tips.SetText("You just pressed [-1]!~")
		}),

		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-3]", func() {
			signal.ExecuteSignal3()
			tips.SetText("You just pressed [-3]!~")
		}),
	)






	//对窗口进行设置内容，设置内容的时候先新建一个垂直对齐的盒子
	windows.SetContent(container.NewVBox(
		tips,
		code,
	))


	// 显示并运行
	windows.ShowAndRun()

}
