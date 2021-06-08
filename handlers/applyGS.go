package handlers

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/ArthurPaivaT/mergi/effects"
	"github.com/ArthurPaivaT/mergi/utils"
)

func ApplyGS(w http.ResponseWriter, r *http.Request) {

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

	grayImg := effects.ToGray(img)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	png.Encode(w, grayImg)
}
