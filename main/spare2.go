package main

import (
	"fmt"
	"github.com/gookit/color"
)

func main2() {
	// 简单快速的使用，跟 fmt.Print* 类似
	color.Red.Println("Simple to use color")
	color.Green.Print("Simple to use color")
	color.Cyan.Printf("Simple to use %s\n", "color")
	color.Yellow.Printf("Simple to use %s\n", "color")

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// 自定义颜色
	color.New(color.FgWhite, color.BgBlack).Println("custom color style")

	// 也可以:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

	// internal style:
	color.Info.Println("message")
	color.Warn.Println("message")
	color.Error.Println("message")

	// 使用颜色标签
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// apply a style tag
	color.Tag("info").Println("info style text")

	// prompt message
	color.Info.Prompt("prompt style message")
	color.Warn.Prompt("prompt style message")

	// tips message
	color.Info.Tips("tips style message")
	color.Warn.Tips("tips style message")


	q := "0"
	fmt.Scan(&q)
}