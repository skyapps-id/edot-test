package main

import (
	"github.com/skyapps-id/edot-test/order-service/container"
	"github.com/skyapps-id/edot-test/order-service/server/http"
)

func main() {
	container := container.Setup()

	http.StartHTTP(container)
}
