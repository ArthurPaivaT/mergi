package convert

import (
	"image"
	"image/color"
)

func ToDrawing(img image.Image) image.Image {

	newImg := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			round(&c)
			cm := color.CMYKModel.Convert(c).(color.CMYK)

			newImg.Set(x, y, cm)
		}
	}

	return newImg
}

func round(c *color.NRGBA) {

	rR := c.R % 51

	if rR > 51 || c.R < 51 {
		if c.R >= 204 {
			c.R = 255
		}
		c.R += rR
	} else {
		if c.R <= 51 {
			c.R = 0
		}
		c.R -= rR
	}

	rG := c.G % 51

	if (rG > 51) || c.G < 51 {
		if c.G >= 204 {
			c.G = 255
		}
		c.G += rG
	} else {
		if c.G <= 51 {
			c.G = 0
		}
		c.G -= rG
	}

	rB := c.B % 51

	if (rB > 51) || c.B < 51 {
		if c.B >= 204 {
			c.B = 255
		}
		c.B += rB

	} else {
		if c.B <= 51 {
			c.B = 0
		}
		c.B -= rB
	}

}
