package main

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ADSP-Project/Federation-Service/federation"
	"github.com/ADSP-Project/Federation-Service/globals"
	"github.com/ADSP-Project/Federation-Service/handlers"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var privKey *rsa.PrivateKey

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run shop.go [port] [name] [desc]")
	}

	globals.LoadEnv()

	port := os.Args[1]
	globals.ShopName = os.Args[2]
	shopDescription := os.Args[3]

	router := mux.NewRouter()
	router.HandleFunc("/webhook", handlers.HandleWebhook).Methods("POST")
	router.HandleFunc("/api/v1/partnerships/request", func(w http.ResponseWriter, r *http.Request) {
		handlers.RequestPartnership(w, r, privKey)
	}).Methods("POST")
	router.HandleFunc("/api/v1/partnerships/process", handlers.ProcessPartnership).Methods("POST")
	router.HandleFunc("/api/v1/partnerships/accept", handlers.AcceptPartnership).Methods("POST")
	router.HandleFunc("/api/v1/partnerships/deny", handlers.DenyPartnership).Methods("POST")
	router.HandleFunc("/api/v1/partnerships/notify", handlers.NotifyHandler).Methods("POST")
	router.HandleFunc("/api/v1/partners", handlers.GetPartners).Methods("GET")
	router.HandleFunc("/api/v1/shop", handlers.GetShop).Methods("GET")
	router.HandleFunc("/api/v1/shops", handlers.GetShops).Methods("GET")

	privKey = federation.JoinFederation(globals.ShopName, shopDescription)
	go federation.PollFederationServer()

	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"*"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
