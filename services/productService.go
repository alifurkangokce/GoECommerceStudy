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
	Repo      repository.ProductRepository
	ImageRepo repository.ProductImageRepository
}
type ProductService interface {
	ProductInsert(product models.Product) (*dto.ProductDto, error)
	ProductsGet() ([]models.Product, error)
	ProductDelete(id primitive.ObjectID) (bool, error)
	ProductUpdate(id primitive.ObjectID, product models.Product) (*dto.ProductDto, error)
}

func NewProductService(Repo repository.ProductRepository, ImageRepo repository.ProductImageRepository) DefaultProductService {
	return DefaultProductService{Repo: Repo, ImageRepo: ImageRepo}
}

func (d DefaultProductService) ProductInsert(product models.Product) (*dto.ProductDto, error) {
	var res dto.ProductDto
	if len(product.Name) <= 3 {
		res.Status = false
		return &dto.ProductDto{Status: false}, nil
	}
	result, err := d.Repo.Insert(product)
	if err != nil {
		res.Status = false
		return nil, err
	}
	if product.Images != nil {
		var updateImages []models.ProductImage
		for _, image := range product.Images {
			dto := models.ProductImage{
				ProductId: result,
				Position:  image.Position,
				CreatedAt: time.Now(),
				Width:     image.Width,
				Height:    image.Height,
				Src:       image.Src,
			}
			imageInsertedId, _ := d.ImageRepo.Insert(dto)
			dto.Id = imageInsertedId
			updateImages = append(updateImages, dto)
		}
		d.Repo.Update(result, models.Product{Images: updateImages})
	}
	res = dto.ProductDto{Status: true}
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
