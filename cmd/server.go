package main

import (
	"log"
	"net/http"

	"github.com/vstruk01/testing/corner"
)

func main() {
	http.HandleFunc("/corner", corner.GetCorner)
	// http.HandleFunc("/mandelbrot", mandelbrot.HandleMandelbrot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
