package additional

import (
	"fmt"
	"testing"
)


// go test -run Read2345PicInstallationPath -v
func TestRead2345PicInstallationPath(t *testing.T){
	got := read2345PicInstallationPath()
	fmt.Println(got)
}

// go test -run CleanUp2345Pic
// 净化软件
func TestCleanUp2345Pic(t *testing.T)  {
	CleanUp2345Pic()
}