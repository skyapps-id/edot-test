package auth

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) (uuid.UUID, bool) {
	userID := c.Get("user_id")
	if userID == nil {
		return uuid.UUID{}, false
	}

	id, ok := userID.(uuid.UUID)
	return id, ok
}
