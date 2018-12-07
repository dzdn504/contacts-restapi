package main

/*
APIs:
GET 	/contacts 		- get all contacts
GET 	/contact/{id}	- get a single contact with id
POST 	/contact/{id}	- create a new contact
PUT		/contact/{id}	- update existing contact
DELETE	/contact/{id}	- delete contact
*/

func main() {
	a := App{}

	a.Initialize()
	a.Run(":3000")
}
