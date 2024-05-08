package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("sua_chave_secreta")

type Claims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateToken(userID, username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Token expira em 24 horas

    claims := &Claims{
        UserID:   userID,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("Token inválido")
    }

    return claims, nil
}
