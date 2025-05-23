package product_service

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

type ProductServiceWrapper interface {
	GetProducts(ctx context.Context, req ProductRequest) (resp map[uuid.UUID]ProductResponse, err error)
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

func (w *wrapper) Setup() ProductServiceWrapper {
	w.httpClient = http_client.NewRestClient(w.cfg.HostProductService, w.cfg.TokenProductService)
	return w
}

func (w *wrapper) GetProducts(ctx context.Context, req ProductRequest) (resp map[uuid.UUID]ProductResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "ProductServiceWrapper.GetProducts")
	defer span.End()

	body, status, err := w.httpClient.Post(ctx, "/internal/products/uuids", http.Header{}, req.Json())
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
