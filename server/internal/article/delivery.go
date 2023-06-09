package article

import "github.com/gin-gonic/gin"

type Handlers interface {
    GetArticle() gin.HandlerFunc
    CreateArticle() gin.HandlerFunc
}
