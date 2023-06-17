package auth

import (
	"github.com/antonpodkur/Blog/internal/models"
)

type Usecase interface {
	SignUp(*models.SignUpInput) (*models.UserDBResponse, error)
	SignIn(input *models.SignInInput) (*models.UserDBResponse, error)
	GetUserById(string) (*models.UserDBResponse, error)
	GetUserByEmail(string) (*models.UserDBResponse, error)
}
