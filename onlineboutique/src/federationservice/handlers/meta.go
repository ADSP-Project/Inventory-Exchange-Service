package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ADSP-Project/Federation-Service/database"
	"github.com/ADSP-Project/Federation-Service/globals"
)

func GetShop(w http.ResponseWriter, r *http.Request) {
    shop, err := database.GetShopByName(globals.ShopName)
    if err != nil {
        http.Error(w, "Failed to get shop information", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(shop)
}

func GetShops(w http.ResponseWriter, r *http.Request) {
    shops, err := database.GetAllShops()
    if err != nil {
        http.Error(w, "Failed to get shops information", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(shops)
}