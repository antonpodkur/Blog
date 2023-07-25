package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/antonpodkur/Blog/config"
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/article"
	"github.com/google/uuid"
)

type articleUsecase struct {
	cfg *config.Config
	db  *db.Queries
}

func NewArticleUsecase(cfg *config.Config, db *db.Queries) article.Usecase {
	return &articleUsecase{
		cfg: cfg,
		db:  db,
	}
}

func (au *articleUsecase) GetAllArticles() (*[]db.Article, error) {
	ctx := context.TODO()

	args := &db.ListArticlesParams{Limit: 50, Offset: 0}

	articles, err := au.db.ListArticles(ctx, *args)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}

func (au *articleUsecase) GetArticle(id string) (*db.Article, error) {
	ctx := context.TODO()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	article, err := au.db.GetArticleById(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (au *articleUsecase) CreateArticle(article *db.Article) (*db.Article, error) {
	ctx := context.TODO()

	args := db.CreateArticleParams{
		Title:     article.Title,
		Content:   article.Content,
		UpdatedAt: time.Now(),
		UserID:    article.UserID,
	}

	createdArticle, err := au.db.CreateArticle(ctx, args)
	if err != nil {
		return nil, err
	}

	return &createdArticle, nil
}
