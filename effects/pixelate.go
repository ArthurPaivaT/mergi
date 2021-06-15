package effects

import (
	"image"
	"image/color"
	"math"
)

func Pixelate(img image.Image, pixelSize int) image.Image {

	minY := img.Bounds().Min.Y
	maxY := img.Bounds().Max.Y
	minX := img.Bounds().Min.X
	maxX := img.Bounds().Max.X

	newImg := image.NewNRGBA(image.Rect(minX, minY, maxX, maxY))

	count := 0

	for i := minX; i < maxX; i = int(math.Min(float64(i+pixelSize), float64(maxX))) {
		for j := minY; j < maxY; j = int(math.Min(float64(j+pixelSize), float64(maxY))) {

			blockRectangle := image.Rectangle{
				Min: image.Point{
					X: i,
					Y: j,
				},
				Max: image.Point{
					X: int(math.Min(float64(i+pixelSize), float64(maxX))),
					Y: int(math.Min(float64(j+pixelSize), float64(maxY))),
				},
			}

			blockColor := getAverage(img, blockRectangle)

			setGroup(newImg, blockRectangle, blockColor)
			count++

		}
	}

	return newImg
}

func getAverage(img image.Image, rect image.Rectangle) color.NRGBA {
	var sumR, sumG, sumB int
	pixels := 0
	for i := rect.Min.X; i < rect.Max.X; i++ {
		for j := rect.Min.Y; j < rect.Max.Y; j++ {

			c := color.NRGBAModel.Convert(img.At(i, j)).(color.NRGBA)

			sumR += int(c.R)
			sumG += int(c.G)
			sumB += int(c.B)
			pixels++
		}
	}

	return color.NRGBA{
		A: uint8(255),
		R: uint8(sumR / pixels),
		G: uint8(sumG / pixels),
		B: uint8(sumB / pixels),
	}
}

func setGroup(img *image.NRGBA, rect image.Rectangle, c color.NRGBA) {
	for i := rect.Min.X; i < rect.Max.X; i++ {
		for j := rect.Min.Y; j < rect.Max.Y; j++ {
			img.Set(i, j, c)
		}
	}
}
