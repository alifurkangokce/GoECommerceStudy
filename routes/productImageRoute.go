package routes

import (
	"GoECommerceStudy/app"
	"github.com/gofiber/fiber/v2"
)

func SetProductImageRoutes(app *fiber.App, p app.ProductImageHandler) {
	api := app.Group("/api/productImage")
	api.Post("", p.CreateProductImage)
}
