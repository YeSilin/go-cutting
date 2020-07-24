package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"os"
)

// 初始化主菜单栏
func makeMainMenu() *fyne.MainMenu {
	fileMenu := fyne.NewMenu("文件")
	deitMenu := fyne.NewMenu("编辑")
	codeMenu := fyne.NewMenu("暗号",
		fyne.NewMenuItem("裁剪快捷键", func() { model.StartCode1() }),
		fyne.NewMenuItem("重建新文档", func() { model.StartCode2() }),
		fyne.NewMenuItem("深度清理PSD", func() { model.StartCode3() }),
		fyne.NewMenuItem("快速清理PSD", func() { model.StartCode6() }),
		fyne.NewMenuItem("自动加黑边", func() { model.StartCode7() }),
		fyne.NewMenuItem("查切图历史", func() { model.StartCode9() }),
		fyne.NewMenuItem("快速导出图片", func() { model.StartCode98() }),
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
	inputCode.SetPlaceHolder("使用其他暗号")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "* :", Widget: inputCode},
		},
		//OnCancel: func() {
		//
		//},
		OnSubmit: func() {
			// 调用暗号
			model.RunCodeGui(inputCode.Text)
		},


		SubmitText: "提交",
		CancelText: "重置",
	}

	code := widget.NewGroup("暗号",

		//fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("裁剪快捷键", func() {
			model.StartCode1()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("重建新文档", func() {
			model.StartCode2()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton(" 深度清理PSD ", func() {
			model.StartCode3()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton(" 快速清理PSD ", func() {
			model.StartCode6()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("自动加黑边", func() {
			model.StartCode7()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("查切图历史", func() {
			model.StartCode9()
		}),
		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("快速导出图片", func() {
			model.StartCode98()
		}),
		layout.NewSpacer(),
		form,


	)
	//code.Resize(fyne.NewSize(300, 800))
	return code
}

// 初始化首页
func makeIndex(a fyne.App) fyne.CanvasObject {

	info := widget.NewGroup("软件介绍",
		widget.NewLabelWithStyle(" 一个简单快速的切图软件 ", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("目的是为了帮助大家", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("快速切图", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("~", fyne.TextAlignCenter, fyne.TextStyle{}),


	)

	theme := widget.NewGroup("主题设置",
		fyne.NewContainerWithLayout(layout.NewGridLayout(1),

			widget.NewButton("   深色   ", func() {
				a.Settings().SetTheme(theme.DarkTheme())
			}),
			widget.NewButton("   浅色   ", func() {
				a.Settings().SetTheme(theme.LightTheme())
			}),


		),
	)

	return widget.NewVBox(

		info,
		theme,

	)

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

	// 主题默认设置成白色
	app.Settings().SetTheme(theme.LightTheme())

	// 为应用程序创建新窗口。第一个打开的窗口被认为是“主窗口”，当它关闭时应用程序将退出。
	w := app.NewWindow("GoCutting")
	w.SetIcon(theme.FyneLogo())

	// 设置主菜单栏
	w.SetMainMenu(makeMainMenu())
	// 修改窗口大小
	//w.Resize(fyne.NewSize(500, 70))

	// 新建一个标签集合
	mainTabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon(`首页`, theme.HomeIcon(), makeIndex(app)),
		//widget.NewTabItemWithIcon(`首页`,theme.HomeIcon(), welcomeScreen(app)),

		widget.NewTabItemWithIcon("切图", theme.ContentCutIcon(), makeCut()),


		widget.NewTabItemWithIcon("贴图", theme.ContentPasteIcon(), widget.NewVBox()),


		widget.NewTabItemWithIcon("效果", theme.DocumentCreateIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("套图", theme.MailReplyIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("附加", theme.MailAttachmentIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("设置", theme.SettingsIcon(), widget.NewVBox()),
		widget.NewTabItemWithIcon("帮助", theme.HelpIcon(), widget.NewVBox()),
	)
	// 设置位置为左对齐
	mainTabs.SetTabLocation(widget.TabLocationLeading)

	//// 创建一个新的折叠构件
	//ac :=widget.NewAccordionContainer(
	//	widget.NewAccordionItem("常规座屏", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("左右镂空", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("左右画布", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("上下镂空", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("顶天立地", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("各种折屏", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("多个座屏", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("卷帘座屏", widget.NewLabel("Two")),
	//	widget.NewAccordionItem("不扣补切", widget.NewLabel("Two")),
	//
	//	)

	// 设置此窗口的内容。更改布局为两列 fyne.NewContainerWithLayout(layout.NewGridLayout(2)
	w.SetContent(widget.NewHSplitContainer(
		// 主标签
		mainTabs,

		// 右侧暗号列表
		makeCode(),

		//tabs,
		//ac,

	))

	// 显示并运行
	w.ShowAndRun()
}
