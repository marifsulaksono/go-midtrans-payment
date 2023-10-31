package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const serverPort = 8080
	log.Printf("Server starting at http://localhost:%v", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), routeInit())
	if err != nil {
		log.Fatalf("Error starting server : %+v", err)
	}
}
