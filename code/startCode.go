// 启动暗号的函数汇总
package code

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/autoPicture"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)

// 启动暗号-1
func StartCode1() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/SelectTailor.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	// 每次选择正确的脚本时删除多余备份，最大保留30个
	go tools.DeleteRedundantBackups("Config/JSX/Temp/*", 30)
}

// 启动暗号-2 重建新文档
func StartCode2() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/newDocument.jsx"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}

// 启动暗号-3 深度清除源数据
func StartCode3() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "config/jsx/clearMetadata.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 启动暗号-4 快捷新建当天工作目录
func StartCode4() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now := time.Now().Format("2006-01-02")

	// 高老板
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/全镂空/半透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/全镂空/不透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/无镂空/半透", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/旧厂切图/%s/无镂空/不透", now))

	// 这里的
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/御尚檀", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/御尚檀", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/岚湘", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/岚湘", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/沐兰", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/沐兰", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/华府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/华府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/木韵阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/木韵阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/金樽府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/金樽府", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/怡柟", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/怡柟", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/舍得", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/舍得", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/西厢", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/西厢", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/藏湘阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/藏湘阁", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/阑若", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/阑若", now))

	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/木墨", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/木墨", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/半透/墨屏", now))
	_ = tools.CreateMkdirAll(fmt.Sprintf("D:/切图（请移动至个人目录）/%s/不透/墨屏", now))

	cmd := exec.Command("cmd.exe", "/c", "start D:\\切图（请移动至个人目录）")
	cmd.Run()
}

// 启动暗号-7 为当前文档添加黑边
func StartCode7() {
	// 创建一个协程使用cmd来运行脚本
	dataPath := "config/jsx/blackEdge.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 启动暗号-6 简单清除元数据
func StartCode6() {
	dataPath := "config/jsx/clearMetadataNoPopUp.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 启动暗号-9 查询历史记录文件
func StartCode9() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	fileName := time.Now().Format("2006-01-02")
	now := time.Now().Format("2006-01")

	// 储存历史记录路径
	path := fmt.Sprintf("config/History/%s/%s.txt", now, fileName)

	// 先查看是否有历史记录文件
	exists, _ := tools.IsPathExists(path)
	// 如果找不到文件，就创建文件 头
	if !exists {
		//fmt.Println("\n【错误】找不到今天的切图历史记录，可能今天还未开始切图，已自动打开历史文件夹！")

		exec.Command("cmd.exe", "/c", "start config\\History").Run()
		//tools.PrintLine(2)
		return
	}
	// 创建一个协程使用cmd来运行脚本
	cmd := exec.Command("cmd.exe", "/c", "start "+path)
	go cmd.Run()
}

// 启动暗号-10 快捷另存为jpg
func StartCode10() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/SaveAsJPEG.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}

// 启动暗号-11 快捷另存全部jpg
func StartCode11() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/JSX/saveAllJPEG.jsx"
	go exec.Command("cmd.exe", "/c", "start "+dataPath).Run()
}



// 启动暗号-97
func StartCode97() {
	autoPicture.ReplaceDetailsPage(viper.GetString("picture")) // 替换详情页
}

// 启动暗号-98
func StartCode98() {
	autoPicture.SaveForWebDetailsPage()
}

// 启动暗号-99
func StartCode99() {
	// 创建一个协程使用cmd启动外部程序
	dataPath := "Config/exe/W10DigitalActivation.exe /activate"
	cmd := exec.Command("cmd.exe", "/c", "start "+dataPath)
	go cmd.Run()
}
