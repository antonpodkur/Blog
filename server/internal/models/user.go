package models

import (
	"time"

	db "github.com/antonpodkur/Blog/db/sqlc"
)

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserWithTokenResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}

type UserWithTokens struct {
	User         UserResponse
	AccessToken  string
	RefreshToken string
}

func UserFilteredResponse(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
