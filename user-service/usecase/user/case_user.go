package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/user-service/entity"
	"github.com/skyapps-id/edot-test/user-service/pkg/apperror"
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

	err = uc.userRepository.CreateOrUpdate(ctx, entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})

	resp.Email = req.Email

	return
}
