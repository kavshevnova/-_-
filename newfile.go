package main

type Clients struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

func NewClients(name string, rating int) Clients {
	return Clients{
		Name:   name,
		Rating: rating,
	}
}
