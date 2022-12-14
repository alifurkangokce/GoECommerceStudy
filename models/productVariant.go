package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductVariant struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	Barcode        string             `bson:"barcode" json:"barcode"`
	CompareAtPrice string             `bson:"compareAtPrice" json:"compareAtPrice"`
	Price          string             `bson:"price" json:"price"`
	Position       int                `bson:"position" json:"position"`
	Sku            string             `bson:"sku" json:"sku"`
	Weight         int                `bson:"weight" json:"weight"`
	WeightUnit     int                `bson:"weightUnit" json:"weightUnit"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
	Title          string             `bson:"title" json:"title"`
}
