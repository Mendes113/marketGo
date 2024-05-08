package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware de autenticação JWT para Fiber
func AuthMiddleware(c *fiber.Ctx) error {
    // Obtenha o token JWT do cabeçalho de autorização
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token de autorização não fornecido"})
    }

    // Valide o token JWT
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verifique se o método de assinatura é válido
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
        }
        // Retorne a chave de assinatura
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
    }

    // Se o token for válido, continue com a próxima manipulador
    return c.Next()
}