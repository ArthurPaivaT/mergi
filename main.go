package main

import (
	"fmt"

	"github.com/ArthurPaivaT/mergi/server"
)

func main() {
	//DIVIDE RGBA BY 257 to get common rgb values

	serverErrChan := make(chan error)

	go server.Start(serverErrChan)

	serverErr := <-serverErrChan
	fmt.Println(fmt.Errorf("Error Serving: %w", serverErr))

	// imgFile, err := os.Open("./images/skyline.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer imgFile.Close()

	// img, err := decode(imgFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// newFile, err := os.Create("./images/result.png")

	// start := time.Now()
	// grayImg := effects.ToDrawing(img)

	// elapsed := time.Since(start)
	// png.Encode(newFile, grayImg)
	// fmt.Println("Took ", elapsed)
	// c := img.At(12, 6)

	// fmt.Println(c)
	// fmt.Printf("%+v", color.YCbCrModel.Convert(c))
	// fmt.Printf("%+v", color.RGBAModel.Convert(c))

	// scales.PrintColors(img)

	// fmt.Println("oxe", reflect.TypeOf(img.At(10, 10)))
	// fmt.Println("oxe", color.NRGBAModel.Convert(img.At(10, 10)))

	return
}
