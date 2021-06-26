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



// 通过两重循环过滤重复元素
func RemoveRepeatedElement(arr []string) []string {
	// 存放结果，空初始化比make定义要快
	result := []string{}
	// 外层循环准备添加到结果的切片
	for i := 0; i < len(arr); i++ {
		// 初始定义该元素不存在，很奇怪，初始在循环里面比先声明在外面之后再赋值要快
		exist := false
		// 这里根据当前切片的长度进行循环，直接使用 len 比初始一个 count变量 记数要快
		for j := 0; j < len(result); j++ {
			// 如果遇到重复提前退出
			if result[j] == arr[i] {
				// 并且说明已存在
				exist = true
				break
			}
		}
		// 如果在 result切片都没有遍历到此元素
		if !exist {
			// 那么就追加到 result
			result = append(result, arr[i])
		}
	}
	return result
}

func main3() {
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
不灭元神 084 完结
不死凡人 025 完结
不死邪王 064 完结
超级吞噬系统 015
大隐于宅 006 一般
道印 009
斗罗大陆 054
斗罗大陆2绝世唐门 011
斗破苍穹 868
都市仙王 032
都市之逆天仙尊 115
都市至尊 056
独步逍遥 082 一般
二初居士 002
废女妖神 012
风起苍岚 085
伏天圣主 047
复仇少爷小甜妻 001
盖世帝尊 023 一般
混沌丹神 125 一般
混沌金乌 083 一般
极武玄帝 077 一般
剑道凌天 048
剑舞 062 一般
解离妖圣 016 一般
近身保镖 013
九阳帝尊 007 一般
九阳神王 010
绝品强少 016 完结
绝世飞刀 128 一般
绝世古尊 036
绝世圣帝 026 一般
绝世武魂 013
狂暴逆袭 009 完结
凌天神帝 156
灵剑尊 319 优秀
灵武帝尊  020
轮回一剑 021 完结
魔耶人间玉 071
末世女友我家后院通末世 086
末世为王 152
你个神棍快走开 035
逆天剑神 097
逆天仙命 001
逆天邪神 132
逆转仙途 041
女子学院的男生 080
盘龙 031
全能至尊 011
人气同桌是只猫 005
神武帝尊 079
神武天尊 292 优秀
踏碎仙河 116
太古狂神 045
天才高手 005
天醒之路 015
歪嘴战神 003 一般
外挂仙尊 004
万古神王 217 优秀
万界仙王 033
万界仙踪 309
王者超神的小兵 004
我把天道修歪了 002
我被封印九亿次 011 优秀
我被困在同一天十万年 014 优秀
我本废材 175 一般
我是大仙尊 198 优秀
我原来是个病娇 163
无敌剑神 004
无敌剑域 053
武道独尊 132
武动乾坤 266
武逆 092 一般
武神天下 113 完结
武神主宰 415
仙帝归来 136
仙帝入侵 011
仙武帝尊 074
仙侠世界 018
星武神诀 274
修罗剑尊 085 一般
修仙狂徒 014
修仙者大战超能力 076 优秀
玄界之门 102
妖神记 318-1 优秀
妖神学院 008
妖者为王 007
一剑独尊 025
一品高手 074
异能少年王 038
姻缘宝典 266
元尊 003
遮天 004
至尊神皇 045 完结
至尊神级系统 146
至尊神魔 002
至尊重生 032 一般
重生八万年 119 优秀
重生都市天尊 036
重生弃少归来 021
重生之都市狂仙 060 完结
重生之剑神归来 019  一般
重生之慕甄 006
诸天至尊 030
主宰三界 101
追天 006
最后的召唤师 043
最强农民工 187
最强升级 154
最强妖孽 008
最强枭雄系统 080
尊上 196 优秀
巅峰预言帝 008 一般
神魂武帝 009 一般
武映三千道 071 一般
开天录 014 一般
超能不良学霸 001`
	s := strings.Split(text, "\n")

	s = RemoveRepeatedElement(s)
	//fmt.Printf("%q\n",s)
	sort.Sort(ByPinyin(s))
	//fmt.Printf("%q\n",s)
	for _, v := range s {
		fmt.Println(v)
	}
}
