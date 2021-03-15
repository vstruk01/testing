package main

import (
	"log"
	"net/http"

	"../corner"
)

func main() {
	http.HandleFunc("/corner", corner.GetCorner)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
