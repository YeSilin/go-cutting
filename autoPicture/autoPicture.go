package autoPicture

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/generate"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
)

// 套图的选择
func Choice() {
	for {
		model.EnglishTitle("Auto picture", 74)
		text := `
:: 自动套图可以正确运行的前提是，款式图片均要放在套图文件夹下方可正确运行！

   [1]套图文件夹                [2]随机重命名                [3]备份文件夹

   [4]家具店主图                [5]通用版主图                [6]屏风店主图

   [7]修改分辨率                [8]替换详情页                [9]导出详情页`
		fmt.Println(text)

		layoutType := model.Input("\n:: 请选择需要使用的功能：", false)
		tools.CallClear() // 清屏
		switch layoutType {
		case "1":
			go exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", viper.GetString("picture"))).Run()
			fmt.Println("\n:: 已打开套图文件夹，请复制正方形的 jpg 或 png 高清图片以备自动套图使用！")
		case "2":
			RandomRename(viper.GetString("picture"))
		case "3":
			go exec.Command("cmd.exe", "/c", "start Config\\Backups").Run()
			fmt.Println("\n:: 已打开备份文件夹，为了避免意外丢失文件，目前备份文件最大为 10 份！")

		case "4":
			furnitureMainPictureChoice()  // 家具店主图
		case "5":
			UniversalMasterGraph(viper.GetString("picture"), viper.GetBool("automaticDeletion")) // 通用主图
		case "6":
			screenMainPictureChoice() // 屏风店主图
		case "7":
			modifyResolution() // 修改分辨率
		case "8":
			generate.ReplaceDetailsPage(viper.GetString("picture")) // 替换详情页
		case"9":
			model.StartCode98() // 详情页导出
		case "-", "--":
			goto FLAG
		default:
			fmt.Printf("\n:: 输入的 [%s] 不是已知的功能选项，请重新输入！\n", model.ColourString(layoutType, ctc.ForegroundGreen))
			continue
		}
	}
FLAG:
}
