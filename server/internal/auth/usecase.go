package auth

import (
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/models"
)

type Usecase interface {
	SignUp(*db.User) (*models.UserResponse, error)
	SignIn(input *models.SignInInput) (*models.UserWithTokens, error)
	GetUserById(string) (*models.UserResponse, error)
	GetUserByEmail(string) (*models.UserResponse, error)
}
