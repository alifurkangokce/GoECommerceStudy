package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductImage struct {
	Id        primitive.ObjectID `json:"_id`
	ProductId primitive.ObjectID `json:"productId,omitempty":"productId"`
	Position  int8               `json:"position,omitempty":"position"`
	CreatedAt time.Time          `json:"createdAt,omitempty":"created_at"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty":"updated_at"`
	Width     int16              `json:"width,omitempty":"width"`
	Height    int16              `json:"height,omitempty":"height"`
	Src       string             `json:"src,omitempty":"src"`
}
