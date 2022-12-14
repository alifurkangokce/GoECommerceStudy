package repository

import (
	"GoECommerceStudy/dto/productVariantDto"
	"GoECommerceStudy/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"strings"
	"time"
)

type ProductVariantRepositoryDb struct {
	ProductVariantCollection *mongo.Collection
}
type ProductVariantRepository interface {
	Insert(variant productVariantDto.ProductVariantInsertDto) (bool, error)
	Update(Id primitive.ObjectID, productVariant productVariantDto.ProductVariantUpdateRequestDto) (bool, error)
	Delete(Id primitive.ObjectID, variantId primitive.ObjectID) (bool, error)
}

func NewProductVariantRepository(dbClient *mongo.Collection) ProductVariantRepositoryDb {
	return ProductVariantRepositoryDb{ProductVariantCollection: dbClient}
}
func (v ProductVariantRepositoryDb) Insert(variant productVariantDto.ProductVariantInsertDto) (bool, error) {
	productVariant := models.ProductVariant{
		Id:             primitive.NewObjectID(),
		Barcode:        variant.Barcode,
		CompareAtPrice: variant.CompareAtPrice,
		Price:          variant.Price,
		Position:       variant.Position,
		Sku:            variant.Sku,
		Weight:         variant.Weight,
		WeightUnit:     variant.WeightUnit,
		CreatedAt:      time.Now(),
		Title:          variant.Title,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	upsert := false
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	result, err := v.ProductVariantCollection.UpdateOne(ctx, bson.M{"_id": variant.ProductId}, bson.M{
		"$push": bson.M{"variants": productVariant},
	}, &opt)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (v ProductVariantRepositoryDb) Update(id primitive.ObjectID, variant productVariantDto.ProductVariantUpdateRequestDto) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Update().SetUpsert(false)

	k := reflect.ValueOf(variant)
	b := bson.M{}
	for i := 0; i < k.NumField(); i++ {
		y := k.Field(i)
		if !y.IsNil() && k.Type().Field(i).Name != "Id" {
			str := fmt.Sprintf("variants.$.%s", k.Type().Field(i).Name)
			b[strings.ToLower(str)] = k.Field(i).Interface()
		}
	}

	result, err := v.ProductVariantCollection.UpdateOne(ctx, bson.M{"_id": id, "variants._id": variant.Id}, bson.M{
		"$set": b,
	}, opt)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil

}
func (v ProductVariantRepositoryDb) Delete(id primitive.ObjectID, variantId primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := v.ProductVariantCollection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{
			"$pull": bson.M{"variants": bson.M{"_id": variantId}},
		})

	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}
