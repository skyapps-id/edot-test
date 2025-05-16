package product_service

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
)

type ProductRequest struct {
	Uuids []uuid.UUID `json:"uuids"`
}

func (t ProductRequest) Json() []byte {
	requestBody, err := json.Marshal(t)
	if err != nil {
		logger.Log.Error("error marsal ProductRequest.Json")
		return []byte{}
	}
	return requestBody
}

type ProductResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Price     float64   `json:"price"`
	ImageURL  *string   `json:"image_url"` // nullable
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
