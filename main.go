package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
APIs:
GET 	/contacts 		- get all contacts
GET 	/contact/{id}	- get a single contact with id
POST 	/contact/{id}	- create a new contact
PUT		/contact/{id}	- update existing contact
DELETE	/contact/{id}	- delete contact
*/

type Contact struct {
	ID         string   `json:"id,omitempty`
	Name       string   `json:"name,omitempty`
	HomeAdress *Address `json:"address,omitempty`
}

type Address struct {
	City    string `json:"city,omitempty`
	State   string `json:"state,omitempty`
	ZipCode string `json:"zip_code,omitempty`
}

var contacts []Contact

func InitializeContacts() {
	contacts = append(contacts, Contact{
		ID: "1", Name: "John Smith",
		HomeAdress: &Address{City: "Tempe", State: "AZ", ZipCode: "85282"}})
	contacts = append(contacts, Contact{
		ID: "2", Name: "Jane Brown",
		HomeAdress: &Address{City: "Mesa", State: "AZ", ZipCode: "85200"}})
	contacts = append(contacts, Contact{
		ID: "3", Name: "Joe Biden"})
}

// GET all contacts
func GetAllContactsEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	json.NewEncoder(responseWriter).Encode(contacts)
}

// GET a single contact with id
func GetContactEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(responseWriter, "not implemented yet !")
}

// POST a contact
func CreateContactEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(responseWriter, "not implemented yet !")
}

// PUT - update existing contact
func UpdateContactEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(responseWriter, "not implemented yet !")
}

// DELETE contact
func DeleteContactEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(responseWriter, "not implemented yet !")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contacts", GetAllContactsEndPoint).Methods("GET")
	router.HandleFunc("/contact/{id}", GetContactEndPoint).Methods("GET")
	router.HandleFunc("/contact/{id}", CreateContactEndPoint).Methods("POST")
	router.HandleFunc("/contact/{id}", UpdateContactEndPoint).Methods("PUT")
	router.HandleFunc("/contact/{id}", DeleteContactEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}

	InitializeContacts()
}
