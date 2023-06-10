package article

import "github.com/antonpodkur/Blog/internal/models"


type Usecase interface {
    GetAllArticles() ([]*models.ArticleDbResponse, error)
    GetArticle(id string) (*models.ArticleDbResponse, error)
    CreateArticle(article *models.Article) (*models.ArticleDbResponse, error)
}
