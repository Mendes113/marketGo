// models/entities.go
package models

import "github.com/golang/protobuf/ptypes/timestamp"



type Market struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Location string `json:"location"`
	Owner string `json:"owner"`
	Email string `json:"email"`
	Products []Product `json:"products"`
}


type Product struct {
	UUID string `json:"uuid_product"`
	Name string `json:"product_name"`
	Price float64 `json:"price"`
	MarketUUID string `json:"market_uuid"`
	Quantity int `json:"quantity"`

}

type Invoice struct {
	UUID string `json:"uuid_invoice"`
	Product Product `json:"product"`
	Quantity int `json:"quantity"`
	Total float64 `json:"total_quantity"`
}

type InvoiceRequest struct {
	Product_uuid string `json:"uuid_product"`
	Quantity int `json:"total_quantity"`
}

type InvoiceResponse struct {
	UUID string `json:"uuid_invoice"`
	Product Product `json:"product"`
	Quantity int `json:"quantity"`
	Total float64 `json:"total_quantity"`
	timestamp.Timestamp `json:"timestamp"`
}





