package worker

import (
	"github.com/skyapps-id/edot-test/order-service/container"
)

func StartWorker(container *container.Container) error {

	SetupTask(container)

	worker := container.Worker.NewWorker("machinery_worker", 3)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil

}
