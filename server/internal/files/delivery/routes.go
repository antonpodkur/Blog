package delivery

import (
	"github.com/antonpodkur/Blog/internal/files"
	"github.com/antonpodkur/Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapFilesRoutes(filesGroup *gin.RouterGroup, h files.Handlers, mw *middleware.MiddlewareManager) {
    filesGroup.GET("/:id", h.Retrieve())
	filesGroup.Use(mw.AuthJwtMiddleware())
    filesGroup.POST("/", h.Upload())
}
