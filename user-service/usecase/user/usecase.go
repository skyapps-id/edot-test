package user

import (
	"context"

	"github.com/skyapps-id/edot-test/user-service/repository"
)

type UserUsecase interface {
	Register(ctx context.Context, req RegisterRequest) (resp RegisterResponse, err error)
}

type usecase struct {
	userRepository repository.User
}

func NewUsecase(userRepository repository.User) UserUsecase {
	return &usecase{
		userRepository: userRepository,
	}
}
