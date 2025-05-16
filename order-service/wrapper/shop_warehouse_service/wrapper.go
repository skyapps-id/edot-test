package shop_warehouse_service

import (
	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/pkg/http_client"
)

type ShopWarehousetServiceWrapper interface {
}

type wrapper struct {
	cfg        config.Config
	httpClient http_client.RestClient
}

func NewWrapper(cfg config.Config) ShopWarehousetServiceWrapper {
	return wrapper{
		cfg: cfg,
	}
}

func (w *wrapper) Setup() *wrapper {
	w.httpClient = http_client.NewRestClient(w.cfg.HostShopWarehouseService, w.cfg.TokenShopWarehouseService)
	return w
}
