package worker

import (
	"log"

	"github.com/skyapps-id/edot-test/order-service/container"
	"github.com/skyapps-id/edot-test/order-service/task"
	"go.uber.org/zap"
)

func SetupTask(container *container.Container) {
	handler := task.SendWebhook

	err := container.Worker.RegisterTasks(map[string]interface{}{
		"send_webhook": handler,
	})
	if err != nil {
		log.Fatal("Failed to register tasks:", zap.Error(err))
	}
}
