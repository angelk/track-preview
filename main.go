package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	fmt.Println("test")
	myImage := image.NewRGBA(image.Rect(0, 0, 10, 20))

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	fmt.Println(red)

	myImage.Set(0, 0, red)
	myImage.Set(5, 0, red)
	myImage.Set(9, 0, red)
	myImage.Set(2, 1, red)

	rectangle := image.Rect(2, 2, 7, 15)
	draw.Draw(myImage, rectangle, &image.Uniform{blue}, image.ZP, draw.Src)

	drawLine(myImage, 0, 0, 10, 20, red)

	outputFile, err := os.Create("test.png")
	if err != nil {
		log.Fatalln(err)
	}

	png.Encode(outputFile, myImage)
	outputFile.Close()
}

// buggy
func drawLine(i *image.RGBA, x, y, x1, y1 int, c color.RGBA) {
	if x == x1 && y == y1 {
		return
	}

	newPoint := x
	if x > x1 {
		newPoint = x - 1
	} else if x < x1 {
		newPoint = x + 1
	}

	newY := y
	if y > y1 {
		newY = y - 1
	} else if y < y1 {
		newY = y + 1
	}

	i.Set(newPoint, newY, c)
	drawLine(i, newPoint, newY, x1, y1, c)
}
