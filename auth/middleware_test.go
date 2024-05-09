package auth

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func TestAuthMiddlewareValidToken(t *testing.T) {
	// Defina um token JWT válido
	tokenString, _ := GenerateToken("user123", "john_doe")
	validToken := "Bearer " + tokenString

	// Crie uma instância do Fiber e defina o middleware de autenticação JWT
	app := fiber.New()
	app.Use(AuthMiddleware)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rota protegida")
	})

	// Faça uma solicitação HTTP com o token JWT válido
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", validToken)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação: %v", err)
	}
	defer resp.Body.Close()

	// Verifique se a resposta está correta (código de status 200)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperava-se status OK, mas obteve %d", resp.StatusCode)
	}

	// Verifique o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler corpo da resposta: %v", err)
	}
	expectedBody := "Rota protegida"
	if string(body) != expectedBody {
		t.Errorf("Corpo da resposta incorreto. Esperado: %s, Obtido: %s", expectedBody, string(body))
	}
}

func TestAuthMiddlewareInvalidToken(t *testing.T) {
	// Defina um token JWT inválido
	invalidToken := "token_inválido"

	// Crie uma instância do Fiber e defina o middleware de autenticação JWT
	app := fiber.New()
	app.Use(AuthMiddleware)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rota protegida")
	})

	// Faça uma solicitação HTTP com o token JWT inválido
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", invalidToken)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação: %v", err)
	}
	defer resp.Body.Close()

	// Verifique se a resposta está correta (código de status 401)
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Esperava-se status Unauthorized, mas obteve %d", resp.StatusCode)
	}
}

func TestAuthMiddlewareMissingToken(t *testing.T) {
	// Crie uma instância do Fiber e defina o middleware de autenticação JWT
	app := fiber.New()
	app.Use(AuthMiddleware)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rota protegida")
	})

	// Faça uma solicitação HTTP sem o token JWT
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer solicitação: %v", err)
	}
	defer resp.Body.Close()

	// Verifique se a resposta está correta (código de status 401)
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Esperava-se status Unauthorized, mas obteve %d", resp.StatusCode)
	}

	// Verifique o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erro ao ler corpo da resposta: %v", err)
	}
	expectedError := "Token de autorização não fornecido"
	if !strings.Contains(string(body), expectedError) {
		t.Errorf("Erro esperado não encontrado no corpo da resposta. Esperado: %s", expectedError)
	}
}

func TestAuthMiddlewareExpiredToken(t *testing.T) {
    jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

    // Gere um token JWT expirado
    expirationTime := time.Now().Add(-1 * time.Hour) // Token expirado há 1 hora
    claims := &Claims{
        UserID:   "user123",
        Username: "john_doe",
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := expiredToken.SignedString(jwtKey)
    expiredTokenString := "Bearer " + tokenString

    // Crie uma instância do Fiber e defina o middleware de autenticação JWT
    app := fiber.New()
    app.Use(AuthMiddleware)

    // Rota de teste
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Rota protegida")
    })

    // Faça uma solicitação HTTP com o token JWT expirado
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    req.Header.Set("Authorization", expiredTokenString)
    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Erro ao fazer solicitação: %v", err)
    }
    defer resp.Body.Close()

    // Verifique se a resposta está correta (código de status 401)
    if resp.StatusCode != http.StatusUnauthorized {
        t.Errorf("Esperava-se status Unauthorized, mas obteve %d", resp.StatusCode)
    }

    // Verifique o corpo da resposta
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Erro ao ler corpo da resposta: %v", err)
    }
    expectedError := "Token expirado"
    if !strings.Contains(string(body), expectedError) {
        t.Errorf("Erro esperado não encontrado no corpo da resposta. Esperado: %s, Obtido: %s", expectedError, string(body))
    }
}
