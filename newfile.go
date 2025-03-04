package main

type Services struct {
	Place   int    `json:"place"`
	Food    string `json:"food"`
	Baggage string `json:"baggage"`
}

func NewServices(place int, food string, baggage string) Services {
	return Services{
		Place:   place,
		Food:    food,
		Baggage: baggage,
	}
}
