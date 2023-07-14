package federation

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ADSP-Project/Federation-Service/database"
	"github.com/ADSP-Project/Federation-Service/types"
	_ "github.com/lib/pq"
)

func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	PublicKey := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(pubkey)}))
	return PublicKey
}

func ExportPrivateKeyAsPemStr(privatekey *rsa.PrivateKey) string {
	privatekey_pem := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}))
	return privatekey_pem
}

func JoinFederation(shopName string, shopDescription string) *rsa.PrivateKey {
	var privKey *rsa.PrivateKey
	var err error
	privateKeyFile := "private.pem"

	// Check if the key file already exists
	if _, err := os.Stat(privateKeyFile); os.IsNotExist(err) {
		// Key file doesn't exist, so we generate a new private key
		privKey, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Fatal(err)
		}

		privateKeyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		})

		// Save the PEM-encoded private key to a file.
		err = ioutil.WriteFile(privateKeyFile, privateKeyPem, 0600)
		if err != nil {
			log.Fatalf("Error writing key to file: %s", err)
		}

	} else {
		// Key file exists, load the PEM-encoded private key from the file.
		privateKeyPem, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			log.Fatalf("Error reading key from file: %s", err)
		}

		// Parse the PEM block.
		block, _ := pem.Decode(privateKeyPem)
		if block == nil || block.Type != "RSA PRIVATE KEY" {
			log.Fatalf("Failed to decode PEM block containing private key")
		}

		// Parse the private key.
		privKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse private key: %s", err)
		}
	}

	privatekey_pem := ExportPrivateKeyAsPemStr(privKey)
	PublicKey := ExportPublicKeyAsPemStr(&privKey.PublicKey)

	newShop := types.Shop{
		Name:        shopName,
		WebhookURL:  fmt.Sprintf("%s/webhook", os.Getenv("VITE_FEDERATION_SERVICE")),
		PublicKey:   PublicKey,
		Description: shopDescription,
	}

	log.Printf("New Shop Private Key is \n %s", privatekey_pem)
	log.Printf("New Shop Public key is \n %s", newShop.PublicKey)

	log.Printf(os.Getenv("AUTH_SERVER") + "/login")
	resp, err := http.PostForm(os.Getenv("AUTH_SERVER")+"/login", url.Values{"name": {shopName}, "webhookURL": {newShop.WebhookURL}, "publicKey": {newShop.PublicKey}})
	if err != nil {
		log.Fatal("Failed to authenticate with auth server")
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	accessToken := result["access_token"]

	jsonData, _ := json.Marshal(newShop)
	req, err := http.NewRequest("POST", os.Getenv("FEDERATION_SERVER")+"/shops", bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", accessToken)

	resp, err = http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Failed to join federation: %v\n", err)
		return privKey
	}
	defer resp.Body.Close()

	fmt.Println("Shop joined the federation")

	return privKey
}

func PollFederationServer() {
	log.Printf("PollFederationServer")
	db := database.DbConn()
	defer db.Close()

	for {
		time.Sleep(10 * time.Second)

		resp, err := http.Get(os.Getenv("FEDERATION_SERVER") + "/shops")
		if err != nil {
			log.Printf("Failed to poll federation server: %v\n", err)
			continue
		}

		var shops []types.Shop
		json.NewDecoder(resp.Body).Decode(&shops)

		var shopsDisplay []types.ShopDisplay
		for _, shop := range shops {
			insForm, err := db.Prepare("INSERT INTO shops(name, webhookURL, publicKey, description) VALUES($1,$2,$3,$4)") // modify this line
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(shop.Name, shop.WebhookURL, shop.PublicKey, shop.Description)
			shopsDisplay = append(shopsDisplay, types.ShopDisplay{Name: shop.Name, WebhookURL: shop.WebhookURL})
		}

		// shopsDisplayJSON, _ := json.MarshalIndent(shopsDisplay, "", "    ")
		// fmt.Printf("Current shops in the federation: \n%s\n", string(shopsDisplayJSON))
	}
}
