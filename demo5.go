package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/yesilin/go-cutting/clib"
	"os"
	"time"
)

func main5() {
	// 输出初始画面
	//36*
	_, _ = fmt.Fprintln(os.Stderr,
		`
 GoCutting v1.0.38，试用期为 29 天！
------------------------------------
 切图 贴图 效果 套图 附加 设置 帮助


      欢迎使用命令行切图软件，

       请使用左右方向键控制！






------------------------------------
`)

	//clib.GotoPostion(3,4)
	//color.LightCyan.Println("【请尽量不关闭软件避免断网时无法使用！")

	tutu := 1

	clib.GotoPostion(1, 3)
	color.LightCyan.Println("切图")

	func() {
		for {
			switch clib.Direction() {
			//根据键盘输入设置方向
			//上
			case 72, 87, 119:

				//下
			case 80, 83, 115:
				return
				//左
			case 65, 75, 97:
				if tutu > 0 {
					tutu--
				}
				switch tutu {
				case 1:
					clib.GotoPostion(6, 3)
					fmt.Println("贴图")

					clib.GotoPostion(1, 3)
					color.LightCyan.Println("切图")

				case 2:
					clib.GotoPostion(11, 3)
					fmt.Println("效果")

					clib.GotoPostion(6, 3)
					color.LightCyan.Println("贴图")

				case 3:
					clib.GotoPostion(16, 3)
					fmt.Println("套图")

					clib.GotoPostion(11, 3)
					color.LightCyan.Println("效果")
				case 4:
					clib.GotoPostion(21, 3)
					fmt.Println("附加")

					clib.GotoPostion(16, 3)
					color.LightCyan.Println("套图")
				case 5:
					clib.GotoPostion(26, 3)
					fmt.Println("设置")

					clib.GotoPostion(21, 3)
					color.LightCyan.Println("附加")
				case 6:
					clib.GotoPostion(31, 3)
					fmt.Println("帮助")

					clib.GotoPostion(26, 3)
					color.LightCyan.Println("设置")
				}


			// 右
			case 68, 77, 100:
				if tutu < 7 {
					tutu++
				}
				switch tutu {
				case 1:
					clib.GotoPostion(1, 3)
					color.LightCyan.Println("切图")

				case 2:
					clib.GotoPostion(1, 3)
					fmt.Println("切图")

					clib.GotoPostion(6, 3)
					color.LightCyan.Println("贴图")

				case 3:
					clib.GotoPostion(6, 3)
					fmt.Println("贴图")

					clib.GotoPostion(11, 3)
					color.LightCyan.Println("效果")
				case 4:
					clib.GotoPostion(11, 3)
					fmt.Println("效果")

					clib.GotoPostion(16, 3)
					color.LightCyan.Println("套图")
				case 5:
					clib.GotoPostion(16, 3)
					fmt.Println("套图")

					clib.GotoPostion(21, 3)
					color.LightCyan.Println("附加")
				case 6:
					clib.GotoPostion(21, 3)
					fmt.Println("附加")

					clib.GotoPostion(26, 3)
					color.LightCyan.Println("设置")
				case 7:
					clib.GotoPostion(26, 3)
					fmt.Println("设置")

					clib.GotoPostion(31, 3)
					color.LightCyan.Println("帮助")
				}
			}
		}
	}()

	clib.GotoPostion(6, 12)
	q := "0"
	fmt.Scan(&q)
	time.Sleep(100 * time.Second)
}
