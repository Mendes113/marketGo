package main

//connect do mongodb
import (
	"fmt"
	"marketgo/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Erro ao carregar o arquivo .env:", err)
        os.Exit(1)
    }
    app := fiber.New()
    routes.Setup(app)
  
}
