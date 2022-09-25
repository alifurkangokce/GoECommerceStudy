package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id          primitive.ObjectID `bson:"_id"`
	Active      bool               `bson:"active" json:"active"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
	Description string             `bson:"description" json:"description"`
	Livemode    bool               `bson:"livemode" json:"livemode"`
	Name        string             `bson:"name" json:"name,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
	CategoryId  primitive.ObjectID `bson:"categoryId" json:"categoryId"`
	Images      []ProductImage     `bson:"images" json:"images"`
	Variants    []ProductVariant   `bson:"variants" json:"variant"`
}
