package files

import "github.com/gin-gonic/gin"

type Handlers interface {
    Upload() gin.HandlerFunc
    Retrieve() gin.HandlerFunc
}
