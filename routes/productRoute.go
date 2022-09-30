package routes

import (
	"GoECommerceStudy/app"
	"github.com/gofiber/fiber/v2"
)

func SetProductRoutes(app *fiber.App, p app.ProductHandler) {
	api := app.Group("/api/products")
	api.Post("", p.CreateProduct)
	api.Post("/:id", p.ProductUpdate)
	api.Get("", p.GetAllProducts)
	api.Delete("/:id", p.DeleteProduct)
}
