package auth

import (
    "testing"
    "time"
)

func TestGenerateAndParseToken(t *testing.T) {
    // Defina as informações do usuário para o teste
    userID := "123"
    username := "john_doe"

    // Gere um token
    tokenString, err := GenerateToken(userID, username)
    if err != nil {
        t.Errorf("Erro ao gerar token: %v", err)
        return
    }

    // Parse o token
    claims, err := ParseToken(tokenString)
    if err != nil {
        t.Errorf("Erro ao fazer parse do token: %v", err)
        return
    }

    // Verifique se as reivindicações (claims) são válidas
    if claims.UserID != userID {
        t.Errorf("UserID não corresponde: esperado %s, obtido %s", userID, claims.UserID)
    }
    if claims.Username != username {
        t.Errorf("Username não corresponde: esperado %s, obtido %s", username, claims.Username)
    }

    // Verifique a expiração do token
    expirationTime := time.Unix(claims.StandardClaims.ExpiresAt, 0)
    expectedExpirationTime := time.Now().Add(24 * time.Hour)
    if expirationTime.After(expectedExpirationTime) {
        t.Errorf("Tempo de expiração do token é após o esperado")
    }
}

func TestInvalidToken(t *testing.T) {
    // Teste um token inválido
    invalidToken := "token_inválido"
    _, err := ParseToken(invalidToken)
    if err == nil {
        t.Error("Esperava-se um erro ao fazer parse de um token inválido")
    }
}
