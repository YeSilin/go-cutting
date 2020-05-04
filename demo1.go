package main

import (
	"fmt"
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