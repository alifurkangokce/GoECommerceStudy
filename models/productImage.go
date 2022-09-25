package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductImage struct {
	Id        primitive.ObjectID `json:"_id`
	ProductId primitive.ObjectID `json:"productId":"productId"`
	Position  int8               `json:"position":"position"`
	CreatedAt time.Time          `json:"createdAt":"created_at"`
	UpdatedAt time.Time          `json:"updatedAt":"updated_at"`
	Width     int16              `json:"width":"width"`
	Height    int16              `json:"height":"height"`
	Src       string             `json:"src":"src"`
}
