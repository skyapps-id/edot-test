package main

import (
	"github.com/skyapps-id/edot-test/product-service/container"
	"github.com/skyapps-id/edot-test/product-service/server/http"
)

func main() {
	container := container.Setup()

	http.StartHTTP(container)
}
