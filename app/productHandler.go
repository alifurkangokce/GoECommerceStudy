package app

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProductHandler struct {
	Service services.ProductService
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	fmt.Println("girdi1")
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.ProductInsert(product)

	if err != nil || result.Status == false {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}

	return ctx.Status(http.StatusCreated).JSON(true)
}
func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	fmt.Println("girdi2")
	result, err := h.Service.ProductsGet()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	fmt.Println("girdi3")
	productId := ctx.Params("id")
	cnvId, _ := primitive.ObjectIDFromHex(productId)
	result, err := h.Service.ProductDelete(cnvId)
	if err != nil || result == false {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"State": true})

}
func (h ProductHandler) ProductUpdate(ctx *fiber.Ctx) error {
	fmt.Println("girdi4")
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	Id := ctx.Params("id")
	_Id, _ := primitive.ObjectIDFromHex(Id)
	result, err := h.Service.ProductUpdate(_Id, product)
	if err != nil || result.Status == false {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}
