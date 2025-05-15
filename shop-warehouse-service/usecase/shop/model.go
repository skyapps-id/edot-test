package shop

import (
	"time"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/response"
	"gopkg.in/guregu/null.v4"
)

type CreateShopRequest struct {
	Name    string      `json:"name" validate:"required"`
	Address null.String `json:"address"`
}

type CreateShopResponse struct {
	Name string `json:"name"`
}

type GetShopsRequest struct {
	Name    null.String `query:"name"`
	Page    null.Int    `query:"page"`
	PerPage null.Int    `query:"per_page"`
	Sort    string      `query:"sort" validate:"oneof='ASC' 'DESC'"`
}

type DataShops struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetShopsResponse struct {
	List       []DataShops         `json:"list"`
	Pagination response.Pagination `json:"pagination"`
}

type GetShopRequest struct {
	UUID uuid.UUID `param:"uuid"`
}

type GetShopResponse struct {
	UUID      uuid.UUID   `json:"uuid"`
	Name      string      `json:"name"`
	Address   null.String `json:"address"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
