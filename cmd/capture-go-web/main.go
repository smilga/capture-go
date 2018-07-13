package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smilga/capture-go/http/handler"
)

func main() {

	r := httprouter.New()

	r.GET("/", handler.Index)
	r.GET("/url/capture", handler.Capture)
	//r.POST("/base64/compress", handler.Compress)

	log.Fatal(http.ListenAndServe(":8888", r))
}
