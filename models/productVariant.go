package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductVariant struct {
	Id             primitive.ObjectID `json:"_id`
	ProductId      primitive.ObjectID `json:"productId,omitempty":"productId"`
	Barcode        string             `json:"barcode,omitempty":"barcode"`
	CompareAtPrice string             `json:"compareAtPrice,omitempty":"compareAtPrice"`
	Price          string             `json:"price,omitempty":"price"`
	Position       int                `json:"position,omitempty":"position"`
	Sku            string             `json:"sku,omitempty" :"sku"`
	Weight         int                `json:"weight,omitempty" :"weight"`
	WeightUnit     int                `json:"weightUnit,omitempty" :"weightUnit"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" :"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt,omitempty" :"updatedAt"`
	Title          string             `json:"title,omitempty":"title"`
}
