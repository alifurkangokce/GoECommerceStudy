package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductImage struct {
	Id        primitive.ObjectID `bson:"_id"`
	ProductId primitive.ObjectID `bson:"productId" json:"productId" validate:"required"`
	Position  int8               `bson:"position" json:"position"`
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updated_at"`
	Width     int16              `bson:"width" json:"width"`
	Height    int16              `bson:"height" json:"height"`
	Src       string             `bson:"src" json:"src" validate:"required"`
}
