package shop_warehouse_service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/skyapps-id/edot-test/product-service/config"
	"github.com/skyapps-id/edot-test/product-service/pkg/http_client"
	"github.com/skyapps-id/edot-test/product-service/pkg/tracer"
)

type ShopWarehousetServiceWrapper interface {
	GetProductStock(ctx context.Context, req GetProductStockRequest) (resp GetProductStockResponse, err error)
}

type wrapper struct {
	cfg        config.Config
	httpClient http_client.RestClient
}

func NewWrapper(cfg config.Config) *wrapper {
	return &wrapper{
		cfg: cfg,
	}
}

func (w *wrapper) Setup() ShopWarehousetServiceWrapper {
	w.httpClient = http_client.NewRestClient(w.cfg.HostShopWarehouseService, w.cfg.TokenShopWarehouseService)
	return w
}

func (w *wrapper) GetProductStock(ctx context.Context, req GetProductStockRequest) (resp GetProductStockResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopWarehousetServiceWrapper.GetProductStock")
	defer span.End()

	body, status, err := w.httpClient.Get(ctx, "/internal/warehouses/product-stock/"+req.ProductUUID.String(), http.Header{})
	if err != nil {
		err = fmt.Errorf("request api status %d error %w", status, err)
		return
	}

	var raw map[string]json.RawMessage
	if err = json.Unmarshal(body, &raw); err != nil {
		err = fmt.Errorf("resp body unmarshal raw error: %w", err)
		return
	}

	if err = json.Unmarshal(raw["data"], &resp); err != nil {
		err = fmt.Errorf("resp body unmarshal data error: %w", err)
		return
	}

	return
}
