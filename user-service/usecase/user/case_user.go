package user

import (
	"context"

	"github.com/skyapps-id/edot-test/user-service/entity"
	"go.opentelemetry.io/otel"
)

func (uc *usecase) Register(ctx context.Context, req RegisterRequest) (resp RegisterResponse, err error) {
	tracer := otel.Tracer("UserUsecase")
	ctx, span := tracer.Start(ctx, "Register")
	defer span.End()

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}

	err = uc.userRepository.CreateOrUpdate(ctx, user)

	resp.UUID = user.UUID

	return
}
