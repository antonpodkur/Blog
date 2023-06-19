package server

import (
	"context"

	articleDelivery "github.com/antonpodkur/Blog/internal/article/delivery"
	articleUC "github.com/antonpodkur/Blog/internal/article/usecase"
	authDelivery "github.com/antonpodkur/Blog/internal/auth/delivery"
	authUC "github.com/antonpodkur/Blog/internal/auth/usecase"
	filesDelivery "github.com/antonpodkur/Blog/internal/files/delivery"
	"github.com/antonpodkur/Blog/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func (s *Server) MapHandlers(router *gin.Engine) error {
    ctx := context.TODO()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"}
    corsConfig.AllowCredentials = true

    router.Use(cors.New(corsConfig))

    //init usecases
    authUsecase := authUC.NewAuthUsecase(s.cfg, s.mongoClient, ctx)
    articleUsecase := articleUC.NewArticleUsecase(s.cfg, s.mongoClient)

    // init handlers
    authHandlers := authDelivery.NewAuthHandlers(s.cfg, authUsecase)
    articleHandlers := articleDelivery.NewArticleHandlers(articleUsecase)
    filesHandlers := filesDelivery.NewFilesHandlers()

    
    // init middleware manager
    mw := middleware.NewMiddlewareManager(s.cfg, authUsecase)

    // create router groups
    v1 := router.Group("api/v1")

    authGroup := v1.Group("auth")
    articleGroup := v1.Group("article")
    filesGroup := v1.Group("files")


    // map routes
    authDelivery.MapAuthRoutes(authGroup, authHandlers, mw)
    articleDelivery.MapArticleRoutes(articleGroup, articleHandlers, mw)
    filesDelivery.MapFilesRoutes(filesGroup, filesHandlers, mw)

    return nil
}
