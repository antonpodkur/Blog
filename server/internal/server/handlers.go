package server

import (
	articleDelivery "github.com/antonpodkur/Blog/internal/article/delivery"
	filesDelivery "github.com/antonpodkur/Blog/internal/files/delivery"
	articleUC "github.com/antonpodkur/Blog/internal/article/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func (s *Server) MapHandlers(router *gin.Engine) error {

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"}
    corsConfig.AllowCredentials = true

    router.Use(cors.New(corsConfig))

    //init usecases
    articleUsecase := articleUC.NewArticleUsecase(s.cfg, s.mongoClient)

    // init handlers
    articleHandlers := articleDelivery.NewArticleHandlers(articleUsecase)
    filesHandlers := filesDelivery.NewFilesHandlers()


    v1 := router.Group("api/v1")

    articleGroup := v1.Group("article")
    filesGroup := v1.Group("files")


    articleDelivery.MapArticleRoutes(articleGroup, articleHandlers)
    filesDelivery.MapFilesRoutes(filesGroup, filesHandlers)

    return nil
}
