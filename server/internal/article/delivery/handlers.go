package delivery

import (
	"net/http"

	"github.com/antonpodkur/Blog/internal/article"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/gin-gonic/gin"
)


type articleHandlers struct {
    articleUsecase article.Usecase
}

func NewArticleHandlers(articleUsecase article.Usecase) article.Handlers {
    return &articleHandlers{
        articleUsecase: articleUsecase,
    }
}

func (ah *articleHandlers) GetAllArticles() gin.HandlerFunc {
    return func(c *gin.Context) {
        articles, err := ah.articleUsecase.GetAllArticles()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error retreiving articles"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success", "data": articles})
    }
}

func (ah *articleHandlers) GetArticle() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        article, err := ah.articleUsecase.GetArticle(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error finding article"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "success", "data": article})
    }
}

func (ah *articleHandlers) CreateArticle() gin.HandlerFunc {
    return func(c *gin.Context) {
        var article *models.Article  

        if err := c.ShouldBindJSON(&article); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid model"})
            return
        }

        newArticle, err := ah.articleUsecase.CreateArticle(article)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error creating article"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newArticle})
    } 
}
