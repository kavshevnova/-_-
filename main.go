package main

import (
	"encoding/json"
	"net/http"
)

type Ticket struct {
	Otkuda string  `json:"otkuda"`
	Kuda   string  `json:"kuda"`
	Price  float64 `json:"price"`
}

func ticketHandler(w http.ResponseWriter, r *http.Request) {
	ticket := Ticket{
		Otkuda: "Москва",
		Kuda:   "Санкт-Петербург",
		Price:  4499.99,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	service := NewServices(23, "Включено", "Ручная кладь")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func main() {
	http.HandleFunc("/ticket", ticketHandler)
	http.HandleFunc("/service", serviceHandler)

	http.ListenAndServe(":8080", nil)
}
