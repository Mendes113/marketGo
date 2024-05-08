package models
import (
	"context"
	"log"
	"marketgo/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)
// //create invoice
func CreateInvoice(InvoiceResponse InvoiceRequest) error {
    validate = validator.New()
    err := validate.Struct(InvoiceResponse)
    if err != nil {
        return err
    }
	log.Print("CreateInvoice")
	log.Print(InvoiceResponse)
    product, err := GetProduct(InvoiceResponse.Product_uuid)
    if err != nil {
        return err
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
        return err
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
        return err
    }
	log.Print("Invoice created")
	log.Print(invoice)
	log.Print(InvoiceResponse.Product_uuid)
    // Atualizar a quantidade do produto após a compra
    newQuantity := product.Quantity - InvoiceResponse.Quantity
    err = UpdateProductQuantity(InvoiceResponse.Product_uuid, newQuantity)
    if err != nil {
        return err
    }
    return nil
}
func calcTotal(product Product, quantity int) float64 {
	return product.Price * float64(quantity)
}
func createUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uuid.String()
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