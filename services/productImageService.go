package services

import (
	"GoECommerceStudy/dto/productImageDto"
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
	ProductImageInsert(dto productImageDto.ProductImageInsertDto) (bool, error)
	ProductImageDelete(id primitive.ObjectID, imageId primitive.ObjectID) (bool, error)
	ProductImageUpdate(id primitive.ObjectID, dto productImageDto.ProductImageUpdateDto) (bool, error)
}

func NewProductImageService(Repo repository.ProductImageRepository) DefaultProductImageService {
	return DefaultProductImageService{Repo: Repo}
}
func (pI DefaultProductImageService) ProductImageInsert(dto productImageDto.ProductImageInsertDto) (bool, error) {
	v := validator.New()
	if err := v.Struct(dto); err != nil {
		log.Fatalln(err)
		return false, err
	}
	result, err := pI.Repo.Insert(dto)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	return result, nil
}

func (pI DefaultProductImageService) ProductImageDelete(id primitive.ObjectID, imageId primitive.ObjectID) (bool, error) {
	result, err := pI.Repo.Delete(id, imageId)
	if err != nil {
		log.Fatalln(err)
		return result, err
	}
	return result, nil

}
func (pI DefaultProductImageService) ProductImageUpdate(id primitive.ObjectID, dto productImageDto.ProductImageUpdateDto) (bool, error) {
	result, err := pI.Repo.Update(id, dto)
	if err != nil {
		log.Fatalln(err)
		return result, err
	}
	return result, nil
}
