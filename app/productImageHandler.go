package app

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if err != nil || result.IsZero() {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return ctx.Status(http.StatusCreated).JSON(true)
}

func (h ProductImageHandler) GetAllProductImages(ctx *fiber.Ctx) error {
	result, err := h.Service.ProductImagesGet()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
func (h ProductImageHandler) GetProductImageById(ctx *fiber.Ctx) error {
	Id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(Id)
	result, err := h.Service.GetImageById(_Id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
func (h ProductImageHandler) ProductImageUpdate(ctx *fiber.Ctx) error {
	var productImage models.ProductImage
	if err := ctx.BodyParser(&productImage); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	Id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(Id)
	result, err := h.Service.ProductImageUpdate(_Id, productImage)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
func (h ProductImageHandler) DeleteProductImage(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")
	cnvId, _ := primitive.ObjectIDFromHex(productId)
	result, err := h.Service.ProductImageDelete(cnvId)
	if err != nil || result == false {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"State": true})

}
