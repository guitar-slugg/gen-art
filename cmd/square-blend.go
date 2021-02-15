package main

import (
	"image"
	"image/color"
	"math/rand"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func main() {

	width := 700
	height := 1000

	//gen.TimeBasedSeed()
	rand.Seed(1613419965343466000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 250, 250, 250})

	for i := 0; i < 50; i++ {

		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		siz := gen.RandInt(1, 200)

		col := color.RGBA{
			gen.RandU8(0),
			gen.RandU8(0),
			gen.RandU8(0),
			255,
		}

		gen.SquareBlend(img, image.Point{locX, locY}, siz, col)
	}

	gen.WritePng("../gallery/square-blend.png", img)
}
