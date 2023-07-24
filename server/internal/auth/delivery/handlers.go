package delivery

import (
	"fmt"
	"net/http"

	"github.com/antonpodkur/Blog/config"
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/auth"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/antonpodkur/Blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

type authHandlers struct {
	cfg         *config.Config
	authUsecase auth.Usecase
}

func NewAuthHandlers(cfg *config.Config, authUsecase auth.Usecase) auth.Handlers {
	return &authHandlers{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}

func (h *authHandlers) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *db.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		newUser, err := h.authUsecase.SignUp(user)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": newUser}})
	}
}

func (h *authHandlers) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials *models.SignInInput

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		userWithTokens, err := h.authUsecase.SignIn(credentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.SetCookie("access_token", userWithTokens.AccessToken, h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("refresh_token", userWithTokens.RefreshToken, h.cfg.Jwt.RefreshTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("logged_in", "true", h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, false)

		userWithToken := models.UserWithTokenResponse{User: userWithTokens.User, AccessToken: userWithTokens.AccessToken}

		c.JSON(http.StatusOK, gin.H{"status": "success", "data": userWithToken.User})
	}
}

func (h *authHandlers) RefreshAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := "could not refresh access token"

		cookie, err := c.Cookie("refresh_token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
			return
		}

		sub, err := utils.ValidateToken(cookie, h.cfg.Jwt.RefreshTokenPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := h.authUsecase.GetUserById(fmt.Sprint(sub))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		access_token, err := utils.CreateToken(h.cfg.Jwt.AccessTokenExpiresIn, user.ID, h.cfg.Jwt.AccessTokenPrivateKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.SetCookie("access_token", access_token, h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("logged_in", "true", h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, false)

		c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
	}
}

func (a *authHandlers) LogOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
		c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func (a *authHandlers) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(*models.UserResponse)

		c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": currentUser}})
	}
}

func (a *authHandlers) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "I am fine")
	}
}
