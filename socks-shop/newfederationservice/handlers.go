package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type InternalProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	ImageURL1   string  `json:"picture1"`
	ImageURL2   string  `json:"picture2"`
}

type ExternalProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
}

type ExternalMoney struct {
	CurrencyCode string
	Units        int64
	Nanos        int32
}

type ExternalAddress struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"`
}

type ExternalOrderItem struct {
	ID   string        `json:"item"`
	Cost ExternalMoney `json:"cost"`
}

type ExternalOrderData struct {
	OrderId            string              `json:"order_id"`
	ShippingTrackingId string              `json:"shipping_tracking_id"`
	ShippingCost       ExternalMoney       `json:"shipping_cost"`
	ShippingAddress    ExternalAddress     `json:"shipping_address"`
	Items              []ExternalOrderItem `json:"items"`
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	log.Println("root:fake_password@tcp(catalogue-db:3306)/socksdb")
	db, err := sql.Open("mysql", "root:fake_password@tcp(catalogue-db:3306)/socksdb")
	if err != nil {
		log.Println("Error connecting to DB:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM sock WHERE sock_id = ?", productID)
	var product InternalProduct
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.ImageURL1, &product.ImageURL2)
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

	var products []ExternalProduct
	for rows.Next() {
		var product InternalProduct
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.ImageURL1, &product.ImageURL2)
		if err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		var externalProduct ExternalProduct
		externalProduct.ID = "SKSH:" + product.ID
		externalProduct.Name = product.Name
		externalProduct.Description = strings.Replace(product.Description, "'", `'\''`, -1)
		externalProduct.Price = product.Price
		externalProduct.Picture = "/static" + product.ImageURL1

		products = append(products, externalProduct)
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
	log.Println(r.Body)
	var orderData ExternalOrderData
	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(orderData)

	// Process the order data...

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Order processed successfully",
	}

	json.NewEncoder(w).Encode(response)
}
