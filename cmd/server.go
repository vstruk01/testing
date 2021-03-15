package main

import (
	"log"
	"net/http"

	// "github.com/vstruk01/testing/corner"
	"../corner"
)

func main() {
	http.HandleFunc("/corner", corner.GetCorner)
	http.HandleFunc("/mandelbrot", mandelbrot.handleMandelbrot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
