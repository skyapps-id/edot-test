package warehouse

import (
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
	"gopkg.in/guregu/null.v4"
)

type CreateWarehouseRequest struct {
	Name     string    `json:"name" validate:"required"`
	Address  string    `json:"address" validate:"required"`
	ShopUUID uuid.UUID `json:"shop_uuid" validate:"required"`
}

type CreateWarehouseResponse struct {
	Name string `json:"name"`
}

type GetWarehousesRequest struct {
	Name    null.String `query:"name"`
	Page    null.Int    `query:"page"`
	PerPage null.Int    `query:"per_page"`
	Sort    string      `query:"sort" validate:"oneof='ASC' 'DESC'"`
}

type DataWarehouses struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetWarehousesResponse struct {
	List       []DataWarehouses    `json:"list"`
	Pagination response.Pagination `json:"pagination"`
}

type GetWarehouseRequest struct {
	UUID uuid.UUID `param:"uuid"`
}

type GetWarehouseResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateWarehouseProductRequest struct {
	WarehouseUUID uuid.UUID `json:"warehouse_uuid" validate:"required"`
	ProductUUID   uuid.UUID `json:"product_uuid" validate:"required"`
	Quantity      int       `json:"quantity" validate:"required"`
}

type CreateWarehouseProductResponse struct {
	WarehouseUUID uuid.UUID `json:"warehouse_uuid"`
	ProductUUID   uuid.UUID `json:"product_uuid"`
	Quantity      int       `json:"quantity"`
}

type GetWarehouseProductRequest struct {
	ProductUUIDs []uuid.UUID `json:"product_uuids" validate:"required"`
}

type GetWarehouseProductResponse struct {
	UUID          uuid.UUID `json:"uuid"`
	ShopUUID      uuid.UUID `json:"shop_uuid"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid"`
	ProductUUID   uuid.UUID `json:"product_uuid"`
	Quantity      int       `json:"quantity"`
}

type DataProductStock struct {
	ProductUUID   uuid.UUID `json:"product_uuid" validate:"required"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid" validate:"required"`
	Quantity      int       `json:"quantity" validate:"required,gt=0"`
}
type ProductStockAdditionRequest struct {
	Products []DataProductStock `json:"products" validate:"required,dive"`
}

type ProductStockAdditionResponse struct {
	ProductUUIDs []uuid.UUID `json:"product_uuids"`
}

type ProductStockReductionRequest struct {
	Products []DataProductStock `json:"products" validate:"required,dive"`
}

type ProductStockReductionResponse struct {
	ProductUUIDs []uuid.UUID `json:"product_uuids"`
}

type GetProductStockRequest struct {
	ProductUUID uuid.UUID `param:"uuid"`
}

type GetProductStockResponse struct {
	ProductUUID uuid.UUID `json:"product_uuid"`
	Quantity    int       `json:"quantity"`
}

type WarehouseUpdateActiveRequest struct {
	UUID   uuid.UUID `param:"uuid"`
	Status string    `param:"status"`
}

type WarehouseUpdateActiveResponse struct {
	UUID uuid.UUID `json:"uuid"`
}

type ProductRestockReqeust struct {
	ProductUUID   uuid.UUID `json:"product_uuid" validate:"required"`
	WarehouseUUID uuid.UUID `json:"warehouse_uuid" validate:"required"`
	Quantity      int       `json:"quantity" validate:"required,gt=0"`
}

type ProductRestockResponse struct {
	UUID uuid.UUID `json:"uuid"`
}
