package guimini

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"io"
	"os"
)

// 绑定资源到此变量
var logo = &fyne.StaticResource{
	StaticName:    "logo7-256.png",
	StaticContent: toStaticContent("config/static/img/logo7-256.png"),
}

// 打开文件转静态资源 "config/static/img/logo7-256.png"
func toStaticContent(path string) (result []byte) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer f.Close()

	// 读取内容保存到这个切片
	var content []byte
	buf := make([]byte, 4096)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("f.Read err:", err)
			return
		}
		if err == io.EOF {
			break
		}
		// 合并切片
		content = append(content, buf[:n]...)
	}
	return content
}

// 返回框架专属资源
func Logo() fyne.Resource {
	return logo
}


func Start() {
	// 没有说自启直接退出函数
	if !viper.GetBool("gui") {
		return
	}

	// 设置字体的环境变量
	//os.Setenv("FYNE_FONT", "./config/苹方常规体.ttf")
	//defer os.Unsetenv("FYNE_FONT")

	// New返回一个具有默认驱动程序且没有唯一ID的新应用程序实例
	app := app.New()

	// 主题默认设置成深色
	app.Settings().SetTheme(theme.DarkTheme())

	// 为应用程序创建新窗口。第一个打开的窗口被认为是“主窗口”，当它关闭时应用程序将退出。
	w := app.NewWindow("GoCutting")

	// 设置一个默认图标
	w.SetIcon(Logo() )


	// 修改窗口大小
	w.Resize(fyne.NewSize(250, 80))

	//hello := widget.NewLabel("Fast operation code!~")
	hello := widget.NewLabelWithStyle("Quick operation code!~", fyne.TextAlignCenter, fyne.TextStyle{})
	//info := widget.NewLabelWithStyle("Mini Edition", fyne.TextAlignCenter,fyne.TextStyle{})

	code := fyne.NewContainerWithLayout(layout.NewGridLayout(2),

		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-1]", func() {
			model.StartCode1()
		}),

		// 新建一个按钮，点击后执行匿名函数
		widget.NewButton("[-3]", func() {
			model.StartCode3()
		}),

	)

	// 设置画布内容，先新建一个垂直对齐的盒子
	w.SetContent(widget.NewVBox(
		hello,

		code,

	))

	// 显示并运行
	w.ShowAndRun()

}
