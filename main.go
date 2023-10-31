package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load file .env : %v", err)
	}

	const serverPort = 8080
	log.Printf("Server starting at http://localhost:%v", serverPort)
	err = http.ListenAndServe(fmt.Sprintf(":%v", serverPort), routeInit())
	if err != nil {
		log.Fatalf("Error starting server : %+v", err)
	}
}
