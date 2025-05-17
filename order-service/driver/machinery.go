package driver

import (
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/skyapps-id/edot-test/order-service/task"
	"go.uber.org/zap"
)

func GetMachineryServer() *machinery.Server {
	cnf := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		Broker:          "redis://localhost:6379",
		ResultBackend:   "redis://localhost:6379",
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

	err = server.RegisterTasks(map[string]interface{}{
		"send_webhook": task.SendWebhook,
	})
	if err != nil {
		log.Fatal("Failed to register tasks:", zap.Error(err))
	}

	return server
}
