package main

import (
	"log"
	"net/http"

	"crm_backend/api"
)

func main() {
	router := api.Api_Routes()
	http.Handle("/", router)

	// Start the server
	log.Println("Starting crm_Backend server on port 3000 ")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
