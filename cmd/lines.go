package main

import (
	"image"
	"image/color"
	"math/rand"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func Lines() {

	width := 700
	height := 1000

	//gen.TimeBasedSeed()
	rand.Seed(1613419965343466000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	for i := 0; i < 100; i++ {

		col := color.RGBA{
			gen.RandU8(0),
			gen.RandU8(0),
			gen.RandU8(0),
			255,
		}

		locY := gen.RandInt(0, height)
		locY2 := gen.RandInt(0, height)
		siz := gen.RandInt(1, 10)

		gen.Line(
			img,
			image.Point{-10, locY},
			image.Point{width + 10, locY2},
			siz,
			col,
		)
	}

	gen.WritePng("../gallery/lines.png", img)
}
