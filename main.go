package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func main() {
	targetWidth := 400
	targetHeight := 0

	f, err := os.Open("image-02.jpeg")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err.Error())
	}

	scaledImg := scaleImage(img, targetWidth, targetHeight)
	asciiArt, grayscaleImg := imageToASCIIArt(scaledImg)

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

func scaleImage(img image.Image, targetWidth int, targetHeight int) image.Image {
	bounds := img.Bounds()
	originalWidth, originalHeight := bounds.Max.X, bounds.Max.Y

	// Calculate the missing dimension to preserve aspect ratio
	if targetWidth == 0 && targetHeight > 0 {
		// Calculate width based on height
		targetWidth = int(float64(targetHeight) * float64(originalWidth) / float64(originalHeight))
	} else if targetHeight == 0 && targetWidth > 0 {
		// Calculate height based on width
		targetHeight = int(float64(targetWidth) * float64(originalHeight) / float64(originalWidth))
	}

	// Create a new image with the target dimensions
	scaledImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Calculate scaling ratios
	xRatio := float64(originalWidth) / float64(targetWidth)
	yRatio := float64(originalHeight) / float64(targetHeight)

	// Map pixels from the original image to the scaled image
	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetWidth; x++ {
			// Calculate the corresponding pixel in the original image
			srcX := int(float64(x) * xRatio)
			srcY := int(float64(y) * yRatio)

			// Get the pixel color from the original image
			pixel := img.At(srcX, srcY)

			// Set the pixel in the scaled image
			scaledImg.Set(x, y, pixel)
		}
	}

	return scaledImg
}

func imageToASCIIArt(img image.Image) (string, image.Image) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	//asciiChars := []byte(" .,:;i1tfLCG08@")
	//asciiChars := []byte(" .:;+xX$&")
	//asciiChars := reverse([]byte("@%#*+=-:. "))
	asciiChars := reverse([]byte("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "))

	var asciiArt string
	grayscaleImg := image.NewGray(bounds)

	// Iterate over each pixel and map to ASCII
	for y := bounds.Min.Y; y < height; y += 2 { // Skip every other row for better aspect ratio
		for x := bounds.Min.X; x < width; x++ {
			// Get the pixel color
			pixel := img.At(x, y)
			gray := color.GrayModel.Convert(pixel).(color.Gray)
			grayscaleImg.Set(x, y, gray)

			// Map the grayscale value to an ASCII character
			charIndex := float64(gray.Y) / 255 * float64(len(asciiChars)-1)
			asciiArt += string(asciiChars[int(charIndex)])
		}
		asciiArt += "\n"
	}

	return asciiArt, grayscaleImg
}

func reverse(input []byte) []byte {
	reversed := make([]byte, len(input))

	for i, j := 0, len(input)-1; i < len(input); i, j = i+1, j-1 {
		reversed[i] = input[j]
	}

	return reversed
}
