package pkg

import (
	"image"
	"image/color"
)

func Square(img *image.RGBA, a image.Point, siz int, col color.Color) {

	limX := img.Bounds().Size().X
	limY := img.Bounds().Size().Y

	for x := a.X; x < a.X+siz; x++ {
		for y := a.Y; y < a.Y+siz; y++ {
			if x < limX && x > -1 && y < limY && y > -1 {
				img.Set(x, y, col)
			}
		}
	}
}

func SquareBlend(img *image.RGBA, a image.Point, siz int, col color.Color) {

	limX := img.Bounds().Size().X
	limY := img.Bounds().Size().Y

	for x := a.X; x < a.X+siz; x++ {
		for y := a.Y; y < a.Y+siz; y++ {
			if x < limX && x > -1 && y < limY && y > -1 {

				r, g, b, _ := img.At(x, y).RGBA()
				r2, g2, b2, _ := col.RGBA()
				R := uint8(r + r2)
				G := uint8(g + g2)
				B := uint8(b + b2)

				img.Set(x, y, color.RGBA{R, G, B, 255})
			}
		}
	}
}

func Line(img *image.RGBA, a, b image.Point, siz int, col color.Color) {

	limX := img.Bounds().Size().X
	limY := img.Bounds().Size().Y

	//start from point a
	dx := b.X - a.X
	dy := b.Y - a.Y

	for x := a.X; x < b.X; x++ {
		y := a.Y + dy*(x-a.X)/dx

		//fill in thickness
		for i := 0; i < siz; i++ {
			if x < limX && x > -1 && (y+i) < limY && (y+i) > -1 {
				img.Set(x, y+i, col)
			}
		}
		for i := 0; i < siz; i++ {
			if x < limX && x > -1 && (y-i) < limY && (y-i) > -1 {
				img.Set(x, y-i, col)
			}
		}
	}

}
