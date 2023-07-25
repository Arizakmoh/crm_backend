package database

import "crm_backend/dtos"

// Customer 1
var Customers = []dtos.Customer{
	{
		ID:        "1",
		Name:      "Abdirizak Abdullahi",
		Role:      "Manager",
		Email:     "abdirizak@noemail.com",
		Phone:     "612-555-5551",
		Contacted: true,
	},
	// Customer 2
	{
		ID:        "2",
		Name:      "Smith John",
		Role:      "Employee",
		Email:     "smith@noemail.com",
		Phone:     "612-555-5552",
		Contacted: false,
	},
	// Customer 3
	{
		ID:        "3",
		Name:      "Alex Morgam",
		Role:      "Customer",
		Email:     "alex@noemail.com",
		Phone:     "612-555-5553",
		Contacted: true,
	},
	// Customer 4
	{
		ID:        "4",
		Name:      "Jenefer Lps",
		Role:      "Customer",
		Email:     "Jenefer@noemail.com",
		Phone:     "612-555-5554",
		Contacted: false,
	},

	// Customer 5
	{
		ID:        "5",
		Name:      "Moha Really",
		Role:      "Customer",
		Email:     "moha@noemail.com",
		Phone:     "612-555-5554",
		Contacted: false,
	},
}
