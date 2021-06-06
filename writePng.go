package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func createPNG(img image.Image) error {

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	return err
}
