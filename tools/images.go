package tools

import (
	"fmt"
	"github.com/cenkalti/dominantcolor"
	"github.com/disintegration/imaging"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/noelyahan/mergi"
	"image"
	"log"
	"os"
	"path/filepath"
)

// 调整图像大小，源路径 和 保存的目标路径，如果宽度或高度之一为0，则图像比率保持不变。
func ImageResize(srcPath, dstPath string, width, height,quality int) {
	// 打开原始图像
	src, err := imaging.Open(srcPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// 获取源图片的长度和宽度
	b := src.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y

	// 如果裁剪为正方形，并且源图像大小又不是正方形时，才使用内容感知自动裁剪
	if width == height && srcWidth != srcHeight {
		// 创建解析器
		analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
		// 裁剪图像长宽比为1500-1500
		topCrop, _ := analyzer.FindBestCrop(src, 1500, 1500)

		// 更新至结构体
		type SubImager interface {
			SubImage(r image.Rectangle) image.Image
		}
		src = src.(SubImager).SubImage(topCrop)
	}

	// 将裁切后的图像调整为 ...，以保留宽高比。imaging.Lanczos 是最高清的
	dst := imaging.Resize(src, width, height, imaging.Lanczos)

	// 保存图片，扩展名可以不一致，质量99，避免超500k
	err = imaging.Save(dst, dstPath, imaging.JPEGQuality(quality))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

// 给图片加水印，保存路径，源图像，水印图像，左上角的定位 x,y，
func ImageWatermark(savePath, originalPath, watermarkPath string, x, y int) {
	// 打开源图像
	originalImage, err1 := imaging.Open(originalPath)
	if err1 != nil {
		log.Fatalf("failed to open image: %v", err1)
	}

	// 打开水印图像
	watermarkImage, err2 := imaging.Open(watermarkPath)
	if err2 != nil {
		log.Fatalf("failed to open image: %v", err2)
	}

	resultImage, _ := mergi.Watermark(watermarkImage, originalImage, image.Pt(x, y))

	// Save the resulting image as JPEG.
	err3 := imaging.Save(resultImage, savePath, imaging.JPEGQuality(99))
	if err3 != nil {
		log.Fatalf("failed to save image: %v", err3)
	}
}

// 修改图片大小并且加水印；
// savePath 保存路径；
// originalPath 源图像路径，宽和高；
// watermarkPath 水印图像路径，左上角的定位 x,y，
func ImageResizeAndWatermark(savePath, originalPath string, width, height int, watermarkPath string, x, y int) {
	// 打开源图像
	originalImage, err1 := imaging.Open(originalPath)
	if err1 != nil {
		log.Fatalf("failed to open image: %v", err1)
	}

	// 获取源图片的长度和宽度
	b := originalImage.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y

	// 如果裁剪为正方形，并且源图像大小又不是正方形时，才使用内容感知自动裁剪
	if width == height && srcWidth != srcHeight {
		// 创建解析器
		analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
		// 裁剪图像长宽比为1比1
		topCrop, _ := analyzer.FindBestCrop(originalImage, 1, 1)

		// 更新至结构体
		type SubImager interface {
			SubImage(r image.Rectangle) image.Image
		}
		originalImage = originalImage.(SubImager).SubImage(topCrop)
	}

	// 修改源图像的大小。imaging.Lanczos 是最高清的
	originalImage = imaging.Resize(originalImage, width, height, imaging.Lanczos)

	// 打开水印图像
	watermarkImage, err2 := imaging.Open(watermarkPath)
	if err2 != nil {
		log.Fatalf("failed to open image: %v", err2)
	}

	// 将水印图像叠加在源图像上，水印在左，源在右，pt调整的是水印的位置
	resultImage, _ := mergi.Watermark(watermarkImage, originalImage, image.Pt(x, y))

	// 保存图片，根据文件扩展名保存
	err3 := imaging.Save(resultImage, savePath, imaging.JPEGQuality(99))
	if err3 != nil {
		log.Fatalf("failed to save image: %v", err3)
	}
}

// 检测这张图片的主要颜色
func FindDomiantColor(fileInput string) (string, error) {
	f, err := os.Open(fileInput)
	defer f.Close()
	if err != nil {
		fmt.Println("File not found:", fileInput)
		return "", err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}



	return dominantcolor.Hex(dominantcolor.Find(img)), nil
}


// 检测单张 是否是白底图，返回 图像大小
func IsWhiteBackground(fileInput string) (bool, int64, error) {
	fInfo, err := os.Stat(fileInput)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return false, 0, err
	}

	// 获取文件大小
	fSize := fInfo.Size()

	// 打开原始图像
	img, err := imaging.Open(fileInput)
	if err != nil {
		fmt.Println("imaging.Open err:", err)
		return false, fSize, err
	}

	// 获取图片的长度和宽度
	bou := img.Bounds()
	imgWidth := bou.Max.X
	imgHeight := bou.Max.Y

	// 封装一个临时函数，用于判断是不是白色
	isWB := func(r, g, b uint32) bool {
		if r == 65535 && g == 65535 && b == 65535 {
			return true
		}
		return false
	}

	// 左上角
	r1, g1, b1, _ := img.At(0+1, 0+1).RGBA()
	if !isWB(r1, g1, b1) {
		return false, fSize, err
	}

	// 右上角
	r2, g2, b2, _ := img.At(imgWidth-1, 0+1).RGBA()
	if !isWB(r2, g2, b2) {
		return false, fSize, err
	}

	// 左下角
	r3, g3, b3, _ := img.At(0+1, imgHeight-1).RGBA()
	if !isWB(r3, g3, b3) {
		return false, fSize, err
	}

	// 右下角
	r4, g4, b4, _ := img.At(imgWidth-1, imgHeight-1).RGBA()
	if !isWB(r4, g4, b4) {
		return false, fSize, err
	}

	// 没有被上面筛选掉就是白底图
	return true, fSize, err
}


// 得到一个文件夹所有白底图的，最小白底图
func MinWhiteBackground(pattern string) (string, bool) {
	// 可能的白底图
	var probablyName []string
	var probablySize []int64

	// 并发避免提前退出
	exit := make(chan bool)

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(pattern)
	for _, v := range jpgSlice {
		go func() {
			b, size, err := IsWhiteBackground(v)
			if err != nil {
				fmt.Println("tools.IsWhiteBackground err:", err)
				return
			}
			// 如果是白底图
			if b {
				probablyName = append(probablyName, v)
				probablySize = append(probablySize, size)
			}
			exit <- true
		}()
	}

	// 通道防止主进程提前退出
	for i := 0; i < len(jpgSlice); i++ {
		<-exit
	}

	// 如果一个白底图也没有检查到就返回
	if len(probablyName) == 0 {
		return "", false
	}

	minIdx := 0
	for i := 1; i < len(probablySize); i++ {
		if probablySize[minIdx] > probablySize[i] {
			minIdx = i
		}
	}

	return probablyName[minIdx], true
}