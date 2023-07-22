package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Api_Routes() *mux.Router {
	router := mux.NewRouter()
	// Get All Customer
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	// Get Specific Customer
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	// Post a Customer
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	// Put a Customer
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	// Delete a Customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	// Inial static Page
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	return router
}
