package routes

import (
	"GoECommerceStudy/app"
	"github.com/gofiber/fiber/v2"
)

func SetProductVariantRoute(app *fiber.App, pv app.ProductVariantHandler) {
	api := app.Group("/api/productVariant")
	api.Post("", pv.CreateProductVariant)
	api.Post("/:id", pv.UpdateProductVariant)
	api.Delete("/:id", pv.DeleteProductVariant)
}
