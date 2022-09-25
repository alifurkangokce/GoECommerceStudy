package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductVariant struct {
	Id             primitive.ObjectID `json:"_id`
	ProductId      primitive.ObjectID `json:"productId":"productId"`
	Barcode        string             `json:"barcode":"barcode"`
	CompareAtPrice string             `json:"compareAtPrice":"compareAtPrice"`
	Price          string             `json:"price":"price"`
	Position       int                `json:"position":"position"`
	Sku            string             `json:"sku" :"sku"`
	Weight         int                `json:"weight" :"weight"`
	WeightUnit     int                `json:"weightUnit" :"weightUnit"`
	CreatedAt      time.Time          `json:"createdAt" :"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" :"updatedAt"`
	Title          string             `json:"title":"title"`
}
