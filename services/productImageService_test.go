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

var imageService ProductImageService
var mockImageRepo *repository.MockProductImageRepository
var ProductImagesFakeData = []models.ProductImage{
	{
		Id:        primitive.NewObjectID(),
		ProductId: primitive.NewObjectID(),
		Src:       gofakeit.ImageURL(640, 480),
	},
}

func productImageSetup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()
	mockImageRepo = repository.NewMockProductImageRepository(ct)
	imageService = NewProductImageService(mockImageRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultProductImageService_ProductImagesGet(t *testing.T) {
	td := productImageSetup(t)
	defer td()
	mockImageRepo.EXPECT().GetAll().Return(ProductImagesFakeData, nil)
	result, err := imageService.ProductImagesGet()
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
func TestDefaultProductImageService_ProductImageGet(t *testing.T) {
	td := productImageSetup(t)
	defer td()
	mockImageRepo.EXPECT().GetById(ProductImagesFakeData[0].Id).Return(&ProductImagesFakeData[0], nil)
	result, err := imageService.GetImageById(ProductImagesFakeData[0].Id)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
func TestDefaultProductImageService_ProductImageInsert(t *testing.T) {
	td := productImageSetup(t)
	defer td()
	mockImageRepo.EXPECT().Insert(ProductImagesFakeData[0]).Return(ProductImagesFakeData[0].Id, nil)
	result, err := imageService.ProductImageInsert(ProductImagesFakeData[0])
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
func TestDefaultProductImageService_ProductImageDelete(t *testing.T) {
	td := productImageSetup(t)
	defer td()
	mockImageRepo.EXPECT().Delete(ProductImagesFakeData[0].Id).Return(true, nil)
	result, err := imageService.ProductImageDelete(ProductImagesFakeData[0].Id)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
func TestDefaultProductImageService_ProductImageUpdate(t *testing.T) {
	td := productImageSetup(t)
	defer td()
	mockImageRepo.EXPECT().Update(ProductImagesFakeData[0].Id, ProductImagesFakeData[0]).Return(true, nil)
	result, err := imageService.ProductImageUpdate(ProductImagesFakeData[0].Id, ProductImagesFakeData[0])
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}
