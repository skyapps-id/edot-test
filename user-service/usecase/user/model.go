package user

import "github.com/google/uuid"

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
