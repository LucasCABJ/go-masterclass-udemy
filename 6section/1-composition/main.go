package main

import "fmt"

// Composition --> has-a relationship
// Example: Car has a Engine, Car has a Wheels, Car has a Doors, etc.

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	return a.Street + ", " + a.City + ", " + a.State + " " + a.ZipCode
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address // embedded
	ShippingAddress Address // embedded
}

func (c Customer) PrintDetails() {
	fmt.Println("Customer ID: ", c.CustomerID)
	fmt.Println("Name: ", c.Name)
	fmt.Println("Email: ", c.Email)
	fmt.Println("Billing Address: ", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address: ", c.ShippingAddress.FullAddress())
}

func main() {

	customer := Customer{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "john@john.com",
		BillingAddress: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
		},
		ShippingAddress: Address{
			Street:  "456 Elm St",
			City:    "Othertown",
			State:   "NY",
			ZipCode: "67890",
		},
	}
	customer.PrintDetails()

	// -- customer with same billing and shipping address

	address := Address{
		Street:  "789 Oak St",
		City:    "Sometown",
		State:   "TX",
		ZipCode: "54321",
	}
	customer2 := Customer{
		CustomerID:      2,
		Name:            "Jane Smith",
		Email:           "jane@jane.com",
		BillingAddress: address,
		ShippingAddress: address,
	}
	customer2.PrintDetails()
}
