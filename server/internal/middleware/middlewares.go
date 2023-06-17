package middleware

import (
	"github.com/antonpodkur/Blog/config"
	"github.com/antonpodkur/Blog/internal/auth"
)

type MiddlewareManager struct {
    cfg *config.Config
    authUsecase auth.Usecase
}

func NewMiddlewareManager(cfg *config.Config, authUsecase auth.Usecase) *MiddlewareManager {
    return &MiddlewareManager{
        cfg: cfg,
        authUsecase: authUsecase,
    }
}
