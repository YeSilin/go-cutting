package quickCipher

import (
	"fmt"
	"github.com/yesilin/go-cutting/tools"
	"os/exec"
	"time"
)

func Work() {
	//获取当前时间，进行格式化 2006-01-02 15:04:05
	now:= time.Now().Format("2006-01-02")

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
