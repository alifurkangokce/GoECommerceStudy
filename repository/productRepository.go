package repository

import (
	"GoECommerceStudy/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockProductRepository.go -package=repository GoECommerceStudy/repository ProductRepository
type ProductRepositoryDB struct {
	ProductCollection *mongo.Collection
}
type ProductRepository interface {
	Insert(product models.Product) (bool, error)
	GetAll() ([]models.Product, error)
	Delete(id primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, product models.Product) (bool, error)
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
func (t ProductRepositoryDB) GetAll() ([]models.Product, error) {
	var product models.Product
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.ProductCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&product); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func (t ProductRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.ProductCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}
func (t ProductRepositoryDB) Update(id primitive.ObjectID, product models.Product) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := t.ProductCollection.UpdateByID(ctx, id, bson.M{
		"$set": bson.M{
			"name":       product.Name,
			"updated_at": time.Now(),
		},
	})
	if err != nil {
		return false, nil
	}

	return true, nil
}
