package oldGui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func makeHelpCode(w fyne.Window) fyne.CanvasObject {
	//layout.NewSpacer(),

	return fyne.NewContainerWithLayout(layout.NewGridLayout(5),
		widget.NewButton("[-]", func() {
			dialog.ShowInformation("返回上一级", "这是一个全局通用的返回功能，不管在任何界面都将返回上一级菜单。", w)
		}),

		widget.NewButton("[--]", func() {
			dialog.ShowInformation("返回上一次输入", "例如镂空大小输错，返回重新输入镂空大小。", w)
		}),
		widget.NewButton("[-1]", func() {
			dialog.ShowInformation("裁剪快捷键 ", "优化另存快速裁剪，是主要优化的一个功能。", w)
		}),

		widget.NewButton("[-2]", func() {
			dialog.ShowInformation("重建新文档", "只能重建前一次使用本软件所创建的ps文档。", w)
		}),
		widget.NewButton("[-3]", func() {
			dialog.ShowInformation("深度清理PSD", "顾名思义就是清除垃圾并且不损失清晰度。", w)
		}),
		widget.NewButton("[-4]", func() {
			dialog.ShowInformation("快捷文件夹", "快捷文件夹功能，此功能创建的文件夹在桌面可以看到。", w)
		}),

		widget.NewButton("[-5]", func() {
			dialog.ShowInformation("复制其他层", "复制其他文档的全部图层并关闭其他文档，关闭时不保存修改。", w)
		}),
		widget.NewButton("[-6]", func() {
			dialog.ShowInformation("快速清理PSD", "仅清理当前文档的普通图层与图层组，不清理智能对象。", w)
		}),
		widget.NewButton("[-7]", func() {
			dialog.ShowInformation("自动加黑边", "自动加黑边功能，现在不管是打开任何ps文档都支持加黑边了。", w)
		}),
		widget.NewButton("[-8]", func() {
			dialog.ShowInformation("清屏快捷键", "在黑框命令行面板下的快捷清屏命令，不常用。", w)
		}),
		widget.NewButton("[-9]", func() {
			dialog.ShowInformation("到切图历史", "查询当天的切图历史记录，若没有记录会打开文件夹。", w)
		}),
		widget.NewButton("[-10]", func() {
			dialog.ShowInformation("单文档另存", "另存为JPG前进行一次元数据清除，并自动选择最佳参数。", w)
		}),
		widget.NewButton("[-11]", func() {
			dialog.ShowInformation("全文档另存", "另存全部打开的文档为JPG，文件名带白底图三字会多保存一张PNG。", w)
		}),
		widget.NewButton("[-12]", func() {
			dialog.ShowInformation("全文档关存", "保存并关闭全部打开的文档，请勿在有打开原图的情况下使用。", w)
		}),
		widget.NewButton("[-97]", func() {
			dialog.ShowInformation("替换详情页", "自动替换文档中DP或DE开头的智能对象，其中DE代表细节，\n修改文件名前缀为DE即可。", w)
		}),
		widget.NewButton("[-98]", func() {
			dialog.ShowInformation("导出详情页", "快速导出为web格式图片到自动套图文件夹的主图中。", w)
		}),
		widget.NewButton("[-99]", func() {
			dialog.ShowInformation("激活win10系统", "激活win10系统，这是一个集成大神之作的功能，只能激活win10系统。", w)
		}),
		widget.NewButton("[ ]", func() {
			dialog.ShowInformation("功能未开发", "功能未开发，占个位置。", w)
		}),
		widget.NewButton("[ ]", func() {
			dialog.ShowInformation("功能未开发", "功能未开发，占个位置。", w)
		}),
		widget.NewButton("[ ]", func() {
			dialog.ShowInformation("功能未开发", "功能未开发，占个位置。", w)
		}),
	)
}

// 初始化帮助
func makeHelp(w fyne.Window) fyne.CanvasObject {
	const helpText = `1. 切图的单位是厘米，分辨率是100像素/英寸，
   颜色模式是 CMYK；

2. 颜色配置文件是 工作中的CMYK:
   Japan color 2011 Coated；

3. 切图时动物和文字都不能被切到，
   并且太阳和月亮这种正圆不能变形；

4. 四周是纯白底色的时候要加0.5厘米的黑色描边，
   快捷键是 Alt+Ctrl+C；

5. 切图时半透或透光不透影最大148的宽，
   不透最大180的宽；

6. 切图遇到不透画布并且双面图案不一样时，
   每张需额外备注：“印一张”；

7. 目前软件公式中，旧厂订布预留是 5.00 厘米。`

	helpInfo := widget.NewMultiLineEntry()
	helpInfo.SetText(helpText)

	// 分别放到两个组
	helpInfoGroup := widget.NewGroup("切图规则", helpInfo)
	//HelpCodeGroup := widget.NewGroup("快捷暗号", makeHelpCode(w))

	return widget.NewVSplitContainer(
		helpInfoGroup,
		//HelpCodeGroup,
		makeHelpCode(w),
	)
}
