package main

//connect do mongodb
import (
	"fmt"
	"log"
	"marketgo/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
func main() {
	log.Print("Iniciando a aplicação")
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Erro ao carregar o arquivo .env:", err)
        os.Exit(1)
    }

    app := fiber.New()
    routes.Setup(app)
  
}
