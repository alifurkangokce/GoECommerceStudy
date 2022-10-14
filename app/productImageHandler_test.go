package app

import (
	services "GoECommerceStudy/mocks/service"
	"GoECommerceStudy/models"
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

var pih ProductImageHandler
var mockProductImageService *services.MockProductImageService
var FakeDataForProductImageHandler = []models.ProductImage{
	{
		Id:        primitive.NewObjectID(),
		ProductId: primitive.NewObjectID(),
		Src:       gofakeit.ImageURL(640, 480),
	},
}

func setupProductImage(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockProductImageService = services.NewMockProductImageService(ctrl)
	pih = ProductImageHandler{
		Service: mockProductImageService,
	}
	return func() {
		defer ctrl.Finish()
	}
}
func TestGetAllProductImages_Handler(t *testing.T) {
	trd := setupProductImage(t)
	defer trd()

	router := fiber.New()
	router.Get("/api/productImage", pih.GetAllProductImages)
	mockProductImageService.EXPECT().ProductImagesGet().Return(FakeDataForProductImageHandler, nil)
	req := httptest.NewRequest("GET", "/api/productImage", nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)
}
func TestGetProductImages_Handler(t *testing.T) {
	trd := setupProductImage(t)
	defer trd()
	id, _ := primitive.ObjectIDFromHex("63302901533dcab4951b9b6b")
	router := fiber.New()
	router.Get("/api/productImage/:id", pih.GetProductImageById)
	mockProductImageService.EXPECT().GetImageById(id).Return(&FakeDataForProductImageHandler[0], nil)
	req := httptest.NewRequest("GET", "/api/productImage/63302901533dcab4951b9b6b", nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)
}
func TestCreateProductImage_Handler(t *testing.T) {
	trd := setupProductImage(t)
	defer trd()

	router := fiber.New()
	router.Post("/api/productImage", pih.CreateProductImage)
	mockProductImageService.EXPECT().ProductImageInsert(FakeDataForProductImageHandler[0]).Return(FakeDataForProductImageHandler[0].Id, nil)
	jsonBytes, _ := json.Marshal(FakeDataForProductImageHandler[0])
	req := httptest.NewRequest("POST", "/api/productImage", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := router.Test(req)
	assert.Equal(t, 201, resp.StatusCode)
}
func TestUpdateProductImage_Handler(t *testing.T) {
	trd := setupProductImage(t)
	defer trd()

	id, _ := primitive.ObjectIDFromHex("63302901533dcab4951b9b6b")
	FakeDataForProductImageHandler[0].Id = id

	router := fiber.New()
	router.Post("/api/productImage/:id", pih.ProductImageUpdate)
	mockProductImageService.EXPECT().ProductImageUpdate(id, FakeDataForProductImageHandler[0]).Return(true, nil)
	jsonBytes, _ := json.Marshal(FakeDataForProductImageHandler[0])
	req := httptest.NewRequest("POST", "/api/productImage/63302901533dcab4951b9b6b", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
func TestDeleteProductImage_Handler(t *testing.T) {
	trd := setupProductImage(t)
	defer trd()

	id, _ := primitive.ObjectIDFromHex("63302901533dcab4951b9b6b")

	router := fiber.New()
	router.Delete("/api/productImage/:id", pih.DeleteProductImage)
	mockProductImageService.EXPECT().ProductImageDelete(id).Return(true, nil)
	req := httptest.NewRequest("DELETE", "/api/productImage/63302901533dcab4951b9b6b", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
