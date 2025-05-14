package container

import (
	"github.com/skyapps-id/edot-test/user-service/config"
	"github.com/skyapps-id/edot-test/user-service/driver"
	"github.com/skyapps-id/edot-test/user-service/repository"
	"github.com/skyapps-id/edot-test/user-service/usecase/user"
)

type Container struct {
	Config      config.Config
	UserUsecase user.UserUsecase
}

func Setup() *Container {
	// Load Config
	config := config.Load()

	// Database
	database := driver.NewGormDatabase(config)

	// Repository
	repo_user := repository.NewUserRepository(database)

	// Usecase
	userUsecase := user.NewUsecase(repo_user)

	return &Container{
		Config:      config,
		UserUsecase: userUsecase,
	}
}
