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
	"sync"
)

// ImageSave 保存图像文件到本地
//保存路径；要保存的图像；保存的图像质量1~100
func ImageSave(img image.Image, dstPath string, quality int) bool {
	//保存图片，扩展名可以不一致，质量99，避免超500k
	err := imaging.Save(img, dstPath, imaging.JPEGQuality(quality))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
		return false
	}
	return true
}

// ImageFit 返回最佳裁剪比例，不改变比例的情况下，剪去多余的尺寸并且尽量保持最大的像素
// 原图的宽；原图的高；希望得到的宽；希望得到的高
func ImageFit(srcWidth, srcHeight, dstWidth, dstHeight int) (fitWidth, fitHeight int) {
	// 先求出倍数，用于得到最大可保留的裁剪尺寸
	fit := float64(srcWidth) / float64(dstWidth)
	//使用倍数求出合适的尺寸
	fitWidth = int(float64(dstWidth) * fit)
	fitHeight = int(float64(dstHeight) * fit)
	//如果叠加倍数后高和宽都没有超出原始尺寸，那说明之前用宽求的倍数正确，否则就再用高来算倍数
	if fitWidth <= srcWidth && fitHeight <= srcHeight {
		return
	}

	//使用宽求的倍数不正确，因此重新利用高求倍数
	fit = float64(srcWidth) / float64(dstWidth)
	//使用倍数求出合适的尺寸
	fitWidth = int(float64(dstWidth) * fit)
	fitHeight = int(float64(dstHeight) * fit)
	return
}

// ImageSmartCrop 对图像进行智能裁剪，返回裁剪后的图像
// 要裁剪的图像；希望得到的宽；希望得到的高
func ImageSmartCrop(img image.Image, width, height int) image.Image {
	// 创建解析器
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	// 裁剪图像长宽比为最小的原图宽或高
	topCrop, _ := analyzer.FindBestCrop(img, width, height)
	// 为各种类型图像的 SubImage 统一做成接口调用
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	// 调用方法，其实这个方法早在标准库实现
	return img.(SubImager).SubImage(topCrop)
}

// ImageResize 裁剪图像大小，不改变比例，多出的部分会被删掉，返回是否裁剪成功
// 源路径；保存路径；宽度；高度；裁剪模式：1智能 2居中 3居上 4居下 5居左 6居右；保存质量1~100
func ImageResize(srcPath, dstPath string, width, height, mode, quality int) bool {
	// 打开原始图像
	src, err := imaging.Open(srcPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return false
	}

	// 获取源图片的长度和宽度
	b := src.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y

	// 如果图片比例一致就直接缩小好了，这里必须转成浮点数
	if float64(width)/float64(height) == float64(srcWidth)/float64(srcHeight) {
		//对图像进行缩小，imaging.Lanczos 是最高清的
		dst := imaging.Resize(src, width, height, imaging.Lanczos)
		//保存图片，扩展名可以不一致，建议质量99，避免超500k
		return ImageSave(dst, dstPath, quality)
	}

	//执行相应的模式进行裁剪
	switch mode {
	case 1: //智能
		// 先处理需要裁剪成正方形的需求
		if width == height {
			// 判断原图的宽高中，是不是宽度的值最小
			isWidthMin := srcWidth<srcHeight
			// 用于保存缩放裁剪后的结果
			var dst image.Image
			// 如果宽最小，就按照宽最小的来等比例缩放图片
			if isWidthMin {
				dst = imaging.Resize(src, width, 0, imaging.Lanczos)
			}else {
				dst = imaging.Resize(src,0 , height, imaging.Lanczos)
			}
			// 再进行智能裁剪
			dst = ImageSmartCrop(dst, width, height)
			//保存图片，扩展名可以不一致，质量99，避免超500k
			return ImageSave(dst, dstPath, quality)

		} else {// 再处理需要裁剪成长方形的需求
			//先得到最佳缩放的宽和高
			fitWidth, fitHeight := ImageFit(srcWidth, srcHeight, width, height)
			//对图像进行智能裁剪并缩放至需求尺寸
			dst := ImageSmartCrop(src, fitWidth, fitHeight)
			dst = imaging.Resize(dst, width, height, imaging.Lanczos)
			//保存图片，扩展名可以不一致，质量99，避免超500k
			return ImageSave(dst, dstPath, quality)
		}

	case 2: //居中
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 3: //居上
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Top, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 4: //居下
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Bottom, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 5: //居左
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Left, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 6: //居右
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Right, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)
	}
	return false
}

