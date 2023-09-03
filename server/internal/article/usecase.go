package article

import (
	db "github.com/antonpodkur/Blog/db/sqlc"
)

type Usecase interface {
	GetAllArticles() (*[]db.ListArticlesRow, error)
	GetArticle(id string) (*db.GetArticleByIdRow, error)
	CreateArticle(article *db.Article) (*db.Article, error)
}
