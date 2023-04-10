package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/kbinani/screenshot"
)

func main() {
	num := screenshot.NumActiveDisplays()
	if num <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < num; i++ {
		//device, _ := screenshot.CaptureDisplay(i)
		bounds := screenshot.GetDisplayBounds(i)
		all = all.Union(bounds)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("test/%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		save(img, fileName)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	// Capture all desktop region into an image.
	fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	save(img, "test/all.png")
}

func save(img *image.RGBA, fileName string) {
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
