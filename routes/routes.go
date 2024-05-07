//	@title			MarketGo API
//	@version		1.0
//	@description	API para gerenciar mercados e produtos
//	@host			localhost:3000
//	@BasePath		/
//	@schemes		http
//	@produce		json
//	@consumes		json
package routes

import (
	"marketgo/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_"marketgo/docs"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

    /**
    * @Summary Get list of markets
    * @Description Get list of markets
    * @ID get-markets
    * @Produce json
    * @Success 200 {array} Market
    * @Router /markets [get]
    */
    app.Get("/markets", handlers.GetMarkets)


	/**
	* @Summary Create a new market
	* @Description Create a new market
	* @ID create-market
	* @Accept json
	* @Produce json
	* @Param body body Market true "Market object that needs to be added"
	* @Success 201 {object} Market
	* @Router /markets [post]
	*/
	app.Post("/markets", handlers.InsertMarket)

	/**
	* @Summary Get a market by ID
	* @Description Get a market by ID
	* @ID get-market
	* @Produce json
	* @Param id path int true "Market ID"
	* @Success 200 {object} Market
	* @Router /markets/{id} [get]
	*/
	app.Get("/markets/:id", handlers.GetMarket)

	/**
	* @Summary Update a market by ID
	* @Description Update a market by ID
	* @ID update-market
	* @Accept json
	* @Produce json
	* @Param id path int true "Market ID"
	* @Param body body Market true "Market object that needs to be updated"
	* @Success 200 {object} Market
	* @Router /markets/{id} [put]
	*/
	app.Put("/markets/:id", handlers.UpdateMarket)

	/**
	* @Summary Delete a market by ID
	* @Description Delete a market by ID
	* @ID delete-market
	* @Param id path int true "Market ID"
	* @Success 204
	* @Router /markets/{id} [delete]
	*/
	app.Delete("/markets/:id/products/:product_id", handlers.DeleteProduct)
	
	// app.Delete("/markets/:id", handlers.DeleteMarket)

	/**
	* @Summary Get list of products
	* @Description Get list of products
	* @ID get-products
	* @Produce json
	* @Success 200 {array} Product
	* @Router /products [get]
	*/
	app.Get("/products", handlers.GetProducts)

	/**
	* @Summary Create a new product
	* @Description Create a new product
	* @ID create-product
	* @Accept json
	* @Produce json
	* @Param body body Product true "Product object that needs to be added"
	* @Success 201 {object} Product
	* @Router /products [post]
	*/
	app.Post("/products", handlers.InsertProduct)
	const productRoute = "/products/:id"
	
	
	app.Get(productRoute, handlers.GetProduct)
	//product bought
	app.Put(productRoute, handlers.UpdateProduct)
	app.Delete(productRoute, handlers.DeleteProduct)

	

	app.Get(productRoute, handlers.GetProduct)
	//product bought
	app.Put("/products/:id", handlers.UpdateProduct)
	app.Post("/products/invoice", handlers.CreateInvoice)
	app.Delete("/products/:id", handlers.DeleteProduct)

	app.Listen(":3000")
}
