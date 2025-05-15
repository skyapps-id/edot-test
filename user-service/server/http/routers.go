package http

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/user-service/container"
	"github.com/skyapps-id/edot-test/user-service/handler/health"
	"github.com/skyapps-id/edot-test/user-service/handler/user"
)

func Router(server *echo.Echo, container *container.Container) {
	// Handler
	health := health.NewHandler()
	user := user.NewHandler(container.UserUsecase)

	server.GET("/", health.Health)
	server.POST("/register", user.Register)
	server.POST("/login", user.Login)
}
