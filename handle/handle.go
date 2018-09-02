package handle

import (
	"net/http"
)

type Handle struct {
}

func (Handle) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! What would you like to buy?"))
}
