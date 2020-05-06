package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/tools"

	"os/exec"
	"strconv"
)

// 贴图框架的选择
func MapFrameChoice() {
	for {
		EnglishTitle("3ds Max map frame", 74)
		text := `
:: 下方所有框架的切图单位均是像素，小数点位之后值均会被舍弃，并非四舍五入！

   [1]常规座屏贴图             [2]左右镂空贴图             [3]左右画布贴图

   [4]上下镂空贴图             [5]顶天立地贴图             [6]各种折屏贴图

   [7]多个座屏贴图             [8]卷帘座屏贴图             [9]不扣补切贴图`
		fmt.Println(text)

		frameType , info:= Input("\n:: 请选择上方的边框类型：", false,true)
		tools.CallClear() // 清屏
		switch frameType {
		case "1":
			mapFrame1() // 小座屏
		case "2":
			fmt.Println("未开发") // 左右镂空
		case "3":
			fmt.Println("未开发") // 中间大两边小
		case "4":
			fmt.Println("未开发") // 上下镂空
		case "5":
			fmt.Println("未开发") // 顶天立地
		case "6":
			mapFrame6() // 常规折屏
		case "7":
			fmt.Println("未开发") // 多座屏
		case "8":
			fmt.Println("未开发") // 补切画布
		case "-":
			goto FLAG
		case "cls":
			// 收到清屏命令
			if len(info) != 0 {
				fmt.Println(info)
			}
			continue
		default:
			continue
		}
	}
FLAG:
}

// 贴图小座屏
func mapFrame1() {
	for {
		ChineseTitle("当前框架常规座屏贴图 ", 74) // 请注意切图的工厂与框架的选择

		widthStr, _ := Input("\n:: 请输入常规座屏的宽：", true,false)
		if widthStr == "-" {
			break
		}
		heightStr, _ := Input("\n:: 请输入常规座屏的高：", true,false)
		if heightStr == "-" {
			break
		}
		reserveStr, _ := Input("\n:: 请输入要减去的单边框大小(一般为40)：", false,false)
		if reserveStr == "-" {
			break
		}

		// 字符串转 int
		width, _ := strconv.Atoi(widthStr)
		height, _ := strconv.Atoi(heightStr)
		reserve, _ := strconv.Atoi(reserveStr)

		// 去掉边框
		width -= reserve * 2
		height -= reserve * 2

		fmt.Printf("\n:: 常规座屏：宽 %d pixels，高 %d pixels", width, height)
		generate.MaxCanvas(float64(width)/10, float64(height)/10)

		generate.NewDocumentForMap(width, height, "常规座屏贴图") // 生成创建ps文档脚本
		if viper.GetBool("openPs") {                        // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}
		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}

// 贴图折屏
func mapFrame6() {
	for {
		ChineseTitle("当前框架各种折屏贴图", 74) // 请注意切图的工厂与框架的选择

		widthStr, _ := Input("\n:: 请输入折屏单扇的宽：", true,false)
		if widthStr == "-" {
			break
		}
		heightStr , _:= Input("\n:: 请输入折屏单扇的高：", true,false)
		if heightStr == "-" {
			break
		}
		upperHollowOutStr, _ := Input("\n:: 请输入上镂空的大小：", false,false)
		if heightStr == "-" {
			break
		}

		downHollowOutStr , _:= Input("\n:: 请输入下镂空的大小：", false,false)
		if heightStr == "-" {

			break
		}

		numberStr , _:= Input("\n:: 请输入共拥有几扇：", false,false)
		if numberStr == "-" {
			break
		}

		reserveStr, _ := Input("\n:: 请输入要减去的单边框大小(一般为40)：", false,false)
		if reserveStr == "-" {
			break
		}

		// 字符串转 int
		width, _ := strconv.Atoi(widthStr)
		height, _ := strconv.Atoi(heightStr)
		upperHollowOut, _ := strconv.Atoi(upperHollowOutStr)
		downHollowOut, _ := strconv.Atoi(downHollowOutStr)
		number, _ := strconv.Atoi(numberStr)
		reserve, _ := strconv.Atoi(reserveStr)

		// 计算边框
		width -= reserve * 2         // 单扇的宽
		totalWidth := width * number // 总宽
		height -= reserve * 2        // 单扇的高

		if upperHollowOut > 0 { // 如果有上镂空的话
			height -= upperHollowOut + reserve
		}
		if downHollowOut > 0 { // 如果有下镂空的话
			height -= downHollowOut + reserve
		}

		fmt.Printf("\n:: 常规折屏：宽 %d pixels，高 %d pixels", totalWidth, height)
		generate.MaxCanvas(float64(width)/10, float64(height)/10)

		//获取当前时间，进行格式化 2006-01-02 15:04:05
		now := nowTime()

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_折屏贴图_%dx%d", now, totalWidth, height)

		// 定义单片名字
		singleName := fmt.Sprintf("%s_折屏贴图", now)

		generate.NewDocumentForMap(totalWidth, height, frameName)               // 生成创建ps文档脚本
		generate.Line3DMapJs6(width, number)                                    // 生成专属参考线
		go generate.TailorForMap6(width, height, number, frameName, singleName) // 生成暗号【-1】可以用的另存脚本
		if viper.GetBool("openPs") {                                            // 是否自动新建ps文档
			// 创建一个协程使用cmd来运行脚本
			dataPath := "config/jsx/newDocument.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
		}

		if !viper.GetBool("memory") { // 是否记忆框架
			break
		}
	}
}
