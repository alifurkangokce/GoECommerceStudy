package services

import (
	"GoECommerceStudy/dto"
	"GoECommerceStudy/models"
	"GoECommerceStudy/repository"
)

type DefaultProductService struct {
	Repo repository.ProductRepository
}
type ProductService interface {
	ProductInsert(product models.Product) (*dto.ProductDto, error)
}

func NewProductService(Repo repository.ProductRepository) DefaultProductService {
	return DefaultProductService{Repo: Repo}
}

func (d DefaultProductService) ProductInsert(product models.Product) (*dto.ProductDto, error) {
	var res dto.ProductDto
	if len(product.Name) <= 3 {
		res.Status = false
		return &res, nil
	}
	result, err := d.Repo.Insert(product)
	if err != nil || result == false {
		res.Status = false
		return nil, err
	}
	res = dto.ProductDto{Status: result}
	return &res, nil
}
