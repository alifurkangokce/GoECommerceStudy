package app

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProductVariantHandler struct {
	Service services.ProductVariantService
}

func (s ProductVariantHandler) CreateProductVariant(ctx *fiber.Ctx) error {
	var productVariant models.ProductVariant
	if err := ctx.BodyParser(&productVariant); err != nil {
		return err
	}
	result, err := s.Service.ProductVariantInsert(productVariant)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": false})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"InsertedId": result})
}
func (s ProductVariantHandler) UpdateProductVariant(ctx *fiber.Ctx) error {
	var productVariant models.ProductVariant
	if err := ctx.BodyParser(&productVariant); err != nil {
		return err
	}
	id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(id)
	result, err := s.Service.ProductVariantUpdate(_Id, productVariant)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": false})
	}
	return ctx.Status(http.StatusNoContent).JSON(fiber.Map{"Status": result})
}
func (s ProductVariantHandler) DeleteProductVariant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(id)
	result, err := s.Service.ProductVariantDelete(_Id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"Status": false})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"Status": result})
}
