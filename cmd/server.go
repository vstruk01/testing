package main

import (
	"log"
	"net/http"

	"github.com/vstruk01/testing/corner"
)

func main() {
	http.HandleFunc("/corner", corner.GetCorner)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
