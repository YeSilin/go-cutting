package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/spf13/viper"
)

// 初始化首页
func makeIndex(a fyne.App) fyne.CanvasObject {

	info := widget.NewGroup("软件介绍",
		widget.NewLabelWithStyle(" 一个简单快速的切图软件 ", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("目的是为了帮助大家", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("快速切图", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("~", fyne.TextAlignCenter, fyne.TextStyle{}),


	)

	start := widget.NewGroup("自启设置",
		widget.NewButton("启用", func() {
			viper.Set("gui", true)
			// 保存最新配置
			go viper.WriteConfig()
		}),
		widget.NewButton("关闭", func() {
			viper.Set("gui", false)
			// 保存最新配置
			go viper.WriteConfig()
		}),
	)

	theme := widget.NewGroup("主题设置",
		fyne.NewContainerWithLayout(layout.NewGridLayout(1),

			widget.NewButton("   深色   ", func() {
				a.Settings().SetTheme(theme.DarkTheme())
				viper.Set("theme", "darkTheme")
				// 保存最新配置
				go viper.WriteConfig()
			}),
			widget.NewButton("   浅色   ", func() {

				a.Settings().SetTheme(theme.LightTheme())
				viper.Set("theme", "lightTheme")
				// 保存最新配置
				go viper.WriteConfig()
			}),
		),
	)

	return widget.NewVBox(

		info,
		start,
		layout.NewSpacer(),
		theme,
		layout.NewSpacer(),
	)

}
