package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/marifsulaksono/go-midtrans-payment/config"
)

const serverPort = 8080

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load file .env : %v", err)
	}

	conf := config.GetDBConfig()
	conn, err := config.Connect(conf)
	if err != nil {
		log.Fatalf("Connection failed : %+v", err)
	}

	log.Printf("Server starting at http://localhost:%v", serverPort)
	err = http.ListenAndServe(fmt.Sprintf(":%v", serverPort), routeInit(conn))
	if err != nil {
		log.Fatalf("Error starting server : %+v", err)
	}
}
