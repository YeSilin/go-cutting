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
		EnglishTitle("3ds Max map frame", 79)
		text := `
【贴图】[1]常规座屏贴图            [2]左右镂空贴图            [3]左右画布贴图

【贴图】[4]上下镂空贴图            [5]顶天立地贴图            [6]各种折屏贴图

【贴图】[7]多个座屏贴图            [8]卷帘座屏贴图            [9]不扣补切贴图`
		fmt.Println(text)

		frameType := Input("\n【贴图】请选择上方的边框类型：", false)

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
		default:
			continue
		}
	}
FLAG:
}

// 贴图小座屏
func mapFrame1() {
	for {
		tools.PrintLine(3) // 请注意切图的工厂与框架的选择
		widthStr := Input("\n【贴图】请输入常规座屏的宽（默认120）：", true)
		if widthStr == "-" {
			break
		}
		heightStr := Input("\n【贴图】请输入常规座屏的高（默认180）：", true)
		if heightStr == "-" {
			break
		}

		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)

		width -= 8
		height -= 8



		fmt.Printf("\n【贴图】常规座屏：宽 %.0f pixels，高 %.0f pixels", width*10, height*10)
		generate.MaxCanvas(width, height)
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
		tools.PrintLine(3) // 请注意切图的工厂与框架的选择
		widthStr := Input("\n【贴图】请输入折屏单扇的宽（默认45）：", true)
		if widthStr == "-" {
			break
		}
		heightStr := Input("\n【贴图】请输入折屏单扇的高（默认180）：", true)
		if heightStr == "-" {
			break
		}
		upperHollowOutStr := Input("\n【贴图】请输入上镂空的大小（默认0）：", false)
		if heightStr == "-" {
			break
		}

		downHollowOutStr := Input("\n【贴图】请输入下镂空的大小（默认0）：", false)
		if heightStr == "-" {

			break
		}

		numberStr := Input("\n【贴图】请输入共拥有几扇：", false)
		if numberStr == "-" {
			break
		}

		width, _ := strconv.ParseFloat(widthStr, 64)
		height, _ := strconv.ParseFloat(heightStr, 64)
		number, _ := strconv.ParseFloat(numberStr, 64)
		upperHollowOut, _ := strconv.ParseFloat(upperHollowOutStr, 64)
		downHollowOut, _ := strconv.ParseFloat(downHollowOutStr, 64)

		width -= 8 // 单扇的宽

		totalWidth := width * number // 总宽
		height -= 8                  // 单扇的高

		if upperHollowOut > 0 { // 如果有上镂空的话
			height -= upperHollowOut + 4
		}
		if downHollowOut > 0 { // 如果有下镂空的话
			height -= downHollowOut + 4
		}

		fmt.Printf("\n【贴图】常规折屏：宽 %.0f pixels，高 %.0f pixels", totalWidth*10, height*10)
		generate.MaxCanvas(width, height)

		//获取当前时间，进行格式化 2006-01-02 15:04:05
		now := nowTime()

		// 为当前框架指定名字
		frameName := fmt.Sprintf("%s_折屏贴图_%.0fx%.0f", now, totalWidth*10, height*10)

		// 定义单片名字
		singleName := fmt.Sprintf("%s_折屏贴图", now)

		generate.NewDocumentForMap(totalWidth, height, frameName)             // 生成创建ps文档脚本
		generate.Line3DMapJs6(width, number)                                  // 生成专属参考线
		go generate.TailorForMap6(width, height, number,frameName,singleName) // 生成暗号【-1】可以用的另存脚本
		if viper.GetBool("openPs") {                                          // 是否自动新建ps文档
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
