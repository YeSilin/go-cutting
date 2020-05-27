package tools

import (
	"reflect"
	"testing"
)

// go test -run IncludeChinese1 指定测试

func TestIncludeChinese1(t *testing.T) {
	// 程序输出的结果
	got := IncludeChinese1("我在讲中文")
	want := true
	if want != got { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func TestIncludeChinese2(t *testing.T) {
	// 程序输出的结果
	got := IncludeChinese2("我在讲中文")
	want := true
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
