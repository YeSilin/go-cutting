package layout

import (
	"fmt"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

// 判断用户输入，是否有白底图
func isWhiteBackgroundMap() string {
	model.EnglishTitle("Is there a white background?", 79)
	for {
		fmt.Print("\n【套图】是否有白底图？[y]是，[n]否：")
		var str string
		_, _ = fmt.Scanln(&str)
		switch str {
		case "y":
			return "y"
		case "n":
			return "n"
		case "-", "--":
			tools.CallClear() // 清屏
			return "-"
		}
	}
}

// 6 专属主图自动化
func picture() {
	for {
		model.EnglishTitle("Layout to picture", 79)
		text := `
【主图】专属版主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与LOGO！

【主图】[1]沐兰主图                [2]怡柟家具                [3]御尚家具

【主图】[4]御尚屏风                [5]木韵主图                [6]华府主图

【主图】[7]金尊主图                [8]暂未开发                [9]暂未开发`
		fmt.Println(text)

		layoutType := model.Input("\n【主图】请选择需要使用的功能：", false)

		switch layoutType {
		case "1", "01":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/mulan.png")
		case "2", "02":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/yinanjj.png")
		case "3", "03":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/yushantanjj.png")
		case "4", "04":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/yushantanpf.png")
		case "5", "05":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/muyunge.png")
		case "6", "06":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/huafu.png")

		case "7", "07":
			tools.CallClear() // 清屏
			WatermarkMasterGraph("config/img/jinzunfu.png")
		case "8", "08":

		case "-", "--":
			goto FLAG
		default:
			tools.CallClear() // 清屏
			continue
		}
	}
FLAG:
}

// 7 修改图片分辨率
func modifyResolution() {
	for {
		model.EnglishTitle("Layout to resolution", 79)
		text := `
【提示】在确保原有像素尺寸不变的情况下，将分辨率强行修改至指定大小，请尽量少用！

【提示】[1]修改全部jpg为 72ppi                         [2]修改全部jpg为 300ppi`
		fmt.Println(text)
		layoutType := model.Input("\n【主图】请选择需要使用的功能：", false)

		switch layoutType {
		case "1":
			tools.CallClear()
			PixelsPerInchChangedTo72()                                      // 改为72ppi
			go exec.Command("cmd.exe", "/c", "start Config\\Picture").Run() // 打开套图文件夹
		case "2":
			tools.CallClear()
			PixelsPerInchChangedTo300()
			go exec.Command("cmd.exe", "/c", "start Config\\Picture").Run() // 打开套图文件夹
		case "-", "--":
			goto FLAG
		default:
			tools.CallClear() // 清屏
			continue
		}
	}
FLAG:
}

// 套图的选择
func Choice() {
	for {
		model.EnglishTitle("Layout", 79)
		text := `
【套图】[1]套图文件夹              [2]随机重命名              [3]备份文件夹

【套图】[4]替换详情页              [5]通用版主图              [6]专属版主图

【套图】[7]修改分辨率              [8]功能未开发              [9]功能未开发`
		fmt.Println(text)

		layoutType := model.Input("\n【套图】请选择需要使用的功能：", false)

		switch layoutType {
		case "1":
			tools.CallClear() // 清屏
			go exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
			fmt.Println("\n【提示】已打开套图文件夹，请复制正方形的 jpg 或 png 高清图片以备自动套图使用！")
		case "2":
			tools.CallClear() // 清屏
			Rename()
			// 打开套图文件夹
			exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
		case "3":
			tools.CallClear() // 清屏
			go exec.Command("cmd.exe", "/c", "start Config\\Backups").Run()
			fmt.Println("\n【提示】已打开备份文件夹，为了避免意外丢失文件，目前备份文件最大为 10 份！")

		case "4":
			tools.CallClear()             // 清屏
			ReplaceDetailsPage() // 替换详情页
		case "5":
			tools.CallClear()      // 清屏
			UniversalMasterGraph() // 通用主图
			// 打开套图文件夹
			exec.Command("cmd.exe", "/c", "start Config\\Picture").Run()
		case "6":
			tools.CallClear() // 清屏
			picture()         // 主图自动化
		case "7":
			tools.CallClear()  // 清屏
			modifyResolution() // 修改分辨率
		case "-", "--":
			goto FLAG
		default:
			tools.CallClear() // 清屏
			continue
		}
	}
FLAG:
}
