package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
)

type GuiView struct {
	app         fyne.App    //新应用程序实例
	mainWindow  fyne.Window //主窗口
	setupWindow fyne.Window // 设置窗口
}

//Logo 返回实现资源接口的logo图标
func logo() fyne.Resource {
	// 绑定资源到此变量
	var logo = &fyne.StaticResource{
		StaticName:    "logo8-32.png",
		StaticContent: tools.ReadFileEasy("resources/static/img/logo8-32.png"),
	}
	return logo
}

func NewGuiView() *GuiView {
	//New返回一个具有默认驱动程序且没有唯一ID的新应用程序实例
	application := app.New()
	//根据配置设置主题
	if viper.GetBool("darkTheme") {
		application.Settings().SetTheme(theme.DarkTheme())
	} else {
		application.Settings().SetTheme(theme.LightTheme())
	}

	//为应用程序创建新窗口。第一个打开的窗口被认为是“主窗口”，当它关闭时应用程序将退出。
	windows := application.NewWindow("GoCutting")
	//设置一个默认图标
	windows.SetIcon(logo())
	//修改主窗口大小
	windows.Resize(fyne.NewSize(250, 85))

	return &GuiView{
		app:        application,
		mainWindow: windows,
	}
}

//选择暗号，同时修改一个标签的文字
func (v *GuiView) choiceCode(tips *widget.Label) fyne.CanvasObject {
	//实例一个容器对象，此容器每行可以放两个画布对象
	return container.New(layout.NewGridLayout(2),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-1]", func() {
			//只有在Photoshop已经打开的情景下运行脚本
			if presenter.IsPhotoshopRun() {
				presenter.Command1()
				tips.SetText("You just pressed [-1]!~")
			} else {
				dialog.ShowInformation("Info", "Photoshop is not running!", v.mainWindow)
			}
		}),

		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-3]", func() {
			//只有在Photoshop已经打开的情景下运行脚本
			if presenter.IsPhotoshopRun() {
				presenter.Command3()
				tips.SetText("You just pressed [-3]!~")
			} else {
				dialog.ShowInformation("Info", "Photoshop is not running!", v.mainWindow)
			}
		}),
	)
}

//ShowAndRun 显示并运行
func (v *GuiView) ShowAndRun() {

	//实例一个标签对象，用于文本提示
	tips := widget.NewLabelWithStyle("Quick operation code!~", fyne.TextAlignCenter, fyne.TextStyle{})

	//对窗口进行设置内容，设置内容的时候先新建一个垂直对齐的盒子
	v.mainWindow.SetContent(container.NewVBox(
		//tips,
		v.choiceCode(tips),
		v.settings(),

		//v.switchTheme(),
	))

	// 显示并运行
	v.mainWindow.ShowAndRun()

}

//文件编辑帮助的主菜单
func mainMenu() *fyne.MainMenu {

	fileMenu := fyne.NewMenu("File")

	editMenu := fyne.NewMenu("Edit")
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {

		}))

	return fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fileMenu,
		editMenu,
		helpMenu,
	)
}
