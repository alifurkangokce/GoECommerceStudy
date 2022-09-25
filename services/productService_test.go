package services

import (
	"GoECommerceStudy/mocks/repository"
	"GoECommerceStudy/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var mockRepo *repository.MockProductRepository
var service ProductService
var FakeData = []models.Product{
	{Id: primitive.NewObjectID(), Name: "Test"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()
	mockRepo = repository.NewMockProductRepository(ct)
	service = NewProductService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultProductService_ProductsGet(t *testing.T) {
	td := setup(t)
	defer td()
	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.ProductsGet()
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)

}
