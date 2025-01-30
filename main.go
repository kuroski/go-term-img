package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"math"
	"os"
)

func main() {
	f, err := os.Open("image.jpeg")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err.Error())
	}

	bounds := img.Bounds()

	//asciiChars := []byte(" .,:;i1tfLCG08@")
	//asciiChars := []byte(" .:;+xX$&")
	//asciiChars := reverse([]byte("@%#*+=-:. "))
	asciiChars := reverse([]byte("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "))
	var asciiArt string

	grayscaleImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 { // Skip every other row for better aspect ratio
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pxColor := img.At(x, y)
			pxGrayscaleColor := color.GrayModel.Convert(pxColor).(color.Gray)
			grayscaleImg.Set(x, y, pxGrayscaleColor)

			charIndex := int(math.Ceil(float64(len(asciiChars)-1) * float64(pxGrayscaleColor.Y/255)))
			asciiArt += string(asciiChars[charIndex])
		}

		asciiArt += "\n"
	}

	fmt.Println(asciiArt)

	file, err := os.Create("tmp.jpeg")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	if err = jpeg.Encode(file, grayscaleImg, nil); err != nil {
		panic(err.Error())
	}
}

func clampDimmensions(bounds image.Rectangle) {
	originalWidth := bounds.Max.X
	originalHeight := bounds.Max.Y

	targetWidth := 200
	targetHeight := (originalHeight * targetWidth) / originalWidth

	scaledImg := image.NewNRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	fmt.Println(scaledImg)
}

func reverse(input []byte) []byte {
	reversed := make([]byte, len(input))

	for i, j := 0, len(input)-1; i < len(input); i, j = i+1, j-1 {
		reversed[i] = input[j]
	}

	return reversed
}
