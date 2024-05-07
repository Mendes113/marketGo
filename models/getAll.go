package models

import (
	"context"
	"marketgo/db"
    "encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Message string `json:"message"`
}





// GetAll retorna todos os documentos da coleção como JSON
func GetAll() ([]byte, error) {
    // Abrir conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        return nil, err
    }
    defer db.Close(client)

    // Buscar documentos da coleção
    var documents []bson.M
    collection := client.Database("market").Collection("markets")
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    // Iterar sobre os documentos retornados pelo cursor
    for cursor.Next(context.Background()) {
        var document bson.M
        if err := cursor.Decode(&document); err != nil {
            return nil, err
        }
        documents = append(documents, document)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    // Serializar os documentos para JSON
    jsonData, err := json.Marshal(documents)
    if err != nil {
        return nil, err
    }

    return jsonData, nil
}




func InsertMarket(market Market) error {
	client, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close(client)

	collection := client.Database("market").Collection("markets")
	_, err = collection.InsertOne(context.Background(), bson.M{
		"UUID":     market.UUID,
		"Name":     market.Name,
		"Location": market.Location,
		"Owner":    market.Owner,
		"Email":    market.Email,
	})

	return err
}

func GetMarket(uuid string) (error, *Market) {
    client, err := db.OpenConnection()
    if err != nil {
        return err, nil
    }
    defer db.Close(client)

    var document Market
    collection := client.Database("market").Collection("markets")
    err = collection.FindOne(context.Background(), bson.M{"UUID": uuid}).Decode(&document)
    if err != nil {
        return err, nil
    }

    return nil, &document
}

func UpdateMarket(uuid string, market Market) error {
    client, err := db.OpenConnection()
    if err != nil {
        return err
    }
    defer db.Close(client)

    collection := client.Database("market").Collection("markets")
    _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
        "$set": bson.M{
            "Name":     market.Name,
            "Location": market.Location,
            "Owner":    market.Owner,
            "Email":    market.Email,
        },
    })

    return err
}




func RemoveProductFromMarket(uuid string, productUUID string) error {
    client, err := db.OpenConnection()
    if err != nil {
        return err
    }
    defer db.Close(client)

    collection := client.Database("market").Collection("markets")
    _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
        "$pull": bson.M{"Products": bson.M{"UUID": productUUID}},
    })

    return err
}



//get Products from a market
//uses the market uuid to get the products

func GetProductsFromMarket(uuid string) ([]byte, error) {
    client, err := db.OpenConnection()
    if err != nil {
        return nil, err
    }
    defer db.Close(client)

    var document Market
    collection := client.Database("market").Collection("products")
    err = collection.FindOne(context.Background(), bson.M{"market_uuid": uuid}).Decode(&document)
    if err != nil {
        return nil, err
    }

    jsonData, err := json.Marshal(document.Products)
    if err != nil {
        return nil, err
    }

    

    return jsonData, nil
}

//append product to market

func AppendProductToMarket(uuid string, product Product) error {
    client, err := db.OpenConnection()
    if err != nil {
        return err
    }
    defer db.Close(client)

    collection := client.Database("market").Collection("markets")
    _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
        "$push": bson.M{"Products": product},
    })

    return err
}























// UUID string `json:"uuid"`
// 	Name string `json:"name"`
// 	Location string `json:"location"`
// 	Owner string `json:"owner"`
// 	Email string `json:"email"
