package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"marketgo/models"
	

	"github.com/gofiber/fiber/v2"
)


var ERROR_PARSE_REQUEST_BODY = "Error parsing request body:"


// GetProducts is a handler that returns all products
//@Tags products
//@Summary Get list of products
//@Description Get list of products
//@ID get-products
//@Produce json
//@Success 200 {array} models.Product
//@Router /products [get]
func GetProducts(c *fiber.Ctx) error {
	  productsData, err := models.GetAllProdutcts()
  if err != nil {
	  log.Println(ERROR_PARSE_REQUEST_BODY, err)
	  return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  // Decodificar os dados do produto em um formato JSON
  var products []models.Product 
  if err := json.Unmarshal(productsData, &products); err != nil {
	  log.Println("Error decoding JSON:", err)
	  return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  // Se não houve erro, retorne os produtos em formato JSON
  return c.JSON(products)

}

// InsertProduct is a handler that inserts a new product
//@Tags products
//@Summary Create a new product
//@Description Create a new product
//@ID create-product
//@Accept json
//@Produce json
//@Param body body models.Product true "Product object that needs to be added"
//@Success 201 {object} models.Product
func InsertProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		log.Println(ERROR_PARSE_REQUEST_BODY, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	err := models.InsertProduct(product)
	if err != nil {
		log.Println("Error inserting product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error inserting product: %s.You can try updating the product if needed", err),
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product":  product,
	})
}


// GetProduct is a handler that returns a product by ID
//@Tags products
//@Summary Get a product by ID
//@Description Get a product by ID
//@ID get-product
//@Produce json
//@Param id path int true "Product ID"
//@Success 200 {object} models.Product
//@Router /products/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	  id := c.Params("id")
  err, product := models.GetProduct(id)
  if err != nil {
	log.Println("Error getting product:", err)
	return c.SendStatus(fiber.StatusInternalServerError)
  }
  
  return c.JSON(product)
}


// DeleteProduct is a handler that deletes a product by ID
//@Tags products
//@Summary Delete a product by ID
//@Description Delete a product by ID
//@ID delete-product
//@Param id path int true "Product ID"
//@Success 204
//@Router /products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	err := models.DeleteProduct(id)
	if err != nil {
		log.Println("Error deleting product", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}


// UpdateProduct is a handler that updates a product by ID
//@Tags products
//@Summary Update a product by ID
//@Description Update a product by ID
//@ID update-product
//@Accept json
//@Produce json
//@Param id path int true "Product ID"
//@Param body body models.Product true "Product object that needs to be updated"	
//@Success 200
//@Router /products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		log.Println(ERROR_PARSE_REQUEST_BODY, err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	err := models.UpdateProduct(id, product)
	if err != nil {
		log.Println("Error updating product:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	
	return c.SendStatus(fiber.StatusOK)
}


// func BuyProduct(c *fiber.Ctx) error {
// 	name := c.Params("id")
// 	var product models.Product
// 	if err := c.BodyParser(&product); err != nil {
// 		log.Println(ERROR_PARSE_REQUEST_BODY, err)
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	err := models.BuyProduct(name, product.Quantity)
// 	if err != nil {
// 		log.Println("Error buying product:", err)
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.SendStatus(fiber.StatusOK)
// }


