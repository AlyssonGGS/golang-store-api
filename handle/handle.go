package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alyssonggs/golang-store-api/repository"
)

type Handle struct {
	Storage repository.ProductStorage
	Logger  *log.Logger
}

func (h *Handle) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! What would you like to buy?"))
}

func (h *Handle) GetProductInfo(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")

	if id == "" {
		fmt.Fprintln(w, "Please, specify a product id")
		return
	}

	product, err := h.Storage.GetProduct(id)
	if err != nil {
		http.Error(w, "Error while getting product", http.StatusInternalServerError)
		h.Logger.Println("Error while getting product", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		fmt.Fprintln(w, "Error on recovering product. Try again later!")
		h.Logger.Println("Error while parsing product", "error", err)
	}
}
