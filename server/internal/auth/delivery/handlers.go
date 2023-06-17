package delivery

import (
	"fmt"
	"github.com/antonpodkur/Blog/config"
	"github.com/antonpodkur/Blog/internal/auth"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/antonpodkur/Blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
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
		var user *models.SignUpInput

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if user.Password != user.PasswordConfirm {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
			return
		}

		newUser, err := h.authUsecase.SignUp(user)

		if err != nil {
			if strings.Contains(err.Error(), "email already exist") {
				c.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
				return
			}
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": models.UserFilteredResponse(newUser)}})

	}
}

func (h *authHandlers) SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials *models.SignInInput

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := h.authUsecase.GetUserByEmail(credentials.Email)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
			return
		}

		// Generate Tokens
		access_token, err := utils.CreateToken(h.cfg.Jwt.AccessTokenExpiresIn, user.ID, h.cfg.Jwt.AccessTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		refresh_token, err := utils.CreateToken(h.cfg.Jwt.RefreshTokenExpiresIn, user.ID, h.cfg.Jwt.RefreshTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.SetCookie("access_token", access_token, h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("refresh_token", refresh_token, h.cfg.Jwt.RefreshTokenMaxAge*60, "/", "localhost", false, true)
		c.SetCookie("logged_in", "true", h.cfg.Jwt.AccessTokenMaxAge*60, "/", "localhost", false, false)

		c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
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
		currentUser := c.MustGet("currentUser").(*models.UserDBResponse)

		c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.UserFilteredResponse(currentUser)}})
	}
}

func (a *authHandlers) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "I am fine")
	}
}
