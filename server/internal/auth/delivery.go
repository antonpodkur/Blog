package auth

import "github.com/gin-gonic/gin"

type Handlers interface {
	SignUp() gin.HandlerFunc
	SignIn() gin.HandlerFunc
	RefreshAccessToken() gin.HandlerFunc
	LogOut() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	Test() gin.HandlerFunc
}
