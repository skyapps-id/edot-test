package shop_warehouse_service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/pkg/http_client"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

type ShopWarehousetServiceWrapper interface {
	GetProductStock(ctx context.Context, req ProductStockRequest) (resp map[uuid.UUID]ProductStockResponse, err error)
	ProductStockAddition(ctx context.Context, req ProductStockAdditionRequest) (resp ProductStockAdditionResponse, err error)
	ProductStockReduction(ctx context.Context, req ProductStockReductionRequest) (resp ProductStockReductionResponse, err error)
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

func (w *wrapper) GetProductStock(ctx context.Context, req ProductStockRequest) (resp map[uuid.UUID]ProductStockResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopWarehousetServiceWrapper.GetProductStock")
	defer span.End()

	body, status, err := w.httpClient.Post(ctx, "/internal/warehouses/product-stock", http.Header{}, req.Json())
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

func (w *wrapper) ProductStockAddition(ctx context.Context, req ProductStockAdditionRequest) (resp ProductStockAdditionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopWarehousetServiceWrapper.ProductStockAddition")
	defer span.End()

	body, status, err := w.httpClient.Post(ctx, "/internal/warehouses/product-stock-addition", http.Header{}, req.Json())
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

func (w *wrapper) ProductStockReduction(ctx context.Context, req ProductStockReductionRequest) (resp ProductStockReductionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ShopWarehousetServiceWrapper.ProductStockReduction")
	defer span.End()

	body, status, err := w.httpClient.Post(ctx, "/internal/warehouses/product-stock-reduction", http.Header{}, req.Json())
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
