package main

import (
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

func main() {
	router := mux.NewRouter()
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
