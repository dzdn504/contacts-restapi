package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
}

var contacts []Contact
var numContacts int

func newContactID() int {
	numContacts++
	return numContacts
}

// InitializeContacts ...
func InitializeContacts() {
	numContacts = 0
	contacts = append(contacts, Contact{
		ID: strconv.Itoa(newContactID()), Name: "John Smith",
		HomeAdress: &Address{City: "Tempe", State: "AZ", ZipCode: "85282"}})
	contacts = append(contacts, Contact{
		ID: strconv.Itoa(newContactID()), Name: "Jane Brown",
		HomeAdress: &Address{City: "Mesa", State: "AZ", ZipCode: "85200"}})
	contacts = append(contacts, Contact{
		ID: strconv.Itoa(newContactID()), Name: "Joe Biden"})
}

// GetAllContactsEndPoint ...
// GET all contacts
func GetAllContactsEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

// GetContactEndPoint ...
// GET a single contact with id
func GetContactEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, contact := range contacts {
		if contact.ID == params["id"] {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	http.Error(w, fmt.Sprintf("Invalid contact id: %s", params["id"]), http.StatusBadRequest)
}

// CreateContactEndPoint ...
// POST a contact
func CreateContactEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact.ID = strconv.Itoa(newContactID())
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// UpdateContactEndPoint ...
// PUT - update existing contact
func UpdateContactEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteContactEndPoint ...
// DELETE contact
func DeleteContactEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// Initialize ...
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/contacts", GetAllContactsEndPoint).Methods("GET")
	a.Router.HandleFunc("/contact/{id}", GetContactEndPoint).Methods("GET")
	a.Router.HandleFunc("/contact", CreateContactEndPoint).Methods("POST")
	a.Router.HandleFunc("/contact/{id}", UpdateContactEndPoint).Methods("PUT")
	a.Router.HandleFunc("/contact/{id}", DeleteContactEndPoint).Methods("DELETE")

	InitializeContacts()
}

// Run ...
func (a *App) Run(addr string) {

	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal(err)
	}
}
