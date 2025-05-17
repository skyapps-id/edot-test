package main

import (
	"os"

	"github.com/skyapps-id/edot-test/order-service/container"
	"github.com/skyapps-id/edot-test/order-service/server/http"
	"github.com/skyapps-id/edot-test/order-service/server/worker"
	"github.com/urfave/cli"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
}

func main() {
	container := container.Setup()

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the server that takes task input",
			Action: func(c *cli.Context) {
				http.StartHTTP(container)
			},
		},
		{
			Name:  "worker",
			Usage: "Run the worker that will consume tasks",
			Action: func(c *cli.Context) {
				worker.StartWorker(container.Worker)
			},
		},
	}
	app.Run(os.Args)
}
