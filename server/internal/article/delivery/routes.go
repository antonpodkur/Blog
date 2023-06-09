package delivery

import (
	"github.com/antonpodkur/Blog/internal/article"
	"github.com/gin-gonic/gin"
)

func MapArticleRoutes(articleGroup *gin.RouterGroup, h article.Handlers) {
    articleGroup.GET("/post", h.GetArticle())
}
