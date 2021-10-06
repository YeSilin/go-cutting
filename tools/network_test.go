package tools

import "testing"

func TestIsNetwork(t *testing.T) {
	// 程序输出的结果
	got := IsNetwork()
	want := true
	if want != got {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}



func TestNetWorkStatus(t *testing.T) {
	// 程序输出的结果
	got := NetWorkStatus()
	want := true
	if want != got {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}