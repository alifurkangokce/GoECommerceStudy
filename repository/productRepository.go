package repository

import (
	"GoECommerceStudy/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProductRepositoryDB struct {
	ProductCollection *mongo.Collection
}
type ProductRepository interface {
	Insert(product models.Product) (bool, error)
}

func NewProductRepository(dbClient *mongo.Collection) ProductRepositoryDB {
	return ProductRepositoryDB{ProductCollection: dbClient}
}

func (t ProductRepositoryDB) Insert(product models.Product) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product.Id = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	result, err := t.ProductCollection.InsertOne(ctx, product)

	if result.InsertedID == nil || err != nil {
		errors.New("product add failed")
		return false, err
	}
	return true, nil
}
