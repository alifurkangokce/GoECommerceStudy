package services

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/repository"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//go:generate mockgen -destination=../mocks/service/mockProductImageService.go -package=services GoECommerceStudy/services ProductService
type DefaultProductImageService struct {
	Repo repository.ProductImageRepository
}

func (pI DefaultProductImageService) ProductImagesGet() ([]models.ProductImage, error) {
	//TODO implement me
	panic("implement me")
}

func (pI DefaultProductImageService) GetImageById(id primitive.ObjectID) (models.ProductImage, error) {
	//TODO implement me
	panic("implement me")
}

func (pI DefaultProductImageService) ProductImageDelete(id primitive.ObjectID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (pI DefaultProductImageService) ProductImageUpdate(id primitive.ObjectID, productImage models.ProductImage) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type ProductImageService interface {
	ProductImageInsert(productImage models.ProductImage) (bool, error)
	ProductImagesGet() ([]models.ProductImage, error)
	GetImageById(id primitive.ObjectID) (models.ProductImage, error)
	ProductImageDelete(id primitive.ObjectID) (bool, error)
	ProductImageUpdate(id primitive.ObjectID, productImage models.ProductImage) (bool, error)
}

func NewProductImageService(Repo repository.ProductImageRepository) DefaultProductImageService {
	return DefaultProductImageService{Repo: Repo}
}
func (pI DefaultProductImageService) ProductImageInsert(productImage models.ProductImage) (bool, error) {
	v := validator.New()
	if err := v.Struct(productImage); err != nil {
		log.Fatalln(err)
		return false, err
	}
	result, err := pI.Repo.Insert(productImage)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	return result, nil
}
