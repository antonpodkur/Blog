package delivery

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/antonpodkur/Blog/internal/files"
	"github.com/antonpodkur/Blog/pkg/image"
	"github.com/gin-gonic/gin"
)

type filesHandlers struct {
}

func NewFilesHandlers() files.Handlers {
    return &filesHandlers{}
}

func (fh *filesHandlers) Upload() gin.HandlerFunc {
    return func(c *gin.Context) {
        _, err := os.Stat("data")    
        if os.IsNotExist(err) {
            errDir := os.MkdirAll("data", 0755)
            if errDir != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "internal system error raised"})
                return
            }
        }

        fileheader, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
            return
        }


        file, err := fileheader.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
            return
        }
        defer file.Close()

        buffer, err := io.ReadAll(file)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
            return
        }

        filename, err := image.ImageProcessing(buffer, 40, "data")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": "Error processing file"})
            return
        }

        fileUrl := "http://localhost:4000/api/v1/files/" + filename
        c.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"file": fileUrl}})
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
