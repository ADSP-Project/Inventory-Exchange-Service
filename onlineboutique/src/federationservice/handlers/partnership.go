package handlers

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ADSP-Project/Federation-Service/database"
	"github.com/ADSP-Project/Federation-Service/globals"
	"github.com/ADSP-Project/Federation-Service/types"
	"github.com/golang-jwt/jwt/v5"
)

func ProcessPartnership(w http.ResponseWriter, r *http.Request) {
	var request types.PartnershipRequest

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString) // prints the actual request body

	// you need to put back the body content into the request
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		http.Error(w, "No Authorization header provided", http.StatusBadRequest)
		return
	}

	splitToken := strings.Split(authorizationHeader, "Bearer ")
	if len(splitToken) != 2 {
		http.Error(w, "Malformed Authorization header", http.StatusBadRequest)
		return
	}

	tokenString := splitToken[1] // Here is your token
	shopName := request.ShopName

	publicKeyStr, err := getPublicKeyFromDB(shopName)
	if err != nil {
		detailedError := fmt.Errorf("Failed to retrieve public key: %w", err)
		fmt.Printf("%+v\n", detailedError)
		http.Error(w, detailedError.Error(), http.StatusInternalServerError)
		return
	}

	publicKeyBlock, rest := pem.Decode([]byte(publicKeyStr))
	if publicKeyBlock == nil {
		detailedError := fmt.Errorf("Failed to decode public key. Remaining data: %s", string(rest))
		fmt.Printf("%+v\n", detailedError)
		http.Error(w, "Failed to decode public key", http.StatusInternalServerError)
		return
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Printf("Failed to parse public key: %v\n", err)
		http.Error(w, "Failed to parse public key", http.StatusInternalServerError)
		return
	}

	// Now we can use publicKey in jwt.Parse
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		log.Printf("Error while validating the token: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		db := database.DbConn()
		shopId := claims["shopId"].(string)
		// partnerId := claims["partnerId"].(string)

		partner, err := database.GetShopByName(shopName)
		if err != nil {
			log.Printf("Failed to get shop with id %s: %v\n", shopId, err)
			http.Error(w, "Failed to process partnership", http.StatusInternalServerError)
			return
		}

		// Insert new partnership entry
		sqlStatement := `
			INSERT INTO partners (shopid, shopname, canearncommission, canshareinventory, cansharedata, cancopromote, cansell, requeststatus)
			VALUES ($1, $2, $3, $4, $5, $6, $7, 'pending')
		`
		fmt.Print(request.Rights)
		_, err = db.Exec(sqlStatement, partner.Id, partner.Name, request.Rights.CanEarnCommission, request.Rights.CanShareInventory, request.Rights.CanShareData, request.Rights.CanCoPromote, request.Rights.CanSell)
		if err != nil {
			log.Printf("Failed to insert new partnership: %v\n", err)
			http.Error(w, "Failed to process partnership", http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("Invalid token")
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}

func sendNotification(shopName string, w http.ResponseWriter, accepted string) {
	db := database.DbConn()
	defer db.Close()

	log.Printf("Creating notify request...")
	log.Printf("Getting webhook url of partner shop %s", shopName)
	var partnerWebhookURL string
	err := db.QueryRow("SELECT webhookurl FROM shops WHERE name = $1", shopName).Scan(&partnerWebhookURL)
	if err != nil {
		http.Error(w, "Shop not found", http.StatusBadRequest)
		log.Printf("Shop not found")
		return
	}

	url, err := url.Parse(partnerWebhookURL)
	if err != nil {
		http.Error(w, "Error parsing the partner webhook URL", http.StatusInternalServerError)
		log.Printf("Error parsing webhook URLL")
		return
	}

	// removes the '/webhook' part
	url.Path = ""

	newURL := url.String()

	var StatusNotify types.PartnerStatus
	StatusNotify.ShopName = globals.ShopName
	StatusNotify.Accept = accepted
	log.Printf("making json from %s and request status", StatusNotify.ShopName)
	jsonData, err := json.Marshal(StatusNotify)
	if err != nil {
		http.Error(w, "Failed to create JSON body", http.StatusInternalServerError)
		log.Printf("Failed to create JSON body")
		return
	}

	log.Printf("Sending POST notification to partner at webhook %s", newURL)
	req, err := http.NewRequest("POST", newURL+"/api/v1/partnerships/notify", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create notification", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to send partnership accept notification", http.StatusInternalServerError)
		return
	}
}

func AcceptPartnership(w http.ResponseWriter, r *http.Request) {
	log.Printf("Changing status of partnership...")

	db := database.DbConn()
	defer db.Close()

	var request types.PartnerName

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	//fmt.Printf("Accepting partnership with %s", request.ShopId)

	sqlStatement := `
	UPDATE partners
	SET requestStatus = 'accepted'
	WHERE shopName = $1;
	`
	_, err = db.Exec(sqlStatement, request.ShopName)
	if err != nil {
		log.Printf("Failed to update partnership status: %v\n", err)
		http.Error(w, "Failed to process acceptPartnership", http.StatusInternalServerError)
		return
	}

	//Create http request to notify other shop
	accepted := "true"
	sendNotification(request.ShopName, w, accepted)

	fmt.Fprintln(w, "Partnership successfully accepted")
	log.Printf("Partnership successfully accepted")

}

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received acceptance...Changing status of partnership...")

	db := database.DbConn()
	defer db.Close()

	var request types.PartnerStatus

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	if request.Accept == "true" {
		fmt.Printf("Changing partnership status with %s to accept", request.ShopName)

		sqlStatement := `
		UPDATE partners
		SET requestStatus = 'accepted'
		WHERE shopName = $1;
		`
		_, err = db.Exec(sqlStatement, request.ShopName)
		if err != nil {
			log.Printf("Failed to update partnership status: %v\n", err)
			http.Error(w, "Failed to process acceptPartnership", http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Printf("Changing partnership status with %s to denied", request.ShopName)

		sqlStatement := `
		UPDATE partners
		SET requestStatus = 'denied'
		WHERE shopName = $1;
		`
		_, err = db.Exec(sqlStatement, request.ShopName)
		if err != nil {
			log.Printf("Failed to update partnership status: %v\n", err)
			http.Error(w, "Failed to process acceptPartnership", http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintln(w, "Partnership Status successfully updated")
	log.Printf("Partnership status successfully updated")
}

func DenyPartnership(w http.ResponseWriter, r *http.Request) {
	log.Printf("Changing status of partnership...")

	db := database.DbConn()
	defer db.Close()
	var request types.PartnerName

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	//fmt.Printf("Denying partnership with %s", request.ShopId)

	sqlStatement := `
	UPDATE partners
	SET requestStatus = 'denied'
	WHERE shopName = $1;
	`
	_, err = db.Exec(sqlStatement, request.ShopName)
	if err != nil {
		log.Printf("Failed to update partnership status: %v\n", err)
		http.Error(w, "Failed to process DenyPartnership", http.StatusInternalServerError)
		return
	}

	//Create http request to notify other shop
	accepted := "false"
	sendNotification(request.ShopName, w, accepted)

	fmt.Fprintln(w, "Partnership successfully denied")
	log.Printf("Partnership successfully denied")

}

func RequestPartnership(w http.ResponseWriter, r *http.Request, privKey *rsa.PrivateKey) {
	var request types.PartnershipRequest
	json.NewDecoder(r.Body).Decode(&request)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"shopId":    request.ShopId,
		"partnerId": request.PartnerId,
		"rights":    request.Rights,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(privKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while signing the token: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Generated JWT for partnership request: %s\n", tokenString)

	db := database.DbConn()
	defer db.Close()

	var partnerWebhookURL string
	err = db.QueryRow("SELECT webhookurl FROM shops WHERE id = $1", request.PartnerId).Scan(&partnerWebhookURL)
	if err != nil {
		http.Error(w, "Shop not found", http.StatusBadRequest)
		return
	}

	url, err := url.Parse(partnerWebhookURL)
	if err != nil {
		http.Error(w, "Error parsing the partner webhook URL", http.StatusInternalServerError)
		return
	}

	// removes the '/webhook' part
	url.Path = ""

	newURL := url.String()

	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Failed to create JSON body", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", newURL+"/api/v1/partnerships/process", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)
	fmt.Printf("Parsed request: %+v\n", request)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to send partnership request", http.StatusInternalServerError)
		return
	}

	partner, err := database.GetShopById(request.PartnerId)
	sqlStatement := `
	INSERT INTO partners (shopid, shopname, canearncommission, canshareinventory, cansharedata, cancopromote, cansell, requeststatus)
	VALUES ($1, $2, $3, $4, $5, $6, $7, 'sent')`
	_, err = db.Exec(sqlStatement, partner.Id, partner.Name, request.Rights.CanEarnCommission, request.Rights.CanShareInventory, request.Rights.CanShareData, request.Rights.CanCoPromote, request.Rights.CanSell)
	if err != nil {
		log.Printf("Failed to insert new partnership: %v\n", err)
		http.Error(w, "Failed to process partnership", http.StatusInternalServerError)
		return
	}

	// Success
	fmt.Fprintln(w, "Partnership request successfully sent")
}

func getPublicKeyFromDB(shopName string) (string, error) {
	db := database.DbConn()
	defer db.Close()

	var publicKey string
	row := db.QueryRow("SELECT publicKey FROM shops WHERE name = $1", shopName)
	err := row.Scan(&publicKey)

	if err != nil {
		return "", fmt.Errorf("error getting public key from DB: %w", err)
	}

	return publicKey, nil
}
