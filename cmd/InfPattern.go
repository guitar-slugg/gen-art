package main

import (
	"image"
	"image/color"
	"math/rand"

	gen "github.com/guitar-slug/gen-art/pkg"
)

func InfPattern() {

	width := 700
	height := 1000

	//gen.TimeBasedSeed()
	rand.Seed(1613973129150296000)

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{200, 200, 200, 250})

	for i := 0; i < 3; i++ {

		col := color.RGBA{
			0,
			0,
			0,
			gen.RandU8(50),
		}

		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		radius := 5
		iters := 50
		spacing := 20
		linewidth := 1

		gen.Concentric(
			img,
			image.Point{locX, locY},
			radius,
			linewidth,
			iters,
			spacing,
			col,
		)
	}

	gen.WritePng("../gallery/inf-pattern.png", img)

}
