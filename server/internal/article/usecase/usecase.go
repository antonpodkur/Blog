package usecase

import (
	"context"
	"time"

	"github.com/antonpodkur/Blog/config"
	"github.com/antonpodkur/Blog/internal/article"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/antonpodkur/Blog/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type articleUsecase struct {
    cfg *config.Config
    mongoClient *mongo.Client
}

func NewArticleUsecase(cfg *config.Config, mongoClient *mongo.Client) article.Usecase {
    return &articleUsecase{
        cfg: cfg,
        mongoClient: mongoClient,
    }
}

func (au *articleUsecase) GetAllArticles() ([]*models.ArticleDbResponse, error) {
    articlesCollection := db.OpenCollection(au.mongoClient, au.cfg, "articles")
    ctx := context.TODO()
    
    var articles []*models.ArticleDbResponse

    cur, err := articlesCollection.Find(ctx, bson.D{})
    if err != nil {
        return nil, err
    }

    for cur.Next(ctx) {
        var article *models.ArticleDbResponse

        err := cur.Decode(&article)
        if err != nil {
            return nil, err
        }

        articles = append(articles, article)
    }
    
    if err := cur.Err(); err != nil {
        return nil, err
    }

    cur.Close(ctx)

    return articles, nil 
}

func (au *articleUsecase) GetArticle(id string) (*models.ArticleDbResponse, error) {
    articlesCollection := db.OpenCollection(au.mongoClient, au.cfg, "articles")
    ctx := context.TODO()

    oid, _ := primitive.ObjectIDFromHex(id)

    var article *models.ArticleDbResponse

    query := bson.M{"_id": oid}
    err := articlesCollection.FindOne(ctx, query).Decode(&article)
    if err != nil {
        return nil, err
    }

    return article, nil
}

func (au *articleUsecase) CreateArticle(article *models.Article) (*models.ArticleDbResponse, error) {
    articlesCollection := db.OpenCollection(au.mongoClient, au.cfg, "articles")
    ctx := context.TODO()

    article.CreatedAt = time.Now()
    article.UpdatedAt = article.CreatedAt

    res, err := articlesCollection.InsertOne(ctx, &article)
    if err != nil {
        return nil, err
    }

    var newArticle *models.ArticleDbResponse
    query := bson.M{"_id": res.InsertedID}

    err = articlesCollection.FindOne(ctx, query).Decode(&newArticle)
    if err != nil {
        return nil, err
    }

    return newArticle, nil
}


