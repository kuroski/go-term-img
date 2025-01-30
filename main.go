package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	f, err := os.Open("image.jpeg")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	image, imageType, err := image.Decode(f)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(image, imageType)
}
