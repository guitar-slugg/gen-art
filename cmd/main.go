package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("need sub-function name")
		os.Exit(-1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "lines":
		Lines()
	case "square-blend":
		SquareBlend()
	case "circles":
		Circles()
	case "circles2":
		Circles2()
	case "spiral":
		Spiral()
	case "inf":
		InfPattern()
	default:
		fmt.Println("no such sub-function: ", os.Args[1])
	}
}
