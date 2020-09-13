package main

import (
	"fmt"
	"path/filepath"
)

func main2() {

	files, _ := filepath.Glob(fmt.Sprintf("%s/*", `E:\淘宝美工\套图汇总`))
	fmt.Println(files)
}
