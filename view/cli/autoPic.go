package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wzshiming/ctc"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/presenter"
	"github.com/yesilin/go-cutting/tools"
	"github.com/yesilin/go-cutting/unclassified"
	"os/exec"
)

// 家具主图选择
func (c *CLI) furnitureMainPictureChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Furniture main picture choice", 74)
		text := `
:: 家具店主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与Logo！

   [1]怡柟家具                   [2]御尚家具                   [3]素梵家具

   [4]棠语家具                   [5]凌轩家具                   [6]通用长图`
		fmt.Println(text)

		key := inputString("\n:: 请选择需要使用的功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/yinanjj.png", viper.GetBool("automaticDeletion"))
		case "2":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/yushantanjj.png", viper.GetBool("automaticDeletion"))
		case "3":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/sufanjj.png", viper.GetBool("automaticDeletion"))
		case "4":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/tangyujj.png", viper.GetBool("automaticDeletion"))
		case "5":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/lingxuan.png", viper.GetBool("automaticDeletion"))
		case "6":
			presenter.UniversalMainImage(viper.GetString("picture"), 1125, 1500, 4, viper.GetBool("automaticDeletion")) // 通用长图
		case "7":

		case "8":

		case "9":

		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 通用主图选择
func (c *CLI) universalMainImageChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Universal Main Image choice", 74)
		text := `
:: 通用主图不会叠加任何水印，并且根据下方的裁剪模式，最终裁剪生成正方形图！

   [1]智能裁剪                   [2]居中裁剪                   [3]居上裁剪

   [4]居下裁剪                   [5]居左裁剪                   [6]居右裁剪`
		fmt.Println(text)

		key := inputString("\n:: 请选择需要使用的功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 1, viper.GetBool("automaticDeletion")) // 通用主图
			return
		case "2":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 2, viper.GetBool("automaticDeletion")) // 通用主图
		case "3":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 3, viper.GetBool("automaticDeletion")) // 通用主图
		case "4":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 4, viper.GetBool("automaticDeletion")) // 通用主图
		case "5":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 5, viper.GetBool("automaticDeletion")) // 通用主图
		case "6":
			presenter.UniversalMainImage(viper.GetString("picture"), 800, 800, 6, viper.GetBool("automaticDeletion")) // 通用主图
		case "7":

		case "8":

		case "9":

		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 屏风主图选择
func (c *CLI) screenMainPictureChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Screen main picture choice", 74)
		text := `
:: 屏风店主图会自动识别白底图，并排除白底图，为其他图片加上专属水印与Logo！

   [1]沐兰主图                   [2]华府主图                   [3]木韵主图

   [4]御尚屏风                   [5]金尊主图                   [6]暂未开发`
		fmt.Println(text)

		key := inputString("\n:: 请选择需要使用的功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/mulan.png", viper.GetBool("automaticDeletion"))
		case "2":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/huafu.png", viper.GetBool("automaticDeletion"))
		case "3":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/muyunge.png", viper.GetBool("automaticDeletion"))
		case "4":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/yushantanpf.png", viper.GetBool("automaticDeletion"))
		case "5":
			presenter.WatermarkMainImage(viper.GetString("picture"), "data/static/img/jinzunfu.png", viper.GetBool("automaticDeletion"))
		case "6":
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 修改图片分辨率的选择
func (c *CLI) modifyResolution() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Modify resolution", 74)
		text := `
:: 在保持原来长宽像素不变的情况下，将分辨率强行修改至指定数值，请尽量少用！

   [1]全部改为72PPI         [2]全部改为300PPI         [3]全部改为自定义PPI.`
		fmt.Println(text)

		key := inputString("\n:: 请选择需要使用的功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			unclassified.PixelsPerInchChangedTo72(viper.GetString("picture")) // 改为72ppi
		case "2":
			unclassified.PixelsPerInchChangedTo300(viper.GetString("picture"))
		case "3":
			fmt.Println("\n:: 因需求不高，所以暂未开发...")
		case "-", "--":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}

// 套图的选择
func (c *CLI) autoPictureChoice() {
OuterLoop:
	for {
		// 先显示通知
		c.showNotice(false)

		tools.EnglishTitle("Auto picture", 74)
		text := `
:: 自动套图可以正确运行的前提是，款式图片均要放在套图文件夹下方可正确运行！

   [1]套图文件夹                [2]随机重命名                [3]备份文件夹

   [4]家具店主图.               [5]通用版主图.               [6]屏风店主图.

   [7]修改分辨率.               [8]替换详情页                [9]导出详情页`
		fmt.Println(text)

		key := inputString("\n:: 请选择需要使用的功能：") // 获取键盘输入
		tools.CallClear()                      // 清屏

		// 如果是暗号就打印暗号传回来的提示
		var ok bool
		if ok, c.info = presenter.SelectCommand(key); ok {
			continue
		}

		switch key {
		case "1":
			presenter.Command41()
			c.info = ":: 已打开套图文件夹，请复制正方形的 jpg 或 png 高清图片以备自动套图使用！"
		case "2":
			presenter.RandomRenameFile(viper.GetString("picture"))
		case "3":
			go exec.Command("cmd.exe", "/c", "start Config\\Backups").Run()
			c.info = ":: 已打开备份文件夹，为了避免意外丢失文件，目前备份文件最大为 15 份！"
		case "4":
			c.furnitureMainPictureChoice() // 家具店主图
		case "5":
			c.universalMainImageChoice() // 通用主图
		case "6":
			c.screenMainPictureChoice() // 屏风店主图
		case "7":
			c.modifyResolution() // 修改分辨率
		case "8":
			presenter.ReplaceSmartObjects(viper.GetString("picture")) // 替换详情页
		case "9":
			model.RunSaveForWeb() // 导出web格式详情页
		case "-":
			break OuterLoop
		case "":
			c.info = ":: 输入的内容为空，请重新输入..."
			continue
		default:
			c.info = fmt.Sprintf(":: 输入的 [%s] 不是已知的功能选项，请重新输入...", tools.ColourString(key, ctc.ForegroundGreen))
		}
	}
}
