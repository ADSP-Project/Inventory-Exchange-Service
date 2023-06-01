package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductService struct {
	products []Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []Product{},
	}
}

func (s *ProductService) AddProduct(product Product) {
	s.products = append(s.products, product)
}

func (s *ProductService) GetProducts() []Product {
	return s.products
}

func main() {
	service := NewProductService()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request body
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the product to the service
		service.AddProduct(product)

		// Return success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Product added successfully")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get the products from the service
		products := service.GetProducts()

		// Encode the response as JSON
		jsonResponse, err := json.Marshal(products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the response
		w.Write(jsonResponse)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
