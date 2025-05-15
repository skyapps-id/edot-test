package user

import (
	"context"

	"github.com/skyapps-id/edot-test/user-service/config"
	"github.com/skyapps-id/edot-test/user-service/repository"
)

type UserUsecase interface {
	Register(ctx context.Context, req RegisterRequest) (resp RegisterResponse, err error)
	Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error)
}

type usecase struct {
	cfg            config.Config
	userRepository repository.User
}

func NewUsecase(cfg config.Config, userRepository repository.User) UserUsecase {
	return &usecase{
		cfg:            cfg,
		userRepository: userRepository,
	}
}
