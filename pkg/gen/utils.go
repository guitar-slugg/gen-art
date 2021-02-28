package gen

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
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

func TimeBasedSeed() {
	seed := time.Now().UnixNano()
	//print out seed in case you get a really nice img, and want to re-use the seed
	fmt.Printf("using seed: %d \n", seed)
	rand.Seed(seed)
}
