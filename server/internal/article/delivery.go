package article

import "github.com/gin-gonic/gin"

type Handlers interface {
    GetAllArticles() gin.HandlerFunc
    GetArticle() gin.HandlerFunc
    CreateArticle() gin.HandlerFunc
}
