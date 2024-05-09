package auth

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware de autenticação JWT para Fiber
func AuthMiddleware(c *fiber.Ctx) error {
    jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

    log.Print("AuthMiddleware")
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token de autorização não fornecido"})
    }

    // Remove "Bearer " prefix from the token string
    tokenString = strings.Replace(tokenString, "Bearer ", "", 1)


    // Parse JWT token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
        }
        if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
            return nil, fmt.Errorf("Método de assinatura inválido: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        log.Printf("Error parsing token: %v", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Erro ao analisar o token"})
    }

    // Check token validity
    if !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
    }

    // Check token expiration
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Erro ao obter reivindicações do token"})
    }

    expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
    if time.Now().Unix() > expirationTime.Unix() {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expirado"})
    }

    return c.Next()
}
