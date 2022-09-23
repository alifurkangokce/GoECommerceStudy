package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	Id          primitive.ObjectID `bson:"_id"`
	Active      bool               `bson:"active" json:"active,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Livemode    bool               `bson:"livemode" json:"livemode,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt,omitempty"`
	CategoryId  primitive.ObjectID `bson:"categoryId" json:"categoryId,omitempty"`
	Images      []ProductImage     `bson:"images" json:"images,omitempty"`
	Variants    []ProductVariant   `bson:"variants" json:"variant,omitempty"`
}
