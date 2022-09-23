package app

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductHandler struct {
	Service services.ProductService
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.ProductInsert(product)
	if err != nil || result.Status == false {
		return err
	}
	return ctx.Status(http.StatusCreated).JSON(true)
}
