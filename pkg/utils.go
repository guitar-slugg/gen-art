package pkg

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func SetBackgroundColor(img *image.RGBA, col color.Color) {
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			img.Set(x, y, col)
		}
	}
}

func WritePng(path string, img *image.RGBA) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
}

func RandU8(min int) uint8 {
	return uint8(rand.Intn(255-min) + min)
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}
