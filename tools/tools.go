package tools

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 打印分割线 1短 2长
func PrintLine(pattern int) {
	switch pattern {
	case 1:
		fmt.Println(strings.Repeat("-", 87)) // 重复打印89个减号
	case 2:
		color.LightCyan.Println(" " + strings.Repeat("-", 78)) // 返回专属的下划线装饰
	case 3:
		color.LightCyan.Println("\n" + strings.Repeat("-", 29) + " 请注意切图的框架选择 " + strings.Repeat("-", 28))
	}
}

// 数字转换中文小写
func Transfer(num int) string {
	chineseMap := []string{"", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	var listNum []int
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}
	n := len(listNum)
	chinese := ""
	//注意这里是倒序的
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])

	}
	//注意替换顺序
	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}

	return chinese
}

//如果画布的高和宽同时大于148则提示
func MaxCanvas(width, height float64) {
	if width > 180 && height > 180 {
		// 这其实是红色
		color.LightBlue.Println("（已超不透最大180cm。）")
	} else if width > 148 && height > 148 {
		color.LightBlue.Println("（已超半透最大148cm。）")
	} else {
		fmt.Println()
	}
	//PrintLine(2) // 每个框架运行完都加个分隔符

	//最大画布字体图层提示
	var jsx = strings.Builder{}

	jsx.WriteString("// 无聊加了个画布大小判断\n")
	jsx.WriteString(fmt.Sprintf("if ((%f>180) && (%f>180)) {\n", width, height))
	jsx.WriteString("     // 生成历史记录\n")
	jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超不透最大180cm\", \"maxCanvas(\\\"注意：已超不透最大180cm。\\\",  \\\"9d2e2d\\\")\");\n")
	jsx.WriteString(fmt.Sprintf("} else if ((%f>148) && (%f>148)) {\n", width, height))
	jsx.WriteString("      // 生成历史记录\n")
	jsx.WriteString("    app.activeDocument.suspendHistory(\"注意：已超半透最大148cm\", \"maxCanvas(\\\"注意：已超半透最大148cm。\\\",  \\\"77bb11\\\")\");\n")
	jsx.WriteString("}\n")
	jsx.WriteString("\n")
	jsx.WriteString("\n")
	jsx.WriteString("function maxCanvas(text, rgbValue){\n")
	jsx.WriteString("    // 在当前文档中添加一个图层。并且用变量 newLayer 记录这个图层。\n")
	jsx.WriteString("    var newLayer = app.activeDocument.artLayers.add();\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //把图层 newLayer 的图层类型变为”文本“ ，图层转换为文本图层。\n")
	jsx.WriteString("    newLayer.kind = LayerKind.TEXT;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //把图层 newLayer 的文本内容类型变为”文本框“。\n")
	jsx.WriteString("    newLayer.textItem.kind = TextType.PARAGRAPHTEXT;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置图层 newLayer 的文本框宽度与高度。\n")
	jsx.WriteString("    newLayer.textItem.width = app.activeDocument.width*0.8;\n")
	jsx.WriteString("    newLayer.textItem.height = app.activeDocument.width*0.1;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置图层 newLayer 的文本框位置，横坐标 50 像素，纵坐标 100 像素。\n")
	jsx.WriteString("    //newLayer.textItem.position= [UnitValue(\"50px\"), UnitValue(\"100px\")]\n")
	jsx.WriteString("    newLayer.textItem.position= [UnitValue(app.activeDocument.width*0.1), UnitValue((app.activeDocument.height*0.5)-(app.activeDocument.width*0.025))];\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置 newLayer 的文本字体大小为“40 点”。\n")
	jsx.WriteString("    newLayer.textItem.size = UnitValue(app.activeDocument.width*0.05);\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置 newLayer 的文本内容。\n")
	jsx.WriteString("    newLayer.textItem.contents= text;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置 newLayer 的文本框对齐方式为居中对齐。\n")
	jsx.WriteString("    newLayer.textItem.justification = Justification.CENTER;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //创建一个色彩变量 c   ，颜色为 #77bb11。\n")
	jsx.WriteString("    var c = new SolidColor();\n")
	jsx.WriteString("    c.rgb.hexValue = rgbValue;\n")
	jsx.WriteString("\n")
	jsx.WriteString("    //设置 newLayer 的文本颜色为 c。\n")
	jsx.WriteString("    newLayer.textItem.color = c;\n")
	jsx.WriteString("}\n")

	// 转成字符串格式
	jsxStr := jsx.String()
	// 追加写入
	WriteFile("Config/jsx/NewDocumentJS.jsx", jsxStr)
}

// 在字符串右侧第 X 位插入字符
func StrRightInsert(str, insert string, num int) string {
	// 在末尾第几位插入插入
	index := len(str) - num

	return str[:index] + insert + str[index:]
}

// 这是一个将切片转换成js数组的函数
func ToJsArray(s []float64) string {
	var str = "new Array("
	for i := 0; i < len(s); i++ {
		str += fmt.Sprintf("%f", s[i])
		if i != len(s)-1 {
			str += ", "
		}
	}
	str += ")"
	return str
}


// 这是一个将切片转换成js数组的函数
func StrToJsArray(s []string) string {
	var str = "new Array("
	for i := 0; i < len(s); i++ {
		str += fmt.Sprintf("\"%s\"", s[i])
		if i != len(s)-1 {
			str += ", "
		}
	}
	str += ")"
	return str
}


// 获取指定天数前后的时间戳  负数是几天前
func AroundTime(day string) int64 {
	// 先转成字符串类型

	// 这里的now是网络时间哦
	now := time.Unix(GetNtpTime(), 0)
	h, _ := time.ParseDuration(fmt.Sprintf("%sh", day))
	//h, _ := time.ParseDuration("1h")

	//   Add 时间相加
	return now.Add(h * 24).Unix()
	// 时间戳转指定格式
	//t := time.Unix(1234567890, 0)
	//fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// 将秒转换成天
func ToDay(second int64) int64 {
	return second/60/60/24 + 1
}

// 距离十八点还有多少秒
func DistanceIsEighteen() (int, bool) {
	t := time.Now()
	//fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	s := (17 - t.Hour()) * 60 * 60 // 时
	f := (59 - t.Minute()) * 60    // 分
	m := 60 - t.Second()           // 秒
	sum := s + f + m
	//fmt.Println(sum)

	if 17-t.Hour() > 0 {
		return sum, true
	} else {
		return sum, false
	}
}

// 获取当前目录
func GetCurrentDirectory() string {
	// 返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// 将\替换成/  如果n<0，则替换次数没有限制
	dir = strings.Replace(dir, "\\", "/", -1)
	// 将:替换成空
	return strings.Replace(dir, ":", "", 1)
}

// 删除多于备份 pattern路径  max保留最大数量
func DeleteRedundantBackups(pattern string, max int) {
	// 获取所有文件名，类型是字符串切片
	files, _ := filepath.Glob(pattern)

	// 求出多出来的文件数量
	count := len(files) - max

	// 多出几个文件数量就循环几次，如果是负数自然就不循环
	for i := 0; i < count; i++ {
		// 删除文件，Go中删除文件和删除文件夹同一个函数
		err := os.RemoveAll(files[i]) // 由于windows 系统获取到的文件名字，默认是升序，于是可以不用排序
		// 打印被删除的文件
		//fmt.Println(files[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}
