package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Article struct {
    Title string `json:"title" bson:"title" validate:"required" binding:"required"`
    Content string `json:"content" bson:"content" validate:"required" binding:"required"`
    CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
    UserId primitive.ObjectID `json:"userId,omitempty" bson:"userId"`
}

type ArticleDbResponse struct {
    ID primitive.ObjectID `json:"id" bson:"_id"`
    Title string `json:"title" bson:"title" validate:"required" binding:"required"`
    Content string `json:"content" bson:"content" validate:"required" binding:"required"`
    CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
    UserId primitive.ObjectID `json:"userId" bson:"userId"`
}
