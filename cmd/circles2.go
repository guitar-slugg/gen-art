package main

import (
	"image"
	"image/color"
	"math/rand"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func Circles2() {

	width := 700
	height := 1000

	gen.TimeBasedSeed()
	rand.Seed(1613969715421272000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	locX := width / 2
	locY := height / 2
	linewidth := 2

	for i := 0; i < height; i += linewidth {

		col := color.RGBA{
			gen.RandU8(0),
			gen.RandU8(0),
			gen.RandU8(0),
			255,
		}

		gen.Circle(
			img,
			image.Point{locX, locY},
			i,
			linewidth,
			col,
		)
	}

	gen.WritePng("../gallery/circles2.png", img)

}
