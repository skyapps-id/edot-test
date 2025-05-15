package util

import (
	"math"

	"github.com/skyapps-id/edot-test/product-service/pkg/response"
)

func Pagination(page, perPage int, count int64) response.Pagination {
	return response.Pagination{
		CurrentPage: page,
		PerPage:     perPage,
		TotalPage:   math.Ceil(float64(count) / float64(perPage)),
		TotalData:   count,
	}
}
