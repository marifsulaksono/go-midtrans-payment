package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverPort := 8080
	log.Printf("Server starting at http://localhost:%v", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	if err != nil {
		log.Fatalf("Error starting server : %+v", err)
	}
}
