package main

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
