package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/image"
)

func Compress(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	img := &capture.Image{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&img)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = image.Compress(img)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccess(w, img)
	return

}
