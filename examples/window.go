// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// 		go run ./examples/window.go /home/ron/Pictures/mcp23017.jpg
//
// +build example

package main

import (
	"os"

	opencv3 ".."
)

func main() {
	filename := os.Args[1]
	window := opencv3.NewWindow("Hello")
	img := opencv3.IMRead(filename, opencv3.IMReadColor)

	for {
		window.IMShow(img)
		opencv3.WaitKey(1)
	}
}
