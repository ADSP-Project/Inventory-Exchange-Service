package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/product/{id}", getProductHandler).Methods("GET")
	router.HandleFunc("/api/v1/products", getProductsHandler).Methods("GET")
	router.HandleFunc("/api/v1/order", orderHandler).Methods("POST")
	router.Handle("/metrics", promhttp.Handler())

	log.Println("Server listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
