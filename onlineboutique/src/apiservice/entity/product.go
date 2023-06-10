package entity

type priceUsd struct {
	CurrencyCode int `json:"currencyCode"`
	Units        int `json:"units"`
	Nanos        int `json:"nanos"`
}

type categories []string

// Product defines a structure for an item in product catalog
type Product struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	PriceUsd    priceUsd   `json:"priceUsd"`
	IsAvailable bool       `json:"isAvailable"`
	Categories  categories `json:"categories"`
}
