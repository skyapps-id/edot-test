package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

type WarehouseHandler interface {
	Create(c echo.Context) (err error)
	Gets(c echo.Context) (err error)
	Get(c echo.Context) (err error)
	CreateWarehouseProduct(c echo.Context) (err error)
	GetWarehouseProduct(c echo.Context) (err error)
}

type handler struct {
	warehouseUsecase warehouse.WarehouseUsecase
}

func NewHandler(usecase warehouse.WarehouseUsecase) WarehouseHandler {
	return &handler{
		warehouseUsecase: usecase,
	}
}
