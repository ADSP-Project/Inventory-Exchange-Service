package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbConn() (db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbDriver := "postgres"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName)

	db, err = sql.Open(dbDriver, dbInfo)
	if err != nil {
		fmt.Println("Error with database")
		fmt.Println(err.Error())
		panic(err.Error())
	}

	return db
}

func GetWebhookURL(shopId string) (string, error) {
	db := DbConn()

	var webhookURL string
	err := db.QueryRow("SELECT webhookurl FROM shops WHERE id = $1", shopId).Scan(&webhookURL)

	db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned - handle according to your requirements
			log.Printf("No shop found with id: %s\n", shopId)
		}
		return "", err
	}

	return webhookURL, nil
}

type ShopInfo struct {
	Id   string
	Name string
}

func GetShopByName(name string) (ShopInfo, error) {
	db := DbConn()
	defer db.Close()

	var shop ShopInfo
	err := db.QueryRow("SELECT id, name FROM shops WHERE name = $1", name).Scan(&shop.Id, &shop.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned - handle according to your requirements
			log.Printf("No shop found with name: %s\n", name)
		}
		return shop, err
	}

	return shop, nil
}

func GetShopById(id string) (ShopInfo, error) {
	db := DbConn()
	defer db.Close()

	var shop ShopInfo
	err := db.QueryRow("SELECT id, name FROM shops WHERE id = $1", id).Scan(&shop.Id, &shop.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned - handle according to your requirements
			log.Printf("No shop found with id: %s\n", id)
		}
		return shop, err
	}

	return shop, nil
}

type Shop struct {
	Id          string
	Name        string
	WebhookURL  string
	PublicKey   string
	Description string
}

func GetAllShops() ([]Shop, error) {
	db := DbConn()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, webhookurl, publickey, description FROM shops")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []Shop
	for rows.Next() {
		var shop Shop
		if err := rows.Scan(&shop.Id, &shop.Name, &shop.WebhookURL, &shop.PublicKey, &shop.Description); err != nil {
			return nil, err
		}
		shops = append(shops, shop)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shops, nil
}
