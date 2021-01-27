package oldGui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/input"
	"github.com/yesilin/go-cutting/signal"
	"os"
)

// 初始化主菜单栏
func makeMainMenu() *fyne.MainMenu {
	fileMenu := fyne.NewMenu("文件")
	deitMenu := fyne.NewMenu("编辑")
	codeMenu := fyne.NewMenu("暗号",
		fyne.NewMenuItem("裁剪快捷键", func() { signal.ExecuteSignal1() }),
		fyne.NewMenuItem("重建新文档", func() { signal.ExecuteSignal2() }),
		fyne.NewMenuItem("深度清理PSD", func() { signal.ExecuteSignal3() }),
		fyne.NewMenuItem("快速清理PSD", func() { signal.ExecuteSignal6() }),
		fyne.NewMenuItem("自动加黑边", func() { signal.ExecuteSignal7() }),
		fyne.NewMenuItem("查切图历史", func() { signal.ExecuteSignal9() }),
		fyne.NewMenuItem("快速导出图片", func() { signal.ExecuteSignal98() }),
	)

	return fyne.NewMainMenu(
		fileMenu,
		deitMenu,
		codeMenu,
	)
}

// 初始化暗号列表
func makeCode() fyne.CanvasObject {
	inputCode := widget.NewEntry()
	inputCode.SetPlaceHolder("使用其他快捷暗号 ~    ")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "* :", Widget: inputCode},
		},
		//OnCancel: func() {
		//
		//},
		OnSubmit: func() {
			// 调用暗号
			input.RunCodeGui(inputCode.Text)
		},


		SubmitText: "提交",
		CancelText: "重置",
	}

	code := widget.NewGroup("暗号",


		//fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("裁剪快捷键", func() {
			signal.ExecuteSignal1()
		}),

		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("重建新文档", func() {
			signal.ExecuteSignal2()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("深度清理PSD", func() {
			signal.ExecuteSignal3()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("复制其他层", func() {
			signal.ExecuteSignal5()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("快速清理PSD", func() {
			signal.ExecuteSignal6()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("自动加黑边", func() {
			signal.ExecuteSignal7()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("查切图历史", func() {
			signal.ExecuteSignal9()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("全文档另存", func() {
			signal.ExecuteSignal11()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("全文档关存", func() {
			signal.ExecuteSignal12()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("快速导出图片", func() {
			signal.ExecuteSignal98()
		}),
		layout.NewSpacer(),
		form,


	)
	//code.Resize(fyne.NewSize(300, 800))
	return code
}

func Start() {
	// 没有说自启直接退出函数
	if !viper.GetBool("gui") {
		return
	}

	// 设置字体的环境变量
	os.Setenv("FYNE_FONT", "./config/苹方常规体.ttf")
	defer os.Unsetenv("FYNE_FONT")

	// New返回一个具有默认驱动程序且没有唯一ID的新应用程序实例
	app := app.New()

	// 获取主题设置
	getTheme := viper.GetString("theme")
	switch {
	case "darkTheme" == getTheme:
		// 主题默认设置成深色
		app.Settings().SetTheme(theme.DarkTheme())
	case "lightTheme" == getTheme:
		// 主题默认设置成白色
		app.Settings().SetTheme(theme.LightTheme())
	}

	// 为应用程序创建新窗口。第一个打开的窗口被认为是“主窗口”，当它关闭时应用程序将退出。
	w := app.NewWindow("GoCutting")
	w.SetIcon(theme.FyneLogo())

	// 设置主菜单栏
	w.SetMainMenu(makeMainMenu())
	// 修改窗口大小
	//w.Resize(fyne.NewSize(500, 1000))

	// 新建一个标签集合
	mainTabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon(`首页介绍`, theme.HomeIcon(), makeIndex(app)),
		//widget.NewTabItemWithIcon(`首页`,theme.HomeIcon(), welcomeScreen(app)),

		widget.NewTabItemWithIcon("快捷切图", theme.ContentCutIcon(), makeCut()),


		widget.NewTabItemWithIcon("快捷贴图", theme.ContentPasteIcon(), widget.NewVBox()),


		widget.NewTabItemWithIcon("快捷效果", theme.DocumentCreateIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("自动套图", theme.MailReplyIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("附加功能", theme.MailAttachmentIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("设置中心", theme.SettingsIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("帮助信息", theme.HelpIcon(), makeHelp(w)),
	)
	// 设置位置为左对齐
	mainTabs.SetTabLocation(widget.TabLocationLeading)

	// 设置此窗口的内容。更改布局为可调节大小的两列  fyne.NewContainerWithLayout(layout.NewGridLayout(2)
	w.SetContent(widget.NewHSplitContainer(
		// 主标签
		mainTabs,

		// 右侧暗号列表
		makeCode(),
	))

	// 显示并运行
	w.ShowAndRun()
}
