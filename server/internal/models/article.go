package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Article struct {
    Content string `json:"content" bson:"content" validate:"required" binding:"required"`
    CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

type ArticleDbResponse struct {
    ID primitive.ObjectID `json:"id" bson:"_id"`
    Content string `json:"content" bson:"content" validate:"required" binding:"required"`
    CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
