package delivery

import (
	"github.com/antonpodkur/Blog/internal/files"
	"github.com/gin-gonic/gin"
)

func MapFilesRoutes(filesGroup *gin.RouterGroup, h files.Handlers) {
    filesGroup.GET("/:id", h.Retrieve())
    filesGroup.POST("/", h.Upload())
}
