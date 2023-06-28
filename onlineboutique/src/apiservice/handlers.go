package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func (fe *apiServer) productsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := fe.getProducts(r.Context())
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(b)
}

type Shop struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type ShopsData struct {
	Shops []Shop `json:"shops"`
}

func getShopMap() (map[string]Shop, error) {
	curdir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current dir:", err)
		return nil, err
	}
	fmt.Println(curdir)
	fileBytes, err := ioutil.ReadFile("shops.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var data ShopsData
	// Unmarshal the JSON data into the ShopData struct
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	shopMap := make(map[string]Shop)
	for _, value := range data.Shops {
		shopMap[value.ID] = value
	}

	return shopMap, nil
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
	} `json:"data"`
}

func externalProductHandler(w http.ResponseWriter, r *http.Request) {
	shops, err := getShopMap()
	fmt.Println(shops)
	if err != nil {
		fmt.Println("Error getting shop map:", err)
		return
	}

	//log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	if id == "" {
		fmt.Println("Wrong path", err)
		return
	}

	s := strings.Split(id, ":")
	store, _ := s[0], s[1]

	// Check if a shop exists among collaborators
	_, exists := shops[store]
	fmt.Println("Exists?", exists)
	if exists {
		// Example request
		// client := &http.Client{}
		// resp, err := client.Get("https://api.example.com/data")
		// if err != nil {
		// 	fmt.Println("Error sending GET request:", err)
		// 	return
		// }
		// defer resp.Body.Close()

		// // Read the response body
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println("Error reading response body:", err)
		// 	return
		// }
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusFound)
		response := Response{
			Status:  "OK",
			Message: "Request processed successfully.",
			Data: struct {
				Result string `json:"result"`
			}{
				Result: "Success",
			},
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		w.Write(jsonData)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

func postExternalOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid request method. Only POST requests are allowed.")
		return
	}

	// Read the request body
	// Assuming the request body is in JSON format
	// You can replace this with your own logic to parse the request body
	// For example, using the encoding/json package
	var requestBody ExternalOrderData
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing request body: %v", err)
		return
	}
	fmt.Println(requestBody)

	// Process the request and generate a response
	response := "Received a POST request"

	// Write the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response)
}
