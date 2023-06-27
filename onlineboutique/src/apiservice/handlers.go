package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (fe *apiServer) productsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := fe.getProducts(r.Context())
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(b)
}
