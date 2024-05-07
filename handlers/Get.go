// marketgo/handlers/Get.go
package handlers

import (
	"encoding/json"
	"log"
	"marketgo/models"

	"github.com/gofiber/fiber/v2"
)




// GetMarkets is a handler that returns all markets
//@Tags markets
//@Summary Get list of markets
//@Description Get list of markets
//@ID get-markets
//@Produce json
//@Success 200 {array} models.Market
//@Router /markets [get]
func GetMarkets(c *fiber.Ctx) error {
  marketsData, err := models.GetAll()
  if err != nil {
      log.Println("Error getting markets:", err)
      return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  // Decodificar os dados do mercado em um formato JSON
  var markets []models.Market 
  if err := json.Unmarshal(marketsData, &markets); err != nil {
      log.Println("Error decoding JSON:", err)
      return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  // Se não houve erro, retorne os mercados em formato JSON
  return c.JSON(markets)
}


// InsertMarket is a handler that inserts a new market
//@Tags markets
//@Summary Create a new market
//@Description Create a new market
//@ID create-market
//@Accept json
//@Produce json 
//@Param body body models.Market true "Market object that needs to be added"
//@Success 201 {object} models.Market
//@Router /markets [post]
func InsertMarket(c *fiber.Ctx) error {
	var market models.Market
	if err := c.BodyParser(&market); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	err := models.InsertMarket(market)
	if err != nil {
		log.Println("Error inserting market:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Market created successfully",
		"market":  market,
	})
}



// GetMarket is a handler that returns a market by ID
//@Tags markets
//@Summary Get a market by ID
//@Description Get a market by ID
//@ID get-market
//@Produce json
//@Param id path int true "Market ID"
//@Success 200 {object} models.Market
//@Router /markets/{id} [get]
func GetMarket(c *fiber.Ctx) error {
  id := c.Params("id")
  err, market := models.GetMarket(id)
  if err != nil {
    log.Println("Error getting market:", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  return c.JSON(market)
}

// DeleteMarket is a handler that deletes a market by ID
//@Tags markets
//@Summary Delete a market by ID
//@Description Delete a market by ID
//@ID delete-market
//@Param id path int true "Market ID"
//@Success 204
//@Router /markets/{id} [delete]
func UpdateMarket(c *fiber.Ctx) error {
  id := c.Params("id")
  var market models.Market
  if err := c.BodyParser(&market); err != nil {
    log.Println("Error parsing request body:", err)
    return c.SendStatus(fiber.StatusBadRequest)
  }
  
  err := models.UpdateMarket(id, market)
  if err != nil {
    log.Println("Error updating market:", err)
    return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  return c.SendStatus(fiber.StatusOK)
}




