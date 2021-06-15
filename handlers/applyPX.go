package handlers

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/ArthurPaivaT/mergi/effects"
	"github.com/ArthurPaivaT/mergi/utils"
)

func ApplyPX(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("file")
	if err != nil {
		err := fmt.Errorf("Missing file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, err := utils.Decode(file)
	if err != nil {
		err := fmt.Errorf("Error Decoding: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pixelatedImg := effects.Pixelate(img, 10)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	png.Encode(w, pixelatedImg)
}
