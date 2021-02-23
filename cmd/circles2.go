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

	//gen.TimeBasedSeed()
	rand.Seed(1614062160148368000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	locX := width / 2
	locY := height / 2
	linewidth := 2

	//lots of rings
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

	//translate boxes
	for i := 0; i < 70; i++ {

		locX1 := gen.RandInt(0, width)
		locY1 := gen.RandInt(0, height)
		locX2 := gen.RandInt(0, width)
		locY2 := gen.RandInt(0, height)

		siz := gen.RandInt(30, 100)

		gen.TranslateBox(img, image.Point{X: locX1, Y: locY1}, image.Point{X: locX2, Y: locY2}, siz)

	}

	gen.WritePng("../gallery/circles2.png", img)

}
