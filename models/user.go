package models

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"marketgo/db"
	

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)



type UserAccount struct {
	UUID string `json:"uuid" validate:"required"`
	Username string `json:"username" validate:"required" min=3,max=50`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}



func CheckCredentials(username, password string) bool {
	log.Print("CheckCredentials")
    user, err := GetUserByUsername(username)
    if err != nil {
		log.Print(err)
        return false
    }
	log.Print(user)
    // Verificar se a senha fornecida corresponde à senha armazenada no banco de dados
    return checkPasswordHash(password, user.Password)
}


func GetUserByUsername(username string) (*UserAccount, error) {
	log.Print("GetUserByUsername")
	log.Print(username)
	client, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close(client)

	var user UserAccount
	collection := client.Database("market").Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"Username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("Usuário não encontrado")
	}
	log.Print(user)
	return &user, nil
}






func CreateUser(user LoginRequest) error {
	// Validar a estrutura do usuário
	log.Print("CreateUser")
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	// Abrir uma conexão com o banco de dados
	client, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close(client)

	// Criptografar a senha do usuário
	password := CryptPassword(user.Password)

	// Gerar um UUID para o usuário
	uuid := createUUID()
	log.Print(uuid)
	// Inserir o novo usuário no banco de dados
	collection := client.Database("market").Collection("users")
	_, err = collection.InsertOne(context.Background(), bson.M{
		"UUID":     uuid,
		"Username": user.Username,
		"Password": password,
		
	})
	if err != nil {
		return err
	}
	log.Print("User created")
	return nil
}






func LoginUser(login LoginRequest) (error, *UserAccount) {
    validate := validator.New()
    err := validate.Struct(login)
    if err != nil {
        return err, nil
    }

    client, err := db.OpenConnection()
    if err != nil {
        return err, nil
    }
    defer db.Close(client)

    var document UserAccount
    collection := client.Database("market").Collection("users")
    err = collection.FindOne(context.Background(), bson.M{"username": login.Username}).Decode(&document)
    if err != nil {
        return err, nil
    }

    // Verificar se a senha fornecida corresponde à senha armazenada
    if !checkPasswordHash(login.Password, document.Password) {
        return errors.New("Credenciais inválidas"), nil
    }

    // Se a senha corresponder, retornar o usuário
    return nil, &document
}

// Função para verificar se a senha fornecida corresponde à senha armazenada
func checkPasswordHash(password, hashFromDatabase string) bool {
    // Calcular o hash da senha fornecida
	log.Print("checkPasswordHash")
    hashedPassword := CryptPassword(password)

    // Comparar o hash da senha fornecida com o hash da senha armazenada no banco de dados
    return hashedPassword == hashFromDatabase
}






func CryptPassword(password string) string {
    // Criar um novo hash SHA-256
	log.Print("CryptPassword")
    hash := sha256.New()

    // Escrever a senha no hash
    hash.Write([]byte(password))

    // Obter o hash finalizado como bytes
    hashedBytes := hash.Sum(nil)

    // Converter os bytes do hash em uma string hexadecimal
    hashedString := hex.EncodeToString(hashedBytes)
	log.Print(hashedString)
    return hashedString
}





//    // Buscar documentos da coleção
//    var documents []bson.M
//    collection := client.Database("market").Collection("products")
//    cursor, err := collection.Find(context.Background(), bson.M{})
//    if err != nil {
// 	   return nil, err
//    }
//    defer cursor.Close(context.Background())

//    // Iterar sobre os documentos retornados pelo cursor
//    for cursor.Next(context.Background()) {
// 	   var document bson.M
// 	   if err := cursor.Decode(&document); err != nil {
// 		   return nil, err
// 	   }
// 	   documents = append(documents, document)
//    }

//    if err := cursor.Err(); err != nil {
// 	   return nil, err
//    }

//    // Serializar os documentos para JSON
//    jsonData, err := json.Marshal(documents)
//    if err != nil {
// 	   return nil, err
//    }

//    return jsonData, nil
// }