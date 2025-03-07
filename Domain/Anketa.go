package Domain

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
