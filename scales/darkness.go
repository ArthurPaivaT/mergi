package scales

import (
	"image"
	"image/color"
)

func Darkness(img image.Image) float64 {
	totalPixels := Width(img) * Height(img)

	darkPixels := 0
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)

			if c.Y >= 100 {
				darkPixels++
			}
		}
	}
	darkness := float64(darkPixels) / float64(totalPixels)

	return darkness
}
