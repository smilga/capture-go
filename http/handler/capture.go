package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/browser"
	"github.com/smilga/capture-go/pkg/image"
)

const (
	AcceptJson = "application/json"
)

// Capture receives url, compression settings and response settings from url query
// and processes that
func Capture(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url, err := capture.NewURL(r.URL.Query().Get("url"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("%s: %s", capture.ErrURLInvalid, err))
		return
	}

	d := &browser.Dimensions{
		Width:  r.URL.Query().Get("width"),
		Height: r.URL.Query().Get("height"),
	}

	img, err := browser.Screenshot(url, d, browser.Device(r.URL.Query().Get("device")))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if compress := r.URL.Query().Get("compress"); compress == "true" {
		img.Compression = &capture.Compression{}
	}

	if err := image.Compress(img); err != nil && err != image.ErrNoCompression {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), AcceptJson) {
		writeSuccess(w, img)
	} else {
		decoded, err := img.Encoded.Decode()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Content-Type", img.Mime.String())
		fmt.Fprint(w, decoded)
	}
}
