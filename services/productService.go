package services

import (
	"GoECommerceStudy/dto"
	"GoECommerceStudy/models"
	"GoECommerceStudy/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//go:generate mockgen -destination=../mocks/service/mockProductService.go -package=services GoECommerceStudy/services ProductService
type DefaultProductService struct {
	Repo repository.ProductRepository
}
type ProductService interface {
	ProductInsert(product models.Product) (*dto.ProductDto, error)
	ProductsGet() ([]models.Product, error)
	ProductDelete(id primitive.ObjectID) (bool, error)
	ProductUpdate(id primitive.ObjectID, product models.Product) (*dto.ProductDto, error)
}

func NewProductService(Repo repository.ProductRepository) DefaultProductService {
	return DefaultProductService{Repo: Repo}
}

func (d DefaultProductService) ProductInsert(product models.Product) (*dto.ProductDto, error) {

	var productImageArr []models.ProductImage
	if product.Images != nil {
		for _, v := range product.Images {
			productImage := models.ProductImage{
				Id:        primitive.NewObjectID(),
				Position:  v.Position,
				CreatedAt: time.Now(),
				Width:     v.Width,
				Height:    v.Height,
				Src:       v.Src,
			}
			productImageArr = append(productImageArr, productImage)
		}
	}
	product.Images = productImageArr
	result, err := d.Repo.Insert(product)
	if err != nil {
		return nil, err
	}
	res := dto.ProductDto{Status: !result.IsZero()}
	return &res, nil
}
func (d DefaultProductService) ProductsGet() ([]models.Product, error) {
	result, err := d.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (d DefaultProductService) ProductDelete(id primitive.ObjectID) (bool, error) {
	result, err := d.Repo.Delete(id)
	if err != nil || result == false {
		return false, err
	}
	return result, nil

}
func (d DefaultProductService) ProductUpdate(id primitive.ObjectID, product models.Product) (*dto.ProductDto, error) {
	var res dto.ProductDto
	result, err := d.Repo.Update(id, product)
	if err != nil || result == false {
		res.Status = false
		return &res, err
	}
	res = dto.ProductDto{Status: result}
	return &res, nil
}
