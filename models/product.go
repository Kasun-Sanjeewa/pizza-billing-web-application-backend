package models

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Barcode string  `json:"barcode"`
	Price   float64 `json:"price"`
	Img     string  `json:"img"`
}
