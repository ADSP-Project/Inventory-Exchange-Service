package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	ImageURL    string  `json:"picture"`
}

type ExternalOrderData struct {
	OrderID            string              `json:"orderId"`
	ShippingTrackingID string              `json:"shippingTrackingId"`
	ShippingCost       ExternalMoney       `json:"shippingCost"`
	ShippingAddress    ExternalAddress     `json:"shippingAddress"`
	Items              []ExternalOrderItem `json:"items"`
}

type ExternalMoney struct {
	CurrencyCode string `json:"CurrencyCode"`
	Units        int    `json:"Units"`
	Nanos        int    `json:"Nanos"`
}

type ExternalAddress struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int    `json:"zipCode"`
}

type ExternalOrderItem struct {
	Item string        `json:"item"`
	Cost ExternalMoney `json:"cost"`
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]

	db, err := sql.Open("mysql", "root:fake_password@tcp(catalogue-db:3306)/socksdb")
	if err != nil {
		log.Println("Error connecting to DB:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM sock WHERE sock_id = ?", productID)
	var product Product
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			log.Println("Error querying DB:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	log.Println("Current product count:", product.Count)

	if product.Count >= 1 {
		_, err = db.Exec("UPDATE sock SET count = count - 1 WHERE sock_id = ?", productID)
		if err != nil {
			log.Println("Error updating count:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		log.Println("Count updated successfully")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Product not available", http.StatusBadRequest)
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:fake_password@tcp(catalogue-db:3306)/socksdb")
	if err != nil {
		log.Println("Error connecting to DB:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM sock")
	if err != nil {
		log.Println("Error querying DB:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.ImageURL)
		if err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error iterating over rows:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	var orderData ExternalOrderData
	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Process the order data...

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Order processed successfully",
	}

	json.NewEncoder(w).Encode(response)
}
