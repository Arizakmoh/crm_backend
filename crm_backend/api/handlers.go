package api

import (
	database "crm_backend/database"
	"crm_backend/dtos"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// API 1
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.Customers)
}

// API 2
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	customerID := params["id"]

	fmt.Println("Received customerID:", customerID)

	for _, customer := range database.Customers {
		fmt.Println("Checking customer with ID:", customer.ID)
		if customer.ID == customerID {
			fmt.Println("Found matching customer:", customer)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "Customer not found"}`)
}

// API 3
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Received Create Request")
	var newCustomer dtos.Customer

	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid request data"}`)
		return
	}

	//newID := generateNewID()
	//newCustomer.ID = newID

	database.Customers = append(database.Customers, newCustomer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func generateNewID() string {
	uniqueID := fmt.Sprintf("%d", len(database.Customers)+1)
	return uniqueID
}

// API 4
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	customerID := params["id"]
	fmt.Println("Received Update Request", customerID)

	// Fech Customer from Database
	for i, customer := range database.Customers {
		if customer.ID == customerID {
			var updatedCustomer dtos.Customer
			err := json.NewDecoder(r.Body).Decode(&updatedCustomer)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `{"error": "Invalid request data"}`)
				return
			}

			// Update the existing customer in the Database
			database.Customers[i] = updatedCustomer

			// Return the updated customer in the response as json.
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "Customer not found"}`)
}

// API 5
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	customerID := params["id"]

	// Fetch Customer From Database
	for i, customer := range database.Customers {
		if customer.ID == customerID {
			// Remove the customer from the Database
			database.Customers = append(database.Customers[:i], database.Customers[i+1:]...)

			// Return success response.
			fmt.Fprintf(w, `{"message": "Customer deleted successfully"}`)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"error": "Customer not found"}`)
}
