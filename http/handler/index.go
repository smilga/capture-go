package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index used to test server
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Server working")
}
