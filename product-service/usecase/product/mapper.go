package product

import (
	"github.com/skyapps-id/edot-test/product-service/entity"
)

func (uc *usecase) productsMapper(results []entity.Product) (data []DataProducts) {
	for _, row := range results {
		product := DataProducts{
			Name:     row.Name,
			SKU:      row.SKU,
			Price:    row.Price,
			ImageURL: row.ImageURL,
		}
		data = append(data, product)
	}
	return
}
