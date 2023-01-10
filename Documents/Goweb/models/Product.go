package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type Products struct {
	Products []Product `json:"products"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
