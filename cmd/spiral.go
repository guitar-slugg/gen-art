package main

import (
	"image"
	"image/color"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func Spiral() {

	width := 700
	height := 1000

	gen.TimeBasedSeed()
	//rand.Seed(1613970822467996000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	for i := 0; i < 50; i++ {

		col := color.RGBA{
			0,
			gen.RandU8(0),
			gen.RandU8(0),
			gen.RandU8(50),
		}

		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		siz := gen.RandInt(10, 100)
		iters := gen.RandInt(1, 50)
		linewidth := gen.RandInt(2, 5)

		gen.Spiral(
			img,
			image.Point{locX, locY},
			siz,
			linewidth,
			iters,
			0.025,
			col,
		)
	}

	gen.WritePng("../gallery/spiral.png", img)

}
