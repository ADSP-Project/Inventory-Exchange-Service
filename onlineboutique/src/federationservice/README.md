# Federation Service Setup

This repository contains the code for setting up a federated marketplace using Go and PostgreSQL.

## Prerequisites

- Go

## Getting Started

1. Clone the repository:

   `git clone https://github.com/ADSP-Project/Federation-Service.git`

   `cd Federation-Service`

2. Install dependencies:

   - Make sure to initialize go.mod to manage dependencies:
   
      `go mod init github.com/ADSP-Project/Federation-Service`

   - Fetch and arrange them into newly generated go.mod:
   
      `go mod tidy`

   - Finally, install the required dependencies with the following command:
   
      `go mod download`


3. Configure environment variables:

   - Rename the .env.example file to .env.
   - Update the database credentials in the .env file to match your PostgreSQL setup.

4. If you set already Hub up, then you have your user already with a password. Connect to Postgres with `psql` and create new database 'federation_service':

      `CREATE DATABASE federation_service;`
   
   Grant rights to your user:

      `GRANT ALL PRIVILEGES ON DATABASE federation_service TO your_username;`

5. Now connect as your user to DB for creating tables:
   
      `psql -d federation_service -U your_username`

   Create tables `shops` and `partners`:

      `CREATE TABLE shops ( id SERIAL PRIMARY KEY, name VARCHAR(255) UNIQUE, description VARCHAR(255), webhookURL VARCHAR(255), publicKey VARCHAR(1024));`

      `CREATE TABLE partners ( shopId SERIAL PRIMARY KEY, shopName VARCHAR(255), canEarnCommission BOOLEAN, canShareInventory BOOLEAN, canShareData BOOLEAN, canCoPromote BOOLEAN, canSell BOOLEAN, requestStatus VARCHAR(1024));`

6. Simulate a shop joining the federation:
   - To simulate a shop joining the federation, open a new terminal and run the following command:

     `go run shop.go [port] [name] [description]`

     Replace `[port]` with the desired port number and `[name]` with the name of the shop.

     This will start a shop server that automatically joins the federation by sending a POST request to the federation server.

     **Important:** Hub and AuthServer from Federation-Hub should be running so that shop can join Federation.

7. Additional Notes:
   - You can run multiple instances of the shop server by providing different port numbers and shop names.
