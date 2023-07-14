package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type Shop struct {
	Name        string `json:"name"`
	WebhookURL  string `json:"webhookURL"`
	PublicKey   string `json:"publicKey"`
	Description string `json:"description"`
}

type Partner struct {
	ShopId   string `json:"shopId"`
	ShopName string `json:"shopName"`
	Rights   struct {
		CanEarnCommission bool `json:"canEarnCommission"`
		CanShareInventory bool `json:"canShareInventory"`
		CanShareData      bool `json:"canShareData"`
		CanCoPromote      bool `json:"canCoPromote"`
		CanSell           bool `json:"canSell"`
	} `json:"rights"`
	RequestStatus string `json:"requestStatus"`
}

type Rights struct {
	CanEarnCommission bool `json:"canEarnCommission"`
	CanSell           bool `json:"canSell"`
	CanShareData      bool `json:"canShareData"`
	CanShareInventory bool `json:"canShareInventory"`
	CanCoPromote      bool `json:"canCoPromote"`
}

type ShopDisplay struct {
	Name       string `json:"name"`
	WebhookURL string `json:"webhookURL"`
}

type PartnershipRequest struct {
	ShopId    string `json:"shopId"`
	ShopName  string `json:"shopName"`
	PartnerId string `json:"partnerId"`
	Rights    Rights `json:"rights"`
}

type PartnerName struct {
	ShopName string `json:"shopName"`
}

type PartnerStatus struct {
	ShopName string `json:"shopName"`
	Accept   string `json:"accept"`
}

type tokenClaims struct {
	ShopId    string   `json:"shopId"`
	PartnerId string   `json:"partnerId"`
	Rights    []string `json:"rights"`
	jwt.RegisteredClaims
}
