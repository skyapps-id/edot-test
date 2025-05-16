package warehouse

import (
	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
)

func (uc *usecase) warehousesMapper(results []entity.Warehouse) (resp []DataWarehouses) {
	for _, row := range results {
		warehouse := DataWarehouses{
			UUID:      row.UUID,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
		resp = append(resp, warehouse)
	}
	return
}

func (uc *usecase) warehouseMapper(result entity.Warehouse) (resp GetWarehouseResponse) {
	return GetWarehouseResponse{
		UUID:      result.UUID,
		Name:      result.Name,
		Address:   result.Address,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}

func (uc *usecase) warehouseProductMapper(req GetWarehouseProductRequest, results []entity.WarehouseProduct) (resp map[uuid.UUID]GetWarehouseProductResponse) {
	resp = make(map[uuid.UUID]GetWarehouseProductResponse)
	mapData := make(map[uuid.UUID]entity.WarehouseProduct)

	for _, row := range results {
		mapData[row.ProductUUID] = row
	}

	for _, productUUID := range req.ProductUUIDs {
		if record, ok := mapData[productUUID]; ok {
			resp[productUUID] = GetWarehouseProductResponse{
				UUID:          record.UUID,
				ShopUUID:      record.ShopUUID,
				WarehouseUUID: record.WarehouseUUID,
				ProductUUID:   record.ProductUUID,
				Quantity:      record.Quantity,
			}
		}
	}

	return
}
