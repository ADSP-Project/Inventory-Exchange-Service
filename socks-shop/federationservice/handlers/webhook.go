package handlers

import (
	"github.com/ADSP-Project/Federation-Service/types"
	"fmt"
	"encoding/json"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var newShop types.Shop
	json.NewDecoder(r.Body).Decode(&newShop)

	fmt.Printf("New shop joined the federation: %s\n", newShop.Name)

	// fmt.Printf("Public Key: %s", newShop.PublicKey)
}
