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
	appRoute.Post("/api/products/:id", p.ProductUpdate)
	appRoute.Get("api/products", p.GetAllProducts)
	appRoute.Delete("api/products/:id", p.DeleteProduct)
	appRoute.Listen(":8080")
}
