package main

import (
	"encoding/json"
	"net/http"
)

type Anketa struct {
	Name        string `json:"name"`
	Id          int    `json:"id"`
	City        string `json:"city"`
	Age         int    `json:"age"`
	Weight      int    `json:"weight"`
	Height      int    `json:"height"`
	Boobs       int    `json:"bobs"`
	HairColor   string `json:"hair_color"`
	Nationality string `json:"nationality"`
	District    string `json:"district"`
	Price       int    `json:"price"`
}

func anketaHandler(w http.ResponseWriter, r *http.Request) {
	anketa := Anketa{
		Name:        "Мишель",
		Id:          123,
		City:        "Moscow",
		Age:         25,
		Weight:      49,
		Height:      175,
		Boobs:       2,
		HairColor:   "Блондинка",
		Nationality: "Русская",
		District:    "ЦСКА",
		Price:       20000,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anketa)
}

func clientsHandler(w http.ResponseWriter, r *http.Request) {
	client := NewClients("Банановый чизкейк", 80)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func main() {

	http.HandleFunc("/anketa", anketaHandler)
	http.HandleFunc("/client", clientsHandler)

	db := NewDatabase()
	anketaService := NewAnketaService(db)
	anketaController := NewAnketaController(anketaService)
	http.HandleFunc("/anketa/create", anketaController.CreateAnketaHandler)
	http.HandleFunc("/anketa/delete", anketaController.DeleteAnketaHandler)
	http.HandleFunc("/anketa/get", anketaController.GetAnketaHandler)
	bd := NewDatabase_clients()
	clientService := NewClientService(bd)
	clientController := NewClientController(clientService)
	http.HandleFunc("/client/create", clientController.CreateClientHandler)
	http.HandleFunc("/anketa/delete", clientController.DeleteClientHandler)
	http.HandleFunc("/anketa/get", clientController.GetClientHandler)

	http.ListenAndServe(":8080", nil)

}
