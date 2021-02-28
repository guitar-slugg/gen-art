package gen

import "image/color"

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
var CYAN = color.RGBA{100, 200, 200, 0xff}
var LIGHT_BLUE = color.RGBA{102, 153, 255, 0xff}
var RED_DARK = color.RGBA{153, 0, 0, 0xff}
var ORANGE = color.RGBA{153, 77, 0, 0xff}
var GREEN = color.RGBA{0, 153, 51, 0xff}
var BLUE_DARK = color.RGBA{0, 38, 153, 0xff}
var VIOLET = color.RGBA{153, 0, 153, 0xff}
var OLIVE = color.RGBA{153, 153, 0, 0xff}
var BLACK = color.RGBA{0, 0, 0, 0xff}
var WHITE = color.RGBA{255, 255, 255, 0xff}
var GREY = color.RGBA{200, 200, 200, 0xff}

func RandomColorMap1() color.RGBA {

	COLORMAP := []color.RGBA{
		CYAN,
		LIGHT_BLUE,
		RED_DARK,
		ORANGE,
		GREEN,
		BLUE_DARK,
		VIOLET,
		OLIVE,
	}

	return COLORMAP[RandInt(0, len(COLORMAP))]
}

func BlackWhite() color.RGBA {

	COLORMAP := []color.RGBA{
		BLACK,
		WHITE,
	}

	return COLORMAP[RandInt(0, len(COLORMAP))]
}

func UsaUsa() color.RGBA {

	COLORMAP := []color.RGBA{
		RED_DARK,
		WHITE,
		BLUE_DARK,
	}

	return COLORMAP[RandInt(0, len(COLORMAP))]
}
