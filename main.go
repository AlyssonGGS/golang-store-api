package main

import (
	"net/http"

	"github.com/alyssonggs/golang-store-api/handle"
	"github.com/alyssonggs/golang-store-api/repository/memory"
)

func main() {
	memStorage := &memory.MemoryStorage{}
	handler := &handle.Handle{Storage: memStorage}

	http.HandleFunc("/store", handler.Hello)
	http.HandleFunc("/store/product", handler.GetProductInfo)

	http.ListenAndServe(":8080", nil)
}
