package server

import (
	articleDelivery "github.com/antonpodkur/Blog/internal/article/delivery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func (s *Server) MapHandlers(router *gin.Engine) error {
    // ctx : context.TODO()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5173"}
    corsConfig.AllowCredentials = true

    router.Use(cors.New(corsConfig))

    // init handlers
    articleHandlers := articleDelivery.NewArticleHandlers()


    v1 := router.Group("api/v1")

    articleGroup := v1.Group("article")

    articleDelivery.MapArticleRoutes(articleGroup, articleHandlers)

    return nil
}
