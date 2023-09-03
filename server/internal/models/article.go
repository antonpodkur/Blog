package models

import (
	"time"

	db "github.com/antonpodkur/Blog/db/sqlc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleResponse struct {
    ID        string `json:"id,omitempty"` 
    Title     string `json:"title,omitempty"`
    Content   string `json:"content,omitempty"` 
    CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
    UserID    string  `json:"userId,omitempty"`
}

func ArticleToArticleReponse(article *db.Article) ArticleResponse {
    return ArticleResponse{
        ID: article.ID.String(),
        Title: article.Title,
        Content: article.Content,
        CreatedAt: article.CreatedAt,
        UpdatedAt: article.UpdatedAt,
        UserID: article.UserID.String(), 
    }
}

func ArticlesToArticleResponses(articles *[]db.Article) []ArticleResponse {
    responses := make([]ArticleResponse, len(*articles))
    for i, article := range *articles {
        responses[i] = ArticleToArticleReponse(&article)
    }
    return responses
}

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
