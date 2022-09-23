package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	Id            primitive.ObjectID `json:"_id"`
	Active        bool               `json:"active,omitempty" :"active"`
	Available     bool               `json:"available,omitempty" :"available"`
	Name          string             `json:"name,omitempty" :"name"`
	SubCategories []Category         `json:"subCategories" :"subCategories"`
	Description   string             `json:"description,omitempty" :"description"`
	ParentId      int                `json:"parentId,omitempty" :"parentId"`
	CreatedAt     time.Time          `json:"createdAt,omitempty":"created_at"`
	UpdatedAt     time.Time          `json:"updatedAt,omitempty":"updated_at"`
}
