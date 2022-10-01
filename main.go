package main

import (
	"GoECommerceStudy/app"
	"GoECommerceStudy/configs"
	"GoECommerceStudy/repository"
	"GoECommerceStudy/routes"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	productDbClient := configs.GetCollection(configs.DB, "product")
	ProductRepository := repository.NewProductRepository(productDbClient)
	product := app.ProductHandler{Service: services.NewProductService(ProductRepository)}
	routes.SetProductRoutes(appRoute, product)

	productImageDbClient := configs.GetCollection(configs.DB, "productImage")
	ProductImageRepository := repository.NewProductImageRepository(productImageDbClient)
	productImage := app.ProductImageHandler{
		Service: services.NewProductImageService(ProductImageRepository),
	}
	routes.SetProductImageRoutes(appRoute, productImage)

	if err := appRoute.Listen(":8080"); err != nil {
		log.Fatalln("Error listening")
	}
}
