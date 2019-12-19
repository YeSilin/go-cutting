package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"sort"
	"strings"
)

// 拼音排序算法
type ByPinyin []string

func (s ByPinyin) Len() int      { return len(s) }
func (s ByPinyin) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPinyin) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i])
	b, _ := UTF82GBK(s[j])
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func main6() {
	//fmt.Println("排序开始=======>")
	//b := []string{"最后","哈", "呼", "嚯", "ha", ",","恐龙","的看的","刘","张三","曾哥","12","da","55","---"}
	//
	//sort.Strings(b)
	////output: [, ha 呼 哈 嚯]
	//fmt.Println("Default sort: ", b)
	//
	//sort.Sort(ByPinyin(b))
	////output: [, ha 哈 呼 嚯]
	//fmt.Println("By Pinyin sort: ", b)

	text := `最强农民工 187
最后的召唤师 043
最强枭雄系统 080
最强妖孽 008
武神主宰 329
武神天下 101
武动乾坤 253
武道独尊 095
斗破苍穹 798
妖神记 280
妖神学院 008
神武天尊 030
外挂仙尊 004
不败战神 002
不灭元神 058
逆天邪神 112
诸天至尊 030
天醒之路 015
天才高手 005
解离妖圣 013
盘龙 031
我本废材 169
我是大仙尊 061
太古狂魔 041
灵剑尊 146
修仙狂徒 014
星武神诀 274
人气同桌是只猫 005
追天 006
末世为王 117
你个神棍快走开 035
女子学院的男生 080
姻缘宝典 251
尊上 039
逆天剑神 076
逆转仙途 041
都市至尊 041
魔耶人间玉 003
至尊神级系统 083
异能少年王 038
九阳神王 010
绝世古尊 007
一品高手 074
元尊 003
绝世飞刀 119
灵武帝尊  015
最强升级 046
二初居士 002
仙帝归来 030
万界仙踪 219
玄界之门 091
至尊神魔 002
王者超神的小兵 004
万古神王 042`
	s := strings.Split(text, "\n")
	//fmt.Printf("%q\n",s)
	sort.Sort(ByPinyin(s))
	//fmt.Printf("%q\n",s)
	for _,v := range s {
		fmt.Println(v)
	}
}