package delivery

import (
	"github.com/antonpodkur/Blog/internal/auth"
	"github.com/antonpodkur/Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.SignUp())
	authGroup.POST("/login", h.SignIn())
	authGroup.GET("/refresh", h.RefreshAccessToken())
	authGroup.Use(mw.AuthJwtMiddleware())
	authGroup.GET("/me", h.GetMe())
	authGroup.GET("/logout", h.LogOut())
}
