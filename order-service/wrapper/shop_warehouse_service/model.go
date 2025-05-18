package shop_warehouse_service

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
)

type ProductStockRequest struct {
	Uuids []uuid.UUID `json:"product_uuids"`
}

func (t ProductStockRequest) Json() []byte {
	requestBody, err := json.Marshal(t)
	if err != nil {
		logger.Log.Error("error marsal ProductRequest.Json")
		return []byte{}
	}
	return requestBody
}

type ProductStockResponse struct {
	UUID          string    `json:"uuid"`
	ShopUUID      uuid.UUID `json:"shop_uuid"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid"`
	ProductUUID   string    `json:"product_uuid"`
	Quantity      int       `json:"quantity"`
}

type DataProductStock struct {
	ProductUUID   uuid.UUID `json:"product_uuid"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid"`
	Quantity      int       `json:"quantity"`
}

type ProductStockAdditionRequest struct {
	Products []DataProductStock `json:"products"`
}

func (t ProductStockAdditionRequest) Json() []byte {
	requestBody, err := json.Marshal(t)
	if err != nil {
		logger.Log.Error("error marsal ProductStockAdditionRequest.Json")
		return []byte{}
	}
	return requestBody
}

type ProductStockAdditionResponse struct {
	ProductUUIDs []uuid.UUID `json:"product_uuids"`
}

type ProductStockReductionRequest struct {
	Products []DataProductStock `json:"products"`
}

func (t ProductStockReductionRequest) Json() []byte {
	requestBody, err := json.Marshal(t)
	if err != nil {
		logger.Log.Error("error marsal ProductStockReductionRequest.Json")
		return []byte{}
	}
	return requestBody
}

type ProductStockReductionResponse struct {
	ProductUUIDs []uuid.UUID `json:"product_uuids"`
}
