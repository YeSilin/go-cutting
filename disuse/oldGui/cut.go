package oldGui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/yesilin/go-cutting/model"
)

// 初始化小座屏的输入界面
func makeFame1() fyne.Widget {
	widthStr := widget.NewEntry()
	widthStr.SetPlaceHolder("请输入常规座屏的宽")

	heightStr := widget.NewEntry()
	heightStr.SetPlaceHolder("请输入常规座屏的高")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "宽度：", Widget: widthStr},
			{Text: "高度：", Widget: heightStr},
		},
		//OnCancel: func() {
		//
		//},
		OnSubmit: func() {
			// 调用框架公式
			model.FormulaFrame1(widthStr.Text, heightStr.Text)
		},


		SubmitText: "提交",
		CancelText: "重置",
	}

	return form
}

// 初始化切图
func makeCut() fyne.CanvasObject {
	// 要显示的框架
	//var showFame fyne.Widget

	//frameSelect := widget.NewSelect([]string{"常规座屏", "左右镂空", "左右画布", "上下镂空", "顶天立地", "各种折屏", "多个座屏", "卷帘座屏", "不扣补切"}, func(s string) {
	//	switch s{
	//	case "常规座屏" :
	//		//showFame = makeFame1()
	//	}
	//})
	//frameSelect.Selected = "左右镂空"

	// 新建一个标签集合
	tabs := widget.NewTabContainer(

		widget.NewTabItem(`常规座屏`, makeFame1()),
		widget.NewTabItem("左右镂空", widget.NewVBox()),
		widget.NewTabItem("左右画布", widget.NewVBox()),
		widget.NewTabItem("上下镂空", widget.NewVBox()),
		widget.NewTabItem("顶天立地", widget.NewVBox()),
		widget.NewTabItem("各种折屏", widget.NewVBox()),
		widget.NewTabItem("多个座屏", widget.NewVBox()),
		widget.NewTabItem("卷帘座屏", widget.NewVBox()),
		widget.NewTabItem("不扣补切", widget.NewVBox()),
	)
	// 设置位置为左对齐
	tabs.SetTabLocation(widget.TabLocationLeading)
	//tabs.Resize(fyne.NewSize(500, 1000))
	return widget.NewGroup("框架选择",  tabs)
}
