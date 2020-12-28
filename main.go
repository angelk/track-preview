package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"math"
)

func main() {
	imageX := 200
	imageY := 200
	myImage := image.NewRGBA(image.Rect(0, 0, imageX, imageY))

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	white := color.RGBA{255, 255, 255, 255}

	rectangle := image.Rect(0, 0, imageX, imageY)
	draw.Draw(myImage, rectangle, &image.Uniform{white}, image.ZP, draw.Src)

	fmt.Println(red)

	drawLine(myImage, 0, 0, 50, 50, blue)
	drawLine(myImage, 0, 0, 10, 20, red)
	drawLine(myImage, 50, 50, 100, 0, red)

	drawLine(myImage, 50, 50, 0, 10, blue)
	drawLine(myImage, 0, 10, 50, 50, blue)

	outputFile, err := os.Create("test.png")
	if err != nil {
		log.Fatalln(err)
	}

	png.Encode(outputFile, myImage)
	outputFile.Close()
}

// draw line between two points using Bresenham's line algorithm
func drawLine(i *image.RGBA, x, y, x1, y1 int, c color.RGBA) {
	if x == x1 && y == y1 {
		return
	}

	if (x > x1) {
		x, x1 = x1, x
		y, y1 = y1, y
	}

	deltaX := x1 - x
	deltaY := y1 - y

	deltaErr := math.Abs(float64(deltaY) / float64(deltaX))

	var realError float64 = 0

	tmpY := float64(y)

	for tmpX := x; tmpX <= x1; tmpX++ {
		i.Set(tmpX, int(tmpY), c)

		realError += deltaErr
		if realError >= 0.5 {
			if (y1 > y) {
				tmpY += deltaErr
			} else {
				tmpY -= deltaErr
			}
			realError -= 1
		}
	}
}
