// 此模块专门用来限制软件的使用
package model

import (
	"encoding/json"
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os"
)

// 创建存放使用限制的私密文件夹
func createPrivateMkdir() {
	// 先获取用户主目录
	privatePath, _ := tools.Home()
	// 在主目录后面补上自定义目录
	privatePath = fmt.Sprintf("%s\\Documents\\Adobe", privatePath)
	//fmt.Println(privatePath)
	// 创建私密目录
	tools.CreateMkdirAll(privatePath)
}

/////////////////////////////////////// 旧版本验证
// 验证版本的函数 传入私密路径、当前版本、当前从网络获取的时间、限制最长使用时间
func RestrictingSoftwareUse(path string, version float64, time, expire int64) (bool, string) {
	// 定义结构体
	type VersionDate struct {
		Version float64
		Date    int64
	}

	//预定义解码结果
	var oldSoftware VersionDate

	// 预定义编码信息
	var newSoftware = VersionDate{version, time}

	// 获取 expire 天前的时间戳，+1 是为了往后推一秒
	minTime := tools.AroundTime(fmt.Sprintf("-%d", expire)) + 1

FLAG: // 这个定位给解码失败用的

	// 先判断文件是否存在
	f, _ := tools.IsPathExists(path)
	if !f { // 如果没有文件
		// 先创建私密文件夹
		createPrivateMkdir()

		//创建文件（并打开）
		filePtr, err := os.Create(path)
		if err != nil {
			//fmt.Println("创建文件失败，err=", err)
			return false, ""
		}
		defer filePtr.Close()

		//创建基于文件的JSON编码器
		encoder := json.NewEncoder(filePtr)

		//将新版本信息实例编码到文件中
		err = encoder.Encode(newSoftware)
		if err != nil {
			//fmt.Println("编码失败，err=", err)
			return false, ""
		} else {
			//fmt.Println("编码成功")
			// 求出剩余时间
			day := tools.ToDay(newSoftware.Date - minTime)
			return true, fmt.Sprintf("\n【验证】欢迎首次使用本切图软件，输入数字回车运行，此版本剩余可使用时间为 %d 天！", day)
		}
	} else {
		//打开文件
		filePtr, _ := os.Open(path)
		defer filePtr.Close()

		//创建该文件的json解码器
		decoder := json.NewDecoder(filePtr)
		//把解码的结果存在software的地址中
		err := decoder.Decode(&oldSoftware)
		if err != nil {
			//fmt.Println("解码失败1，err=", err)
			_ = filePtr.Close()    // 提前关闭文件
			del := os.Remove(path) // 删除文件
			if del != nil {
				//fmt.Println(del)
			}
			goto FLAG // 跳到重新判断是否有配置文件
			//return false

		} else if oldSoftware.Version == version { // 先判断是不是一个版本
			//fmt.Printf("解码成功:%#v\n", software)
			if oldSoftware.Date > minTime {
				// 求出剩余时间
				day := tools.ToDay(oldSoftware.Date - minTime)
				return true, fmt.Sprintf("\n【验证】欢迎再次使用本切图软件，输入数字回车运行，此版本剩余可使用时间为 %d 天！", day)

			} else {
				return false, "\n【验证】很遗憾此版本已过期，请在切图软件问题反馈群里，下载并安装最新版切图软件！"
			}

			// 文件版本小于软件版本
		} else if oldSoftware.Version < version {
			//创建文件（并打开）
			filePtr, err := os.Create(path)
			if err != nil {
				//fmt.Println("创建文件失败，err=", err)
				return false, ""
			}
			defer filePtr.Close()

			//创建基于文件的JSON编码器
			encoder := json.NewEncoder(filePtr)
			//将新版本信息实例编码到文件中
			err = encoder.Encode(newSoftware)
			if err != nil {
				//fmt.Println("编码失败，err=", err)
				return false, ""
			} else {
				//fmt.Println("编码成功")
				//fmt.Println("\n【验证】新版本已更新成功，欢迎使用最新版切图软件，此版本剩余可使用时间为 21 天！")
				// 求出剩余时间
				day := tools.ToDay(newSoftware.Date - minTime)
				return true, fmt.Sprintf("\n【验证】新版本已更新成功，欢迎使用最新版切图软件，此版本剩余可使用时间为 %d 天！", day)
			}

			// 如果版本直接相差大于5个小版本，就不让使用
		} else if (oldSoftware.Version - version) > 0.000005 { // 文件的版本大于当前版本说明新版本已存在
			return false, "\n【验证】请继续使用新版本！此旧版本已被强制停用，若有出现Bug请及时在群里反馈！"

			// 如果版本直接相差小于5个小版本，就不让使用
		} else if (oldSoftware.Version - version) < 0.000005 {
			if oldSoftware.Date > minTime {
				// 求出剩余时间
				day := tools.ToDay(oldSoftware.Date - minTime)
				return true, fmt.Sprintf("\n【验证】欢迎再次使用本切图软件，建议更新至最新版，此版本剩余可使用时间为 %d 天！", day)

			} else {
				return false, "\n【验证】很遗憾此版本已过期，请在切图软件问题反馈群里，下载并安装最新版切图软件！"
			}

		} else { // 文件的版本小于当前版本 那么就修改版本号
			//创建文件（并打开）
			filePtr, err := os.Create(path)
			if err != nil {
				//fmt.Println("创建文件失败，err=", err)
				return false, ""
			}
			defer filePtr.Close()

			//创建基于文件的JSON编码器
			encoder := json.NewEncoder(filePtr)

			//将新版本信息实例编码到文件中
			err = encoder.Encode(newSoftware)
			if err != nil {
				//fmt.Println("编码失败，err=", err)
				return false, ""
			} else {
				//fmt.Println("编码成功")
				//fmt.Println("\n【验证】新版本已更新成功，欢迎使用最新版切图软件，此版本剩余可使用时间为 21 天！")
				// 求出剩余时间
				day := tools.ToDay(newSoftware.Date - minTime)
				return true, fmt.Sprintf("\n【验证】新版本已更新成功，欢迎使用最新版切图软件，此版本剩余可使用时间为 %d 天！", day)
			}
		}
	}
}
/////////////////////////////////////// 旧版本验证

