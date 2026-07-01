package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" && a.State == "" && a.ZipCode == "" {
		return "No address provided"
	}
	return a.Street + ", " + a.City + ", " + a.State + " " + a.ZipCode
}

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) FullContactInfo() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) PrintDetails() {
	fmt.Println("Company Name: ", c.Name)
	fmt.Println("Business Type: ", c.BusinessType)
	fmt.Println("Address street:", c.Street)
	fmt.Println("Address: ", c.FullAddress())
	fmt.Println("Contact Info: ", c.FullContactInfo())
}

func main() {

}
