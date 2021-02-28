package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"strings"

	"github.com/guitar-slug/gen-art/pkg/gen"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("need example name")
		os.Exit(-1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "lines":
		lines()
	case "rings":
		rings()
	case "circles":
		circles()
	case "inf":
		infPattern()
	default:
		fmt.Println("no such example: ", os.Args[1])
	}
}

func defaultCanvas() (*image.RGBA, int, int) {

	width := 700
	height := 1000

	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	gen.SetBackgroundColor(img, color.RGBA{100, 100, 100, 250})

	return img, width, height
}

func rings() {

	img, width, height := defaultCanvas()

	//gen.TimeBasedSeed()
	rand.Seed(1614555718659077000)

	//create a bunch of rings
	for i := 0; i < 100; i++ {

		col := gen.UsaUsa()
		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		siz := gen.RandInt(10, 200)
		linewidth := gen.RandInt(2, 10)

		gen.Circle(
			img,
			image.Point{locX, locY},
			siz,
			linewidth,
			col,
		)
	}

	//translate boxes
	for i := 0; i < 5; i++ {

		locX1 := gen.RandInt(0, width)
		locY1 := gen.RandInt(0, height)
		locX2 := gen.RandInt(0, width)
		locY2 := gen.RandInt(0, height)

		siz := gen.RandInt(30, 500)

		gen.TranslateBox(img, image.Point{X: locX1, Y: locY1}, image.Point{X: locX2, Y: locY2}, siz)

	}

	gen.WritePng("../gallery/rings.png", img)

}

func circles() {

	img, width, height := defaultCanvas()

	//gen.TimeBasedSeed()
	rand.Seed(1614554583621770000)

	locX := width / 2
	locY := height / 2

	//lots of rings
	i := 0
	for i < height {

		col := gen.BlackWhite()
		linewidth := gen.RandInt(1, 10)

		gen.Circle(
			img,
			image.Point{locX, locY},
			i,
			linewidth,
			col,
		)

		i += linewidth
	}

	//translate boxes
	for i := 0; i < 50; i++ {

		locX1 := gen.RandInt(0, width)
		locY1 := gen.RandInt(0, height)
		locX2 := gen.RandInt(0, width)
		locY2 := gen.RandInt(0, height)

		siz := gen.RandInt(30, 100)

		gen.TranslateBox(img, image.Point{X: locX1, Y: locY1}, image.Point{X: locX2, Y: locY2}, siz)

	}

	gen.WritePng("../gallery/circles.png", img)

}

func infPattern() {

	img, width, height := defaultCanvas()

	//gen.TimeBasedSeed()
	rand.Seed(1614554762803500000)

	for i := 0; i < 4; i++ {

		col := gen.UsaUsa()

		locX := gen.RandInt(0, width)
		locY := gen.RandInt(0, height)
		radius := 5
		iters := 50
		spacing := 20
		linewidth := 2

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

func lines() {

	img, width, height := defaultCanvas()
	gen.SetBackgroundColor(img, color.RGBA{0, 0, 0, 250})

	//gen.TimeBasedSeed()
	rand.Seed(1614553922480794000)

	for i := 0; i < 120; i++ {

		col := gen.RandomColorMap1()

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
