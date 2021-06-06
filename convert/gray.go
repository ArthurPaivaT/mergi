package convert

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
