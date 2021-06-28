package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
)

//进行切换主题
func (g *GuiView) switchTheme() *fyne.Container {
	//返回一个容器，布局是每行只能放一个元素
	return container.New(layout.NewGridLayout(1),

		//添加一个带图标的按钮
		widget.NewButtonWithIcon("SwitchTheme", theme.SettingsIcon(), func() {
			//如果是深色主题就切换成浅色
			if viper.GetBool("darkTheme") {
				g.app.Settings().SetTheme(theme.LightTheme())
				viper.Set("darkTheme", false)
			} else {
				g.app.Settings().SetTheme(theme.DarkTheme())
				viper.Set("darkTheme", true)
			}
			//保存最新配置
			viper.WriteConfig()
		}),
	)
}


//单独的设置窗口
func (g *GuiView) settings() *fyne.Container {
	return container.New(layout.NewGridLayout(1),
		//添加一个带图标的按钮
		widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
			w := g.app.NewWindow("Settings")
			//设置一个默认图标
			w.SetIcon(logo())
			w.Resize(fyne.NewSize(220, 40))

			//对窗口布置内容
			w.SetContent(container.NewVBox(
				g.switchTheme(),
				//这里放各种功能对象

			))

			w.Show()

		}),

	)
}
