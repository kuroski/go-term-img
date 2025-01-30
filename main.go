package main

import (
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
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

	grayscaleImg := image.NewGray(bounds)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pxColor := img.At(x, y)
			pxGrayscaleColor := color.GrayModel.Convert(pxColor)
			grayscaleImg.Set(x, y, pxGrayscaleColor)
		}

		//fmt.Println()
	}

	file, err := os.Create("tmp.jpeg")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	if err = jpeg.Encode(file, grayscaleImg, nil); err != nil {
		panic(err.Error())
	}
	//fmt.Println(bounds)
	//fmt.Println(img.At(0, 0))
	//fmt.Println(img.At(100, 100))
}
