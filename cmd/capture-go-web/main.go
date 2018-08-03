package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smilga/capture-go/http/handler"
)

func main() {

	r := httprouter.New()

	r.GET("/capture/url", handler.Capture)
	r.POST("/compress/image", handler.Compress)

	fmt.Println("Server started on 8888 port")
	log.Fatal(http.ListenAndServe(":8888", r))
}
