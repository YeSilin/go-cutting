package tools

// 汉字转拼音

import (
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"unicode"
)

// IncludeChinese1 判断是否包含中文 正则
func IncludeChinese1(s string) bool {
	for _, c := range s {
		regexp := regexp.MustCompile("[\u4e00-\u9fa5]")
		if ok := regexp.MatchString(string(c)); ok {
			return true
		}
	}
	return false
}

// IncludeChinese2 判断是否包含中文 直接使用 unicode
func IncludeChinese2(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// ToPinyin 汉字转拼音
func ToPinyin(str string) []string {
	// 默认
	a := pinyin.NewArgs()

	// 包含声调
	a.Style = pinyin.Tone3

	return pinyin.LazyPinyin(str, a)
}
