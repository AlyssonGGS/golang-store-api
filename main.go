package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alyssonggs/golang-store-api/handle"
	"github.com/alyssonggs/golang-store-api/repository/memory"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "store-api", 0)
}

func main() {
	memStorage := &memory.MemoryStorage{}
	handler := &handle.Handle{Storage: memStorage, Logger: logger}

	http.HandleFunc("/store", handler.Hello)
	http.HandleFunc("/store/product", handler.GetProductInfo)

	http.ListenAndServe(":8080", nil)
}
