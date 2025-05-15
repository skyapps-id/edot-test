package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/product-service/pkg/response"
	"gopkg.in/guregu/null.v4"
)

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	SKU         string  `json:"sku" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

type CreateProductResponse struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type GetProductsRequest struct {
	Name    null.String `query:"name"`
	Page    null.Int    `query:"page"`
	PerPage null.Int    `query:"per_page"`
	Sort    string      `query:"sort" validate:"oneof='ASC' 'DESC'"`
}

type DataProducts struct {
	UUID      uuid.UUID   `json:"uuid"`
	Name      string      `json:"name"`
	SKU       string      `json:"sku"`
	Price     float64     `json:"price"`
	ImageURL  null.String `json:"image_url"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type GetProductsResponse struct {
	List       []DataProducts      `json:"list"`
	Pagination response.Pagination `json:"pagination"`
}

type GetProductRequest struct {
	UUID uuid.UUID `param:"uuid"`
}

type GetProductResponse struct {
	UUID      uuid.UUID   `json:"uuid"`
	Name      string      `json:"name"`
	SKU       string      `json:"sku"`
	Price     float64     `json:"price"`
	ImageURL  null.String `json:"image_url"`
	Stock     int         `json:"stock"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
