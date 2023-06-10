package server

import (
	articleDelivery "github.com/antonpodkur/Blog/internal/article/delivery"
	articleUC "github.com/antonpodkur/Blog/internal/article/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func (s *Server) MapHandlers(router *gin.Engine) error {
    // ctx : context.TODO()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5173"}
    corsConfig.AllowCredentials = true

    router.Use(cors.New(corsConfig))

    //init usecases
    articleUsecase := articleUC.NewArticleUsecase(s.cfg, s.mongoClient)

    // init handlers
    articleHandlers := articleDelivery.NewArticleHandlers(articleUsecase)


    v1 := router.Group("api/v1")

    articleGroup := v1.Group("article")

    articleDelivery.MapArticleRoutes(articleGroup, articleHandlers)

    return nil
}
