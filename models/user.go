package models

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"marketgo/db"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)



type UserAccount struct {
	UUID string `json:"uuid" validate:"required"`
	Username string `json:"username" validate:"required" min=3,max=50`
	Password string `json:"password" validate:"required"`
	LastLogin time.Time `json:"last_login" db:"last_login"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email string `json:"email"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
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

func GetUserByEmail(email string) (*UserAccount, error) {
	client, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close(client)

	var user UserAccount
	collection := client.Database("market").Collection("users")
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("usuário nao encontrado")
	}

	return &user, nil
}


func UpdatePassword(username, password string) error {
	client, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close(client)

	collection := client.Database("market").Collection("users")
	_, err = collection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"password": password}})
	if err != nil {
		return err
	}

	return nil
}


func (user *UserAccount) Save() error {
    // Abrir uma conexão com o banco de dados
    client, err := db.OpenConnection()
    if err != nil {
        return err
    }
    defer db.Close(client)

    // Atualizar o documento do usuário no banco de dados
    collection := client.Database("market").Collection("users")
    _, err = collection.UpdateOne(
        context.Background(),
        bson.M{"uuid": user.UUID},
        bson.M{"$set": bson.M{
            "username":     user.Username,
            "password":     user.Password,
            "last_login":   user.LastLogin,
        }},
    )
    if err != nil {
        return err
    }

    log.Print("User updated successfully")
    return nil
}