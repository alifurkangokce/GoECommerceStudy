package repository

import (
	"GoECommerceStudy/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockProductImageRepository.go -package=repository GoECommerceStudy/repository ProductRepository
type ProductImageRepositoryDB struct {
	ProductImageCollection *mongo.Collection
}

func (t ProductImageRepositoryDB) GetById(id primitive.ObjectID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t ProductImageRepositoryDB) GetAll() ([]models.ProductImage, error) {
	//TODO implement me
	panic("implement me")
}

func (t ProductImageRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t ProductImageRepositoryDB) Update(id primitive.ObjectID, product models.ProductImage) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type ProductImageRepository interface {
	Insert(productImage models.ProductImage) (bool, error)
	GetById(id primitive.ObjectID) (bool, error)
	GetAll() ([]models.ProductImage, error)
	Delete(id primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, product models.ProductImage) (bool, error)
}

func NewProductImageRepository(dbClient *mongo.Collection) ProductImageRepositoryDB {
	return ProductImageRepositoryDB{ProductImageCollection: dbClient}
}

func (t ProductImageRepositoryDB) Insert(productImage models.ProductImage) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	productImage.Id = primitive.NewObjectID()
	productImage.CreatedAt = time.Now()
	result, err := t.ProductImageCollection.InsertOne(ctx, productImage)
	if err != nil || result.InsertedID == nil {
		log.Fatalln(err)
		return false, err
	}
	return true, nil
}
