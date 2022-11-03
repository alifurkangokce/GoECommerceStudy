package productVariantDto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductVariantUpdateRequestDto struct {
	Id             *primitive.ObjectID `bson:"_id" validate:"required"`
	ProductId      *primitive.ObjectID `bson:"productId" json:"productId"`
	Barcode        *string             `bson:"barcode" json:"barcode"`
	CompareAtPrice *string             `bson:"compareAtPrice" json:"compareAtPrice"`
	Price          *string             `bson:"price" json:"price"`
	Position       *int                `bson:"position" json:"position"`
	Sku            *string             `bson:"sku" json:"sku"`
	Weight         *int                `bson:"weight" json:"weight"`
	WeightUnit     *int                `bson:"weightUnit" json:"weightUnit"`
	UpdatedAt      *time.Time          `bson:"updatedAt" json:"updatedAt"`
	Title          *string             `bson:"title" json:"title"`
}
