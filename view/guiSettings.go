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
func (v *GuiView) switchTheme() fyne.CanvasObject {
	//返回一个容器，布局是每行只能放一个画布对象
	return container.New(layout.NewGridLayout(1),

		//添加一个带图标的按钮
		widget.NewButtonWithIcon("SwitchTheme", theme.ViewRefreshIcon(), func() {
			//如果是深色主题就切换成浅色
			if viper.GetBool("darkTheme") {
				v.app.Settings().SetTheme(theme.LightTheme())
				viper.Set("darkTheme", false)
			} else {
				v.app.Settings().SetTheme(theme.DarkTheme())
				viper.Set("darkTheme", true)
			}
			//保存最新配置
			viper.WriteConfig()
		}),
	)
}



//设置自动描边颜色
func (v *GuiView) strokeColor() fyne.CanvasObject {
	//返回一个容器，布局是每行只能放一个画布对象
	return container.New(layout.NewGridLayout(1),
		//添加一个带图标的按钮
		widget.NewButtonWithIcon("StrokeColor", theme.SettingsIcon(), func() {

		}),

	)
}

//单独的设置窗口
func (v *GuiView) settings() fyne.CanvasObject {
	return container.New(layout.NewGridLayout(1),
		//添加一个带图标的按钮
		widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
			//新建一个设置窗口并且赋值给结构体
			v.setupWindow = v.app.NewWindow("Settings")
			//设置一个默认图标
			v.setupWindow.SetIcon(logo())
			v.setupWindow.Resize(fyne.NewSize(220, 40))

			//对窗口布置内容
			v.setupWindow.SetContent(container.NewVBox(
				v.switchTheme(),
				//这里放各种功能对象
				v.strokeColor(),
			))

			v.setupWindow.Show()

		}),

	)
}
