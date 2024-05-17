package routes

import (
	"marketgo/auth"
	"marketgo/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_"marketgo/docs"
)

func Setup(app *fiber.App) {

	// Rota para login
	app.Get("/login", handlers.LoginHandler)

	// Rota para registro de usuário
	app.Post("/register", handlers.Register)


	// Rotas protegidas com autenticação JWT
	protectedRoutes := app.Group("/api/", auth.AuthMiddleware)

	// Rota para obter uma lista de mercados
	/**
	* @Summary Get list of markets
	* @Description Get list of markets
	* @ID get-markets
	* @Produce json
	* @Success 200 {array} Market
	* @Router /api/markets [get]
	*/
	protectedRoutes.Get("/markets", handlers.GetMarkets)

	// Rota para criar um novo mercado
	/**
	* @Summary Create a new market
	* @Description Create a new market
	* @ID create-market
	* @Accept json
	* @Produce json
	* @Param body body Market true "Market object that needs to be added"
	* @Success 201 {object} Market
	* @Router /api/markets [post]
	*/
	protectedRoutes.Post("/markets", handlers.InsertMarket)

	

	// Rota para atualizar um mercado por ID
	/**
	* @Summary Update a market by ID
	* @Description Update a market by ID
	* @ID update-market
	* @Accept json
	* @Produce json
	* @Param id path int true "Market ID"
	* @Param body body Market true "Market object that needs to be updated"
	* @Success 200 {object} Market
	* @Router /api/markets/{id} [put]
	*/
	protectedRoutes.Put("/markets/:id", handlers.UpdateMarket)

	// Rota para excluir um mercado por ID
	/**
	* @Summary Delete a market by ID
	* @Description Delete a market by ID
	* @ID delete-market
	* @Param id path int true "Market ID"
	* @Success 204
	* @Router /api/markets/{id} [delete]
	*/
	protectedRoutes.Delete("/markets/:id", handlers.DeleteMarket)

	// Rota para obter uma lista de produtos
	/**
	* @Summary Get list of products
	* @Description Get list of products
	* @ID get-products
	* @Produce json
	* @Success 200 {array} Product
	* @Router /api/products [get]
	*/
	protectedRoutes.Get("/products", handlers.GetProducts)

	// Rota para criar um novo produto
	/**
	* @Summary Create a new product
	* @Description Create a new product
	* @ID create-product
	* @Accept json
	* @Produce json
	* @Param body body Product true "Product object that needs to be added"
	* @Success 201 {object} Product
	* @Router /api/products [post]
	*/
	protectedRoutes.Post("/products", handlers.InsertProduct)

	// Rota para obter um produto por ID
	/**
	* @Summary Get a product by ID
	* @Description Get a product by ID
	* @ID get-product
	* @Produce json
	* @Param id path int true "Product ID"
	* @Success 200 {object} Product
	* @Router /api/products/{id} [get]
	*/
	protectedRoutes.Get("/products/:id", handlers.GetProduct)

	// Rota para atualizar um produto por ID
	/**
	* @Summary Update a product by ID
	* @Description Update a product by ID
	* @ID update-product
	* @Accept json
	* @Produce json
	* @Param id path int true "Product ID"
	* @Param body body Product true "Product object that needs to be updated"
	* @Success 200 {object} Product
	* @Router /api/products/{id} [put]
	*/
	protectedRoutes.Put("/products/:id", handlers.UpdateProduct)

	// Rota para excluir um produto por ID
	/**
	* @Summary Delete a product by ID
	* @Description Delete a product by ID
	* @ID delete-product
	* @Param id path int true "Product ID"
	* @Success 204
	* @Router /api/products/{id} [delete]
	*/
	protectedRoutes.Delete("/products/:id", handlers.DeleteProduct)

	// Rota para criar uma fatura de produto
	/**
	* @Summary Create product invoice
	* @Description Create product invoice
	* @ID create-invoice
	* @Accept json
	* @Produce json
	* @Success 200 {object} Invoice
	* @Router /api/products/invoice [post]
	*/
	protectedRoutes.Post("/products/invoice", handlers.CreateInvoice)

	// Rota padrão para a documentação Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Configuração personalizada para documentação Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:            "http://example.com/doc.json",
		DeepLinking:    false,
		DocExpansion:   "none",
		OAuth: &swagger.OAuthConfig{
			AppName:    "OAuth Provider",
			ClientId:   "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	// Inicia o servidor na porta 3000
	app.Listen(":3002")
}
