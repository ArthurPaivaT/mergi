package utils

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func Decode(file io.Reader) (image.Image, error) {

	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "\x47\x49\x46\x38\x39\x61", gif.Decode, gif.DecodeConfig)

	img, _, err := image.Decode(file)
	if err == nil {

		return img, nil
	}

	return nil, image.ErrFormat
}
