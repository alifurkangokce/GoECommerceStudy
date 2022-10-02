package routes

import (
	"GoECommerceStudy/app"
	"github.com/gofiber/fiber/v2"
)

func SetProductImageRoutes(app *fiber.App, p app.ProductImageHandler) {
	api := app.Group("/api/productImage")
	api.Post("", p.CreateProductImage)
	api.Post("/:id", p.ProductImageUpdate)
	api.Get("/:id", p.GetProductImageById)
	api.Delete("/:id", p.DeleteProductImage)
	api.Get("", p.GetAllProductImages)
}
