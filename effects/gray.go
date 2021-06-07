package effects

import (
	"image"
	"image/color"
)

func ToGray(img image.Image) image.Image {
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X

	newImg := image.NewGray(image.Rect(minX, minY, maxX, maxY))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			newImg.Set(x, y, c)
		}
	}

	return newImg
}

func ToGrayCb(img image.Image) image.Image {
	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X

	newImg := image.NewGray(image.Rect(minX, minY, maxX, maxY))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.YCbCrModel.Convert(img.At(x, y)).(color.YCbCr)
			c.Y = c.Cr
			c.Cb = 0
			c.Cr = 0
			c.Y = c.Y / 2
			newImg.Set(x, y, c)

		}
	}

	return newImg
}
