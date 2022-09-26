package services

import (
	"GoECommerceStudy/mocks/repository"
	"GoECommerceStudy/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var mockRepo *repository.MockProductRepository
var service ProductService
var ProductsGetFakeData = []models.Product{
	{Id: primitive.NewObjectID(), Name: gofakeit.Name()},
}
var ProductRequest = models.Product{
	Id:   primitive.NewObjectID(),
	Name: gofakeit.Name(),
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
	mockRepo.EXPECT().GetAll().Return(ProductsGetFakeData, nil)
	result, err := service.ProductsGet()
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}

func TestDefaultProductService_ProductInsert(t *testing.T) {
	td := setup(t)
	defer td()
	mockRepo.EXPECT().Insert(ProductRequest).Return(true, nil)
	result, err := service.ProductInsert(ProductRequest)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, result.Status, true)
}
func TestDefaultProductService_ProductDelete(t *testing.T) {
	td := setup(t)
	defer td()
	Id := ProductsGetFakeData[0].Id
	mockRepo.EXPECT().Delete(Id).Return(true, nil)
	result, err := service.ProductDelete(Id)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, result, true)
}

func TestDefaultProductService_ProductUpdate(t *testing.T) {
	td := setup(t)
	defer td()
	Id := ProductsGetFakeData[0].Id
	mockRepo.EXPECT().Update(Id, ProductsGetFakeData[0]).Return(true, nil)
	result, err := service.ProductUpdate(Id, ProductsGetFakeData[0])

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, result.Status, true)
}
