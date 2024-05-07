package models

import (
	"context"
	"log"
	"marketgo/db"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// //create invoice
func TestCreateInvoice(t *testing.T)  {
    InvoiceResponse := InvoiceRequest{
        Product_uuid: "1",
        Quantity:     2,
    }


	log.Print("CreateInvoice")
	log.Print(InvoiceResponse)
    product, err := GetProduct(InvoiceResponse.Product_uuid)
    if err != nil {
       t.Fatal("Erro ao buscar produto")
        
    }

    // Criar fatura
    invoice := Invoice{
        UUID:     createUUID(),
        Product:  *product,
        Quantity: InvoiceResponse.Quantity,
        Total:    calcTotal(*product, InvoiceResponse.Quantity),
    }

    // Abrir conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        
    }
    defer db.Close(client)
	
    // Inserir fatura na coleção "invoices"
    collection := client.Database("market").Collection("invoices")
    _, err = collection.InsertOne(context.Background(), bson.M{
        "UUID":     invoice.UUID,
        "Product":  invoice.Product,
        "Quantity": invoice.Quantity,
        "Total":    invoice.Total,
    })
    if err != nil {
        
    }
	log.Print("Invoice created")
	log.Print(invoice)
	log.Print(InvoiceResponse.Product_uuid)
    // Atualizar a quantidade do produto após a compra
    newQuantity := product.Quantity - InvoiceResponse.Quantity
    err = UpdateProductQuantity(InvoiceResponse.Product_uuid, newQuantity)
    if err != nil {
        
    }

   assert.NoError(t, err)
}



func TestCalcTotal(t *testing.T) {
    product := Product{
        UUID:     "1",
        Name:     "Produto de Teste",
        Price:    10.0,
        Quantity: 100,
    }
    quantity := 2

	total := product.Price * float64(quantity)
    assert.Equal(t, total, calcTotal(product, quantity))
}


func TestCreateUUID(t *testing.T) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Fatal("Erro ao criar UUID")
	}
	uuidString := uuid.String()
    assert.NotEmpty(t, uuidString)
}




//  func InsertMarket(market Market) error {
// 	client, err := db.OpenConnection()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close(client)

// 	collection := client.Database("market").Collection("markets")
// 	_, err = collection.InsertOne(context.Background(), bson.M{
// 		"UUID":     market.UUID,
// 		"Name":     market.Name,
// 		"Location": market.Location,
// 		"Owner":    market.Owner,
// 		"Email":    market.Email,
// 	})

// 	return err
// }