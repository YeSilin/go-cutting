package model

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"testing"
)


// go test -run Read2345PicInstallationPath -v
func TestRead2345PicInstallationPath(t *testing.T){
	got := tools.Get2345PicInstallPath()
	fmt.Println(got)
}

// go test -run CleanUp2345Pic
// 净化软件
func TestCleanUp2345Pic(t *testing.T)  {
	CleanUp2345Pic()
}