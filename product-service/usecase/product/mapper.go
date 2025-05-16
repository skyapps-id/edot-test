package product

import (
	"github.com/skyapps-id/edot-test/product-service/entity"
)

func (uc *usecase) productsMapper(results []entity.Product) (resp []DataProducts) {
	for _, row := range results {
		product := DataProducts{
			UUID:      row.UUID,
			Name:      row.Name,
			SKU:       row.SKU,
			Price:     row.Price,
			ImageURL:  row.ImageURL,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
		resp = append(resp, product)
	}
	return
}

func (uc *usecase) productMapper(result entity.Product) (resp GetProductResponse) {
	return GetProductResponse{
		UUID:      result.UUID,
		Name:      result.Name,
		SKU:       result.SKU,
		Price:     result.Price,
		ImageURL:  result.ImageURL,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}

func (uc *usecase) productByUUidsMapper(results []entity.Product) (resp []GetProductResponse) {
	for _, row := range results {
		product := GetProductResponse{
			UUID:      row.UUID,
			Name:      row.Name,
			SKU:       row.SKU,
			Price:     row.Price,
			ImageURL:  row.ImageURL,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
		resp = append(resp, product)
	}
	return
}
