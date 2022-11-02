package services

import (
	"GoECommerceStudy/dto"
	"GoECommerceStudy/repository"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultProductVariantService struct {
	Repo repository.ProductVariantRepository
}

func NewProductVariantService(Repo repository.ProductVariantRepository) DefaultProductVariantService {
	return DefaultProductVariantService{Repo: Repo}
}

type ProductVariantService interface {
	ProductVariantInsert(variant dto.ProductVariantInsertDto) (bool, error)
	ProductVariantUpdate(id primitive.ObjectID, variant dto.ProductVariantUpdateRequestDto) (bool, error)
	ProductVariantDelete(id primitive.ObjectID, variantId primitive.ObjectID) (bool, error)
}

func (v DefaultProductVariantService) ProductVariantInsert(variant dto.ProductVariantInsertDto) (bool, error) {
	valid := validator.New()
	if err := valid.Struct(variant); err != nil {
		return false, err
	}
	_, err := v.Repo.Insert(variant)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (v DefaultProductVariantService) ProductVariantUpdate(id primitive.ObjectID, variant dto.ProductVariantUpdateRequestDto) (bool, error) {

	result, err := v.Repo.Update(id, variant)
	valid := validator.New()
	if err := valid.Struct(variant); err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	return result, nil
}
func (v DefaultProductVariantService) ProductVariantDelete(id primitive.ObjectID, variantId primitive.ObjectID) (bool, error) {
	result, err := v.Repo.Delete(id, variantId)
	if err != nil && result {
		return false, err
	}
	return result, nil
}
