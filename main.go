
package main

//connect do mongodb
import (
	"marketgo/routes"

	"github.com/gofiber/fiber/v2"
)
func main() {
   
    app := fiber.New()
    routes.Setup(app)
  
}
