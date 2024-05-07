package handlers

import (
	"log"
	"marketgo/models"

	"github.com/gofiber/fiber/v2"
)



// CreateInvoice is a handler that inserts a new invoice
//@Tags invoices
//@Summary Create a new invoice
//@Description Create a new invoice
//@ID create-invoice
//@Accept json
//@Produce json
//@Param body body models.InvoiceRequest true "Invoice object that needs to be added"
//@Success 201 {object} models.InvoiceRequest
//@Router /invoices [post]
func CreateInvoice(c *fiber.Ctx) error {
	
	var invoice models.InvoiceRequest
	if err := c.BodyParser(&invoice); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := models.CreateInvoice(invoice)
	if err != nil {
		log.Println("Error inserting invoice:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error inserting invoice",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Invoice created successfully",
		"invoice":  invoice,
	})

}
