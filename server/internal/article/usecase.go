package article

import (
	db "github.com/antonpodkur/Blog/db/sqlc"
)

type Usecase interface {
	GetAllArticles() (*[]db.Article, error)
	GetArticle(id string) (*db.Article, error)
	CreateArticle(article *db.Article) (*db.Article, error)
}
