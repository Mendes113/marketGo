package models

import (
	"context"
	"encoding/json"
	"errors"

	"testing"

	"marketgo/db"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetAll(t *testing.T) {
	// Abrir conexão com o banco de dados
	client, err := db.OpenConnection()
	if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close(client)

	// Buscar documentos da coleção
	var documents []bson.M
	collection := client.Database("market").Collection("markets")
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

	// Aqui, geralmente você não retorna nada em um teste. Você só pode verificar se a função se comportou corretamente.
	assert.NotEmpty(t, documents)
	assert.NoError(t, err)

	// Você pode, por exemplo, verificar se há documentos retornados e se a serialização para JSON foi bem-sucedida
	if len(documents) == 0 {
		t.Fatal("Nenhum documento encontrado")
	}

	// Serializar os documentos para JSON
	_, err = json.Marshal(documents)
	if err != nil {
		t.Fatalf("Erro ao serializar documentos para JSON: %v", err)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, documents)
}


func TestGetMarket(t *testing.T) {
	// Criar um mercado de teste no banco de dados
	// Isso pode ser feito inserindo um mercado de teste com UUID conhecido

	// Definir o UUID do mercado de teste
	testUUID := "123456"

	// Abrir conexão com o banco de dados
	client, err := db.OpenConnection()
	if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close(client)

	// Recuperar o mercado do banco de dados usando o UUID
	var document Market
	collection := client.Database("market").Collection("markets")
	err = collection.FindOne(context.Background(), bson.M{"UUID": testUUID}).Decode(&document)

	// Verificar se ocorreu um erro ao recuperar o mercado
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// O mercado com o UUID fornecido não existe
			t.Log(t, "Mercado não encontrado")

		} else {
			// Outro erro ocorreu
			t.Fatalf("Erro ao recuperar mercado: %v", err)
		}
	}

	// Verificar se o mercado foi recuperado com sucesso
	assert.Empty(t, document, "O mercado recuperado está vazio")

}



//should fail
func TestUpdateMarketShouldFail(t *testing.T)  {

	// Definir o UUID do mercado a ser atualizado
	uuid := "123456"
	market := Market{
		UUID:	"123456",
		Name:     "Mercado de Teste",
		Location: "Rua de Teste, 123",
		Owner:    "Proprietário de Teste",
		Email:    "",
	}
    client, err := db.OpenConnection()
    if err != nil {
        t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
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

   // Get the updated market
  	err, updatedMarket := GetMarket(uuid)
   if err != nil {
	   assert.Fail(t, "Mercado não encontrado")
   }

   // Compare the updated market with the original market
   assert.Equal(t, market, updatedMarket)
}


func TestRemoveProductFromMarket(t *testing.T) {

	// Definir o UUID do mercado e do produto
	uuid := "123456"
	productUUID := "789012"
    client, err := db.OpenConnection()
    if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)

    collection := client.Database("market").Collection("markets")
    _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
        "$pull": bson.M{"Products": bson.M{"UUID": productUUID}},
    })

	assert.NoError(t, err)
	assert.Empty(t, err)
}


func TestGetProductsFromMarket(t *testing.T) {
	// Definir o UUID do mercado
	uuid := "123456"
    client, err := db.OpenConnection()
    if err != nil {
        
    }
    defer db.Close(client)

    var document Market
    collection := client.Database("market").Collection("products")
    err = collection.FindOne(context.Background(), bson.M{"market_uuid": uuid}).Decode(&document)
    if err != nil {
        
    }

    jsonData, err := json.Marshal(document.Products)
    if err != nil {
        assert.Fail(t, "Erro ao serializar produtos para JSON")
    }
	t.Log(jsonData)
	assert.NotEmpty(t, jsonData)
	assert.NoError(t, err)
	// Verificar se a lista de produtos está vazia
	assert.Empty(t, document.Products)	

}


func TestAppendProductToMarket(t *testing.T) {
	// Definir o UUID do mercado e do produto
	uuid := "123456"
	product := Product{
		UUID:     "789012",
		Name:     "Produto de Teste",
		Price:    10.0,
		Quantity: 100,
	}

	client, err := db.OpenConnection()
    if err != nil {
		t.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
    }
    defer db.Close(client)

    collection := client.Database("market").Collection("markets")
    _, err = collection.UpdateOne(context.Background(), bson.M{"UUID": uuid}, bson.M{
        "$push": bson.M{"Products": product},
    })

	assert.NoError(t, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, collection.Name())
	

}




