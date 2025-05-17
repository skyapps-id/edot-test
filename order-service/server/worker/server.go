package worker

import "github.com/RichardKnop/machinery/v1"

func StartWorker(taskserver *machinery.Server) error {

	worker := taskserver.NewWorker("machinery_worker", 3)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil

}
