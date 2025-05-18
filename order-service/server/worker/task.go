package worker

import (
	"log"

	"github.com/skyapps-id/edot-test/order-service/container"
	"github.com/skyapps-id/edot-test/order-service/task"
	"go.uber.org/zap"
)

func SetupTask(container *container.Container) {
	paymentTask := task.NewWrapper(container.OrderUsecase)

	err := container.Worker.RegisterTasks(map[string]interface{}{
		"send_webhook": paymentTask.OrderCheck,
	})
	if err != nil {
		log.Fatal("Failed to register tasks:", zap.Error(err))
	}
}
