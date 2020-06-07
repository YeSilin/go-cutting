package main

import "fmt"


func a(s string) {
	fmt.Println(len(s))
}
func main00() {
	fmt.Println(len(":: "))
	fmt.Println(len("："))

	fmt.Println(len(":: 请输入常规座屏的宽："))
	b:=":: 请输入常规座屏的宽："
	a(b)

}