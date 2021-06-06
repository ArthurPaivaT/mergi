package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/ArthurPaivaT/mergi/convert"
)

func main() {

	imgFile, err := os.Open("./imageFiles/inputPic.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	img, err := decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("./imageFiles/result.png")

	start := time.Now()
	grayImg := convert.ToGray(img)

	elapsed := time.Since(start)
	png.Encode(newFile, grayImg)
	fmt.Println("Took ", elapsed)

	return
}
