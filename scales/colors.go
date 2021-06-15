package scales

import (
	"fmt"
	"image"
	"image/color"
)

func PrintColors(img image.Image) {

	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			rgb := color.NRGBAModel.Convert(img.At(y, x)).(color.NRGBA)

			fmt.Println("AA", rgb.R, rgb.G, rgb.B)

		}
	}

}

func Width(i image.Image) int {
	return i.Bounds().Max.X - i.Bounds().Min.X
}
func Height(i image.Image) int {
	return i.Bounds().Max.Y - i.Bounds().Min.Y
}
