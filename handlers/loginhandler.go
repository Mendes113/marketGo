package handlers

import (
	"log"
	"marketgo/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// LoginHandler é um manipulador de rota para o processo de login
// @Summary Processo de login
// @Description Processo de login
// @ID login
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Credenciais de login"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(c *fiber.Ctx) error {
    // Parse do corpo da requisição para obter as credenciais do usuário
    var loginData models.LoginRequest
    if err := c.BodyParser(&loginData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Falha ao analisar os dados de login"})
    }

    // Verificar se as credenciais do usuário são válidas (por exemplo, consultando um banco de dados)
    if isValid := models.CheckCredentials(loginData.Username, loginData.Password); !isValid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciais inválidas"})
    }

    // Se as credenciais forem válidas, crie e assine um token JWT
    // Obtenha a chave secreta da variável de ambiente
    jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
    if len(jwtKey) == 0 {
        // Se a variável de ambiente não estiver definida, retorne um erro
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Chave secreta JWT não definida"})
    }
    //remover antes do commit
    log.Print(jwtKey)

    expirationTime := time.Now().Add(time.Hour * 24) // Token válido por 24 horas
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": loginData.Username,
        "exp": expirationTime.Unix(),
    })
    
    // Assinar o token com a chave secreta e obter a string do token
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao gerar o token"})
    }

    // Retorne o token JWT para o cliente
    return c.JSON(fiber.Map{"token": tokenString})
}

// Register é um manipulador de rota para o processo de registro
// @Summary Processo de registro
// @Description Processo de registro
// @ID register
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Credenciais de registro"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /register [post]
func Register(c *fiber.Ctx) error {
    // Parse do corpo da requisição para obter os dados do novo usuário
    var userData models.LoginRequest
    if err := c.BodyParser(&userData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Falha ao analisar os dados de registro"})
    }

	log.Print(userData)

   
	existingUser, err := models.GetUserByUsername(userData.Username)
    
    log.Print(existingUser)
	if existingUser != nil {
		// Se o usuário já existir, retorne um erro de solicitação inválida
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username Já Em Uso"})
	}else{
        models.CreateUser(userData)
    }

    // Se houver um erro ao verificar se o usuário já existe, registre o erro e retorne um erro interno do servidor
    if err == nil {
        log.Print("Error checking user existence:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao verificar a existência do usuário"})
    }
   

    log.Print(err)

    // Criptografar a senha do novo usuário
    userData.Password = models.CryptPassword(userData.Password)

   

    // Se o registro for bem-sucedido, crie e retorne um token JWT para o novo usuário
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = userData.Username
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token válido por 24 horas

    // Obtenha a chave secreta da variável de ambiente
    secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        // Se a variável de ambiente não estiver definida, retorne um erro
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Chave secreta JWT não definida"})
    }

    // Assinar o token com a chave secreta e obter a string do token
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao gerar o token"})
    }

     // Inserir o novo usuário no banco de dados
     if err := models.CreateUser(userData); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Falha ao registrar o usuário"})
    }

    // Retorne o token JWT para o novo usuário
    return c.JSON(fiber.Map{"token": tokenString})
}
