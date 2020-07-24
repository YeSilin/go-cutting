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

// UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

// GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func main33() {
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

	text := `傲世九重天 009
不败战神 002
不灭元神 066
不死邪王 064 完结
道印 009
斗破苍穹 838
都市仙王 032
都市之逆天仙尊 059
都市至尊 056
独步逍遥 064
二初居士 002
复仇少爷小甜妻 001
盖世帝尊 023 一般
剑舞051
解离妖圣 013
九阳神王 010
绝品强少 016 完结
绝世飞刀 127
绝世古尊 036
狂暴逆袭 009 完结
灵剑尊 298
灵武帝尊  020
魔耶人间玉 071
末世为王 152
你个神棍快走开 035
逆天剑神 097
逆天邪神 132
逆转仙途 041
女子学院的男生 080
盘龙 031
人气同桌是只猫 005
神武天尊 139 优秀
踏碎仙河 047
太古狂神 045
天才高手 005
天醒之路 015
外挂仙尊 004
万古神王 114 优秀
万界仙踪 219
王者超神的小兵 004
我把天道修歪了 002
我本废材 175 一般
我是大仙尊 155 优秀
我原来是个病娇 163
无敌剑神 004
武道独尊 132
武动乾坤 266
武神天下 105
武神主宰 371
仙帝归来 074
仙武帝尊 039
仙侠世界 018
星武神诀 274
修仙狂徒 014
玄界之门 102
妖神记 343
妖神学院 008
妖者为王 007
一品高手 074
异能少年王 038
姻缘宝典 251
元尊 003
遮天 004
至尊神级系统 083
至尊神魔 002
重生八万年 082
重生都市天尊 036
重生弃少归来 021
重生之都市狂仙 060 完结
重生之慕甄 006
诸天至尊 030
追天 006
最后的召唤师 043
最强农民工 187
最强升级 075
最强妖孽 008
最强枭雄系统 080
尊上 102 优秀
伏天圣主 009
修罗剑尊 075
万界仙王 030
极武玄帝 070 
轮回一剑 018
近身保镖 013
混沌金乌 014
不死凡人 025
斗罗大陆 054`
	s := strings.Split(text, "\n")
	//fmt.Printf("%q\n",s)
	sort.Sort(ByPinyin(s))
	//fmt.Printf("%q\n",s)
	for _, v := range s {
		fmt.Println(v)
	}
}