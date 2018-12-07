package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
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

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/contacts", GetAllContactsEndPoint).Methods("GET")
	a.Router.HandleFunc("/contact/{id}", GetContactEndPoint).Methods("GET")
	a.Router.HandleFunc("/contact/{id}", CreateContactEndPoint).Methods("POST")
	a.Router.HandleFunc("/contact/{id}", UpdateContactEndPoint).Methods("PUT")
	a.Router.HandleFunc("/contact/{id}", DeleteContactEndPoint).Methods("DELETE")

	InitializeContacts()

}

func (a *App) Run(addr string) {

	if err := http.ListenAndServe(addr, a.Router); err != nil {
		log.Fatal(err)
	}
}
