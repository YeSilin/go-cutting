package main

import (
	"fmt"
	"time"
)

func a(){
	fmt.Println("a开始")
	go func() {
		for i := 0; i <10;i++{
			fmt.Println("a",i)
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("a结束")

}


func main11() {


	a()

	time.Sleep(30*time.Second)

}
