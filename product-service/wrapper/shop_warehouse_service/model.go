package shop_warehouse_service

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/product-service/pkg/logger"
)

type GetProductStockRequest struct {
	ProductUUID uuid.UUID `json:"uuid"`
}

func (t GetProductStockRequest) Json() []byte {
	requestBody, err := json.Marshal(t)
	if err != nil {
		logger.Log.Error("error marsal ProductStockReductionRequest.Json")
		return []byte{}
	}
	return requestBody
}

type GetProductStockResponse struct {
	ProductUUID uuid.UUID `json:"product_uuid"`
	Quantity    int       `json:"quantity"`
}
