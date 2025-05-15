package shop

import (
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
)

func (uc *usecase) shopsMapper(results []entity.Shop) (data []DataShops) {
	for _, row := range results {
		shop := DataShops{
			UUID:      row.UUID,
			Name:      row.Name,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
		data = append(data, shop)
	}
	return
}

func (uc *usecase) shopMapper(result entity.Shop) (data GetShopResponse) {
	return GetShopResponse{
		UUID:      result.UUID,
		Name:      result.Name,
		Address:   result.Address,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}
