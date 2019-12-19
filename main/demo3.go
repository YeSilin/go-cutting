package main

import "fmt"

func main3() {
	a := []int{1, 2, 3}

	// 追加元素
	a = append(a, 4, 5, 6)
	fmt.Println(a) // [1 2 3 4 5 6]

	// 要追加的数据
	b := []int{7, 8, 9}
	// 合并切片
	a = append(a, b...)
	fmt.Println(a) // [1 2 3 4 5 6 7 8 9]
}
