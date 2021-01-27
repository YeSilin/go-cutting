package nested

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// 获取指定目录下的所有jpg图片，并且排除白底图
func getAllImages(path string) (images []string, err error) {
	// 返回绝对路径
	path, err = filepath.Abs(path)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 全部换成正斜杠，这里只是修改输入的路径
	path = strings.Replace(path, "\\", "/", -1)

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	images, err = filepath.Glob(fmt.Sprintf("%s/*.jpg", path))
	if err != nil {
		logrus.Error(err)
		return
	}
	//fmt.Println(images)
	// 如果jpg小于一张就不执行
	if len(images) < 1 {
		fmt.Println("\n:: 脚本注入失败，因为 Picture 文件夹下没有 jpg 格式图片！")
		err = fmt.Errorf("文件夹下没有 jpg 格式图片")
		// 打开套图文件夹
		exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s", viper.GetString("picture"))).Run()
		return
	}

	// 获取白底图
	minImage, exist := tools.MinWhiteBackground(fmt.Sprintf("%s/*.jpg", path))
	// 如果没有白底图就直接返回已经收集到的图片切片
	if !exist {
		return
	}
	//fmt.Println(minImage)
	// 索引计数
	index := 0
	for i := 0; i < len(images); i++ {
		// 如果是白底图就忽略
		if images[i] == minImage {
			continue
		}
		images[index] = images[i]
		index++
	}

	// 对切片进行截取
	images = images[:index]
	return
}

// 拆分细节图片，将一个切片分成两个切片
func splitDetails(images []string) (dp, de []string) {
	// 先分配内存空间
	dp = make([]string, len(images))
	// de 是细节
	de = make([]string, len(images))

	// 先设置两个的索引
	dpi := 0
	dei := 0

	for i := 0; i < len(images); i++ {
		// 先全部转换成小写
		temp := strings.ToLower(images[i])
		// 返回路径的最后一个元素
		temp = filepath.Base(temp)

		// 是否包含de前缀
		if strings.HasPrefix(temp, "de") {
			de[dei] = images[i]
			dei++
			continue
		}
		// 其他全部算普通
		dp[dpi] = images[i]
		dpi++
	}
	// 最后截取
	dp = dp[:dpi]
	de = de[:dei]
	return
}

// 生成详情页替换智能对象的脚本
func ReplaceDetailsPage(path string) {
	// 获取指定目录下的所有jpg图片
	images, err := getAllImages(path)
	if err != nil {
		return
	}

	go func() {

		// 修改成js脚本可以看懂的路径
		for i := 0; i < len(images); i++ {
			images[i] = strings.Replace(images[i], "\\", "/", -1)
			images[i] = strings.Replace(images[i], ":", "", 1)
			images[i] = "/" + images[i]
		}

		// 将一个切片分成两个切片
		dp, de := splitDetails(images)

		// 定义一个匿名结构体，给模板使用，属性必须大写，不然无权调用
		info := struct {
			DPStr string
			DEStr string
		}{tools.StrToJsArray("dpArray", dp), tools.StrToJsArray("deArray", de)}

		// 解析指定文件生成模板对象
		tmpl, err := template.ParseFiles("config/jsx/template/replaceDetailsPage.gohtml")
		if err != nil {
			logrus.Error(err)
			return
		}

		// 创建文件，返回两个值，一是创建的文件，二是错误信息
		f, err := os.Create("config/jsx/replaceDetailsPage.jsx")
		if err != nil { // 如果有错误，打印错误，同时返回
			logrus.Error(err)
			return
		}

		// 利用给定数据渲染模板，并将结果写入f
		_ = tmpl.Execute(f, info)

		// 关闭文件
		f.Close()

		// 创建一个协程使用cmd来运行脚本
		dataPath := "Config/jsx/ReplaceDetailsPage.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

	}()
	fmt.Println("\n:: 脚本注入成功，正在自动替换详情页中名字以 [DP] 开头的智能对象图层！")
}



// 导出web格式详情页
func SaveForWebDetailsPage() {
	go func() {
		// 自动套图工作路径
		picturePath := viper.GetString("picture")
		// 创建套图文件夹
		_ = tools.CreateMkdirAll(fmt.Sprintf("%s/主图", picturePath))
		// 创建一个协程使用cmd来运行脚本
		dataPath := "Config/JSX/SaveForWeb.jsx"
		exec.Command("cmd.exe", "/c", "start "+dataPath).Run()

		time.Sleep(time.Second) // 停一秒
		// 如果存在images就打开
		if ok, _ := tools.IsPathExists(fmt.Sprintf("%s/主图/images", picturePath)); ok {
			exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图\\images", picturePath)).Run()
		} else {
			exec.Command("cmd.exe", "/c", fmt.Sprintf("start %s\\主图", picturePath)).Run()
		}
	}()
}
