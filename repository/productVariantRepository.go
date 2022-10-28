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

type ProductVariantRepositoryDb struct {
	ProductVariantCollection *mongo.Collection
}
type ProductVariantRepository interface {
	Insert(productVariant models.ProductVariant) (primitive.ObjectID, error)
	Update(Id primitive.ObjectID, productVariant models.ProductVariant) (bool, error)
	Delete(Id primitive.ObjectID) (bool, error)
}

func NewProductVariantRepository(dbClient *mongo.Collection) ProductVariantRepositoryDb {
	return ProductVariantRepositoryDb{ProductVariantCollection: dbClient}
}
func (v ProductVariantRepositoryDb) Insert(variant models.ProductVariant) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	variant.Id = primitive.NewObjectID()
	variant.CreatedAt = time.Now()
	result, err := v.ProductVariantCollection.InsertOne(ctx, variant)
	if err != nil {
		log.Fatalln(err)
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), err
}
func (v ProductVariantRepositoryDb) Update(id primitive.ObjectID, variant models.ProductVariant) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := v.ProductVariantCollection.UpdateByID(ctx, id, bson.M{
		"$set": bson.M{
			"Barcode":   variant.Barcode,
			"Price":     variant.Price,
			"UpdatedAt": time.Now(),
		},
	})
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil

}
func (v ProductVariantRepositoryDb) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := v.ProductVariantCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}