// 定义结构体
type versionDate struct {
	Version float64
	Date    int64
}

// 定义版本编码器
func versionEncode(path string, newSoftware versionDate) (err error) {
	// 创建文件（并打开）
	filePtr, err := os.Create(path)
	if err != nil {
		//fmt.Println("创建文件失败，err=", err)
		return err
	}
	defer filePtr.Close()

	// 创建基于文件的JSON编码器
	encoder := json.NewEncoder(filePtr)

	// 将新版本信息实例编码到文件中
	err2 := encoder.Encode(newSoftware)
	if err2 != nil {
		// 编码失败
		fmt.Println(err2)
		return err2
	}
	return err
}

// 定义版本解码器
func versionDecode(path string) (oldSoftware versionDate, err error) {
	// 先判断文件是否存在
	f, err := tools.IsPathExists(path)
	// 如果没有文件
	if !f {
		return oldSoftware, err
	}

	// 打开文件获得文件指针
	filePtr, err2 := os.Open(path)
	if err2 != nil {
		return oldSoftware, err2
	}
	defer filePtr.Close()

	// 创建基于文件的JSON解码器
	decoder := json.NewDecoder(filePtr)

	// 把解码的结果存在software的地址中
	err3 := decoder.Decode(&oldSoftware)
	if err3 != nil { // 解码失败
		return oldSoftware, err3
	}
	return
}

// 验证版本的函数 传入私密路径、当前版本、当前从网络获取的时间、限制最长使用时间
func RestrictingSoftwareUse2(path string, version float64, time, expire int64) (bool, string) {
	// 预定义编码信息
	var newSoftware = versionDate{version, time}

	// 获取 expire 天前的时间戳，+1 是为了往后推一秒
	var minTime = tools.AroundTime(fmt.Sprintf("-%d", expire)) + 1

	// 先判断文件是否存在
	f, _ := tools.IsPathExists(path)
	// 如果没有文件
	if !f {
		// 创建私密文件夹
		createPrivateMkdir()
		// 编码新的信息
		_ = versionEncode(path, newSoftware)
		// 求出剩余时间
		day := tools.ToDay(newSoftware.Date - minTime)
		return true, fmt.Sprintf("\n【验证】欢迎首次使用本切图软件，输入数字回车运行，此版本剩余可使用时间为 %d 天！", day)
	}

	// 定义解码结果，把解码的结果存在software的地址中
	oldSoftware, err := versionDecode(path)
	// 如果解码出错，直接删了文件算了
	if err != nil {
		del := os.Remove(path) // 删除文件
		if del != nil {
			//fmt.Println(del)
		}
		// 然后重新编码新的信息
		_ = versionEncode(path, newSoftware)
		// 求出剩余时间
		day := tools.ToDay(newSoftware.Date - minTime)
		return true, fmt.Sprintf("\n【验证】欢迎首次使用本切图软件，输入数字回车运行，此版本剩余可使用时间为 %d 天！", day)
	}

	// 到这一步说明文件已经解码成功，优先对比版本号了
	var versionCompare float64 = oldSoftware.Version - version // 求出文件版本比当前版本相差多少
	var dateCompare bool = oldSoftware.Date > minTime          // true 表示没有过期

	// 前面已经把重要的变量都求出来了，现在开始对比（两个版本相同）
	if versionCompare == 0 {
		if dateCompare {
			// 求出剩余时间
			day := tools.ToDay(oldSoftware.Date - minTime)
			return true, fmt.Sprintf("\n【验证】欢迎再次使用本切图软件，输入数字回车运行，此版本剩余可使用时间为 %d 天！", day)
		} else {
			return false, "\n【验证】很遗憾此版本已过期，请在切图软件问题反馈群里，下载并安装最新版切图软件！"
		}
	}

	// 文件版本小于软件版本，文件版本 - 当前版本 = 小于0（安装了新版本）
	if versionCompare < 0 {
		// 将新版本信息实例编码到文件中
		_ = versionEncode(path, newSoftware)
		// 求出剩余时间
		day := tools.ToDay(newSoftware.Date - minTime)
		return true, fmt.Sprintf("\n【验证】新版本已更新成功，欢迎使用最新版切图软件，此版本剩余可使用时间为 %d 天！", day)
	}

	// 如果版本直接相差大于5个小版本，就不让使用（安装了新版，但又重新安装旧版）
	if versionCompare > 0.000005 { // 文件的版本大于当前版本说明新版本已存在
		return false, "\n【验证】请继续使用新版本！此旧版本已被强制停用，若有出现Bug请及时在群里反馈！"
	} else { // 如果版本直接相差小于或等于5个小版本，可以使用，但还是会查水表 （安装了新版，但又重新安装旧版）
		if dateCompare {
			// 求出剩余时间
			day := tools.ToDay(oldSoftware.Date - minTime)
			return true, fmt.Sprintf("\n【验证】欢迎再次使用本切图软件，建议更新至最新版，此版本剩余可使用时间为 %d 天！", day)
		} else {
			return false, "\n【验证】很遗憾此版本已过期，请在切图软件问题反馈群里，下载并安装最新版切图软件！"
		}
	}
}
