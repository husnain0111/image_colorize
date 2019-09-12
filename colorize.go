package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// ... All your code will go here
	f, err := os.Open("test_image.JPG")
	check(err)
	defer f.Close()
	img, format, err := image.Decode(f)
	check(err)
	if format != "jpeg" {
		log.Fatalln("Only jpeg images are supported")
	}
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)
	// loop though all the x
	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			// Offset colors a little, adjust it to your taste
			r := float64(originalColor.R) * 0.92126
			g := float64(originalColor.G) * 0.97152
			b := float64(originalColor.B) * 0.90722
			// average
			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			wImg.Set(x, y, c)
		}
	}
	ext := filepath.Ext("imgPath.jpg")
	name := strings.TrimSuffix(filepath.Base("imgPath.jpg"), ext)
	newImagePath := fmt.Sprintf("%s/%s_gray%s", filepath.Dir("imgPath.jpg"), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}
