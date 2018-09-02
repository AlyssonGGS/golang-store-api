package main

import (
	"net/http"

	"github.com/alyssonggs/golang-store-api/handle"
)

func main() {
	handler := &handle.Handle{}

	http.HandleFunc("/store", handler.Hello)

	http.ListenAndServe(":8080", nil)
}
