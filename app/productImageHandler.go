package app

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductImageHandler struct {
	Service services.ProductImageService
}

func (h ProductImageHandler) CreateProductImage(ctx *fiber.Ctx) error {
	var productImage models.ProductImage
	if err := ctx.BodyParser(&productImage); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.ProductImageInsert(productImage)

	if err != nil || result == false {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return ctx.Status(http.StatusCreated).JSON(true)
}
