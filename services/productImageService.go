package services

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/repository"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//go:generate mockgen -destination=../mocks/service/mockProductImageService.go -package=services GoECommerceStudy/services ProductImageService
type DefaultProductImageService struct {
	Repo repository.ProductImageRepository
}
type ProductImageService interface {
	ProductImageInsert(productImage models.ProductImage) (primitive.ObjectID, error)
	ProductImagesGet() ([]models.ProductImage, error)
	GetImageById(id primitive.ObjectID) (*models.ProductImage, error)
	ProductImageDelete(id primitive.ObjectID) (bool, error)
	ProductImageUpdate(id primitive.ObjectID, productImage models.ProductImage) (bool, error)
}

func NewProductImageService(Repo repository.ProductImageRepository) DefaultProductImageService {
	return DefaultProductImageService{Repo: Repo}
}
func (pI DefaultProductImageService) ProductImageInsert(productImage models.ProductImage) (primitive.ObjectID, error) {
	v := validator.New()
	if err := v.Struct(productImage); err != nil {
		log.Fatalln(err)
		return primitive.ObjectID{}, err
	}
	result, err := pI.Repo.Insert(productImage)
	if err != nil {
		log.Fatalln(err)
		return primitive.ObjectID{}, err
	}
	return result, nil
}
func (pI DefaultProductImageService) ProductImagesGet() ([]models.ProductImage, error) {
	result, err := pI.Repo.GetAll()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return result, nil
}
func (pI DefaultProductImageService) GetImageById(id primitive.ObjectID) (*models.ProductImage, error) {
	result, err := pI.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (pI DefaultProductImageService) ProductImageDelete(id primitive.ObjectID) (bool, error) {
	result, err := pI.Repo.Delete(id)
	if err != nil {
		log.Fatalln(err)
		return result, err
	}
	return result, nil

}
func (pI DefaultProductImageService) ProductImageUpdate(id primitive.ObjectID, productImage models.ProductImage) (bool, error) {
	result, err := pI.Repo.Update(id, productImage)
	if err != nil {
		log.Fatalln(err)
		return result, err
	}
	return result, nil
}
