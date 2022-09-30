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
	dbClient := configs.GetCollection(configs.DB, "product")
	ProductRepository := repository.NewProductRepository(dbClient)
	p := app.ProductHandler{Service: services.NewProductService(ProductRepository)}
	routes.SetProductRoutes(appRoute, p)

	if err := appRoute.Listen(":8080"); err != nil {
		log.Fatalln("Error listening")
	}
}
