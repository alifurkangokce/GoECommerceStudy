package repository

import (
	"GoECommerceStudy/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

//go:generate mockgen -destination=../mocks/repository/mockProductImageRepository.go -package=repository GoECommerceStudy/repository ProductImageRepository
type ProductImageRepositoryDB struct {
	ProductImageCollection *mongo.Collection
}

func (t ProductImageRepositoryDB) Insert(productImage models.ProductImage) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	productImage.Id = primitive.NewObjectID()
	productImage.CreatedAt = time.Now()
	result, err := t.ProductImageCollection.InsertOne(ctx, productImage)
	if err != nil || result.InsertedID == nil {
		log.Fatalln(err)
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (t ProductImageRepositoryDB) GetById(id primitive.ObjectID) (*models.ProductImage, error) {
	var productImage models.ProductImage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := t.ProductImageCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&productImage)
	if err != nil {
		return nil, err
	}
	return &productImage, nil

}

func (t ProductImageRepositoryDB) GetAll() ([]models.ProductImage, error) {
	var allProductImages []models.ProductImage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.ProductImageCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = result.All(ctx, &allProductImages); err != nil {
		return nil, err
	}
	return allProductImages, nil

}

func (t ProductImageRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.ProductImageCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func (t ProductImageRepositoryDB) Update(id primitive.ObjectID, product models.ProductImage) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := t.ProductImageCollection.UpdateByID(ctx, id, bson.M{
		"$set": bson.M{
			"src":        product.Src,
			"updated_at": time.Now(),
			"position":   product.Position,
		},
	})
	if err != nil {
		return false, nil
	}

	return true, nil
}

type ProductImageRepository interface {
	Insert(productImage models.ProductImage) (primitive.ObjectID, error)
	GetById(id primitive.ObjectID) (*models.ProductImage, error)
	GetAll() ([]models.ProductImage, error)
	Delete(id primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, product models.ProductImage) (bool, error)
}

func NewProductImageRepository(dbClient *mongo.Collection) ProductImageRepositoryDB {
	return ProductImageRepositoryDB{ProductImageCollection: dbClient}
}
