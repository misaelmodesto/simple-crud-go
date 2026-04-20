package model

type Product struct {
	ID    int     `json:"id_prooduct"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
