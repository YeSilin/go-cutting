package model

import (
	"fmt"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/clib"
	"github.com/yesilin/go-cutting/model/quickCipher"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 判断是否为数字，并根据指定值提供指定的全局功能
func Input(text string, canvasMode bool) string {
	var num string

	for {
		// 用户输入提示，获取键盘输入
		fmt.Print(text)

		// 获取光标位置
		x, y := clib.WhereXY()

		_, err := fmt.Scanln(&num) // Scanln 换行也是一次输入结束
		if err != nil {
			if !canvasMode { // 不是画布模式就输出0
				// 重新指定xy位置
				clib.GotoPostion(x-1, y-1)
				fmt.Println("0")
				return "0"
			} else {
				// 重新指定xy位置
				clib.GotoPostion(0, y-2)
				continue
			}
		}

		// 二级返回
		if num == "--" {
			return "--"
		}

		// 在字符串中最后出现位置的索引，如果返回 -1 表示字符串不包含要检索的字符串
		lastIndex := strings.LastIndex(num, "-")
		// 如果减号出现在最后一位
		if lastIndex == len(num)-1 {
			tools.CallClear() // 清屏
			//fmt.Println(strings.Repeat("-", 36) + " Return " + strings.Repeat("-", 35) + "\n")
			return "-"
		}

		// 开始指定功能
		switch num {
		case "":
			continue
		case "-1":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/SelectTailor.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

			// 每次选择正确的脚本时删除多余备份，最大保留30个
			go tools.DeleteRedundantBackups("Config/JSX/Temp/*", 30)
			continue
		case "-2":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/NewDocumentJS.jsx"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
			fmt.Println("\n【注意】已重建新文档，参数来自最近一次切图记录！")
			tools.PrintLine(2)
			continue
		case "-3": // 深度清除源数据
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/JSX/ClearMetadataJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

			continue
		case "-4":
			quickCipher.Work() // 工作目录
		case "-5":
			// 将矩形选框转换为标记测量标志
			dataPath := "Config/JSX/SizeMarks.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
		case "-6":
			// 简单清除元数据
			dataPath := "Config/JSX/ClearMetadataNoPopUpJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
			continue
		case "-7":
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/BlackEdgeJS.jsx"
			go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
			continue
		case "-8":
			tools.CallClear() // 清屏
			continue
		case "-9":
			//获取当前时间，进行格式化 2006-01-02 15:04:05
			fileName := time.Now().Format("2006-01-02")
			now := time.Now().Format("2006-01")

			// 储存历史记录路径
			path := fmt.Sprintf("Config/History/%s/%s.txt", now, fileName)

			// 先查看是否有历史记录文件
			exists, _ := tools.IsPathExists(path)
			// 如果找不到文件，就创建文件 头
			if !exists {
				fmt.Println("\n【错误】找不到今天的切图历史记录，可能今天还未开始切图，已自动打开历史文件夹！")

				exec.Command("cmd.exe", "/c", "start Config\\History").Run()
				tools.PrintLine(2)
				continue
			}
			// 创建一个协程使用cmd来运行脚本
			cmd := exec.Command("cmd.exe", "/c", "start "+path)
			go cmd.Run()

			continue
		case "-11":
			OldFrame1()
			continue
		case "-12":
			OldFrame2()
			continue
		case "-13":
			OldFrame3()
			continue
		case "-14":
			OldFrame4()
			continue
		case "-15":
			OldFrame5()
			continue
		case "-16":
			OldFrame6()
			continue
		case "-17":
			OldFrame7()
			continue
		case "-18":
			OldFrame8()
			continue
		case "-19":
			OldFrame9()
			continue
		case "-98":
			// 创建套图文件夹
			_ = tools.CreateMkdirAll("Config/Picture/主图")
			// 创建一个协程使用cmd来运行脚本
			dataPath := "Config/JSX/DetailsPage.jsx"
			exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

			time.Sleep(time.Second) // 停一秒
			// 如果存在images就打开
			if ok, _ := tools.IsPathExists("Config/Picture/主图/images"); ok {
				go exec.Command("cmd.exe", "/c", "start Config\\Picture\\主图\\images").Run()
			} else {
				go exec.Command("cmd.exe", "/c", "start Config\\Picture\\主图").Run()
			}
			continue
		case "-99":
			// 创建一个协程使用cmd启动外部程序
			dataPath := "Config/exe/W10DigitalActivation.exe /activate"
			cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
			go cmd.Run()
			continue
		}

		//strings.Replace(num,".","",0)  // 删掉全部小数点
		// 如果小数点多于一个就循环
		if strings.Count(num, ".") > 1 {
			continue
		}

		// 如果不是整数或浮点数就循环
		integer, _ := regexp.MatchString(`^(\-|\+)?\d+(\.\d+)?$`, num)
		if !integer {
			continue
		}

		// 如果输入的数字小于6就循环
		tempNum, _ := strconv.ParseFloat(num, 64)
		if tempNum < 6 && canvasMode {
			// 查找当前提示信息中是否已包含 插入的提示信息
			if !strings.Contains(text, "（尺寸不可小于6厘米）") {
				// 没有就插入
				text = tools.StrRightInsert(text, ColourString("（尺寸不可小于6厘米）", ctc.ForegroundRed), 3)
			}
			continue
		}
		return num
	}
}
