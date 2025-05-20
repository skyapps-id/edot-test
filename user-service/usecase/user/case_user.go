package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/user-service/entity"
	"github.com/skyapps-id/edot-test/user-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/user-service/pkg/auth"
	"github.com/skyapps-id/edot-test/user-service/pkg/tracer"
	"gorm.io/gorm"
)

func (uc *usecase) Register(ctx context.Context, req RegisterRequest) (resp RegisterResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "UserUsecase.Register")
	defer span.End()

	_, err = uc.userRepository.FindByEmailOrPhone(ctx, req.Email, req.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("fail to get user"))
		return
	}
	if errors.Is(err, nil) {
		err = apperror.New(http.StatusBadRequest, fmt.Errorf("email or phone is used"))
		return
	}

	hashed, err := auth.HashPassword(req.Password)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, err)
	}

	err = uc.userRepository.CreateOrUpdate(ctx, entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashed,
	})

	resp.Email = req.Email

	return
}

func (uc *usecase) Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "UserUsecase.Login")
	defer span.End()

	user, err := uc.userRepository.FindByEmailOrPhone(ctx, req.ID, req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("fail to get user"))
		return
	}

	valid := auth.CheckPassword(user.Password, req.Password)
	if !valid {
		err = apperror.New(http.StatusUnauthorized, fmt.Errorf("wrong password"))
		return
	}

	token, err := auth.GenerateJWT(user.UUID.String(), uc.cfg.JwtSecret)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, err)
	}

	resp.UUID = user.UUID
	resp.Token = token

	return
}
