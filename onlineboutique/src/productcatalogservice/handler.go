package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// type priceUsd struct {
// 	CurrencyCode string `json:"currencyCode"`
// 	Units        int    `json:"units"`
// 	Nanos        int    `json:"nanos"`
// }

//type categories []string

// Product defines a structure for an item in product catalog
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	//PriceUsd    priceUsd   `json:"priceUsd"`
	//Categories  categories `json:"categories"`
}

type ProductsCatalog struct {
	Products []Product `json:"products"`
}

// CreateProductHandler is used to create a new product and add to our product store.
func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var products []Product
		err = json.Unmarshal(data, &products)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		fmt.Println(products)
		// Load existing products and append the data to product list
		var catalog ProductsCatalog
		data, err = ioutil.ReadFile("products.json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(data, &catalog)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			fmt.Println("An error happended with unmarshal")
			return
		}
		fmt.Println(catalog)
		catalog.Products = append(catalog.Products, products...)
		updatedData, err := json.Marshal(catalog)
		if err != nil {
			fmt.Println("An error happended with appending")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(updatedData)
		err = ioutil.WriteFile("products.json", updatedData, os.ModePerm)
		if err != nil {
			fmt.Println("Could not write to file")
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
		fmt.Println("Added products")
		reloadCatalog = true
		return
	}
}
