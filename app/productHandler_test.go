package app

import (
	"GoECommerceStudy/dto"
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

var td ProductHandler
var mockService *services.MockProductService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = services.NewMockProductService(ctrl)
	td = ProductHandler{
		Service: mockService,
	}
	return func() {
		defer ctrl.Finish()
	}
}

func TestGetAllProducts_Handler(t *testing.T) {
	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("api/products", td.GetAllProducts)
	var FakeDataForHandler = []models.Product{
		{Id: primitive.NewObjectID(), Name: "Test 1"},
	}
	mockService.EXPECT().ProductsGet().Return(FakeDataForHandler, nil)
	req := httptest.NewRequest("GET", "/api/products", nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)
}
func TestCreateProduct_Handler(t *testing.T) {
	trd := setup(t)
	defer trd()

	fake := models.Product{
		Name: gofakeit.Name(),
	}
	firstReturn := dto.ProductDto{
		Status: true,
	}

	router := fiber.New()
	router.Post("/api/products", td.CreateProduct)
	mockService.EXPECT().ProductInsert(fake).Return(&firstReturn, nil)
	jsonBytes, _ := json.Marshal(fake)
	req := httptest.NewRequest("POST", "/api/products", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := router.Test(req)
	assert.Equal(t, 201, resp.StatusCode)
}
func TestUpdateProduct_Handler(t *testing.T) {
	trd := setup(t)
	defer trd()

	id, _ := primitive.ObjectIDFromHex("63302901533dcab4951b9b6b")
	fake := models.Product{
		Id:   id,
		Name: gofakeit.Name(),
	}
	firstReturn := dto.ProductDto{
		Status: true,
	}

	router := fiber.New()
	router.Post("/api/products/:id", td.ProductUpdate)
	mockService.EXPECT().ProductUpdate(id, fake).Return(&firstReturn, nil)
	jsonBytes, _ := json.Marshal(fake)
	req := httptest.NewRequest("POST", "/api/products/63302901533dcab4951b9b6b", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
