package pkg

import (
	"image"
	"image/color"
)

func TranslateBox(img *image.RGBA, pnt1, pnt2 image.Point, siz int) {
	for i := -siz; i < siz; i++ {
		for j := -siz; j < siz; j++ {
			SwapColors(img, image.Point{X: pnt1.X + i, Y: pnt1.Y + j}, image.Point{X: pnt2.X + i, Y: pnt2.Y + j})
		}
	}
}

func Smudge(img *image.RGBA, pnt image.Point) {

	//row above point
	for i := pnt.X - 1; i < pnt.X+1; i++ {
		Avg(img, pnt, image.Point{X: i, Y: pnt.Y - 1})
	}

	//row below point
	for i := pnt.X - 1; i < pnt.X+1; i++ {
		Avg(img, pnt, image.Point{X: i, Y: pnt.Y + 1})
	}

	//left of point
	Avg(img, pnt, image.Point{X: pnt.X - 1, Y: pnt.Y})

	//right of point
	Avg(img, pnt, image.Point{X: pnt.X + 1, Y: pnt.Y})

}

func Avg(img *image.RGBA, pnt1, pnt2 image.Point) {
	limX := img.Bounds().Size().X
	limY := img.Bounds().Size().Y

	if pnt2.X < 0 || pnt2.X > limX {
		return
	}
	if pnt2.Y < 0 || pnt2.Y > limY {
		return
	}
	if pnt1.X < 0 || pnt1.X > limX {
		return
	}
	if pnt1.Y < 0 || pnt1.Y > limY {
		return
	}

	//Get the color of point 1
	r1, g1, b1, a1 := img.At(pnt1.X, pnt1.Y).RGBA()

	//Get the color of point 2
	r2, g2, b2, a2 := img.At(pnt2.X, pnt2.Y).RGBA()

	rR := uint8((r1 + r2) / 2)
	rG := uint8((g1 + g2) / 2)
	rB := uint8((b1 + b2) / 2)
	rA := uint8((a1 + a2) / 2)

	newCol := color.RGBA{
		R: rR,
		G: rG,
		B: rB,
		A: rA,
	}

	img.Set(pnt2.X, pnt2.Y, newCol)
}

func SwapColors(img *image.RGBA, pnt1, pnt2 image.Point) {

	limX := img.Bounds().Size().X
	limY := img.Bounds().Size().Y

	if pnt2.X < 0 || pnt2.X > limX {
		return
	}
	if pnt2.Y < 0 || pnt2.Y > limY {
		return
	}
	if pnt1.X < 0 || pnt1.X > limX {
		return
	}
	if pnt1.Y < 0 || pnt1.Y > limY {
		return
	}

	img.Set(pnt1.X, pnt1.Y, img.At(pnt2.X, pnt2.Y))
}
