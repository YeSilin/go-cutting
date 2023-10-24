package presenter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yesilin/go-cutting/model"
	"github.com/yesilin/go-cutting/tools"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
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

// ReplaceSmartObjects 生成详情页替换智能对象的脚本
func ReplaceSmartObjects(path string) {
	const script = `// 生成进度条函数
function progressBar() {
    // 进度条调用清除元数据函数
    app.doForcedProgress("正在替换智能对象... ", "replaceSmartObjects(doc)")  // 添加进度条
}


// 替换智能对象
function replaceContents(newFile, theSO) {
    app.activeDocument.activeLayer = theSO;
    // =======================================================
    var idplacedLayerReplaceContents = stringIDToTypeID("placedLayerReplaceContents");
    var desc3 = new ActionDescriptor();
    var idnull = charIDToTypeID("null");
    desc3.putPath(idnull, new File(newFile));
    var idPgNm = charIDToTypeID("PgNm");
    desc3.putInteger(idPgNm, 1);
    executeAction(idplacedLayerReplaceContents, desc3, DialogModes.NO);
    return app.activeDocument.activeLayer
}

// 判断当前图层名字是DE开头还是DP开头
function judgeDeOrDp(layerName) {
    // 得到图层名字后，统一转成大写
    var upperCase = layerName.toUpperCase();

    // 再判断字符串是否以 DE 或DP 开头
    var isDE = upperCase.indexOf("DE")
    if (isDE == 0) { // 表示t图层名字是以DE开头； 
        return "DE"
    }

    var isDP = upperCase.indexOf("DP")
    if (isDP == 0) { // 表示t图层名字是以DP开头；
        return "DP"
    }
    return "NO"
}

// 替换图层时智能切换资源，选择优先消耗的资源
function resourcesSwitchDE(curLayer) {
    // 替换之前看下资源是否已经用完，相等表示用完了
    if (de == deArray.length) {
        // 替换智能对象，因为用完了所以使用dp补
        replaceContents(new File(dpArray[dp]), curLayer);
        curLayer.name = "DE 细节图层(不够用替补)";
        // 表示有dp图层
        isDPDE = true;
        dp++; // 索引加1
        countBar++; // 进度条加1
        // 更新进度条
        updateProgress(countBar, maxBar);
        return
    }

    // 替换智能对象
    replaceContents(new File(deArray[de]), curLayer);
    curLayer.name = "DE 细节图层";
    // 表示有dp图层
    isDPDE = true;
    de++; // 索引加1
    countBar++; // 进度条加1
    // 更新进度条
    updateProgress(countBar, maxBar);
}

// 替换图层时智能切换资源，选择优先消耗的资源
function resourcesSwitchDP(curLayer) {
    // 替换之前看下资源是否已经用完，相等表示用完了
    if (dp == dpArray.length) {
        // 替换智能对象
        replaceContents(new File(deArray[de]), curLayer);
        curLayer.name = "DP 普通图层(不够用替补)";
        // 表示有dp图层
        isDPDE = true;
        de++; // 索引加1
        countBar++; // 进度条加1
        // 更新进度条
        updateProgress(countBar, maxBar);
        return
    }

    replaceContents(new File(dpArray[dp]), curLayer);
    curLayer.name = "DP 普通图层";
    // 表示有dp图层
    isDPDE = true;
    dp++; // 索引加1
    countBar++; // 进度条加1
    // 更新进度条
    updateProgress(countBar, maxBar);
}


// 改正图层到指定边界大小
function sameSize(modifyLayer, boundsRef) {
    // 求出原本的图层边界
    var boundsCur = modifyLayer.boundsNoEffects

    // 如果相等直接返回
    if (boundsRef.toString() == boundsCur.toString()) {
        return
    }

    // 求出彼此的宽度和高度
    var widthRef = boundsRef[2] - boundsRef[0]
    var heightRef = boundsRef[3] - boundsRef[1]
    var widthCur = boundsCur[2] - boundsCur[0]
    var heightCur = boundsCur[3] - boundsCur[1]

    // 求出百分比
    widthPercentage = widthRef / widthCur * 100
    heightPercentage = heightRef / heightCur * 100

    // 重设大小
    modifyLayer.resize(widthPercentage, heightPercentage);
}


// 递归函数
function replaceSmartObjects(doc) {
    // 这里倒序，因为最上层的图层是0
    for (var i = doc.layers.length - 1; i >= 0; i--) {
        // 如果索引大于源文件数量就不再替换
        if (countBar == maxBar) {
            return
        }

        var curLayer = doc.layers[i];

        // 如果选中的图层类型不是普通图层，即代表是图层组
        if (curLayer.typename != "ArtLayer") {
            replaceSmartObjects(curLayer);
            continue;
        }
        // 如果当前图层是智能对象
        if (curLayer.kind == "LayerKind.SMARTOBJECT") {
            // 得到前面的头
            var frontName = judgeDeOrDp(curLayer.name)

            // 没有指定的头
            if (frontName == "NO") {
                continue
            }

            // 获得参考边界
            var bounds = curLayer.boundsNoEffects

            // 表示是以DE开头；
            if (frontName == "DE") {
                // 替换的时候智能切换
                resourcesSwitchDE(curLayer)
                // 如果边界大小不对就改成一样大
                sameSize(curLayer, bounds)
                continue
            }
            // 表示是以DP开头；
            if (frontName == "DP") {
                // 替换的时候智能切换
                resourcesSwitchDP(curLayer)
                // 如果边界大小不对就改成一样大
                sameSize(curLayer, bounds)
                continue
            }

        }
    }
}


// 没有打开的文档直接略过
if (!documents.length) {
    // alert("没有打开的文档，请打开一个文档来运行此脚本！");
} else {
    try {
        var doc = app.activeDocument;

        // dp源文件
        {{.DPStr}}  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

        // de源文件
        {{.DEStr}}  // 这里传go模板语句！！！！！！！！！！！！！！！！！！！！！

        // 这是源文件索引
        var dp = 0;
        var de = 0;

        // 进度条最大进度
        const maxBar = dpArray.length + deArray.length;
        // 进度条计数
        var countBar = 0

        // 默认当做没有dp de智能对象
        var isDPDE = false;

        // 运行带进度条的图层替换
        progressBar();

        if (!isDPDE) {
            alert("此文档没有以 DP 或 DE 开头命名的智能对象图层！");
        }
    } catch (error) {
        //发生错误执行的代码
    }
}`

	// 获取指定目录下的所有jpg图片
	images, err := getAllImages(path)
	if err != nil {
		return
	}

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

	// 解析字符串生成模板对象
	tmpl, err := template.New("tmpl").Parse(script)
	if err != nil {
		logrus.Error(err)
		return
	}

	// 创建文件，返回两个值，一是创建的文件，二是错误信息
	f, err := os.Create("data/jsx/replaceSmartObjects.jsx")
	if err != nil { // 如果有错误，打印错误，同时返回
		logrus.Error(err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入f
	err = tmpl.Execute(f, info)
	if err != nil {
		logrus.Error(err)
	}

	// 关闭文件不然无法调用脚本
	f.Close()

	// 运行脚本
	model.RunReplaceSmartObjects()
}
