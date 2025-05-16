package warehouse

import (
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
)

func (uc *usecase) warehousesMapper(results []entity.Warehouse) (data []DataWarehouses) {
	for _, row := range results {
		warehouse := DataWarehouses{
			UUID:      row.UUID,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
		data = append(data, warehouse)
	}
	return
}

func (uc *usecase) warehouseMapper(result entity.Warehouse) (data GetWarehouseResponse) {
	return GetWarehouseResponse{
		UUID:      result.UUID,
		Name:      result.Name,
		Address:   result.Address,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}

func (uc *usecase) warehouseProductMapper(result entity.WarehouseProduct) (data GetWarehouseProductResponse) {
	return GetWarehouseProductResponse{
		UUID:          result.UUID,
		WarehouseUUID: result.WarehouseUUID,
		ProductUUID:   result.ProductUUID,
		Quantity:      result.Quantity,
		UpdatedAt:     result.UpdatedAt,
	}
}
