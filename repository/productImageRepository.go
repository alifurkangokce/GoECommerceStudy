package repository

import (
	"GoECommerceStudy/dto/productImageDto"
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

//go:generate mockgen -destination=../mocks/repository/mockProductImageRepository.go -package=repository GoECommerceStudy/repository ProductImageRepository
type ProductImageRepositoryDB struct {
	ProductImageCollection *mongo.Collection
}
type ProductImageRepository interface {
	Insert(dto productImageDto.ProductImageInsertDto) (bool, error)
	Delete(id primitive.ObjectID, imageId primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, dto productImageDto.ProductImageUpdateDto) (bool, error)
}

func NewProductImageRepository(dbClient *mongo.Collection) ProductImageRepositoryDB {
	return ProductImageRepositoryDB{ProductImageCollection: dbClient}
}

func (t ProductImageRepositoryDB) Insert(dto productImageDto.ProductImageInsertDto) (bool, error) {
	productImage := models.ProductImage{
		Id:        primitive.NewObjectID(),
		Position:  dto.Position,
		CreatedAt: time.Now(),
		Width:     dto.Width,
		Height:    dto.Height,
		Src:       dto.Src,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	upsert := false
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	result, err := t.ProductImageCollection.UpdateOne(ctx, bson.M{"_id": dto.ProductId}, bson.M{
		"$push": bson.M{"images": productImage},
	}, &opt)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	return result.ModifiedCount > 0, nil
}
func (t ProductImageRepositoryDB) Delete(id primitive.ObjectID, imageId primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.ProductImageCollection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{
			"$pull": bson.M{"images": bson.M{"_id": imageId}},
		})

	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (t ProductImageRepositoryDB) Update(id primitive.ObjectID, dto productImageDto.ProductImageUpdateDto) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Update().SetUpsert(false)

	k := reflect.ValueOf(dto)
	b := bson.M{}
	for i := 0; i < k.NumField(); i++ {
		y := k.Field(i)
		if !y.IsNil() && k.Type().Field(i).Name != "Id" {
			str := fmt.Sprintf("images.$.%s", k.Type().Field(i).Name)
			b[strings.ToLower(str)] = k.Field(i).Interface()
		}
	}

	result, err := t.ProductImageCollection.UpdateOne(ctx, bson.M{"_id": id, "images._id": dto.Id}, bson.M{
		"$set": b,
	}, opt)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil

}
