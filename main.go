package main

import (
	"ankets_and_clients/Controllers"
	"ankets_and_clients/Databases"
	"ankets_and_clients/Domain"
	"ankets_and_clients/Services"
	"encoding/json"
	"net/http"
)

func anketaHandler(w http.ResponseWriter, r *http.Request) {
	anketa := Domain.Anketa{
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

func NewClients(name string, rating int) Domain.Clients {
	return Domain.Clients{
		Name:   name,
		Rating: rating,
	}
}

func main() {

	http.HandleFunc("/examleanketa", anketaHandler)
	http.HandleFunc("/exampleclient", clientsHandler)

	setapAnketaHandlers()
	setapClientHandler()

	http.ListenAndServe(":8080", nil)

}

func setapAnketaHandlers() {
	db := Databases.NewDatabase()
	anketaService := Services.NewAnketaService(db)
	anketaController := Controllers.NewAnketaController(anketaService)
	http.HandleFunc("/anketa/create", anketaController.CreateAnketaHandler)
	http.HandleFunc("/anketa/delete", anketaController.DeleteAnketaHandler)
	http.HandleFunc("/anketa/get", anketaController.GetAnketaHandler)
}
func setapClientHandler() {
	bd := Databases.NewDatabase_clients()
	clientService := Services.NewClientService(bd)
	clientController := Controllers.NewClientController(clientService)
	http.HandleFunc("/client/create", clientController.CreateClientHandler)
	http.HandleFunc("/client/delete", clientController.DeleteClientHandler)
	http.HandleFunc("/client/get", clientController.GetClientHandler)
}
