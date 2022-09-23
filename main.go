package main

import (
	"GoECommerceStudy/app"
	"GoECommerceStudy/configs"
	"GoECommerceStudy/repository"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "product")
	ProductRepository := repository.NewProductRepository(dbClient)
	p := app.ProductHandler{Service: services.NewProductService(ProductRepository)}

	appRoute.Post("/api/products", p.CreateProduct)
	appRoute.Listen(":8080")
}
