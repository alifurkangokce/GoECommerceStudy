package services

import (
	"GoECommerceStudy/models"
	"GoECommerceStudy/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultProductVariantService struct {
	Repo repository.ProductVariantRepository
}

func NewProductVariantService(Repo repository.ProductVariantRepository) DefaultProductVariantService {
	return DefaultProductVariantService{Repo: Repo}
}

type ProductVariantService interface {
	ProductVariantInsert(variant models.ProductVariant) (primitive.ObjectID, error)
	ProductVariantUpdate(id primitive.ObjectID, variant models.ProductVariant) (bool, error)
	ProductVariantDelete(id primitive.ObjectID) (bool, error)
}

func (v DefaultProductVariantService) ProductVariantInsert(variant models.ProductVariant) (primitive.ObjectID, error) {
	InsertedId, err := v.Repo.Insert(variant)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return InsertedId, nil
}
func (v DefaultProductVariantService) ProductVariantUpdate(id primitive.ObjectID, variant models.ProductVariant) (bool, error) {
	result, err := v.Repo.Update(id, variant)
	if err != nil {
		return false, err
	}
	return result, nil
}
func (v DefaultProductVariantService) ProductVariantDelete(id primitive.ObjectID) (bool, error) {
	result, err := v.Repo.Delete(id)
	if err != nil {
		return false, err
	}
	return result, nil
}
