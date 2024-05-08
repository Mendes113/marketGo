package models

import "github.com/go-playground/validator/v10"

type Market struct {
	UUID     string   `json:"uuid"`
	Name     string   `json:"name" validate:"required,min=3,max=50"`
	Location string   `json:"location" validate:"required"`
	Owner    string   `json:"owner"`
	Email    string   `json:"email" validate:"required,email"`
	Products []Product `json:"products"`
}


type Product struct {
	UUID string `json:"uuid_product" validate:"required"`
	Name string `json:"product_name" validate:"required,min=3,max=50"`
	Price float64 `json:"price" validate:"required"` 
	MarketUUID string `json:"market_uuid" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`

}

type Invoice struct {
	UUID string `json:"uuid_invoice" validate:"required"`
	Product Product `json:"product" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
	Total float64 `json:"total_quantity" validate:"required"`
}

type InvoiceRequest struct {
	Product_uuid string `json:"uuid_product" validate:"required"`
	Quantity int `json:"total_quantity" validate:"required,min=1"`
}

type InvoiceResponse struct {
	UUID string `json:"uuid_invoice" validate:"required"`
	Product Product `json:"product" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
	Total float64 `json:"total_quantity" validate:"required"`
}

var validate *validator.Validate



