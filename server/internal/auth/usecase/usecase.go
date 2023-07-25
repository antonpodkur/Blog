package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/antonpodkur/Blog/config"
	db "github.com/antonpodkur/Blog/db/sqlc"
	"github.com/antonpodkur/Blog/internal/auth"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/antonpodkur/Blog/pkg/utils"
	"github.com/google/uuid"
)

type authUsecase struct {
	cfg *config.Config
	db  *db.Queries
	ctx context.Context
}

func NewAuthUsecase(cfg *config.Config, db *db.Queries, ctx context.Context) auth.Usecase {
	return &authUsecase{
		cfg: cfg,
		db:  db,
		ctx: ctx,
	}
}

func (u *authUsecase) SignUp(user *db.User) (*models.UserResponse, error) {

	hashedPassword, _ := utils.HashPassword(user.Password)

	args := &db.CreateUserParams{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Photo:     "default.webp",
		Verified:  true,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	createdUser, err := u.db.CreateUser(u.ctx, *args)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserFilteredResponse(createdUser)

	return &userResponse, nil
}

func (u *authUsecase) SignIn(input *models.SignInInput) (*models.UserWithTokens, error) {
	user, err := u.db.GetUserByEmail(u.ctx, input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := utils.VerifyPassword(user.Password, input.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	access_token, err := utils.CreateToken(u.cfg.Jwt.AccessTokenExpiresIn, user.ID, u.cfg.Jwt.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refresh_token, err := utils.CreateToken(u.cfg.Jwt.RefreshTokenExpiresIn, user.ID, u.cfg.Jwt.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	userWithTokens := &models.UserWithTokens{
		User:         models.UserFilteredResponse(user),
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	return userWithTokens, nil
}

func (u *authUsecase) GetUserById(id string) (*models.UserResponse, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("user does not exist")
	}
	user, err := u.db.GetUserById(u.ctx, uuid)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	userResponse := models.UserFilteredResponse(user)

	return &userResponse, nil
}

func (u *authUsecase) GetUserByEmail(email string) (*models.UserResponse, error) {
	user, err := u.db.GetUserByEmail(u.ctx, email)
	if err != nil {
		return nil, errors.New("user does not exist")
	}

	userResponse := models.UserFilteredResponse(user)

	return &userResponse, nil

}
