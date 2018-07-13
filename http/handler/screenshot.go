package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/slimer"
)

//func POSTScreenshot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//var screenshot capture.Screenshot

//decoder := json.NewDecoder(r.Body)
//err := decoder.Decode(&screenshot)
//if err != nil {
//writeError(w, http.StatusInternalServerError, err)
//return
//}
//spew.Dump(screenshot)
//writeSuccess(w, screenshot)
//}

func Capture(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url, err := capture.NewURL(r.URL.Query().Get("url"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("%s: %s", capture.ErrURLInvalid, err))
		return
	}

	// TODO read compression settings and set to screenshot

	image, err := slimer.CaptureURL(url)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// To return image to browser
	img, err := base64.StdEncoding.DecodeString(string(image))
	if err != nil {
		fmt.Println("ERRRRR", err)
	}
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprint(w, string(img))
	//writeSuccess(w, image)
}
