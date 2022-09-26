package app

import (
	services "GoECommerceStudy/mocks/service"
	"GoECommerceStudy/models"
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

func TestProductHandler(t *testing.T) {
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
