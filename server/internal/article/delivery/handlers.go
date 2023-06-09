package delivery

import (
	"net/http"
	"time"

	"github.com/antonpodkur/Blog/internal/article"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/gin-gonic/gin"
)


type articleHandlers struct {
}

func NewArticleHandlers() article.Handlers {
    return &articleHandlers{}
}

func (ah *articleHandlers) GetArticle() gin.HandlerFunc {
    return func(c *gin.Context) {
        article := models.Article{
            Content: "<h1>Hello from backend!</h1><h2>I hope i am rendering good.</h2>",
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }
        c.JSON(http.StatusOK, gin.H{"status": "success", "data": article})

    }
}

func (ah *articleHandlers) CreateArticle() gin.HandlerFunc {
	panic("not implemented") // TODO: Implement
}
