package tools_scCapture

import (
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/kbinani/screenshot"
)

// 获取某个显示器的部分区域。 screenshot.NumActiveDisplays()获取所有显示器
// displayIndex 下标，从0开始
func CaptureRect(displayIndex int, rect image.Rectangle) *image.RGBA {
	//device, _ := screenshot.CaptureDisplay(i)
	bounds := screenshot.GetDisplayBounds(displayIndex)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	return img
}

// png 格式保存
func SavePng(img *image.RGBA, fileName string) {
	parent := filepath.Dir(fileName)
	if parent != "" {
		os.MkdirAll(parent, 0666)
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