// ImageResize 裁剪图像大小，不改变比例，多出的部分会被删掉，返回是否裁剪成功
// 源路径；保存路径；宽度；高度；裁剪模式：1智能 2居中 3居上 4居下 5居左 6居右；保存质量1~100
func ImageResizeBK(srcPath, dstPath string, width, height, mode, quality int) bool {
	// 打开原始图像
	src, err := imaging.Open(srcPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return false
	}

	// 获取源图片的长度和宽度
	b := src.Bounds()
	srcWidth := b.Max.X
	srcHeight := b.Max.Y

	// 如果图片比例一致就直接缩小好了，这里必须转成浮点数
	if float64(width)/float64(height) == float64(srcWidth)/float64(srcHeight) {
		//对图像进行缩小，imaging.Lanczos 是最高清的
		dst := imaging.Resize(src, width, height, imaging.Lanczos)
		//保存图片，扩展名可以不一致，建议质量99，避免超500k
		return ImageSave(dst, dstPath, quality)
	}

	//执行相应的模式进行裁剪
	switch mode {
	case 1: //智能
		//先处理正方形需求
		if width == height {
			// 先为智能裁剪提前得到源图中宽和高哪个最小
			minSize := srcWidth
			// 默认宽最小，如果不是那就是高最小
			if srcHeight < srcWidth {
				minSize = srcHeight
			}
			// 对图像进行智能裁剪并缩放至需求尺寸
			dst := ImageSmartCrop(src, minSize, minSize)
			dst = imaging.Resize(dst, width, height, imaging.Lanczos)
			//保存图片，扩展名可以不一致，质量99，避免超500k
			return ImageSave(dst, dstPath, quality)
		} else {
			//先得到最佳缩放的宽和高
			fitWidth, fitHeight := ImageFit(srcWidth, srcHeight, width, height)
			//对图像进行智能裁剪并缩放至需求尺寸
			dst := ImageSmartCrop(src, fitWidth, fitHeight)
			dst = imaging.Resize(dst, width, height, imaging.Lanczos)
			//保存图片，扩展名可以不一致，质量99，避免超500k
			return ImageSave(dst, dstPath, quality)
		}

	case 2: //居中
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 3: //居上
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Top, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 4: //居下
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Bottom, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 5: //居左
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Left, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)

	case 6: //居右
		// 不拉伸的情况下获得正确的纵横比，将对源图像进行裁剪。imaging.Lanczos 是最高清的
		dst := imaging.Fill(src, width, height, imaging.Right, imaging.Lanczos)
		return ImageSave(dst, dstPath, quality)
	}
	return false
}

// ImageWatermark 给图片加水印，保存路径，源图像，水印图像，左上角的定位 x,y，
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

// ImageResizeAndWatermark 修改图片大小并且加水印；
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

// FindDomiantColor 检测这张图片的主要颜色
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

// IsWhiteBackground 检测单张 是否是白底图，返回 图像大小
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
	r1, g1, b1, _ := img.At(0+10, 0+10).RGBA()
	if !isWB(r1, g1, b1) {
		return false, fSize, err
	}

	// 右上角
	r2, g2, b2, _ := img.At(imgWidth-10, 0+10).RGBA()
	if !isWB(r2, g2, b2) {
		return false, fSize, err
	}

	// 左下角
	r3, g3, b3, _ := img.At(0+10, imgHeight-10).RGBA()
	if !isWB(r3, g3, b3) {
		return false, fSize, err
	}

	// 右下角
	r4, g4, b4, _ := img.At(imgWidth-10, imgHeight-10).RGBA()
	if !isWB(r4, g4, b4) {
		return false, fSize, err
	}

	// 没有被上面筛选掉就是白底图
	return true, fSize, err
}

// MinWhiteBackground 得到一个文件夹的白底图的，如果有多个白底图就返回最小的一个白底图
func MinWhiteBackground(pattern string) (result string, exist bool) {
	// 可能的白底图
	var probablyName []string
	var probablySize []int64

	// 并发操作同一个切片时，加个互斥锁
	var add sync.Mutex

	// 并发避免提前退出
	var wg sync.WaitGroup

	// 获取所有扩展名是jpg的文件名，类型是字符串切片
	jpgSlice, _ := filepath.Glob(pattern)

	// 每个图片都使用新的go程去判断是不是白底图
	for i := 0; i < len(jpgSlice); i++ {
		wg.Add(1) // 计时器加1，go程还没创建时就要先加1

		// 使用go并发时记得传参，因为外面是公用变量，在等待的过程中，公用变量可能会发生变化
		go func(i int) {
			defer wg.Done() //计数器减1

			ok, size, err := IsWhiteBackground(jpgSlice[i])
			if err != nil {
				fmt.Println("tools.IsWhiteBackground err:", err)
				return
			}

			// 如果是白底图，就添加进切片
			if ok {
				add.Lock() // 加锁
				probablyName = append(probablyName, jpgSlice[i])
				probablySize = append(probablySize, size)
				add.Unlock() // 解锁
			}
		}(i)
	}

	// 通道防止主进程提前退出
	wg.Wait()

	// 如果一个白底图也没有检查到就返回
	if len(probablyName) == 0 {
		return
	}

	// 只有一个也马上返回
	if len(probablyName) == 1 {
		return probablyName[0], true
	}

	// 如果有多个就得求出最小的白底图
	minIdx := 0
	for i := 1; i < len(probablySize); i++ {
		if probablySize[minIdx] > probablySize[i] {
			minIdx = i
		}
	}

	return probablyName[minIdx], true
}
