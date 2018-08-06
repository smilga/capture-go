package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	capture "github.com/smilga/capture-go"
	"github.com/smilga/capture-go/pkg/image"
	"github.com/smilga/capture-go/pkg/logger"
	"github.com/smilga/capture-go/pkg/slimer"
)

// Capture receives url, compression settings and response settings from url query
// and processes that
func Capture(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url, err := capture.NewURL(r.URL.Query().Get("url"))
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("%s: %s", capture.ErrURLInvalid, err))
		return
	}

	logger.Info("This is info from capture called")

	img, err := slimer.CaptureURL(url)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if compress := r.URL.Query().Get("compress"); compress == "true" {
		img.Compression = &capture.Compression{}
	}

	if err := image.Compress(img); err != nil && err != image.ErrNoCompression {
		logger.Info("Setting compression")
		logger.Error(fmt.Sprintf("Error: compresing error %s", err))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO read Accept header and return response depending on that
	decoded, err := img.Encoded.Decode()
	if err != nil {
		fmt.Println("Error decoding", err)
	}
	w.Header().Set("Content-Type", img.Mime.String())
	fmt.Fprint(w, decoded)

	// To write json response

	//writeSuccess(w, img)
}
