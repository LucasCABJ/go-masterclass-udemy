package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists for: %s\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1 
	fmt.Printf("Contact added: %s\n", name)
}

func findContact(name string) *Contact {
	id, found := contactIndexByName[name]
	if !found {
		return nil
	}
	return &contactList[id]
}

func listContacts() {
	fmt.Println("---- Listing Contacts ----")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for _, contact := range contactList {
		fmt.Printf("Contact: %+v\n", contact)
	}
	fmt.Println()
}

func main() {
	addContact("Alice Wonderland", "alice@alice.com", "123-456-7890")
	addContact("Bob", "bob@bob.com", "987-654-3210")
	addContact("Charlie", "charlie@charlie.com", "555-555-5555")
	addContact("Alice Wonderland", "alice2@alice.com", "111-222-3333")
	listContacts()

	if bob := findContact("Bob"); bob != nil {
		fmt.Printf("Found contact: %s\n", bob.Name)
	} else {
		fmt.Println("Contact not found.")
	}
}
