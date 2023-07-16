package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ADSP-Project/Federation-Service/database"
	"github.com/ADSP-Project/Federation-Service/types"
	_ "github.com/lib/pq"
)

func GetPartners(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()

	rows, err := db.Query("SELECT * FROM partners")
	if err != nil {
		http.Error(w, "Failed to execute DB query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	partners := make([]types.Partner, 0) // Initialize partners as empty slice
	for rows.Next() {
		var p types.Partner
		if err := rows.Scan(&p.ShopId, &p.ShopName, &p.Rights.CanEarnCommission, &p.Rights.CanSell, &p.Rights.CanShareData, &p.Rights.CanShareInventory, &p.Rights.CanCoPromote, &p.RequestStatus); err != nil {
			http.Error(w, "Failed to scan DB result", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		partners = append(partners, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partners)
}
