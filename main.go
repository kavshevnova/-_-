package main

import (
	"encoding/json"
	"net/http"
)

// в этих трех файлах описать структуры
// создание анкеты и удаление анкеты
// хранить объекты в слайсе базы данных
//в анкете сервис создать структуру и интерфейс с методами доступными из внею структура должна реализовывать этот интерфейс.
//	В контроллере описать еще раз этот интерфейс и работать с объектом который будет передаваться в мейн

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

	http.ListenAndServe(":8080", nil)
}
