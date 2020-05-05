package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"time"
)

func main0() {

	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundBright))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundBlack))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundRed))

	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundGreen))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundYellow))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundBlue))

	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundMagenta))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundCyan))
	fmt.Println(model.ColourString("就想看看是什么颜色", ctc.ForegroundWhite))
	time.Sleep(time.Second*1000)
}

func main1() {
	hans := "薄纸"
	// 默认
	a := pinyin.NewArgs()

	// 包含声调
	a.Style = pinyin.Tone3

	fmt.Println(pinyin.LazyPinyin(hans, a))
	// [[zhōng] [guó] [rén]]
}