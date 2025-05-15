package main

import (
	"github.com/skyapps-id/edot-test/shop-warehouse-service/container"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/server/http"
)

func main() {
	container := container.Setup()

	http.StartHTTP(container)
}
