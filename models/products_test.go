package models

import (
	"context"
	"encoding/json"
	
	"marketgo/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)



func TestGetAllProdutcts(t *testing.T) {
    // Abrir conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)

    // Buscar documentos da coleção
    var documents []bson.M
    collection := client.Database("market").Collection("products")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
       t.Fatalf("Erro ao buscar documentos da coleção: %v", err)
    }
    defer cursor.Close(context.Background())

    // Iterar sobre os documentos retornados pelo cursor
    for cursor.Next(context.Background()) {
        var document bson.M
        if err := cursor.Decode(&document); err != nil {
            t.Fatalf("Erro ao decodificar documento: %v", err)
        }
        documents = append(documents, document)
    }

    if err := cursor.Err(); err != nil {
        t.Fatalf("Erro no cursor: %v", err)
    }

    // Serializar os documentos para JSON
    jsonData, err := json.Marshal(documents)
    if err != nil {
        t.Fatalf("Erro ao serializar documentos para JSON: %v", err)
    }

    assert.NotEmpty(t, documents)
    assert.NotEmpty(t, jsonData)
}




func TestInsertProduct(t *testing.T) {
	// Abrir conexão com o banco de dados
    product := Product{
        UUID:        "4",
        Name:        "Produto 1",
        Price:       10.0,
        Quantity:    10,
        MarketUUID:  "1",
    }

	client, err := db.OpenConnection()
	if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close(client)
    
    // Verificar se o produto já existe
    existingProduct, err := GetProduct(product.UUID)
    // log.Println(existingProduct)
    if err == nil && existingProduct != nil {
       t.Fatal("Product already exists")
    }

	// Inserir documento na coleção
	collection := client.Database("market").Collection("products")
	_, err = collection.InsertOne(context.Background(), product)
	if err != nil {
		t.Fatalf("Erro ao inserir documento na coleção: %v", err)
	}

    assert.NotEmpty(t, collection)
    assert.NotEmpty(t, product)
	// // Adicionar produto ao mercado
	// AppendProductToMarket(product.MarketUUID, product)

	assert.NoError(t, err)
    deleteObject(product.UUID)
}


func deleteObject(uuid string) error {
    // Abrir conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        return err
    }
    defer db.Close(client)

    // Deletar documento na coleção
    collection := client.Database("market").Collection("products")
    _, err = collection.DeleteOne(context.Background(), bson.M{"uuid": uuid})
    if err != nil {
        return err
    }

    return nil
}




func TestGetProduct(t *testing.T)  {
	// Abrir conexão com o banco de dados
    uuid := "1asdasdasdxczxasda"
	client, err := db.OpenConnection()
	if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}

    // log.Print(uuid)
    // log.Print("GetProduct")
	defer db.Close(client)
	var document Product
    collection := client.Database("market").Collection("products")
    // log.Println(uuid)
    err = collection.FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&document)
    if err != nil {
        assert.Error(t, err)
    }
    // log.Println(document)

    assert.Empty(t, document)
}



// //count product
func TestCountProduct(t *testing.T){
    // Abrir conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)

    // Contar documentos da coleção
    collection := client.Database("market").Collection("products")
    count, err := collection.CountDocuments(context.Background(), bson.M{})
    if err != nil {
        t.Fatalf("Erro ao contar documentos da coleção: %v", err)
    }

    assert.NotEqual(t, count, 0)
    assert.NoError(t, err)
}

// func UpdateProduct(uuid string, product Product) error {
//     // Abrir conexão com o banco de dados
//     client, err := db.OpenConnection()
//     if err != nil {
//         return err
//     }
//     defer db.Close(client)

//     // Atualizar documento na coleção
//     collection := client.Database("market").Collection("products")
//     _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
//         "$set": bson.M{
//             "product_name":     product.Name,
//             "Price":    product.Price,
//             "quantity": product.Quantity,
//         },
//     })
//     if err != nil {
//         return err
//     }

//     return nil
// }


// // func BuyProduct(name string, quantity int) error {
// //     // Abrir conexão com o banco de dados
// //     client, err := db.OpenConnection()
// //     if err != nil {
// //         return err
// //     }
// //     defer db.Close(client)

// //     // Atualizar documento na coleção
// //     collection := client.Database("market").Collection("products")
// //     _, err = collection.UpdateOne(context.Background(), bson.M{"name": name}, bson.M{
// //         "$inc": bson.M{
// //             "quantity": -quantity,
// //         },
// //     })
// //     if err != nil {
// //         return err
// //     }

// //     return nil
// // }



// func DeleteProduct(uuid string) error {
//     // Abrir conexão com o banco de dados
//     client, err := db.OpenConnection()
//     if err != nil {
//         return err
//     }
//     defer db.Close(client)

//     // Deletar documento na coleção
//     collection := client.Database("market").Collection("products")
//     _, err = collection.DeleteOne(context.Background(), bson.M{"UUID": uuid})
//     if err != nil {
//         return err
//     }

//     return nil
// }


// func UpdateProductQuantity(uuid string, newQuantity int) error {
//     log.Println("UpdateProductQuantity")
//     log.Println(uuid)
//     log.Println(newQuantity)

//     // Verificar se a nova quantidade é negativa
//     if newQuantity < 0 {
//         return errors.New("quantity cannot be negative")
//     }

//     client, err := db.OpenConnection()
//     if err != nil {
//         return err
//     }
//     defer db.Close(client)

//     // Verificar se a quantidade atual é negativa
//     isAvailable, err := ValidateQuantity(uuid, 0)
//     if err != nil {
//         return err
//     }
//     if !isAvailable {
//         return errors.New("Current quantity is negative")
//     }

//     // Verificar se a nova quantidade é disponível
//     isAvailable, err = ValidateQuantity(uuid, newQuantity)
//     if err != nil {
//         return err
//     }

//     productname , err := GetProductNameWithUUID(uuid)
//     if !isAvailable {
//         return fmt.Errorf("Quantity is not available for product: %s", productname)
//     }

//     // Atualizar documento na coleção
//     collection := client.Database("market").Collection("products")
//     _, err = collection.UpdateOne(context.Background(), bson.M{"uuid": uuid}, bson.M{
//         "$set": bson.M{
//             "quantity": newQuantity,
//         },
//     })
//     if err != nil {
//         return err
//     }

//     log.Println("Product quantity updated")
//     return nil
// }




// //validate if quantity is available
func TestValidateQuantity(t *testing.T)  {
    // Abrir conexão com o banco de dados
    uuid := "1"
    quantity := 10
    client, err := db.OpenConnection()
    if err != nil {
        t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)
    // Buscar documento da coleção
    collection := client.Database("market").Collection("products")
    var document Product
    err = collection.FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&document)
    if err != nil {
        t.Fatalf("Erro ao buscar documento da coleção: %v", err)
    }
    // Verificar se a quantidade é suficiente
    if document.Quantity >= quantity {
        assert.True(t, true)
    } else {
        assert.False(t, false)
    }
    assert.NoError(t, err)
}


func TestGetProductNameWithUUID(t *testing.T) {
    
    uuid := "1"
    client, err := db.OpenConnection()
    if err != nil {
        t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)

    // Buscar documento da coleção
    collection := client.Database("market").Collection("products")
    var document Product
    err = collection.FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&document)
    if err != nil {
        t.Fatalf("Erro ao buscar documento da coleção: %v", err)
    }

    assert.NotEmpty(t, document.Name)
    assert.NoError(t, err)
    
}