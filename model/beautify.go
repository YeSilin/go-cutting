package model

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/wzshiming/ctc"
	"strings"
)

// 第一个字符串带颜色，第二个字符串不带颜色
func BsPrintln(cs, ds string) {
	color.Green.Print(cs)
	fmt.Println(ds)
}

func BsPrint(cs, ds string) {
	color.Green.Print(cs)
	fmt.Print(ds)
}

// 中文提示标题; s 中文提示，width 标题宽度
func ChineseTitle(s string, width int) {
	// 求出有几个中文字
	count := len(s) / 3

	// 求出实际可用的 装饰----
	available := width - count*2 - 2

	// 右边的装饰可以少一点，例如9/2=4,
	right := available / 2

	left := available - right

	// 因为起始位置要空一格，所以如果左边大于右边就减1
	if left>right{
		left--
	} else {
		right--
	}

	color.LightCyan.Println("\n " + strings.Repeat("-", left) + fmt.Sprintf(" %s ", s) + strings.Repeat("-", right))
}

// 英文提示标题; s 中文提示，width 标题宽度
func EnglishTitle(s string, width int) {
	// 求出有几个英文字
	count := len(s)

	// 求出实际可用的 装饰----  -2 是因为两边要留空格
	available := width - count - 2

	// 右边的装饰可以少一点
	right := available / 2
	left := available - right

	// 因为起始位置要空一格，所以如果左边大于右边就减1
	if left>right{
		left--
	} else {
		right--
	}
	color.LightCyan.Println("\n " + strings.Repeat("-", left) + fmt.Sprintf(" %s ", s) + strings.Repeat("-", right))
}

// 带颜色的字符串，第一参数 字符串 ，第二参数前景色，第三参数背景色
func ColourString(str string, c ...ctc.Color) string {
	if len(c) == 1 {
		return fmt.Sprintf("%s%s%s",c[0], str, ctc.Reset)
	}
	return fmt.Sprintf("%s%s%s",c[0]|c[1], str, ctc.Reset)
}