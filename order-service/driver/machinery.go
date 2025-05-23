package driver

import (
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	app "github.com/skyapps-id/edot-test/order-service/config"
	"go.uber.org/zap"
)

func GetMachineryServer(cfg app.Config) *machinery.Server {
	cnf := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Broker:          cfg.RedisUrl,
		ResultBackend:   cfg.RedisUrl,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Fatal("Failed to create Machinery server:", zap.Error(err))
	}

	return server
}
