package globals

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ShopName string
var FederationServer string
var AuthServer string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	FederationServer = os.Getenv("FEDERATION_SERVER")
	AuthServer = os.Getenv("AUTH_SERVER")
}