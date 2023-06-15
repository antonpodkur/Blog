package delivery

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/antonpodkur/Blog/internal/files"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type filesHandlers struct {
}

func NewFilesHandlers() files.Handlers {
    return &filesHandlers{}
}

func (fh *filesHandlers) Upload() gin.HandlerFunc {
    return func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
            return
        }

        id, _ := uuid.NewRandom()
        fileId := id.String() 

        dst, err := os.Create(filepath.Join("data", fileId+filepath.Ext(file.Filename)))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
            return
        }

        defer dst.Close()

        openedFile, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
            return
        }

        if _, err := io.Copy(dst, openedFile); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"status": "success", "data": fileId+filepath.Ext(file.Filename)})
    }
}

func (fh *filesHandlers) Retrieve() gin.HandlerFunc {
    return func(c *gin.Context) {
        fileId := c.Param("id")

        filePath := filepath.Join("data", fileId)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
            c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Error finding specified file"})
            return
		}

		file, err := os.Open(filePath)
		if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
            return
		}
		defer file.Close()

        c.Header("Content-Disposition", "attachment; filename="+fileId)
        c.Header("Content-Type", "application/octet-stream")
        c.Header("Content-Transfer-Encoding", "binary")

        c.File(filePath)
    }
}
