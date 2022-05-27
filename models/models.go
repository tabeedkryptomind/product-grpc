package models

type Product struct {
	ID          int64   `json:"id,omitempty"`
	ProductName string  `json:"product_name"`
	Tag         string  `json:"tag"`
	Price       float32 `json:"price"`
}