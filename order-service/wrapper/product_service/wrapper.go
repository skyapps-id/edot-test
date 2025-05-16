package product_service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/skyapps-id/edot-test/order-service/config"
	"github.com/skyapps-id/edot-test/order-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/order-service/pkg/http_client"
)

type ProductServiceWrapper interface {
	GetProducts(ctx context.Context, req ProductRequest) (resp []ProductResponse, err error)
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

func (w *wrapper) GetProducts(ctx context.Context, req ProductRequest) (resp []ProductResponse, err error) {
	body, status, err := w.httpClient.Post(ctx, "/internal/products/uuids", http.Header{}, req.Json())
	if err != nil {
		err = apperror.New(status, err)
		return
	}

	var raw map[string]json.RawMessage
	if err = json.Unmarshal(body, &raw); err != nil {
		err = apperror.New(status, err)
		return
	}

	if err = json.Unmarshal(raw["data"], &resp); err != nil {
		err = apperror.New(status, err)
		return
	}

	return
}
