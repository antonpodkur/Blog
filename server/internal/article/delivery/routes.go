package delivery

import (
	"github.com/antonpodkur/Blog/internal/article"
	"github.com/antonpodkur/Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapArticleRoutes(articleGroup *gin.RouterGroup, h article.Handlers, mw *middleware.MiddlewareManager) {
    articleGroup.GET("/", h.GetAllArticles())
    articleGroup.GET("/:id", h.GetArticle())
    articleGroup.Use(mw.AuthJwtMiddleware())
    articleGroup.POST("/", h.CreateArticle())
}
