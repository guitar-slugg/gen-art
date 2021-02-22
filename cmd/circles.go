package main

import (
	"image"
	"image/color"
	"math/rand"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func Circles() {

	width := 700
	height := 1000

	//gen.TimeBasedSeed()
	rand.Seed(1613969309100889000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	for i := 0; i < 300; i++ {

		col := color.RGBA{
			gen.RandU8(0),
			gen.RandU8(0),
			gen.RandU8(0),
			255,
		}

		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		siz := gen.RandInt(10, 100)
		linewidth := gen.RandInt(2, 5)

		gen.Circle(
			img,
			image.Point{locX, locY},
			siz,
			linewidth,
			col,
		)
	}

	gen.WritePng("../gallery/circles.png", img)

}
