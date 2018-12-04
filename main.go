package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
APIs:
GET 	/contacts 		- get all contacts
GET 	/contacts/{id}	- get a single contact with id
POST 	/contacts/{id}	- create a new contact
PUT		/contacts/{id}	- update existing contact
DELETE	/contacts/{id}	- delete contact
*/

// GET all contacts
func GetAllContactsEndPoint(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(responseWriter, "not implemented yet !")
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
	router.HandleFunc("/contacts/{id}", GetContactEndPoint).Methods("GET")
	router.HandleFunc("/contacts/{id}", CreateContactEndPoint).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContactEndPoint).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContactEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
