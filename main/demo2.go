package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/globa"
	"github.com/yesilin/go-cutting/tools"
	"net/url"
	"os/exec"
	"strconv"
	"time"
)


func welcomeScreen2(a fyne.App) fyne.CanvasObject {
	//logo := canvas.NewImageFromResource(theme.FyneLogo())

	logo := canvas.NewImageFromResource(theme.FyneLogo())

	logo.SetMinSize(fyne.NewSize(128, 128))

	link, err := url.Parse("https://shang.qq.com/wpa/qunwpa?idkey=3f7d1f234ceac52d97dc324874193608e2b7b1690aa46bd2b5531cf854c9ef02")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the GoCutting app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("Feedback", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}



func makeButtonTab() fyne.Widget {
	disabled := widget.NewButton("Disabled", func() {})
	disabled.Disable()

	return widget.NewVBox(
		widget.NewLabel("Text label"),
		widget.NewButton("Text button", func() { fmt.Println("tapped text button") }),
		widget.NewButtonWithIcon("With icon", theme.ConfirmIcon(), func() { fmt.Println("tapped icon button") }),
		disabled,
	)
}

func makeInputTab() fyne.Widget {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")
	entryReadOnly := widget.NewEntry()
	entryReadOnly.SetPlaceHolder("Entry (read only)")
	entryReadOnly.ReadOnly = true

	disabledCheck := widget.NewCheck("Disabled check", func(bool) {})
	disabledCheck.Disable()
	radio := widget.NewRadio([]string{"Radio Item 1", "Radio Item 2"}, func(s string) { fmt.Println("selected", s) })
	radio.Horizontal = true
	disabledRadio := widget.NewRadio([]string{"Disabled radio"}, func(string) {})
	disabledRadio.Disable()

	return widget.NewVBox(
		entry,
		entryReadOnly,
		widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"}, func(s string) { fmt.Println("selected", s) }),
		widget.NewCheck("Check", func(on bool) { fmt.Println("checked", on) }),
		disabledCheck,
		radio,
		disabledRadio,
	)
}

func makeProgressTab() fyne.Widget {
	progress := widget.NewProgressBar()
	infProgress := widget.NewProgressBarInfinite()

	go func() {
		num := 0.0
		for num < 1.0 {
			time.Sleep(100 * time.Millisecond)
			progress.SetValue(num)
			num += 0.01
		}

		progress.SetValue(1)
	}()

	return widget.NewVBox(
		widget.NewLabel("Percent"), progress,
		widget.NewLabel("Infinite"), infProgress)
}

func makeFormTab() fyne.Widget {
	widthInput := widget.NewEntry()
	widthInput.SetPlaceHolder("Please enter widthInput")
	heightInput := widget.NewEntry()
	heightInput.SetPlaceHolder("Please enter heightInput")



	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			// 定义一个预留尺寸
			var reserve = globa.NowSetting.Reserve

			// 强制类型转换成浮点数
			width, _ := strconv.ParseFloat(widthInput.Text, 64)
			height, _ := strconv.ParseFloat(heightInput.Text, 64)

			width = width - 10 + reserve
			height = height - 10 + reserve

			//获取当前时间，进行格式化 2006-01-02 15:04:05
			now := time.Now().Format("0102150405")
			// 为当前框架指定名字
			frameName := fmt.Sprintf("%s_普通座屏_%.0fx%.0f", now, width, height)

			generate.NewDocumentJS(width, height, frameName, true) // 创建ps文档
			go generate.Tailor0(frameName)                         // 生成暗号【-1】可以用的另存脚本
			tools.MaxCanvas(width, height)   // 最大画布判断

			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/jsx/NewDocumentJS.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		},
	}
	form.Append("Width", widthInput)
	form.Append("Height", heightInput)
	return form
}

func makeScrollTab() fyne.Widget {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(320, 320))
	list := widget.NewVBox()
	for i := 1; i <= 20; i++ {
		index := i // capture
		list.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			fmt.Println("Tapped", index)
		}))
	}

	scroll := widget.NewScrollContainer(list)
	scroll.Resize(fyne.NewSize(200, 200))

	return scroll
}


// WidgetScreen shows a panel containing widget demos
func WidgetScreen() fyne.CanvasObject {
	//toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
	//	widget.NewToolbarSeparator(),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.ContentCutIcon(), func() { fmt.Println("Cut") }),
	//	widget.NewToolbarAction(theme.ContentCopyIcon(), func() { fmt.Println("Copy") }),
	//	widget.NewToolbarAction(theme.ContentPasteIcon(), func() { fmt.Println("Paste") }),
	//)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		widget.NewTabContainer(
			widget.NewTabItem("Buttons", makeButtonTab()),
			widget.NewTabItem("Input", makeInputTab()),
			widget.NewTabItem("Progress", makeProgressTab()),
			widget.NewTabItem("Form", makeFormTab()),
			widget.NewTabItem("Scroll", makeScrollTab()),
		),
	)
}


func main02() {
	a := app.New()

	w := a.NewWindow("GoCutting")
	// 添加编辑栏
	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
		// 我们的第一个菜单将附加一个退出项
	),fyne.NewMenu("Edit",
			fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
			fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
			fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	)))


	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcomeScreen2(a)),
		widget.NewTabItemWithIcon("Widgets", theme.ContentCopyIcon(), WidgetScreen()))
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)


	//w.SetContent(widget.NewVBox(
	//
	//	widget.NewLabel("Hello Fyne!"),
	//	widget.NewButton("Quit", func() {
	//		a.Quit()
	//	}),
	//))

	w.ShowAndRun()
}