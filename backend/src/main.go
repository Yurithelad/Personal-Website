package main

import (
	"log"
	"net/http"
)

func main() {

	port := ":8080"
	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error 500: problem starting server", err)
	}
}
