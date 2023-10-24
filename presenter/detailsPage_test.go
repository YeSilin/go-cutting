package presenter

import (
	"testing"
)

// go test -run GetAllImages -v
// -v 是现实详细内容
func TestGetAllImages(t *testing.T) {

	// 程序输出的结果
	got, _ := getAllImages("E:\\Code\\Golang\\go-cutting\\data\\Picture")
	// 期望得到的结果
	//want := []

	dp, de := splitDetails(got)
	t.Log(got)
	t.Log("dp如下：", dp)
	t.Log("de如下：", de)
}
