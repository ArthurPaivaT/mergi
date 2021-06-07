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
	if rR > 25 {
		if (c.R + 51 - rR) > 255 {
			c.R = 255
		} else {
			c.R += 51 - rR
		}
	} else {
		if (c.R - rR) < 0 {
			c.R = 0
		} else {
			c.R -= rR
		}
	}

	rG := c.G % 51
	if rG > 25 {
		if (c.G + 51 - rG) > 255 {
			c.G = 255
		} else {
			c.G += 51 - rG
		}
	} else {
		if (c.G - rG) < 0 {
			c.G = 0
		} else {
			c.G -= rG
		}
	}

	rB := c.B % 51
	if rB > 25 {
		if (c.B + 51 - rB) > 255 {
			c.B = 255
		} else {
			c.B += 51 - rB
		}
	} else {
		if (c.B - rB) < 0 {
			c.B = 0
		} else {
			c.B -= rB
		}
	}

}
