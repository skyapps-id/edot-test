package main

import (
	"github.com/skyapps-id/edot-test/user-service/container"
	"github.com/skyapps-id/edot-test/user-service/server/http"
)

func main() {
	container := container.Setup()

	http.StartHTTP(container)
}
