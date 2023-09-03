package article

import (
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/models"
)

type Usecase interface {
	GetAllArticles() (*[]models.ArticleResponse, error)
	GetArticle(id string) (*models.ArticleResponse, error)
	CreateArticle(article *db.Article) (*models.ArticleResponse, error)
}
