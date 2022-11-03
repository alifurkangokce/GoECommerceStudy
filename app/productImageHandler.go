package app

import (
	"GoECommerceStudy/dto/productImageDto"
	"GoECommerceStudy/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProductImageHandler struct {
	Service services.ProductImageService
}

func (h ProductImageHandler) CreateProductImage(ctx *fiber.Ctx) error {
	var productImage productImageDto.ProductImageInsertDto
	if err := ctx.BodyParser(&productImage); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.ProductImageInsert(productImage)

	if err != nil || !result {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return ctx.Status(http.StatusCreated).JSON(true)
}
func (h ProductImageHandler) ProductImageUpdate(ctx *fiber.Ctx) error {
	var productImage productImageDto.ProductImageUpdateDto
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
	var productImage productImageDto.ProductImageDeleteRequestDto
	if err := ctx.BodyParser(&productImage); err != nil {
		return err
	}

	id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(id)
	_imageId, _ := primitive.ObjectIDFromHex(productImage.ImageId)

	result, err := h.Service.ProductImageDelete(_Id, _imageId)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"Status": false})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"Status": result})

}
