package user

import "github.com/google/uuid"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	ID       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UUID  uuid.UUID `json:"uuid"`
	Token string    `json:"token"`
}
